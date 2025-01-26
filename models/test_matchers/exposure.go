package test_matchers

import (
	"fmt"
	"github.com/sophie-rigg/havs-service/models"
)

func NewExposureMatcher(exposure *models.Exposure) *ExposureMatcher {
	return &ExposureMatcher{exposure: exposure}
}

type ExposureMatcher struct {
	exposure *models.Exposure
}

func (e *ExposureMatcher) Matches(x interface{}) bool {
	exposure, ok := x.(*models.Exposure)
	if !ok {
		return false
	}

	if e.exposure == nil && exposure == nil {
		return true
	}

	if e.exposure.Equipment != nil && exposure.Equipment != nil {
		if e.exposure.Equipment.ID != exposure.Equipment.ID {
			return false
		}

		if e.exposure.Equipment.Name != exposure.Equipment.Name {
			return false
		}

		if e.exposure.Equipment.VibrationalMagnitude != exposure.Equipment.VibrationalMagnitude {
			return false
		}
	} else if e.exposure.Equipment != exposure.Equipment {
		return false
	}

	if e.exposure.Duration != exposure.Duration {
		return false
	}

	if e.exposure.A8 != exposure.A8 {
		return false
	}

	if e.exposure.Points != exposure.Points {
		return false
	}

	if e.exposure.User != nil && exposure.User != nil {
		if e.exposure.User.ID != exposure.User.ID {
			return false
		}

		if e.exposure.User.Name != exposure.User.Name {
			return false
		}
	} else if e.exposure.User != exposure.User {
		return false
	}

	//don't compare CreatedTime

	return true
}

func (e *ExposureMatcher) String() string {
	return fmt.Sprintf("&%v", e.exposure)
}

func NewExposureArrayMatcher(exposures []*models.Exposure) *ExposureArrayMatcher {
	return &ExposureArrayMatcher{exposures: exposures}
}

type ExposureArrayMatcher struct {
	exposures []*models.Exposure
}

func (e *ExposureArrayMatcher) Matches(x interface{}) bool {
	exposures, ok := x.([]*models.Exposure)
	if !ok {
		return false
	}

	if len(e.exposures) != len(exposures) {
		return false
	}

	for i, exposure := range e.exposures {
		if !(&ExposureMatcher{exposure: exposure}).Matches(exposures[i]) {
			return false
		}
	}

	return true
}

func (e *ExposureArrayMatcher) String() string {
	return fmt.Sprintf("&%v", e.exposures)
}
