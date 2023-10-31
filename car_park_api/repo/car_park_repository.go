package repo

import (
	"uc_task/car_park_api/ds"
	"uc_task/car_park_api/dto"
	"uc_task/car_park_api/models"
	"uc_task/car_park_api/utils"

	"gorm.io/gorm"
)

type carParkRepository struct {
	DB *gorm.DB
}

func newCarParkRepository(ds *ds.DataSource) *carParkRepository {
	return &carParkRepository{
		DB: ds.DB,
	}
}

func (r *carParkRepository) FindByID(id uint64) (*models.CarPark, error) {
	carPark := models.CarPark{}
	db := r.DB.Model(&models.CarPark{})
	db.Where("id", id)
	err := db.First(&carPark).Error
	return &carPark, err
}

func (r *carParkRepository) FindAll(req *dto.PaginationRequest) ([]*models.CarPark, error) {
	db := r.DB.Model(&models.CarPark{})
	carParks := []*models.CarPark{}

	db.Scopes(utils.Paginate(req.Page, req.PageSize))
	err := db.Find(&carParks).Error
	return carParks, err
}
