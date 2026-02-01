# Thought REST API

A small REST API for storing and retrieving random thoughts.  
The project is educational, but built with a **production-oriented approach**: clear layering, explicit contracts, database migrations, and graceful shutdown.

---

## Features

- create a thought

- update a thought

- delete a thought

- get a random thought

- readiness endpoint

- PostgreSQL as storage

- SQL migrations via `golang-migrate`


---

## Running the application

### 1. Start PostgreSQL

`docker compose up -d`

PostgreSQL will be available at:

`localhost:20061`

---

### 2. Run database migrations

`migrate \   -path ./migrations \   -database "postgres://postgres:TESTPASSWORD@localhost:20061/postgres?sslmode=disable" \   up`

To check migration status:

`migrate \   -path ./migrations \   -database "postgres://postgres:TESTPASSWORD@localhost:20061/postgres?sslmode=disable" \   version`

---

### 3. Configure the application

Create `config/local.yaml`:

`env: "local"  http_server:   address: "localhost:8888"   timeout: 4s   idle_timeout: 60s  database_url: "postgres://postgres:TESTPASSWORD@localhost:20061/postgres?sslmode=disable"`

Set config path:

`export CONFIG_PATH=./config/local.yaml`

---

### 4. Start the API

`go run cmd/api/main.go`

Expected output:

`Connecting to database Successfully connected to database server started`

---

## API Endpoints

### Readiness check

`GET /ready`

Response:

`200 OK`

---

### Get random thought

`GET /api/v1/thoughts/random`

Response `200 OK`:

`{   "id": 1,   "text": "Better done than perfect",   "author": "Anonymous" }`

If no thoughts exist:

`404 Not Found`

---

### Create thought

`POST /api/v1/thoughts`

Request body:

`{   "text": "If it works, don't touch it",   "author": "Senior Dev" }`

Response `201 Created`:

`{   "id": 2 }`

---

### Update thought

`PUT /api/v1/thoughts/{id}`

Request body:

`{   "text": "Updated thought",   "author": "Author" }`

Response:

`204 No Content`

If the thought does not exist:

`404 Not Found`

---

### Delete thought

`DELETE /api/v1/thoughts/{id}`

Response:

`204 No Content`

---

## Database migrations

Migrations are stored in the `migrations` directory and are executed **outside of the application**.

Apply all migrations:

`migrate -path ./migrations -database "$DATABASE_URL" up`

Rollback one migration:

`migrate -path ./migrations -database "$DATABASE_URL" down 1`
## Design principles

- the application does not manage database schema

- migrations are executed separately

- the service does not start if the database is unavailable

- business logic is isolated from HTTP transport

- domain errors are mapped to HTTP status codes at the boundary


---

## Possible improvements

- OpenAPI / Swagger documentation

- unit tests for use cases

- structured HTTP logging with request id and latency

- database-aware health check

- input length validation

