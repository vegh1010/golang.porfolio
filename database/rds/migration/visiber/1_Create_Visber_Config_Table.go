package main

import (
	"github.com/vegh1010/golang.porfolio/database/rds"
	"gopkg.in/go-pg/migrations.v5"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		databaseName := db_rds.GetDatabaseName("visiber")
		upQuery := `CREATE TABLE ` + databaseName + `.visiber_config (
            "id" 	serial4 NOT NULL primary key,
            "key" 	text,
            "value" text
        );`

		_, err := db.Exec(upQuery)
		return err
	}, func(db migrations.DB) error {
		databaseName := db_rds.GetDatabaseName("visiber")
		downQuery := `DROP TABLE ` + databaseName + `.visiber_config;`
		_, err := db.Exec(downQuery)
		return err
	})
}
