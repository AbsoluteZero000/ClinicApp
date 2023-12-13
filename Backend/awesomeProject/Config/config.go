package Config

import (
	"database/sql"
    "os"
	"fmt"
)

func Connect() *sql.DB {
	// dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "12345678"
	// dbName := "ClinicApp"

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(database:3306)/"+dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	return db
}
