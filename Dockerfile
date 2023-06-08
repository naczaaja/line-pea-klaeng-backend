FROM golang:1.20

WORKDIR /src/app/backend

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy