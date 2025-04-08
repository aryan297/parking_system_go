package models

import "errors"

type Slot struct {
	ID         int
	Vehicle    *Vehicle
	IsOccupied bool
	Distance   int
}

func (s *Slot) Park(vehicle *Vehicle) error {
	if s.IsOccupied {
		return errors.New("slot already occupied")
	}
	s.Vehicle = vehicle
	s.IsOccupied = true
	return nil
}

func (s *Slot) Unpark() {
	s.Vehicle = nil
	s.IsOccupied = false
}
