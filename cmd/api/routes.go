package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/cmd/api/handler"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"
)

func Routes(e *echo.Echo, queries *db.Queries) {

	managementHandler := handler.NewManagementRoutesHandler(queries)
	e.GET("/v1/management_route", managementHandler.GetManagementRoutesAll)
	e.GET("/v1/management_route/:id", managementHandler.GetManagementRoutesById)
	e.POST("/v1/management_route", managementHandler.PostManagementRoutes)
	e.PUT("/v1/management_route/:id", managementHandler.PutManagementRoutes)
	e.DELETE("/v1/management_route/:id", managementHandler.DeleteRoutes)

	managementTraverHandler := handler.NewManagementTravelHandler(queries)
	e.POST("/v1/management_travel", managementTraverHandler.CreateManagementTravel)
	e.GET("/v1/management_travel/:id", managementTraverHandler.GetManagementTravelById)
	e.GET("/v1/management_travel", managementTraverHandler.GetAllManagementTravels)
	e.PUT("/v1/management_travel/:id", managementTraverHandler.UpdateManagementTravel)
	e.DELETE("/v1/management_travel/:id", managementTraverHandler.DeleteManagementTravel)

}
