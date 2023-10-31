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
	cp.cronJob.AddFunc("* * * * *", cp.updateVacancyData) // at every start of the minute
	// cp.cronJob.AddFunc("@every 3s", cp.updateVacancyData) // at every 3s

	cp.cronJob.Start() // Start the cron job scheduler
}

func (cp *cronPool) updateVacancyData() {
	log.Println("HEHE")
	// Todo
	// get data from rest
	// convert into json
	// update in table
}
