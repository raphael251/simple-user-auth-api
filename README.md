# Simple User Auth API

### Introduction
A simple user authentication API written in Golang. This is one of my first projects in Go. The purpose is just to know better the language and then start to build more complex apps.

### Project Support Features
* Users can signup and login to their accounts

### Prerequisites

* [Go](https://go.dev/) language.
* [Docker](https://docs.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).
* [Git](https://git-scm.com/).
* [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) (if you are using VSCode and want to run the requests in the `test/user.http` file).

### Installation Guide
1. Clone this repository.
2. Duplicate the `.env.example` file in the `cmd/server` folder and rename it to `.env`.
3. Run the `go mod tidy` command.


### Usage

1. Enter the `cmd/server` folder and run `go run main.go`. If you don't change the initial `.env` file, the application will start on port 3000.

### API Endpoints
| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| POST | /api/v1/users/ | To sign up a new user account |
| POST | /api/v1/users/login | To login an existing user account |

### Technologies Used

* [testify](github.com/stretchr/testify) from stretchr - used for asserting on tests.
* [uuid](github.com/google/uuid) from google - used to create and parse uuids.
* [crypto](golang.org/x/crypto) from golang - used to hash and compare the user password.
* [chi](github.com/go-chi/chi) from go-chi - used to manage the routes and provide useful middlewares.
* [pq](github.com/lib/pq) from lib - the PostgreSQL driver for golang.
* [viper](github.com/spf13/viper) from spf13 - used to read the .env configs and load it into a variable.
* [swag](github.com/swaggo/swag) from swaggo - used to generate the swagger docs from the code comments.

### Authors
* [Raphael Passos](https://github.com/raphael251)

