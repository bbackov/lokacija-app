package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "6NLG6HBW"
	dbname   = "završni"
	sslmode  = "disable"
)

func Connect() (*sql.DB, error) {
	infosql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", infosql)
	if err != nil {
		log.Println("Greška prilikom otvaranja konekcije:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Greška prilikom pingovanja baze:", err)
		return nil, err
	}

	fmt.Println("Konekcija uspešno ostvarena")
	return db, nil
}
