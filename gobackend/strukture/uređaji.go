package strukture

import (
	"time"
)

type Uređaj struct {
	IDUredaj   int       `json:"id_uredaja"`
	ImeUredaj  string    `json:"ime_uredaja"`
	TipUredaj  string    `json:"tip_uredaja"`
	Status     string    `json:"status_uredaja"`
	IDKorisnik int       `json:"id_korisnik"`
	Aktivnost  time.Time `json:"vrijeme"`
}
