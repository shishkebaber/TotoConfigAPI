### TotoConfigAPI - Test Task
TotoConfigAPI is a Go-based web service designed for managing configuration settings, developed as a test task for a Golang developer role. It leverages MongoDB for data storage, with the database environment fully configured in Docker Compose.

## Features
RESTful API for configuration management.
Automatic country code resolution for requests.
Dockerized setup including MongoDB.

## Requirements
Docker and Docker Compose

## TODO
Implement other CRUD endpoints, with validation, authentication middleware.
Custom logging 
More comprehencive application config, instead of simple .env usage.
More test-cases, unit tests and integration tests using docker-compose.
Proper IP2Country database storing, probably in some cloud storage.

## Quick Start
Clone the repository and navigate into the project directory.

Start the service with Docker Compose:

``` docker-compose up --build ```

The API is now accessible at http://localhost:8080/api/configs

## API Usage
Retrieve configurations example:

``` GET /api/configs?package=com.softinit.iquitos.mainapp ```

## Testing

Run unit tests by executing:

``` go test ./... ```

## Note

This project serves as a test task for a Golang developer role.