package logika

import (
	"database/sql"
	"gobackend/strukture"
)

func DodajUredaj(db *sql.DB, u strukture.Uređaj) (int, error) {

	query := "INSERT INTO uređaj (ime_uredaj, tip_uredaj, status_uredaj, id_korisnik,posljednja_aktivnost) VALUES ($1, $2, $3, $4,NOW()) RETURNING id_uredaj"

	err := db.QueryRow(query, u.ImeUredaj, u.TipUredaj, u.Status, u.IDKorisnik).Scan(&u.IDUredaj)
	if err != nil {
		return 0, err
	}

	return u.IDUredaj, nil
}

func PostaviStatus(db *sql.DB, idUređaj int, status string) error {

	query := "UPDATE uređaj SET status_uredaj=$1 WHERE id_uredaj=$2"
	_, err := db.Exec(query, idUređaj, status)
	if err != nil {
		return err
	}
	return nil
}

func Azuriranjeaktivnosti(db *sql.DB, idUređaj int) error {
	query := "UPDATE uređaj SET posljednja_aktivnost=Now(),status_uredaj='Aktivno' WHERE id_uredaj=$1"
	_, err := db.Exec(query, idUređaj)
	if err != nil {
		return err
	}
	return nil
}

func Offlineuredaj(db *sql.DB) error {
	query := `UPDATE uređaj SET status_uredaj='offline' 
	WHERE posljednja_aktivnost< Now()- INTERVAL '5 minutes' and status_uredaj !='offline'`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func ObrisiUredaj(db *sql.DB, idUredaj int) error {
	query := "DELETE FROM uređaj WHERE id_uredaj = $1"
	_, err := db.Exec(query, idUredaj)
	if err != nil {
		return err
	}
	return nil
}

func GetStatus(db *sql.DB, idUredaj int) (string, error) {
	query := "SELECT status_uredaj FROM uređaj WHERE id_uredaj = $1"

	var status string
	err := db.QueryRow(query, idUredaj).Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}

func GetUređaji(db *sql.DB, idKorisnik int) ([]strukture.Uređaj, error) {
	rows, err := db.Query("SELECT id_uredaj, ime_uredaj, tip_uredaj, status_uredaj, id_korisnik FROM uređaj WHERE id_korisnik=$1", idKorisnik)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var uređaji []strukture.Uređaj

	for rows.Next() {
		var u strukture.Uređaj
		err := rows.Scan(&u.IDUredaj, &u.ImeUredaj, &u.TipUredaj, &u.Status, &u.IDKorisnik)
		if err != nil {
			return nil, err
		}
		uređaji = append(uređaji, u)
	}
	return uređaji, nil
}
