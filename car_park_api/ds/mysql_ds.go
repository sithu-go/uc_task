package ds

import (
	"fmt"
	"log"
	"os"
	"uc_task/car_park_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Error won't pass to invoker if happen, just gonna fatal
func LoadMySqlDB() (*gorm.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	name := os.Getenv("MYSQL_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error(), dsn)
		return nil, err
	}

	log.Println("Successfully connected to MySQL")

	// migrate DB
	err = db.AutoMigrate(
		&models.CarPark{},
		&models.VehicleType{},
		&models.ServiceCategory{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
