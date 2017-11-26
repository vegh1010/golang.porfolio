package main

import (
	"flag"
	"fmt"
	"github.com/vegh1010/golang.porfolio/database/rds"
	"gopkg.in/go-pg/migrations.v5"
)

const verbose = true

//TODO: discuss on changes to voice schema

func main() {
	databaseName := "visiber"

	migrationDB := db_rds.DBConnect(databaseName)
	fmt.Println(flag.Args())
	database := db_rds.GetDatabaseName(databaseName)
	migrationName := database + `.gopg_migrations`
	migrations.SetTableName(migrationName)

	oldVersion, newVersion, err := migrations.Run(migrationDB, flag.Args()...)
	db_rds.Check(err)
	if verbose {
		if newVersion != oldVersion {
			fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
		} else {
			fmt.Printf("version is %d\n", oldVersion)
		}
	}
}
