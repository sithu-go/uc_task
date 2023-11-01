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

func (r *carParkRepository) FindAll(req *dto.CarParkReq) ([]*models.CarPark, error) {
	db := r.DB.Model(&models.CarPark{})

	// Check and apply search criteria for Name
	if req.Name != nil {
		nameSearch := "%" + *req.Name + "%"
		db = db.Where("name_en LIKE ? OR name_tc LIKE ? OR name_sc LIKE ?", nameSearch, nameSearch, nameSearch)
	}

	// Check and apply search criteria for Address
	if req.Address != nil {
		addressSearch := "%" + *req.Address + "%"
		db = db.Where("display_address_en LIKE ? OR display_address_tc LIKE ? OR display_address_sc LIKE ?", addressSearch, addressSearch, addressSearch)
	}

	// Check and apply search criteria for Latitude, Longitude, and Radius
	if req.Lat != nil && req.Lng != nil {
		if req.Radius == nil {
			// default
			radius := float64(10)
			req.Radius = &radius
		}

		// to filter car parks within the specified geographic bounds
		latMin, latMax, lngMin, lngMax := utils.CalculateBounds(*req.Lat, *req.Lng, *req.Radius)

		db.Where("latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?", latMin, latMax, lngMin, lngMax)
	}

	db.Scopes(utils.Paginate(req.Page, req.PageSize))

	carParks := []*models.CarPark{}
	err := db.Find(&carParks).Error
	return carParks, err
}

func (r *carParkRepository) CreateOrUpdateCarParks(carParks []*models.CarPark) error {
	for _, carPark := range carParks {
		// Attempt to find an existing record by its primary key (e.g., ParkID)
		var existingCarPark models.CarPark
		if err := r.DB.Where("park_id = ?", carPark.ParkID).First(&existingCarPark).Error; err != nil {
			// If the record doesn't exist, create a new one
			if err := r.DB.Create(&carPark).Error; err != nil {
				return err
			}
		} else {
			// Update the existing record with the new data
			if err := r.DB.Model(&existingCarPark).Updates(carPark).Error; err != nil {
				return err
			}
		}
	}

	return nil

}

func (r *carParkRepository) GetAllParkIDs() ([]string, error) {
	var parkIDs []string
	if err := r.DB.Model(&models.CarPark{}).Pluck("ParkID", &parkIDs).Error; err != nil {
		return nil, err
	}

	return parkIDs, nil
}
