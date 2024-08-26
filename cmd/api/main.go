package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/cmd/api/handler"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	connStr := "host=postgres port=5432 user=postgres password=postgres dbname=sistema_de_passagem sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(dbConn)

	query := db.New(dbConn)

	Routes(e, query)

	hc := handler.NewHealthCheckerHandler(e)
	e.GET("/v1/healthcheck", hc.HealthChecker)
	e.Logger.Fatal(e.Start(":8000"))
}
