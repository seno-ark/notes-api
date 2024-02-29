# notes-api
Backend Test Submission for Perqara.

This API facilitates CRUD operations for managing notes. 
It is crafted with Golang and PostgreSQL, adhering to clean architecture principles to accentuate code testability. 
The routing is handled using [Chi](https://github.com/go-chi/chi) , and database operations are managed through [Gorm](https://github.com/go-gorm/gorm).

## How to start

Clone this repository:
```git clone https://github.com/seno-ark/notes-api.git```

Install all dependencies:
```
cd notes-api
go mod tidy
```

Make sure to copy `.env-example` file to `.env` and set all values as you need.

To run the application, we need to start a PostgreSQL database first.

For simplicity, we can use this PostgreSQL Alpine image to create the database:
```
docker pull postgres:16.1-alpine
docker run --name postgres-alpine-notes -p 5432:5432 -e POSTGRES_DB=notes -e POSTGRES_USER=dev_user -e POSTGRES_PASSWORD=s3CR*t -d postgres:16.1-alpine
```

Then we need to execute the database migration:
```
make migrate-up
```

Now we can start the application:
```
make go-dev
```

If you prefer running the build version:
```
make go-build
./bin/api
```

## Api Test & Documentation

Available Endpoints:
- `POST /notes ` for create a new note from json body
- `PUT /notes/:note_id` for update an existing note from json body
- `GET /notes/:note_id` for get a specific note
- `GET /notes` for get list of note filtered using query params
- `DELETE /notes/:note_id` for delete an existing note

After running the api server, you can test the api endpoints using swagger from your browser: [http://localhost:9000/swagger/index.html](http://localhost:9000/swagger/index.html)

I've also created a Postman collection here: [perqara-notes-api.postman_collection.json](https://github.com/seno-ark/notes-api/blob/master/perqara-notes-api.postman_collection.json)

## Commands
These commands are available in Makefile:
- go-dev
  Running the application in dev mode
- go-lint:
	Run revive linter. Need install https://github.com/mgechev/revive
- go-test:
	Run unit test
- go-mock:
	Run Mockery to generate mocks for all abstractions. Need install https://github.com/vektra/mockery
- go-swag:
	Run Swaggo to generate swagger documentation. Need install https://github.com/swaggo/swag
- go-build:
	Build application into executable file
- migrate-file:
	Run migrate to generate sql migration file. Need install https://github.com/golang-migrate/migrate
- migrate-up:
	Run migrate to execute database migration.
- migrate-down:
	Run migrate to cancel previous database migration.
