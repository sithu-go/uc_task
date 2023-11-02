
# Car Park Information System Readme

This repository contains a project for collecting and managing data related to car parks, including their basic information, real-time vacancy data. The system is designed to provide information about car parks, check occupancy and changes during specific periods, and identify new car parks.


## Features

- Find Car Parks
    - Enable users to find car parks by name, address, location, and other parameters.
- Check vacancies for parking
- Identify New Car Parks


## Requirements

* mysql
    
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file. You can see example in _./car_park_api/config/.env.example_

`MYSQL_HOST`

`MYSQL_PORT`

`MYSQL_USER`

`MYSQL_PASS`

`MYSQL_NAME`

`APP_PORT`
## API Reference

#### Get information about car parks

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
| `start_date`  | `string`      | For checking recent vacacncy within start_date and end_date e.g. 2022-11-11 or 2022-11-11 10:00:00|
| `end_date`    | `string`      | For checking recent vacacncy within start_date and end_date e.g. 2022-11-13 or 2022-11-13 10:00:00|
| `vehicle_type`| `string`      | **One of** (_P_, _M_, _L_, _H_, _C_, _T_, _B_, _O_, _N_, _P_D_, _M_D_, _L_D_, _H_D_, _C_D_, _T_D_, _B_D_) Type of vehicle.                                    |
| `page`        | `int`     | **Required.** the number of records to skip or page number|
| `page_size`   | `int`     | **Required.** the number of records or page size          |

#### Note:

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


## Authors

- [@sithu-go](https://www.github.com/sithu-go)

