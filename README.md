# go-crud-api

This is a CRUD API written in Go.

## Features

- RESTful API with routing handled by the Gorilla mux router


## Usage

go run main.go

This will start the server on port 8080.

# development
The API will start on port 8000.

### Endpoints

The API exposes the following endpoints:

- `GET /movies` - Get all movies
- `POST /movies` - Create a new entry 
- `GET /movies/{id}` - Get a single movie 
- `PUT /movies/{id}` - Update a movie
- `DELETE /movies/{id}` - Delete a movie

