package ds

import "gorm.io/gorm"

type DataSource struct {
	DB *gorm.DB
}

func NewDataSource() (*DataSource, error) {
	db, err := LoadMySqlDB()
	if err != nil {
		return nil, err
	}

	return &DataSource{
		DB: db,
	}, nil
}
