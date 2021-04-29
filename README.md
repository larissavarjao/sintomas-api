# sintomas-api

## Structure

This project is structured in the main folder with:
- `api`: folder to hold the structure code which follows the structure of a `entity.go` (for model the entity), `service.go` (who put the services and connect with the database);
- `handlers`:  abstractions before the intended code on the database is executed;
- `middlewares`: middlewares where we can run functions in all api structure, like authentication;
- `main.go` file starts the api and hold the structure of the code and the API;
- `go.mod`: file defines the module’s module path, which is also the import path used for the root directory;
- `go.sum`: contain the expected cryptographic hashes of the content of specific module versions.

## To run the project:

```shell
npm i -g nodemon
nodemon --exec go run main.go --signal SIGTERM
```