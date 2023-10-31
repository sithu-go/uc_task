package main

import (
	_ "uc_task/car_park_api/config"

	"log"
	"uc_task/car_park_api/ds"
)

func main() {
	_, err := ds.NewDataSource()
	if err != nil {
		log.Fatal(err)
	}

}
