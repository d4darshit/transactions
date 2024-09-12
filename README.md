# Transactions API

This project is a simple transactions API written in Go. It allows you to manage accounts and transactions using HTTP endpoints. It includes features for creating accounts, retrieving account information, and creating transactions.

## Features

- **Docker Support:** Easily containerize and run the application with Docker.
- **Unit Tests:** Comprehensive tests to ensure the reliability of the application.

## Getting Started

### Prerequisites

- **Go**: Install Go (1.18 or later) from [golang.org](https://golang.org/dl/).
- **Docker**: Install Docker from [docker.com](https://www.docker.com/get-started).

### Running Test Cases

Command - go test ./..

## Running the Application

### Using Docker-Compose
---
This spins up application as well as MySQL DB. Change values in .env as suitable

command -

1. docker compose up --build

### Using Docker
1. Changes values in .env file for Database connection
2. Build the Docker image:
   docker build -t transactions-api .
3. Running Application
   docker run -p 8080:8080 transactions-api


#### Running without docker

1. go mod tidy 
2. go run ./cmd/main.go