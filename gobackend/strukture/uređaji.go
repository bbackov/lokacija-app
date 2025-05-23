package strukture

import (
	"time"
)

type Uređaj struct {
	IDUredaj   int       `json:"id_uredaj"`
	ImeUredaj  string    `json:"ime_uredaj"`
	TipUredaj  string    `json:"tip_uredaj"`
	Status     string    `json:"status_uredaj"`
	IDKorisnik int       `json:"id_korisnik"`
	Aktivnost  time.Time `json:"-"`
}
