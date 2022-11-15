package repository

import (
	"time"

	"github.com/Macple/Bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailibilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailibilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error)
}
