package models

// ExposureSummary represents a summary of an exposure.
type ExposureSummary struct {
	A8     float64 `json:"a8"`
	Points float64 `json:"points"`
	User   *User   `json:"user"`
}

// NewExposureSummary creates a new exposure summary.
func NewExposureSummary(exposures []*Exposure) *ExposureSummary {
	if len(exposures) == 0 {
		return nil
	}

	var summary ExposureSummary
	for i, exposure := range exposures {
		if i == 0 {
			summary.User = exposure.User
		}

		summary.AddExposure(exposure)
	}
	return &summary
}

func (e *ExposureSummary) AddExposure(exposure *Exposure) {
	// If an exposure is not valid, the A8 and points will be 0
	// This is to prevent invalid exposures from affecting the summary
	e.A8 += exposure.A8
	e.Points += exposure.Points
}
