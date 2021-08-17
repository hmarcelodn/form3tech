package test

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "interview_accountapi_user"
	password = "123"
	dbname   = "interview_accountapi"
)

func Truncate() {
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
