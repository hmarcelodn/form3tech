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

func getConnection() (*sql.DB, error) {
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

	return db, nil
}

func Seed() {
	db, err := getConnection()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	insert :=
		`
	INSERT INTO "Account" (id, organisation_id, version, is_deleted, is_locked, created_on, modified_on, record, pagination_id) 
	VALUES (
		'0C879B45-CEEF-4350-946C-D672CDC43FB5',
		'D612C78B-7C54-4985-919F-0E393F034E0D',
		0,
		false,
		false,
		'2021-08-18 04:42:34.931617',
		'2021-08-18 04:42:34.931617',
		'{"bic": "NWBKGB22", "name": ["Marcelo Del Negro"], "bank_id": "400300", "country": "GB", "bank_id_code": "GBDSC", "alternative_bank_account_names": null}',
		0
	);

	INSERT INTO "Account" (id, organisation_id, version, is_deleted, is_locked, created_on, modified_on, record, pagination_id) 
	VALUES (
		'0C879B45-CEEF-4350-946C-D672CDC43FB6',
		'D612C78B-7C54-4985-919F-0E393F034E0E',
		0,
		false,
		false,
		'2021-08-18 04:42:34.931617',
		'2021-08-18 04:42:34.931617',
		'{"bic": "NWBKGB22", "name": ["Hugo Del Negro"], "bank_id": "400300", "country": "GB", "bank_id_code": "GBDSC", "alternative_bank_account_names": null}',
		1
	)
	`

	if _, err := db.Exec(insert); err != nil {
		panic(err)
	}
}

func Truncate() {
	db, connErr := getConnection()

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
