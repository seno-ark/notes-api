// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "notes-api/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// NoteRepository is an autogenerated mock type for the NoteRepository type
type NoteRepository struct {
	mock.Mock
}

// CreateNote provides a mock function with given fields: ctx, payload
func (_m *NoteRepository) CreateNote(ctx context.Context, payload *entity.Note) (string, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateNote")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) (string, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) string); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Note) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNote provides a mock function with given fields: ctx, noteID
func (_m *NoteRepository) DeleteNote(ctx context.Context, noteID string) error {
	ret := _m.Called(ctx, noteID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteNote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, noteID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNote provides a mock function with given fields: ctx, noteID
func (_m *NoteRepository) GetNote(ctx context.Context, noteID string) (*entity.Note, error) {
	ret := _m.Called(ctx, noteID)

	if len(ret) == 0 {
		panic("no return value specified for GetNote")
	}

	var r0 *entity.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Note, error)); ok {
		return rf(ctx, noteID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Note); ok {
		r0 = rf(ctx, noteID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, noteID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNoteList provides a mock function with given fields: ctx, filter
func (_m *NoteRepository) GetNoteList(ctx context.Context, filter *entity.GetNoteListFilter) ([]*entity.Note, int64, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetNoteList")
	}

	var r0 []*entity.Note
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.GetNoteListFilter) ([]*entity.Note, int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.GetNoteListFilter) []*entity.Note); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.GetNoteListFilter) int64); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *entity.GetNoteListFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateNote provides a mock function with given fields: ctx, payload
func (_m *NoteRepository) UpdateNote(ctx context.Context, payload *entity.Note) (string, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for UpdateNote")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) (string, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) string); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Note) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNoteRepository creates a new instance of NoteRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNoteRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *NoteRepository {
	mock := &NoteRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}