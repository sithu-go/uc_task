package utils

import (
	"gorm.io/gorm"
)

type Criteria interface {
	ToORM(db *gorm.DB) (*gorm.DB, error)
}

type Page struct {
	Page, PageSize int
}

func (c *Page) ToORM(db *gorm.DB) (*gorm.DB, error) {
	if c.Page < 1 {
		c.Page = 1
	}
	if c.PageSize < 1 {
		c.PageSize = 10
	}

	offset := (c.Page - 1) * c.PageSize

	return db.Offset(offset).Limit(c.PageSize), nil
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	return func(db *gorm.DB) *gorm.DB {
		crit := &Page{
			Page:     page,
			PageSize: pageSize,
		}
		res, _ := crit.ToORM(db)
		return res
	}
}
