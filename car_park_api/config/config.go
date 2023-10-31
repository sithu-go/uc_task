package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// to get file line and path when log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load env file
	err := godotenv.Load("./car_park_api/config/.env")
	if err != nil {
		log.Println("error opening .env file")
		log.Fatalf(err.Error())
	}
}
