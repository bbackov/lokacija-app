package main

import (
	"encoding/json"
	"fmt"
	connection "gobackend/bazapodataka"
	"gobackend/logika"
	"gobackend/strukture"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func main() {
	fmt.Print("Verzija ova")
	http.HandleFunc("/registracija", RegistracijaHandler)
	http.HandleFunc("/prijava", PrijavaHandler)
	http.HandleFunc("/dodaj_uređaj", DodavanjeUređajaHandler)
	http.HandleFunc("/dodaj_lokaciju", DodajLokacijuHandler)
	http.HandleFunc("/dohvati_10lokacija", Dohvati_10LokacijaHandler)
	http.HandleFunc("/dohvati_zadnjulokaciju", Dohvati_ZadnjuLokacijuHandler)
	http.HandleFunc("/dohvati_vremenskilokacije", Dohvati_VremenskeLokacijeHandler)
	http.HandleFunc("/azuriraj_profil", AzuriranjeprofilaHandler)
	http.HandleFunc("/obrisi_profil", BrisajeProfilaHandler)
	http.HandleFunc("/dohvati_korisnika", GetKorisnikHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/postavi_status", PostaviStatusHandler)
	http.HandleFunc("/obrisi_uredaj", ObrisiUredajHandler)
	http.HandleFunc("/dohvati_status", GetStatusHandler)
	http.HandleFunc("/dohvati_uredaje", GetUređajiHandler)

	go func() {
		ticker := time.NewTicker(24 * time.Hour)

		for range ticker.C {
			db, err := connection.Connect()
			if err != nil {
				fmt.Println("Greška pri konekciji na bazu (automatsko brisanje):", err)
				continue
			}

			err = logika.BrisanjeLokacije(db)
			if err != nil {
				fmt.Println("Greška pri automatskom brisanju lokacija:", err)
			} else {
				fmt.Println("Automatski obrisane stare lokacije")
			}

			db.Close()

		}

	}()

	go func() {
		ticker := time.NewTicker(15 * time.Minute)

		for range ticker.C {
			db, err := connection.Connect()
			if err != nil {
				fmt.Println("Greška pri konekciji na bazu (automatsko brisanje):", err)
				continue
			}
			err = logika.Offlineuredaj(db)
			if err != nil {
				fmt.Println("Greška pri promjeni statusa u offline", err)
			} else {
				fmt.Println("Automatska promjena aktivnosti")
			}

			db.Close()
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Nepoznata ruta:", r.URL.Path)
		fmt.Println(" Nešto je pozvano, ali nije pronađena ruta:", r.URL.Path)
		http.NotFound(w, r)
	})

	fmt.Println("Server pokrenut na 0.0.0.0:8085")
	http.ListenAndServe("0.0.0.0:8085", nil)
}

func jsonError(w http.ResponseWriter, status int, poruka string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"poruka": poruka,
	})
}

func RegistracijaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	var korisnik strukture.Korisnik

	err := json.NewDecoder(r.Body).Decode(&korisnik)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan JSON format")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	id, err := logika.RegistriranjeKorisnik(db, korisnik)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Registracija neuspješna: %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":     "Registracija uspješna",
		"idKorisnik": id,
	})
}

func PrijavaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}
	type podacizaprijau struct {
		Email   string `json:"email"`
		Lozinka string `json:"lozinka"`
	}
	var podaci podacizaprijau

	err := json.NewDecoder(r.Body).Decode(&podaci)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan JSON format")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	korisnik, err := logika.LogiranjeKorisnik(db, podaci.Email, podaci.Lozinka)
	if err != nil {

		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Prijava neuspješna: %s", err.Error()))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_korisnik": korisnik.IDKorisnik,
		"ime":         korisnik.Ime,
		"prezime":     korisnik.Prezime,
		"email":       korisnik.Email,
		"vrijeme":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString(logika.TajniKljuc)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška kod generiranja novog tokena")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":   "uspješno logiranje",
		"token":    tokenString,
		"korisnik": korisnik,
	})

}

func DodavanjeUređajaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	var podaci strukture.Uređaj

	err := json.NewDecoder(r.Body).Decode(&podaci)

	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan JSON format")
		return
	}

	idkorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}
	podaci.IDKorisnik = idkorisnik

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	podaci.IDUredaj, err = logika.DodajUredaj(db, podaci)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dodavanje uređaja %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka": "uspješno dodavanje uređaja",
		"uredaj": podaci,
	})
}

func DodajLokacijuHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handler aktiviran")

	if r.Method != http.MethodPost {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	var podaci strukture.Lokacija

	data, err := io.ReadAll(r.Body)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Greška kod čitanja tijela zahtjeva")
		return
	}
	fmt.Println("Raw JSON body:", string(data))

	err = json.Unmarshal(data, &podaci)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan JSON format")
		return
	}

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Provjeravlasnistva(db, podaci.IDUredaj, idKorisnik)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dodavanje lokacije %s", err.Error()))
		return
	}

	var idlokacija uuid.UUID
	var vrijeme time.Time

	idlokacija, vrijeme, err = logika.DodajLokaciju(db, podaci)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dodavanje lokacije %s", err.Error()))
		return
	}

	err = logika.Azuriranjeaktivnosti(db, podaci.IDUredaj)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dodavanje lokacije %s", err.Error()))
		return
	}

	podaci.IDLokacija = idlokacija
	podaci.Vrijeme = vrijeme

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":   "uspješno dodavanje lokacije",
		"lokacija": podaci,
	})

}

func Dohvati_10LokacijaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	var lokacije []strukture.Lokacija
	var id_uredaja int

	idParam := r.URL.Query().Get("id")
	id_uredaja, err := strconv.Atoi(idParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan id uredaja")
		return
	}

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}

	db, err := connection.Connect()

	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Provjeravlasnistva(db, id_uredaja, idKorisnik)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje lokacije %s", err.Error()))
		return
	}

	lokacije, err = logika.GetLokacija(db, id_uredaja)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje lokacije %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":   "uspješno dohvaćanje lokacije",
		"lokacije": lokacije,
	})

}

func Dohvati_ZadnjuLokacijuHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	var lokacije strukture.Lokacija
	var id_uredaja int

	idParam := r.URL.Query().Get("id")
	id_uredaja, err := strconv.Atoi(idParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan id uredaja")
		return
	}

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}

	db, err := connection.Connect()

	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Provjeravlasnistva(db, id_uredaja, idKorisnik)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje lokacije %s", err.Error()))
		return
	}

	lokacije, err = logika.GetZadnjaLokacija(db, id_uredaja)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje lokacije %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":   "uspješno dohvaćanje lokacije",
		"lokacija": lokacije,
	})

}

func Dohvati_VremenskeLokacijeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	var lokacije []strukture.Lokacija
	var id_uredaja int

	idParam := r.URL.Query().Get("id")
	pocetakParam := r.URL.Query().Get("pocetak")
	krajParam := r.URL.Query().Get("kraj")

	id_uredaja, err := strconv.Atoi(idParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan id uredaja")
		return
	}

	pocetak, err := time.Parse(time.RFC3339, pocetakParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan format za pocetak")
		return
	}

	kraj, err := time.Parse(time.RFC3339, krajParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan format za kraj")
		return
	}

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}

	db, err := connection.Connect()

	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Provjeravlasnistva(db, id_uredaja, idKorisnik)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje lokacije %s", err.Error()))
		return
	}

	lokacije, err = logika.GetVremenskiLokacija(db, id_uredaja, pocetak, kraj)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje lokacije %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":   "uspješno dohvaćanje lokacije",
		"lokacije": lokacije,
	})

}

func AzuriranjeprofilaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPatch {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	korisnikid, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}

	var korisnik strukture.Korisnik

	korisnik.Ime = r.URL.Query().Get("ime")
	korisnik.Prezime = r.URL.Query().Get("prezime")
	korisnik.Email = r.URL.Query().Get("email")
	korisnik.Lozinka = r.URL.Query().Get("lozinka")
	korisnik.IDKorisnik = korisnikid

	db, err := connection.Connect()

	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Azuriranjeprofila(db, korisnik)
	if err != nil {
		http.Error(w, fmt.Sprintf("Neuspješno Ažuriranje %s", err.Error()), http.StatusBadRequest)
		return
	}
	korisnik, err = logika.GetKorisnik(db, korisnikid)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Neuspješno dohvaćanje ažuriranog profila")
		return
	}

	authHeader := r.Header.Get("Authorization")
	stariToken := strings.TrimPrefix(authHeader, "Bearer ")
	logika.PonistiToken(stariToken)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_korisnik": korisnik.IDKorisnik,
		"ime":         korisnik.Ime,
		"prezime":     korisnik.Prezime,
		"email":       korisnik.Email,
		"vrijeme":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString(logika.TajniKljuc)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška kod generiranja novog tokena")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":     "uspješno ažuriranje",
		"korisnik":   korisnik,
		"novi_token": tokenString,
	})
}

func BrisajeProfilaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	korisnikid, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.BrisajeProfila(db, korisnikid)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno brisanje %s", err.Error()))
		return
	}

	authHeader := r.Header.Get("Authorization")
	stariToken := strings.TrimPrefix(authHeader, "Bearer ")
	logika.PonistiToken(stariToken)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka": "uspješno brisanje",
	})

}

func GetKorisnikHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
	}

	korisnikid, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Neispravan token")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	korisnik, err := logika.GetKorisnik(db, korisnikid)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje: %s", err.Error()))
		return
	}

	type PublicKorisnik struct {
		IDKorisnik int    `json:"id_korisnik"`
		Ime        string `json:"ime"`
		Prezime    string `json:"prezime"`
		Email      string `json:"email"`
	}

	pubkorisnik := PublicKorisnik{
		IDKorisnik: korisnik.IDKorisnik,
		Ime:        korisnik.Ime,
		Prezime:    korisnik.Prezime,
		Email:      korisnik.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":   "uspješno dohvaćanje",
		"korisnik": pubkorisnik,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	_, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Nevažeći token")
		return
	}

	authHeader := r.Header.Get("Authorization")
	stariToken := strings.TrimPrefix(authHeader, "Bearer ")
	logika.PonistiToken(stariToken)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka": "uspješni logout",
	})

}

func PostaviStatusHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPatch {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	idParam := r.URL.Query().Get("id")
	id_uredaja, err := strconv.Atoi(idParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan id uredaja")
		return
	}
	status := r.URL.Query().Get("id")

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Nevažeći token")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Provjeravlasnistva(db, id_uredaja, idKorisnik)

	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješna promjena statusa %s", err.Error()))
		return
	}

	err = logika.PostaviStatus(db, id_uredaja, status)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješna promjena statusa %s", err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka": "uspješna promjena statusa",
	})

}

func ObrisiUredajHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	idParam := r.URL.Query().Get("id")
	id_uredaja, err := strconv.Atoi(idParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan id uredaja")
		return
	}

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Nevažeći token")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Provjeravlasnistva(db, id_uredaja, idKorisnik)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno brisanje uređaja %s", err.Error()))
		return
	}

	err = logika.ObrisiUredaj(db, id_uredaja)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno brisanje uređaja %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka": "uspješno brisanje uređaja",
	})

}

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	idParam := r.URL.Query().Get("id")
	id_uredaja, err := strconv.Atoi(idParam)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Neispravan id uredaja")
		return
	}

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Nevažeći token")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	err = logika.Provjeravlasnistva(db, id_uredaja, idKorisnik)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje statusa %s", err.Error()))
		return
	}

	status, err := logika.GetStatus(db, id_uredaja)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje statusa %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka": "uspješno brisanje uređaja",
		"status": status,
	})

}

func GetUređajiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		jsonError(w, http.StatusMethodNotAllowed, "Metoda nije dozvoljena")
		return
	}

	idKorisnik, err := logika.DohvatiIDKorisnika(r)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "Nevažeći token")
		return
	}

	db, err := connection.Connect()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Greška pri konekciji na bazu")
		return
	}
	defer db.Close()

	var uredaji []strukture.Uređaj

	uredaji, err = logika.GetUređaji(db, idKorisnik)
	if err != nil {
		jsonError(w, http.StatusBadRequest, fmt.Sprintf("Neuspješno dohvaćanje uređaja %s", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"poruka":  "uspješno dohvaćanje uređaja",
		"uredaji": uredaji,
	})

}
