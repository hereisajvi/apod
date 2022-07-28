# apod 🌌 ![GitHub](https://img.shields.io/github/license/chiefcake/ergoproxy?style=flat-square) ![GitHub repo size](https://img.shields.io/github/repo-size/chiefcake/ergoproxy?style=flat-square) ![GitHub top language](https://img.shields.io/github/languages/top/chiefcake/ergoproxy?style=flat-square) ![go-version](https://img.shields.io/badge/go--version-v1.18-blue?style=flat-square)

The main idea is to develop REST API service for loading astronomy picture of the day from the APOD API.


<img align="right" width="50%" src="./assets/GopherSpaceMentor.png">

## What has been implemented? 🤔

This project is implemented using REST. The service has two endpoints to retrieve pictures and one shudow worker to retrieve a picture of the day every 12 hours. You can the OpenAPI specification [here](./openapiv2/swagger.yaml).

## How to run it? 🤔

Just run the `make local-run` command. This command builds a Docker image with a Go binary inside and starts the container in the docker-compose.

## How to test it? 🤔

If you want to test it, you can find the Postman collection [here](./postman/).

## Thank you! 💟
