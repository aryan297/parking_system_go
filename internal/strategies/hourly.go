package strategies

import (
	"time"

	models "github.com/aryan297/parking-system/internal/model"
)

type HourlyCostStrategy struct {
	Rates map[models.VehicleType]float64
}

func (h *HourlyCostStrategy) CalculateCost(entry, exit time.Time, vType models.VehicleType) float64 {
	duration := exit.Sub(entry).Hours()
	return h.Rates[vType] * duration
}
