
# Car Park Information System Readme

This repository contains a project for collecting and managing data related to car parks, including their basic information, real-time vacancy data. The system is designed to provide information about car parks, check occupancy and changes during specific periods, and identify new car parks.

## Features

- Find Car Parks
  - Enable users to find car parks by name, address, location, and other parameters.
- Check vacancies for parking
- Identify New Car Parks

## Requirements

- mysql

## Run Locally

Clone the project

```bash
  git clone https://github.com/sithu-go/uc_task
```

Go to the project directory

```bash
  cd uc_task
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run car_park_api/main.go
```

For visualization, change your ip in prometheus.yml

```bash
  docker compose up
```

## Project Structure

The directory structure presented represents "car_park_api". This application manages car park-related data, provides an API for accessing this information, and includes features such as database connectivity, scheduled data collection, and metric monitoring. Key components include configuration settings, route handlers, database models, middleware for CORS support, and utilities for various tasks. Notably, the application uses scheduled cron jobs to periodically collect data from external APIs and update the database with the latest information. This organized structure simplifies the development and maintenance of the car park API.

```tree
car_park_api/
├── main.go                  # Main application entry point
├── config/                  # Configuration files or settings
│   ├── config.go            # Application configuration
├── cronjob/                 # Configuration files or settings
│   ├── collector.go         # Collector for car park information and parking vacancies
│   ├── cronpool.go          # Cron jobs which run at startup, every start of five minutes and 1 am
├── data/                    # Folder for storing data files (e.g., JSON files)
├── ds/                      # Folder for database source
│   ├── mysql_ds.go          # Making database connection and migrating table structure
├── dto/                     # For data transformation and validation
├── handler/                 # Route handlers
│   ├── car_park_handler.go  # handler for car park info and parking vacancy info 
│   ├── metric_handler.go    # handler for metric data powered by prometheus
├── metric/                  # Metrics and monitoring code
│   ├── metrics.go           # Metrics implementation and Metric global variables for global usage
├── middleware/              # Middleware
│   ├── cors_middleware.go   # Allow cross-origin requests while maintaining security
├── models/                  # Database struct/model definitions
│   ├── car_park.go          # GORM model for car park table
│   ├── vehicle_type.go      # GORM model for vehicle type table like which type of car is allowed in car park
│   ├── service_category.go  # GORM model for service category table which has fields like vacancies
├── repo/                    # Database connection setup
│   ├──car_park_repository.go# Repository for making operations on car park table like searches
├── utils/                   # Folder where utility functions are stored
├── README.md                # Documentation
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file. You can see example in _./car_park_api/config/.env.example_

`MYSQL_HOST`

`MYSQL_PORT`

`MYSQL_USER`

`MYSQL_PASS`

`MYSQL_NAME`

`APP_PORT`

## API Reference

### Get information about car parks

```http
  GET /api/carParks
```

| Parameter     | Type      | Description                                                |
| :--------     | :-------  | :----------------------------------------------------------|
| `park_id`     | `string`  | ID of the car park                                         |
| `name`        | `string`  | name of the car park                                       |
| `address`     | `string`  | address of the car park                                    |
| `lat`         | `float64` | latitude of the car park                                   |
| `lng`         | `float64` | longitude of the car park                                  |
| `radius`      | `float64` | within radius of the current location                      |
| `order_by`    | `string`  | **One of** (_ASC_ or _DESC_) to check new car park                                    |
| `page`        | `int`     | **Required.** the number of records to skip or page number |
| `page_size`   | `int`     | **Required.** the number of records or page size           |

#### Get information about vacancy data

```http
  GET /api/carParks/vacancy
```

| Parameter     | Type          | Description                                           |
| :--------     | :-------      | :-----------------------------------------------------|
| `park_id`     | `string`      | ID of the car park                                    |
| `start_date`  | `string`      | For checking recent vacacncy within start_date and end_date e.g. 2022-11-11 or 2022-11-11 10:00:00  |
| `end_date`    | `string`      | For checking recent vacacncy within start_date and end_date e.g. 2022-11-13 or 2022-11-13 10:00:00  |
| `vehicle_type`| `string`      | **One of** (_P_, _M_, _L_, _H_, _C_, _T_, _B_, _O_, _N_, _P_D_, _M_D_, _L_D_, _H_D_, _C_D_, _T_D_, _B_D_) Type of vehicle.  |
| `vacancy_type`| `string`      | **One of** (_A_, _B_, _C_) Availability of parking space|
| `current_vacancy` | `int`     | If vacacncy_type is _A_, _0_ is full, _-1_ is no data. If _B_, _0_ is full, _1_ is available, _-1_ is no data. If _C_, always _0_.          |
| `page`        | `int`         | **Required.** the number of records to skip or page number  |
| `page_size`   | `int`         | **Required.** the number of records or page size          |

#### Note

I don't have route for this because you can search with `park_id`

```http
    GET /api/carParks/${id}
```

```http
    GET /api/carParks/vacancy/${id}
```

#### Get metrics about system information and error reports powered by prometheus

```http
  GET /api/metrics
```

| Metric Name         | Metric Type      | Description                                  | Labels                                       |
| :------------------- | :---------------  | :-------------------------------------------- | :--------------------------------------------------------------    |
| http_requests_total      | CounterVec       | Total number of HTTP requests.                | Labels: method, endpoint, status_code                               |
| http_request_errors_total        | CounterVec       | Total number of HTTP request errors.         | Labels: method, endpoint, status_code                               |
| http_request_duration_seconds     | HistogramVec     | Duration of HTTP requests in seconds.       | Labels: method, endpoint  |
| cron_errors_total    | CounterVec       | Total number of cron job errors and messages.| Labels: job_name, error_message, data(e.g. id)                               |

## Author

- [@sithu-go](https://www.github.com/sithu-go)
