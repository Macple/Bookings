package dbrepo

import (
	"errors"
	"time"

	"github.com/Macple/Bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into a database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// the room id is 2, than fail; otherwise pass
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into a database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some err")
	}
	return nil
}

// SearchAvailibilityByDatesByRoomID return true if availability exists for roomID and false otherwise
func (m *testDBRepo) SearchAvailibilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailibilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailibilityForAllRooms(start time.Time, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("invalid room ID")
	}

	return room, nil
}

// GetUserByID returns a user by ID
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}

// UpdateUser updates a user in the database
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

// Authenticate authenticates a user
func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	var id int
	var hashedPassword string
	return id, hashedPassword, nil
}
