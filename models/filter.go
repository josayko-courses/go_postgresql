package models

import "math"

type Filter struct {
	Page, PageSize int
}

// Limit is the number of items to retrieve
func (f Filter) Limit() int {
	return f.PageSize
}

// Offset is the number of items to skip from the beginning
func (f Filter) Offset() int {
	return (f.Page - 1) * f.PageSize
}

type Metadata struct {
	CurrentPage, PageSize, FirstPage, LastPage, TotalRecords int
}

func ComputeMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}

}
