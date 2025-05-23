package logika

import (
	"database/sql"
	"fmt"
	"gobackend/strukture"
	"time"

	"github.com/google/uuid"
)

func DodajLokaciju(db *sql.DB, l strukture.Lokacija) (uuid.UUID, time.Time, error) {
	fmt.Println("usao u funkciju")
	query := `INSERT INTO lokacija (geografska_sirina, geografska_duzina, preciznost, brzina, pravac, visina, id_uredaj)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id_lokacije, vrijeme`
	err := db.QueryRow(query, l.GeografskaSirina, l.GeografskaDuzina, l.Preciznost, l.Brzina, l.Pravac, l.Visina, l.IDUredaj).Scan(&l.IDLokacija, &l.Vrijeme)
	if err != nil {
		return uuid.Nil, time.Now(), err
	}

	return l.IDLokacija, l.Vrijeme, nil
}

func GetLokacija(db *sql.DB, idUređaj int) ([]strukture.Lokacija, error) {
	rows, err := db.Query(`SELECT id_lokacije,geografska_sirina,geografska_duzina,vrijeme,preciznost,brzina,visina,pravac FROM lokacija WHERE id_uredaj=$1
		ORDER BY vrijeme DESC LIMIT 10`, idUređaj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lokacije []strukture.Lokacija

	for rows.Next() {
		var l strukture.Lokacija
		err := rows.Scan(&l.IDLokacija, &l.GeografskaSirina, &l.GeografskaDuzina, &l.Vrijeme, &l.Preciznost, &l.Brzina, &l.Visina, &l.Pravac)
		if err != nil {
			return nil, err
		}
		lokacije = append(lokacije, l)
	}
	return lokacije, nil
}

func GetZadnjaLokacija(db *sql.DB, idUredaj int) (strukture.Lokacija, error) {
	query := `
		SELECT id_lokacije, geografska_sirina, geografska_duzina, vrijeme, preciznost, brzina, visina, pravac
		FROM lokacija
		WHERE id_uredaj = $1
		ORDER BY vrijeme DESC
		LIMIT 1
	`
	fmt.Println(query)
	var l strukture.Lokacija
	err := db.QueryRow(query, idUredaj).Scan(
		&l.IDLokacija,
		&l.GeografskaSirina,
		&l.GeografskaDuzina,
		&l.Vrijeme,
		&l.Preciznost,
		&l.Brzina,
		&l.Visina,
		&l.Pravac,
	)
	if err != nil {
		return l, err
	}

	return l, nil
}

func GetVremenskiLokacija(db *sql.DB, idUređaj int, pocetak time.Time, kraj time.Time) ([]strukture.Lokacija, error) {
	query := `SELECT id_lokacije,geografska_sirina,geografska_duzina,vrijeme,preciznost,brzina,visina,pravac FROM lokacija where id_uredaj=$1 and 
		vrijeme >= $2 and vrijeme<= $3
		ORDER BY vrijeme`
	rows, err := db.Query(query, idUređaj, pocetak, kraj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lokacije []strukture.Lokacija

	for rows.Next() {
		var l strukture.Lokacija
		err := rows.Scan(&l.IDLokacija, &l.GeografskaSirina, &l.GeografskaDuzina, &l.Vrijeme, &l.Preciznost, &l.Brzina, &l.Visina, &l.Pravac)
		if err != nil {
			return nil, err
		}
		lokacije = append(lokacije, l)
	}
	return lokacije, nil
}

func BrisanjeLokacije(db *sql.DB) error {
	query := "DELETE FROM lokacija WHERE vrijeme<Now() - INTERVAL '1 month'"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
