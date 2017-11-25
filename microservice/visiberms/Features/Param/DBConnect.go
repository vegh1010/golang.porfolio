package visiberms_param

import (
	"database/sql"
	"fmt"
	"os"
)

func (self *Param) DBConnect() (err error) {
	self.DBPrefix = os.Getenv("PGSQL_DB_PREFIX")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("PGSQL_USER"),
		os.Getenv("PGSQL_PASS"),
		"postgres",
		os.Getenv("PGSQL_HOST"),
		os.Getenv("PGSQL_PORT"),
	)
	fmt.Println("PostgreSQL AUTH:", dbinfo)

	self.Postgres, err = sql.Open("postgres", dbinfo)
	if err != nil {
		return
	}
	fmt.Println("Database connection established")

	return
}
