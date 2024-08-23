###############
# stage      #
##############
FROM golang:1.23.0-alpine3.20 AS build
LABEL authors="ylanzey"
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o build cmd/api/main.go

###############
# build       #
##############
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/build .
EXPOSE 8000
ENTRYPOINT ["/app/build"]