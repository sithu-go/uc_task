package cronjob

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"uc_task/car_park_api/dto"
	"uc_task/car_park_api/models"

	"github.com/jinzhu/copier"
)

func CollectCarParkInformation() ([]*models.CarPark, error) {
	// // read from file
	// source := "./car_park_api/data/basic_info_all.json"

	// // Check if the file exists
	// if _, err := os.Stat(source); err != nil {
	// 	return nil, err
	// }

	// // Read data from the local file
	// data, err := os.ReadFile(source)
	// if err != nil {
	// 	return nil, err
	// }

	// get data from rest api
	source := "https://resource.data.one.gov.hk/td/carpark/basic_info_all.json"
	// source := "https://resource.data.one.gov.hk/td/carpark/basic_info_tdc103p2.json"

	// Make an HTTP GET request to the API
	response, err := http.Get(source)
	if err != nil {
		log.Println("http request problem")
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("http response problem")
		return nil, fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	// Read the response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("io read problem")
		return nil, err
	}

	// log.Println("Response body:", string(data))
	// c/68a6690e-83eb-4851-9170-b03edab88b9e
	// BOM is to specify whether the text is encoded as little-endian or big-endian, which is important for correctly interpreting multi-byte characters.
	// UTF-8 byte sequence `EF BB BF` for little-endian
	// UTF-8 byte sequence `FE FF` for big-endian

	// The character 'ï' is often indicative of a byte order mark (BOM) or
	// some other non-printable character that is not part of valid JSON.
	trimmedData := strings.TrimPrefix(string(data), "\xef\xbb\xbf") // Remove UTF-8 BOM

	var carParksData dto.APICarParks
	// Unmarshal the JSON data into the 'carParksData' variable
	// if err := json.Unmarshal(data, &carParksData); err != nil {
	if err := json.Unmarshal([]byte(trimmedData), &carParksData); err != nil {
		log.Println("unmarshalling problem")
		return nil, err
	}

	var carParks []*models.CarPark
	if err := copier.Copy(&carParks, &carParksData.CarPark); err != nil {
		log.Println("Copier problem")
		return nil, err
	}

	return carParks, nil
}
