package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"notes-api/internal"
	"notes-api/internal/entity"
	appErr "notes-api/pkg/error"
	"notes-api/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) internal.NoteRepository {
	return &noteRepository{
		db: db,
	}
}

func (r *noteRepository) CreateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (string, error) {
	model := Note{}
	model.FromDto(payload)
	model.ID, _ = utils.ULID()

	err := r.db.WithContext(ctx).Table("notes").Create(&model).Error
	if err != nil {
		slog.Error("error repository create note", "err", err)
		return "", appErr.NewErrInternalServer("failed to create note")
	}

	return model.ID, nil
}

func (r *noteRepository) UpdateNote(ctx context.Context, payload *entity.CreateUpdateNotePayload) (string, error) {
	model := Note{}
	model.FromDto(payload)

	err := r.db.WithContext(ctx).Table("notes").Model(&model).Updates(model).Error
	if err != nil {
		slog.Error("error repository update note", "err", err)
		return "", appErr.NewErrInternalServer("failed to update note")
	}

	return model.ID, nil
}

func (r *noteRepository) DeleteNote(ctx context.Context, noteID string) error {
	err := r.db.WithContext(ctx).Table("notes").Where("id = ?", noteID).Delete(&Note{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErr.ErrNotFound
		}
		slog.Error("error repository delete note", "err", err)
		return appErr.NewErrInternalServer("failed to delete note")
	}

	return nil
}

func (r *noteRepository) GetNote(ctx context.Context, noteID string) (*entity.Note, error) {
	model := Note{}

	err := r.db.WithContext(ctx).Table("notes").Where("id = ?", noteID).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.NewErrNotFound("note not found")
		}
		slog.Error("error repository get note", "err", err)
		return nil, appErr.NewErrInternalServer("failed to get note")
	}

	return model.ToEntity(), nil
}

func (r *noteRepository) GetNoteList(ctx context.Context, params *entity.GetNoteListParams) ([]*entity.Note, int64, error) {
	var (
		total          int64
		model          = []*Note{}
		order          = "created_at DESC"
		availableSorts = []string{"id", "title", "created_at", "updated_at"}
	)

	if params.Sort != "" {
		validSort := utils.SortValidation(params.Sort, order, availableSorts)
		order = utils.ToSqlSort(validSort)
	}

	query := r.db.WithContext(ctx).Table("notes")

	if params.Search != "" {
		keyword := fmt.Sprintf("%%%v%%", strings.ToLower(params.Search))
		query.Where("LOWER(title) LIKE ? OR LOWER(content) LIKE ?", keyword, keyword)
	}

	err := query.Count(&total).Error
	if err != nil {
		slog.Error("error repository get note list count", "err", err)
		return nil, total, appErr.NewErrInternalServer("failed to get notes")
	}

	if total == 0 {
		return nil, total, nil
	}

	err = query.
		Order(order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&model).Error

	if err != nil {
		slog.Error("error repository get note list data", "err", err)
		return nil, total, appErr.NewErrInternalServer("failed to get notes")
	}

	results := []*entity.Note{}
	for _, m := range model {
		results = append(results, m.ToEntity())
	}

	return results, total, nil
}
