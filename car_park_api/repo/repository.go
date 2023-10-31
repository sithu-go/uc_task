package repo

import "uc_task/car_park_api/ds"

type Repository struct {
	DS      *ds.DataSource
	CarPark *carParkRepository
}

func NewRepository(ds *ds.DataSource) *Repository {
	carParkRepo := newCarParkRepository(ds)

	return &Repository{
		DS:      ds,
		CarPark: carParkRepo,
	}
}
