package handler

import (
	"database/sql"
	"errors"
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

func (s *ManagementRoutesHandler) GetManagementRoutesById(echo echo.Context) error {

	params := echo.Param("id")

	route, err := s.dbHandler.GetRouteById(echo.Request().Context(), params)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return echo.JSON(http.StatusNotFound, map[string]string{
				"message": "Id not found",
			})
		}

		return echo.JSON(http.StatusInternalServerError, map[string]string{})
	}

	dto := &entity.ManagementRoute{
		Id:          route.ID,
		RouteName:   route.RouteName.String,
		Origin:      route.Origin.String,
		Destination: route.Destination.String,
	}

	return echo.JSON(http.StatusOK, dto)
}

func (s *ManagementRoutesHandler) PostManagementRoutes(echo echo.Context) error {
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

func (s *ManagementRoutesHandler) PutManagementRoutes(echo echo.Context) error {
	id := echo.Param("id")

	var newManagementEntity entity.ManagementRoute

	if err := echo.Bind(&newManagementEntity); err != nil {
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
	err := s.dbHandler.UpdateManagementRoute(echo.Request().Context(), params)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update management route",
		})
	}
	return nil
}

func (s *ManagementRoutesHandler) DeleteRoutes(echo echo.Context) error {
	id := echo.Param("id")

	err := s.dbHandler.DeleteManagementRoute(echo.Request().Context(), id)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return echo.NoContent(http.StatusNoContent)
}
