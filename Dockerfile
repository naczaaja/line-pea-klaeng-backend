FROM golang:1.21.0

WORKDIR /src/app/backend

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy