package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ylanzinhoy/sistema-de-reserva-de-passagem/internal/entity"
	db "github.com/ylanzinhoy/sistema-de-reserva-de-passagem/sql"
	"net/http"
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

func (s *ManagementTravelHandler) GetManagementTravelById(c echo.Context) error {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		return c.JSON(500, map[string]interface{}{})
	}

	res, err := s.query.GetManagementTravelById(c.Request().Context(), id)

	if err != nil {
		return c.JSON(500, map[string]interface{}{})
	}

	responseManagementTravel := &entity.ManagementTravel{
		ManagementTravelID: res.ManagementTravelID,
		ManagementRoutesID: res.ManagementRoutesID,
		TicketPrice:        res.TicketPrice,
		TotalSeats:         res.TotalSeats,
		TravelStart:        res.TravelStart,
		TravelFinish:       res.TravelFinish,
		TravelCompany:      res.TravelCompany,
	}

	return c.JSON(200, responseManagementTravel)

}

func (s *ManagementTravelHandler) GetAllManagementTravels(c echo.Context) error {

	values, err := s.query.GetAllManagementTravel(c.Request().Context())

	if err != nil {
		return c.JSON(500, map[string]interface{}{})
	}

	response := make([]entity.ManagementTravel, len(values))

	for k, v := range values {
		response[k] = entity.ManagementTravel{
			ManagementTravelID: v.ManagementTravelID,
			ManagementRoutesID: v.ManagementRoutesID,
			TicketPrice:        v.TicketPrice,
			TotalSeats:         v.TotalSeats,
			TravelStart:        v.TravelStart,
			TravelFinish:       v.TravelFinish,
			TravelCompany:      v.TravelCompany,
		}
		return c.JSON(200, response)
	}

	return nil
}

func (s *ManagementTravelHandler) UpdateManagementTravel(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	var newManagementTravelDto entity.ManagementTravel

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := c.Bind(&entity.ManagementTravel{}); err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	params := db.PutManagementTravelParams{
		ManagementTravelID: id,
		TicketPrice:        newManagementTravelDto.TicketPrice,
		TotalSeats:         newManagementTravelDto.TotalSeats,
		TravelStart:        newManagementTravelDto.TravelStart,
		TravelCompany:      newManagementTravelDto.TravelCompany,
	}

	res, err := s.query.PutManagementTravel(c.Request().Context(), params)

	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}
	responseManagementTravel := &entity.ManagementTravel{
		ManagementTravelID: res.ManagementTravelID,
		ManagementRoutesID: res.ManagementRoutesID,
		TicketPrice:        res.TicketPrice,
		TotalSeats:         res.TotalSeats,
		TravelStart:        res.TravelStart,
		TravelFinish:       res.TravelFinish,
		TravelCompany:      res.TravelCompany,
	}

	return c.JSON(200, responseManagementTravel)

}

func (s *ManagementTravelHandler) DeleteManagementTravel(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(500, map[string]interface{}{})
	}

	err = s.query.DeleteManagementTravel(c.Request().Context(), id)
	if err != nil {
		return c.JSON(500, map[string]interface{}{})
	}
	return c.NoContent(http.StatusNoContent)
}
