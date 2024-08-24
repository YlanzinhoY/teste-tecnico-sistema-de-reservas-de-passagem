package main

import (
	"database/sql"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/cmd/api/handler"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/internal/entity"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"

	_ "github.com/lib/pq"
)

func main() {

	e := echo.New()

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

	managementHandler := handler.NewManagementRoutesHandler(query)

	e.GET("/v1/management_route/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!")
	})

	e.POST("/v1/management_route/", managementHandler.GetManagementRoutes)

	e.PUT("/v1/management_route/:id", func(c echo.Context) error {
		id := c.Param("id")

		var newManagementEntity entity.ManagementRoute

		if err := c.Bind(&newManagementEntity); err != nil {
			return err
		}

		params := db.UpdateManagementRouteParams{
			ID: id,
			RouteName: sql.NullString{
				String: newManagementEntity.RouteName,
				Valid:  newManagementEntity.RouteName != "",
			},
			Destination: sql.NullString{
				String: newManagementEntity.Destination,
				Valid:  newManagementEntity.Destination != "",
			},
			Origin: sql.NullString{
				String: newManagementEntity.Origin,
				Valid:  newManagementEntity.Origin != "",
			},
		}
		err := query.UpdateManagementRoute(c.Request().Context(), params)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update management route",
			})
		}
		return nil

	})

	e.Logger.Fatal(e.Start(":8000"))
}
