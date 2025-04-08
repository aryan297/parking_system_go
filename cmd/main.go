package main

import (
	"fmt"
	"time"

	model "github.com/aryan297/parking-system/internal/model"
	"github.com/aryan297/parking-system/internal/payments"
	"github.com/aryan297/parking-system/internal/services"
	"github.com/aryan297/parking-system/internal/strategies"
)

func main() {
	entrances := []*model.Entrance{
		{ID: 1, Name: "Main Gate"},
	}

	slots := []*model.Slot{
		{ID: 1, Distance: 5},
		{ID: 2, Distance: 10},
	}

	strategy := &strategies.HourlyCostStrategy{
		Rates: map[model.VehicleType]float64{
			model.TwoWheeler:  10,
			model.FourWheeler: 20,
		},
	}

	lot := services.NewParkingLot(slots, entrances, strategy)

	vehicle := &model.Vehicle{Number: "MH12AB1234", Type: model.FourWheeler}
	ticket, err := lot.GenerateTicket(vehicle, entrances[0])
	fmt.Println("Ticket generated:", ticket.ID, "for vehicle:", vehicle.Number)
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second) // Simulate parking time

	payment := &payments.CardPayment{CardNumber: "1234567890123456"}
	if err := lot.Exit(ticket, payment); err != nil {
		panic(err)
	}
}
