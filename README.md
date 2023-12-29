# simple-user-auth-api

A simple user authentication API.

## Packages

- testify from stretchr - used for asserting on tests.
- uuid from google - used to create and parse uuids.
- crypto from golang - used to hash and compare the user password.
- chi from go-chi - used to manage the routes and provide useful middlewares.
- pq from lib - the PostgreSQL driver for golang.
- viper from spf13 - used to read the .env configs and load it into a variable.
- swag from swaggo - used to generate the swagger docs from the code comments.