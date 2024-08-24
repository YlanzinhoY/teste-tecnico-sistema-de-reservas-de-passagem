package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=sistema_de_passagem sslmode=disable"
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
	e.Logger.Fatal(e.Start(":8000"))
}
