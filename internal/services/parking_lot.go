package services

import (
	"errors"
	"fmt"

	models "github.com/aryan297/parking-system/internal/model"

	"time"
)

type ParkingLot struct {
	Slots        []*models.Slot
	Entrances    []*models.Entrance
	CostStrategy models.CostStrategy
	ticketSeq    int
}

func NewParkingLot(slots []*models.Slot, entrances []*models.Entrance, strategy models.CostStrategy) *ParkingLot {
	return &ParkingLot{
		Slots:        slots,
		Entrances:    entrances,
		CostStrategy: strategy,
		ticketSeq:    0,
	}
}

func (p *ParkingLot) FindNearestAvailableSlot(vehicleType models.VehicleType, entrance *models.Entrance) (*models.Slot, error) {
	var nearest *models.Slot
	for _, slot := range p.Slots {
		if !slot.IsOccupied && (nearest == nil || slot.Distance < nearest.Distance) {
			nearest = slot
		}
	}
	if nearest == nil {
		return nil, errors.New("no slot available")
	}
	return nearest, nil
}

func (p *ParkingLot) GenerateTicket(vehicle *models.Vehicle, entrance *models.Entrance) (*models.Ticket, error) {
	slot, err := p.FindNearestAvailableSlot(vehicle.Type, entrance)
	if err != nil {
		return nil, err
	}
	err = slot.Park(vehicle)
	if err != nil {
		return nil, err
	}
	p.ticketSeq++
	return &models.Ticket{
		ID:        p.ticketSeq,
		Vehicle:   vehicle,
		Slot:      slot,
		EntryTime: time.Now(),
	}, nil
}

func (p *ParkingLot) Exit(ticket *models.Ticket, payment models.PaymentMethod) error {
	cost := p.CostStrategy.CalculateCost(ticket.EntryTime, time.Now(), ticket.Vehicle.Type)
	if !payment.Pay(cost) {
		return errors.New("payment failed")
	}

	ticket.Cost = cost
	ticket.Paid = true
	ticket.Payment = payment
	ticket.Slot.Unpark()

	fmt.Printf("Vehicle %s exited. Paid ₹%.2f\n", ticket.Vehicle.Number, cost)
	return nil
}
