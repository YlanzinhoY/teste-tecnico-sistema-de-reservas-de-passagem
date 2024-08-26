###############
# stage      #
##############
FROM golang:1.23.0-alpine3.20 AS build
LABEL authors="ylanzey"
RUN apk add --no-cache git gcc musl-dev make
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o build ./cmd/api




###############
# build       #
##############
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/build .
COPY --from=build /go/bin/migrate /usr/local/bin/migrate
COPY ./sql/migrations /app/sql/migrations
EXPOSE 8000
ENTRYPOINT ["sh", "-c", "migrate -path=/app/sql/migrations -database 'postgres://postgres:postgres@postgres:5432/sistema_de_passagem?sslmode=disable' up && /app/build"]