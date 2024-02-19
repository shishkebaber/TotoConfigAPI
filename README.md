### TotoConfigAPI - Test Task
TotoConfigAPI is a Go-based web service designed for managing configuration settings, developed as a test task for a Golang developer role. It leverages MongoDB for data storage, with the database environment fully configured in Docker Compose.

## Features
RESTful API for configuration management.
Automatic country code resolution for requests.
Dockerized setup including MongoDB.

## Requirements
Docker and Docker Compose

## Quick Start
Clone the repository and navigate into the project directory.

Start the service with Docker Compose:

``` docker-compose up --build ```

The API is now accessible at http://localhost:8080/api/configs

## API Usage
Retrieve configurations:

``` GET /api/configs?package=<package_name> ```

## Testing

Run unit tests by executing:

``` go test ./... ```

## Note

This project serves as a test task for a Golang developer role.