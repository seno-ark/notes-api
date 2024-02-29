package utils

import (
	"fmt"
	"strings"
)

func SortValidation(sortBy, defaultSort string, availableSorts []string) string {

	for _, as := range availableSorts {
		if sortBy == as || sortBy == fmt.Sprintf("-%s", as) {
			return sortBy
		}
	}

	return defaultSort
}

func ToSqlSort(sortBy string) string {
	order := "id ASC"

	if strings.HasPrefix(sortBy, "-") && len(sortBy) > 1 {
		order = fmt.Sprintf("%s DESC", sortBy[1:])
	} else if len(sortBy) > 0 {
		order = fmt.Sprintf("%s ASC", sortBy)
	}

	return order
}
