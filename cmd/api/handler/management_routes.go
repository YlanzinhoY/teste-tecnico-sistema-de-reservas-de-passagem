package handler

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/internal/entity"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"
	"log"
	"net/http"
)

type ManagementRoutesHandler struct {
	dbHandler *db.Queries
}

func NewManagementRoutesHandler(dbHandler *db.Queries) *ManagementRoutesHandler {
	return &ManagementRoutesHandler{dbHandler: dbHandler}
}

func (s *ManagementRoutesHandler) GetManagementRoutes(echo echo.Context) error {
	var newManagementEntity *entity.ManagementRoute

	if err := echo.Bind(&newManagementEntity); err != nil {
		return err
	}

	newManagementEntity.Id = uuid.NewString()

	err := s.dbHandler.CreateManagementRoute(echo.Request().Context(), db.CreateManagementRouteParams{
		ID: newManagementEntity.Id,
		RouteName: sql.NullString{
			String: newManagementEntity.RouteName,
			Valid:  newManagementEntity.RouteName != "",
		},
		Origin: sql.NullString{
			String: newManagementEntity.Origin,
			Valid:  newManagementEntity.Origin != "",
		},
		Destination: sql.NullString{
			String: newManagementEntity.Destination,
			Valid:  newManagementEntity.Destination != "",
		},
	})

	if err != nil {
		return err
	}

	err = echo.JSON(http.StatusCreated, newManagementEntity)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
