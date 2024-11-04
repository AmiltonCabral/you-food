# syntax=docker/dockerfile:1
# youFood backend

FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /youfood-api

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /youfood-api /youfood-api

EXPOSE 8000

USER nonroot:nonroot

CMD ["/youfood-api"]
