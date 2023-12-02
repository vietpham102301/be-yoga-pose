# Stage 1: MySQL setup
FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=viet1234
ENV MYSQL_DATABASE=yoga_support
ENV MYSQL_USER=vietpham102301
ENV MYSQL_PASSWORD=viet1234

WORKDIR /app

COPY ./init.sql /docker-entrypoint-initdb.d/init.sql

# Stage 2: Build Go application
FROM golang:latest AS build
LABEL authors="vietpham1023"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

# Stage 3: Build the final image with MySQL and the Go application binary
FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=viet1234
ENV MYSQL_DATABASE=yoga_support
ENV MYSQL_USER=vietpham102301
ENV MYSQL_PASSWORD=viet1234

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]
