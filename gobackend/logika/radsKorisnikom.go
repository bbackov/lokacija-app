package logika

import (
	"database/sql"
	"errors"
	"gobackend/strukture"

	"golang.org/x/crypto/bcrypt"
)

func RegistriranjeKorisnik(db *sql.DB, k strukture.Korisnik) (int, error) {

	var email string
	err := db.QueryRow("SELECT e_mail FROM korisnik WHERE e_mail = $1", k.Email).Scan(&email)
	if err != sql.ErrNoRows {
		return 0, errors.New("korisnik s tim emailom već postoji")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(k.Lozinka), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO korisnik (ime, prezime, e_mail, lozinka) VALUES ($1, $2, $3, $4) RETURNING id_korisnik"
	err = db.QueryRow(query, k.Ime, k.Prezime, k.Email, string(hashedPassword)).Scan(&k.IDKorisnik)
	if err != nil {
		return 0, err
	}

	return k.IDKorisnik, nil
}

func LogiranjeKorisnik(db *sql.DB, email string, lozinka string) (strukture.Korisnik, error) {

	var korisnik strukture.Korisnik

	err := db.QueryRow("SELECT id_korisnik, ime , prezime,lozinka FROM korisnik WHERE e_mail=$1", email).Scan(&korisnik.IDKorisnik, &korisnik.Ime, &korisnik.Prezime, &korisnik.Lozinka)

	if err != nil {
		if err == sql.ErrNoRows {
			return korisnik, errors.New("neispravan e-mail ili lozinka")
		}
		return korisnik, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(korisnik.Lozinka), []byte(lozinka))
	if err != nil {
		return korisnik, errors.New("neispravan e-mail ili lozinka")
	}

	return korisnik, nil
}

func Azuriranjeprofila(db *sql.DB, k strukture.Korisnik) error {
	var trenutnik strukture.Korisnik
	err := db.QueryRow("SELECT ime, prezime, e_mail, lozinka FROM korisnik WHERE id_korisnik = $1", k.IDKorisnik).
		Scan(&trenutnik.Ime, &trenutnik.Prezime, &trenutnik.Email, &trenutnik.Lozinka)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("korisnik nije pronađen")
		}
		return err
	}

	if k.Ime == "" {
		k.Ime = trenutnik.Ime
	}
	if k.Prezime == "" {
		k.Prezime = trenutnik.Prezime
	}
	if k.Email == "" {
		k.Email = trenutnik.Email
	}
	if k.Lozinka != trenutnik.Lozinka {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(k.Lozinka), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		k.Lozinka = string(hashedPassword)
	}

	querry := "UPDATE korisnik SET ime = $1, prezime = $2, e_mail = $3, lozinka = $4 WHERE id_korisnik = $5"
	_, err = db.Exec(querry, k.Ime, k.Prezime, k.Email, k.Lozinka, k.IDKorisnik)
	if err != nil {
		return err
	}
	return nil
}

func BrisajeProfila(db *sql.DB, id int) error {

	query := "DELETE FROM korisnik WHERE id_korisnik =$1;"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetKorisnik(db *sql.DB, idkorisnik int) (strukture.Korisnik, error) {

	var korisnik strukture.Korisnik
	query := "SELECT id_korisnik, ime, prezime, e_mail, lozinka FROM korisnik WHERE id_korisnik=$1"

	err := db.QueryRow(query, idkorisnik).Scan(
		&korisnik.IDKorisnik,
		&korisnik.Ime,
		&korisnik.Prezime,
		&korisnik.Email,
		&korisnik.Lozinka,
	)
	if err != nil {
		return korisnik, err
	}

	return korisnik, nil
}
