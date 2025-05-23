package logika

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var TajniKljuc = []byte("kljuc")

var PoništeniTokeni = make(map[string]bool)

type ZahtjeviZaToken struct {
	Email   string
	Lozinka string
}

type PrijavaOdgovor struct {
	Token string `json:"token"`
}

func ValidirajToken(tokenString string) (*jwt.Token, error) {
	if PoništeniTokeni[tokenString] {
		return nil, errors.New("token je poništen")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("neispravan token")
		}
		return TajniKljuc, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DohvatiIDKorisnika(r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := ValidirajToken(tokenString)
	if err != nil || !token.Valid {
		return 0, errors.New("nevažeći token")
	}

	claims := token.Claims.(jwt.MapClaims)
	return int(claims["id_korisnik"].(float64)), nil
}

func PonistiToken(tokenString string) {
	PoništeniTokeni[tokenString] = true
}

func Provjeravlasnistva(db *sql.DB, idUredaj int, idKorisnik int) error {
	var vlasnikID int
	err := db.QueryRow("SELECT id_korisnik FROM uređaj WHERE id_uredaj = $1", idUredaj).Scan(&vlasnikID)
	if err != nil {
		return err
	}

	if vlasnikID != idKorisnik {
		return errors.New("nemate pravo dodavati lokaciju za ovaj uređaj")
	}
	return nil
}
