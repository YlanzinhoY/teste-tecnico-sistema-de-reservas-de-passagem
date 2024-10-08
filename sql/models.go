// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ManagementRoute struct {
	ID          string
	RouteName   sql.NullString
	Origin      sql.NullString
	Destination sql.NullString
}

type ManagementTravel struct {
	ManagementTravelID uuid.UUID
	ManagementRoutesID string
	TicketPrice        float64
	TotalSeats         int32
	TravelStart        time.Time
	TravelFinish       time.Time
	TravelCompany      string
}

type ReservationSystem struct {
	ReservationID      uuid.UUID
	ManagementTravelID uuid.NullUUID
	PassengerName      string
	SeatNumber         int32
}
