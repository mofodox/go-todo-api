# Todo app

This is a personal project on creating a todo list API.

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

## Note

To run the application on a different port and environment, we can use the `-port` and `-env` flags:

```bash
go run ./cmd/api -port=3030 -env=production
```
