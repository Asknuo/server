FROM golang:1.24.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app/
CMD ["./main"]