# youFood

## How to connect with the data base

You can define an `URI` or all the separate variables. If you define the `URI`,
the API will use it, if not, you need to define all the other environment
variables.

To connect to the database, you need to define:

- URI

or

- DB_HOST
- DB_PORT
- DB_USER
- DB_PASSWORD
- DB_NAME

To run with docker-compose, one alternative is create an .env file with the
database environment variables, and up the docker-compose passing this file.
`docker-compose --env-file .env up --build`

To run with Kubernetes, you need to define it on `kubernetes/backend-api.yaml`.

## How to run the backend API with go run

```sh
cd api
export $(grep -v '^#' ../.env | xargs) && go run .
```

## Monitoring Kubernetes with Prometheus

[link](https://hackmd.io/@OrALpDtKSY6QetSsHgj-IA/BJXuuNmYP)
