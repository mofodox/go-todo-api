# Todo app

This is a personal project on creating a todo list API.

## To Get Started

> Before you run the project, rename the `env.sample` to `env.local` and include your database dsn. If not, you will encounter an error.

To get started, clone the project onto your local. Once you have cloned the project, use `go mod download` to download the project's dependencies. After the dependencies has been downloaded, you can run the project `go run ./cmd/api`.

## File Structure

```
.
|-- bin // contains our compiled application binaries, ready for deployment to a production server
|-- cmd
|    |-- api // contains the application specific code. This includes the code for running the server, reading and writing http requests and managing authentication
|-- internal // contains various packages used by the API and the code interacting with the database, data validation, sending emails etc.
|-- migrations // contains SQL migration files for the database
|-- remote // contains the configuration files and setup scripts for the production server
|-- go.mod // to declare the project dependencies, versions and module path
|-- Makefile // contains the recipes for automating common admin tasks e.g. audit the Go code, building binaries and executing database migrations
```

## Stack

1. `httprouter` – For routing
2. `pq` – For PostgreSQL database driver
3. `godotenv` – To load environment variables
4. `golang-migrate` – SQL migrations

## Available Endpoints

| Method | URL Pattern         | Handler            | Action                                |
| ------ | ------------------- | ------------------ | ------------------------------------- |
| GET    | /api/v1/healthcheck | healthcheckHandler | Show application information          |
| POST   | /api/v1/todos       | createTodoHandler  | Create a new todo                     |
| GET    | /api/v1/todos/:id   | showTodoHandler    | Show the details of a specific todo   |
| PUT    | /api/v1/todos/:id   | updateTodoHandler  | Update the details of a specific todo |
| DELETE | /api/v1/todos/:id   | deleteTodoHandler  | Delete a specific todo                |

## Note

To run the application on a different port and environment, we can use the `-port` and `-env` flags:

```bash
go run ./cmd/api -port=3030 -env=production
```
