package generic_repo

import (
	"gorm.io/gorm"
	"math"
	pagination_utils "wabustock/pkg/utils/pagination-utils"
)

func Paginate(value interface{}, pagination *pagination_utils.PaginationRequest, paginationResponse *pagination_utils.PaginationResponse, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	paginationResponse.TotalElements = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Rows)))
	paginationResponse.TotalPages = totalPages
	paginationResponse.CurrentPageIndex = pagination.Page

	// since total no. of elements actual elements can't be GET without executing query, we're calculating it with assumptions
	convertedRows := int(totalRows)
	totalExpectedElements := pagination.Page * pagination.Rows
	if totalExpectedElements > convertedRows {
		paginationResponse.NoOfElements = convertedRows - (pagination.Page-1)*pagination.Rows
	} else {
		paginationResponse.NoOfElements = pagination.Rows
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
