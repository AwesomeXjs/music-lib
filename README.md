![Go Report](https://goreportcard.com/badge/github.com/AwesomeXjs/music-lib)
![Repository Top Language](https://img.shields.io/github/languages/top/AwesomeXjs/music-lib)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/AwesomeXjs/music-lib)
![Github Repository Size](https://img.shields.io/github/repo-size/AwesomeXjs/music-lib)
![Github Open Issues](https://img.shields.io/github/issues/evt/rest-api-example)
![GitHub last commit](https://img.shields.io/github/last-commit/AwesomeXjs/music-lib)
![GitHub contributors](https://img.shields.io/github/contributors/AwesomeXjs/music-lib)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)


<div align="center">
    <img align="center" width="50%" src="./assets/echo-logo.svg">
    <img align="right" width="40%" src="./assets/big-gopher.png">
</div>


# Music library test task
# REST API Server



This is Golang REST API server example including the following features:
*   based on minimalist Go web framework - [Echo](https://echo.labstack.com)
*   made with Clean Architecture (Controller => Service => Repository)
*   has services that work with  PostgreSQL database
*   config based on envconfig with [GoDotEnv](<https://github.com/joho/godotenv>)
*   fastest [Zap](<https://github.com/uber-go/zap>) logger
*   swagger documentation by [Swaggo](<https://github.com/swaggo/swag>)
* Implemented classic CRUD with all the requirements, including working with a third-party service when adding music to the library, [Mokky.dev](https://mokky.dev/) for example



##  [How to start project](#start)


1. to start correctly you will need [Docker](https://www.docker.com/products/docker-desktop/) and preferably "Make tools"
```sh
$ make up
```
or
```sh
$ docker compose -f docker-compose.yml up -d
```
2. After assembly, the server will start and Swagger documentation will become available to you at this path:
```sh
$ http://localhost:9999/swagger/
```

<img align="right" width="100%" src="./assets/swagger.jpg">