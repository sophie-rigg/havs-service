package models

type EquipmentItem struct {
	ID                   string  `json:"id"` // uuid
	Name                 string  `json:"name"`
	VibrationalMagnitude float64 `json:"vibrational_magnitude"` // ms/2
}

func NewEquipmentItem(id string) *EquipmentItem {
	return &EquipmentItem{
		ID: id,
	}
}

func (e *EquipmentItem) SetName(name string) {
	e.Name = name
}

func (e *EquipmentItem) SetVibrationalMagnitude(vibrationalMagnitude float64) {
	e.VibrationalMagnitude = vibrationalMagnitude
}
