# syntax=docker/dockerfile:1
# youFood backend

FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-relivery

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /api-relivery /api-relivery

EXPOSE 8001

USER nonroot:nonroot

CMD ["/api-relivery"]
