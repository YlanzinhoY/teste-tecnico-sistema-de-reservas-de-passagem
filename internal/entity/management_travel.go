package entity

import (
	"github.com/google/uuid"
	"time"
)

type ManagementTravel struct {
	ManagementTravelID uuid.UUID `json:"management_travel_id"`
	ManagementRoutesID string    `json:"management_routes_id"`
	TicketPrice        float64   `json:"ticket_price"`
	TotalSeats         int32     `json:"total_seats"`
	TravelStart        time.Time `json:"travel_start"`
	TravelFinish       time.Time `json:"travel_finish"`
	TravelCompany      string    `json:"travel_company"`
}
