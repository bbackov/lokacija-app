package strukture

type Korisnik struct {
	IDKorisnik int    `json:"id_korisnik"`
	Ime        string `json:"ime"`
	Prezime    string `json:"prezime"`
	Email      string `json:"email"`
	Lozinka    string `json:"lozinka"`
}
