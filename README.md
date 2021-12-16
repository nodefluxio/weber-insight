# weber-insight

Weber Dashboard Management and Insight

## Getting Started

## Run without Docker

### Prerequisites

- Golang
- Yarn
- Postgre with database filled from https://github.com/nodefluxio/weber

### How to run

- Clone/download this repository
- Copy .env.example as .env and edit your credentials
- Open views folder and run `yarn install`
- Open the root of this repository and run `go run .`

## Run with Docker

- Clone/download this repository
- Copy .env.example as .env and edit your credentials
- Run `docker-compose up --build`
