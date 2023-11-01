package cronjob

import (
	"log"
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
}

func (cp *cronPool) updateCarParkInformation() {
	log.Println("updateCarParkInformation")
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

}

func (cp *cronPool) updateVacancyData() {
	log.Println("updateVacancyData")

}
