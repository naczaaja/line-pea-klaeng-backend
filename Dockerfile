FROM golang:1.25.1

WORKDIR /src/app/backend

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy