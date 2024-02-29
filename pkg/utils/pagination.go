package utils

import (
	"strconv"
)

const (
	maxPaginationPage      = 500
	maxPaginationCount     = 100
	defaultPaginationCount = 10
)

// Pagination is used for convert and validate "page" & "count" parameter
func Pagination(pageStr, countStr string) (page, count int) {

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	if countStr != "" {
		count, _ = strconv.Atoi(countStr)
	}

	if page < 1 {
		page = 1
	} else if page > maxPaginationPage {
		page = maxPaginationPage
	}

	if count < 1 {
		count = defaultPaginationCount
	} else if count > maxPaginationCount {
		count = maxPaginationCount
	}

	return
}
