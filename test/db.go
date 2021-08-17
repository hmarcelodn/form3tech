package test

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DATABASE_HOST")
	port     = os.Getenv("PSQL_PORT")
	user     = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	dbname   = os.Getenv("DATABASE_NAME")
)

func Truncate() {
	portInt, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)

	db, connErr := sql.Open("postgres", psqlConnectionString)

	if connErr != nil {
		panic(connErr)
	}

	defer db.Close()

	if pingErr := db.Ping(); pingErr != nil {
		panic(pingErr)
	}

	if _, err := db.Exec(`TRUNCATE TABLE "Account"`); err != nil {
		panic(err)
	}
}
