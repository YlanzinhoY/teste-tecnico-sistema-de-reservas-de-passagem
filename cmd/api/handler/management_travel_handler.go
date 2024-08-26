package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/internal/entity"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"
)

type ManagementTravelHandler struct {
	query *db.Queries
}

func NewManagementTravelHandler(query *db.Queries) *ManagementTravelHandler {
	return &ManagementTravelHandler{query: query}
}

func (s *ManagementTravelHandler) CreateManagementTravel(c echo.Context) error {
	var newManagementTravel *entity.ManagementTravel

	if err := c.Bind(&newManagementTravel); err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	newManagementTravel.ManagementTravelID = uuid.New()

	err := s.query.CreateManagementTravel(c.Request().Context(), db.CreateManagementTravelParams{
		ManagementTravelID: newManagementTravel.ManagementTravelID,
		ManagementRoutesID: newManagementTravel.ManagementRoutesID,
		TicketPrice:        newManagementTravel.TicketPrice,
		TotalSeats:         newManagementTravel.TotalSeats,
		TravelStart:        newManagementTravel.TravelStart,
		TravelFinish:       newManagementTravel.TravelFinish,
		TravelCompany:      newManagementTravel.TravelCompany,
	})

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"sucess":   true,
		"response": newManagementTravel,
	})
}

func (s *ManagementTravelHandler) GetManagementTravelById(echo echo.Context) error {
	return nil
}
