package cronjob

import (
	"log"
	"uc_task/car_park_api/models"
	"uc_task/car_park_api/repo"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type cronPool struct {
	cronJob *cron.Cron
	repo    *repo.Repository
	db      *gorm.DB
}

func NewCronPool(repo *repo.Repository) *cronPool {
	cronJob := cron.New()

	return &cronPool{
		cronJob: cronJob,
		repo:    repo,
		db:      repo.DS.DB,
	}
}

func (cp *cronPool) StartCronPool() {
	// cp.cronJob.AddFunc("* * * * *", cp.updateCarParkInformation) // at start of every minute
	// cp.cronJob.AddFunc("@every 3s", cp.updateVacancyData) // at every 3s

	cp.cronJob.AddFunc("* 1 * * *", func() {
		go cp.updateCarParkInformation()
	}) // at start of every 1am
	// cp.cronJob.AddFunc("* * * * *", cp.updateCarParkInformation) // at start of every 1am

	cp.cronJob.AddFunc("*/5 * * * *", cp.updateVacancyData)

	go cp.runAtStartup()

	cp.cronJob.Start() // Start the cron job scheduler
}

func (cp *cronPool) runAtStartup() {
	cp.updateCarParkInformation()
	cp.updateVacancyData()
}

func (cp *cronPool) updateCarParkInformation() {
	log.Println("updateCarParkInformation START!")
	// Todo
	// get data from rest
	carParks, err := CollectCarParkInformation()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// update in table

	if err := cp.repo.CarPark.CreateOrUpdateCarParks(carParks); err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("updateCarParkInformation DONE!")

}

// commit one by one
func (cp *cronPool) updateVacancyData() {
	log.Println("updateVacancyData START!")
	ids, err := cp.repo.CarPark.GetAllParkIDs()
	if err != nil {
		log.Println(err)
		return
	}

	for _, id := range ids {
		vacancyData, err := CollectVacancyInformation(id)
		if err != nil {
			log.Println(err.Error())
			continue // Continue to the next ID
		}

		// Start a transaction
		tx := cp.db.Begin()
		if tx.Error != nil {
			log.Println(tx.Error)
			continue
		}

		for _, carPark := range vacancyData.CarPark {
			for _, vehicleTypeDto := range carPark.VehicleType {
				// Create or update VehicleType
				vehicleType := models.VehicleType{
					CarParkID: carPark.ParkID,
					Type:      models.VehicleTypeEnum(vehicleTypeDto.Type),
				}

				var existingVehicleType models.VehicleType
				if tx.Where("car_park_id = ? AND type = ?", vehicleType.CarParkID, vehicleType.Type).First(&existingVehicleType).Error == nil {
					if err := tx.Model(&existingVehicleType).Updates(&vehicleType).Error; err != nil {
						log.Println(err.Error())
						tx.Rollback() // Rollback the transaction and exit
						continue
					}
					vehicleType = existingVehicleType
				} else {
					if err := tx.Create(&vehicleType).Error; err != nil {
						log.Println(err.Error())
						tx.Rollback() // Rollback the transaction and exit
						continue
					}
				}

				for _, serviceCategoryDto := range vehicleTypeDto.ServiceCategory {
					// Create or update ServiceCategory
					serviceCategory := models.ServiceCategory{
						VehicleTypeID:  vehicleType.ID,
						Category:       models.CategoryEnum(serviceCategoryDto.Category),
						VacancyType:    models.VacancyTypeEnum(serviceCategoryDto.VacancyType),
						CurrentVacancy: serviceCategoryDto.Vacancy,
					}

					var existingServiceCategory models.ServiceCategory
					if tx.Where("vehicle_type_id = ? AND category = ? AND vacancy_type = ?", serviceCategory.VehicleTypeID, serviceCategory.Category, serviceCategory.VacancyType).First(&existingServiceCategory).Error == nil {
						if err := tx.Model(&existingServiceCategory).Updates(&serviceCategory).Error; err != nil {
							log.Println(err.Error())
							tx.Rollback() // Rollback the transaction and exit
							continue
						}
					} else {
						if err := tx.Create(&serviceCategory).Error; err != nil {
							log.Println(err.Error())
							tx.Rollback() // Rollback the transaction and exit
							continue
						}
					}
				}
			}
		}

		// Commit the transaction
		tx.Commit()
	}

	log.Println("updateVacancyData DONE!")

}

// just one whole commit
// func (cp *cronPool) updateVacancyData() {
// 	log.Println("updateVacancyData")
// 	ids, err := cp.repo.CarPark.GetAllParkIDs()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Start a transaction
// 	tx := cp.db.Begin()
// 	if tx.Error != nil {
// 		log.Println(tx.Error)
// 		return
// 	}

// 	for _, id := range ids {
// 		vacancyData, err := CollectVacancyInformation(id)
// 		if err != nil {
// 			log.Println(err.Error())
// 			continue // Continue to the next ID
// 		}

// 		for _, carPark := range vacancyData.CarPark {
// 			for _, vehicleTypeDto := range carPark.VehicleType {
// 				// Create or update VehicleType
// 				vehicleType := models.VehicleType{
// 					CarParkID: carPark.ParkID,
// 					Type:      models.VehicleTypeEnum(vehicleTypeDto.Type),
// 				}

// 				var existingVehicleType models.VehicleType
// 				if tx.Where("car_park_id = ? AND type = ?", vehicleType.CarParkID, vehicleType.Type).First(&existingVehicleType).Error == nil {
// 					if err := tx.Model(&existingVehicleType).Updates(&vehicleType).Error; err != nil {
// 						log.Println(err.Error())
// 						tx.Rollback() // Rollback the transaction and exit
// 						return
// 					}
// 					vehicleType = existingVehicleType
// 				} else {
// 					if err := tx.Create(&vehicleType).Error; err != nil {
// 						// "type": "O", it won't run
// 						// "type": "O",
// 						// [1.196ms] [rows:0] INSERT INTO `vehicle_types` (`car_park_id`,`type`) VALUES ('tdc15p5','O')
// 						// 2023/11/01 14:15:38 cronpool.go:186: Error 1265 (01000): Data truncated for column 'type' at row 1
// 						log.Println(err.Error())
// 						tx.Rollback() // Rollback the transaction and exit
// 						return
// 					}
// 				}

// 				for _, serviceCategoryDto := range vehicleTypeDto.ServiceCategory {
// 					// Create or update ServiceCategory
// 					serviceCategory := models.ServiceCategory{
// 						VehicleTypeID:  vehicleType.ID,
// 						Category:       models.CategoryEnum(serviceCategoryDto.Category),
// 						VacancyType:    models.VacancyTypeEnum(serviceCategoryDto.VacancyType),
// 						CurrentVacancy: serviceCategoryDto.Vacancy,
// 					}

// 					var existingServiceCategory models.ServiceCategory
// 					if tx.Where("vehicle_type_id = ? AND category = ? AND vacancy_type = ?", serviceCategory.VehicleTypeID, serviceCategory.Category, serviceCategory.VacancyType).First(&existingServiceCategory).Error == nil {
// 						if err := tx.Model(&existingServiceCategory).Updates(&serviceCategory).Error; err != nil {
// 							log.Println(err.Error())
// 							tx.Rollback() // Rollback the transaction and exit
// 							return
// 						}
// 					} else {
// 						if err := tx.Create(&serviceCategory).Error; err != nil {
// 							log.Println(err.Error())
// 							tx.Rollback() // Rollback the transaction and exit
// 							return
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	// Commit the transaction
// 	tx.Commit()
// }
