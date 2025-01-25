package utils

import (
	"github.com/google/uuid"
	"math"
)

// PartialExposureA8 calculates the exposure A8 of a user based on the vibration magnitude and trigger time.
func PartialExposureA8(vibrationMagnitude float64, triggerTime int) float64 {
	return vibrationMagnitude * math.Sqrt((float64(triggerTime)/60)/8)
}

// PartialExposurePoints calculates the exposure points of a user based on the vibration magnitude and trigger time.
func PartialExposurePoints(vibrationMagnitude float64, triggerTime int) float64 {
	points := math.Pow(vibrationMagnitude/2.5, 2) * (((float64(triggerTime) / 60) / 8) * 100)
	return math.Round(points)
}

func GenerateUUID() (string, error) {
	val, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return val.String(), nil
}
