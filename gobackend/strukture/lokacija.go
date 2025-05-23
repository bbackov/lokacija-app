package strukture

import (
	"time"

	"github.com/google/uuid"
)

type Lokacija struct {
	IDLokacija       uuid.UUID `json:"id_lokacije"`
	GeografskaSirina float64   `json:"geografska_sirina"`
	GeografskaDuzina float64   `json:"geografska_duzina"`
	Vrijeme          time.Time `json:"vrijeme"`
	Pravac           float64   `json:"pravac"`
	Preciznost       float64   `json:"preciznost"`
	Visina           float64   `json:"visina"`
	Brzina           float64   `json:"brzina"`
	IDUredaj         int       `json:"id_uredaj"`
}
