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
	db.Where("park_id", id)
	err := db.First(&carPark).Error
	return &carPark, err
}

func (r *carParkRepository) FindAll(req *dto.CarParkReq) ([]*models.CarPark, error) {
	db := r.DB.Model(&models.CarPark{})

	if req.ParkID != nil {
		db.Where("park_id", req.ParkID)
	}

	if req.Name != nil {
		nameSearch := "%" + *req.Name + "%"
		db = db.Where("name_en LIKE ? OR name_tc LIKE ? OR name_sc LIKE ?", nameSearch, nameSearch, nameSearch)
	}

	if req.Address != nil {
		addressSearch := "%" + *req.Address + "%"
		db = db.Where("display_address_en LIKE ? OR display_address_tc LIKE ? OR display_address_sc LIKE ?", addressSearch, addressSearch, addressSearch)
	}

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

func (r *carParkRepository) FindVacancyData(req *dto.VacacncyReq) (map[string]any, error) {
	var carParks []*models.CarPark
	db := r.DB.Model(&models.CarPark{})
	// preloading of related associations
	db.Preload("VehicleTypes.ServiceCategories")
	// To get data only related to park_id, not to duplicate park_id
	db.Select("car_parks.park_id, MAX(vehicle_types.type) AS type, MAX(service_categories.category) AS category, SUM(service_categories.current_vacancy) AS current_vacancy, MAX(service_categories.vacancy_type) AS vacancy_type, MAX(service_categories.updated_at) AS updated_at")
	db.Joins("JOIN vehicle_types ON car_parks.park_id = vehicle_types.car_park_id")
	db.Joins("JOIN service_categories ON vehicle_types.id = service_categories.vehicle_type_id")

	if req.ParkID != nil {
		db = db.Where("car_parks.park_id = ?", *req.ParkID)
	}

	if req.VehicleType != nil {
		db = db.Where("vehicle_types.type = ?", *req.VehicleType)
	}

	if req.StartDate != nil && req.EndDate != nil {
		db = db.Where("service_categories.created_at >= ? AND service_categories.created_at <= ?", req.StartDate, req.EndDate)
	}

	db.Group("car_parks.park_id")
	db.Order("MAX(service_categories.updated_at) DESC")

	// this is incompatible with sql_mode=only_full_group_by
	// If you want "ORDER BY service_categories.updated_at DESC"
	// Do below step
	// SET sql_mode = '';
	// db.Order("service_categories.updated_at DESC")

	// Pagination, uncomment if needed
	db = db.Scopes(utils.Paginate(req.Page, req.PageSize))

	if err := db.Debug().Find(&carParks).Error; err != nil {
		return nil, err
	}

	// Organize the data into the desired JSON format
	formattedData := make(map[string]any)
	carParkData := make([]map[string]any, 0)

	for _, carPark := range carParks {
		carParkInfo := make(map[string]any)
		carParkInfo["park_id"] = carPark.ParkID

		vehicleData := make([]map[string]any, 0)
		for _, vehicleType := range carPark.VehicleTypes {
			vehicleInfo := make(map[string]any)
			vehicleInfo["type"] = vehicleType.Type

			serviceCategoryData := make([]map[string]any, 0)
			for _, serviceCategory := range vehicleType.ServiceCategories {
				categoryInfo := make(map[string]any)
				categoryInfo["category"] = serviceCategory.Category
				categoryInfo["vacancy_type"] = serviceCategory.VacancyType
				categoryInfo["vacancy"] = serviceCategory.CurrentVacancy
				categoryInfo["updated_at"] = serviceCategory.UpdatedAt
				serviceCategoryData = append(serviceCategoryData, categoryInfo)
			}

			vehicleInfo["service_category"] = serviceCategoryData
			vehicleData = append(vehicleData, vehicleInfo)
		}

		carParkInfo["vehicle_type"] = vehicleData
		carParkData = append(carParkData, carParkInfo)
	}

	formattedData["car_park"] = carParkData

	return formattedData, nil
}
