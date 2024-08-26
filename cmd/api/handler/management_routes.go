package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/cmd/api/errors"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/internal/entity"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"
)

type ManagementRoutesHandler struct {
	dbHandler *db.Queries
}

func NewManagementRoutesHandler(dbHandler *db.Queries) *ManagementRoutesHandler {
	return &ManagementRoutesHandler{dbHandler: dbHandler}
}

func (s *ManagementRoutesHandler) GetManagementRoutesAll(echo echo.Context) error {

	values, err := s.dbHandler.GetManagementRouteAll(echo.Request().Context())

	if err != nil {
		return echo.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}

	response := make([]entity.ManagementRoute, len(values))

	for i, value := range values {
		response[i] = entity.ManagementRoute{
			Id:          value.ID,
			RouteName:   value.RouteName.String,
			Origin:      value.Origin.String,
			Destination: value.Destination.String,
		}
	}

	return echo.JSON(200, response)

}

func (s *ManagementRoutesHandler) GetManagementRoutesById(echo echo.Context) error {

	params := echo.Param("id")

	route, err := s.dbHandler.GetRouteById(echo.Request().Context(), params)
	if err != nil {
		return errors.IdNotFound(err, echo)
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

	var newManagementEntity *entity.ManagementRoute

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
	value, err := s.dbHandler.UpdateManagementRoute(echo.Request().Context(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.JSON(404, map[string]string{
				"message": err.Error(),
			})
		}
		return echo.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := &entity.ManagementRoute{
		Id:          id,
		RouteName:   value.RouteName.String,
		Origin:      value.Origin.String,
		Destination: value.Destination.String,
	}

	return echo.JSON(http.StatusOK, response)
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
