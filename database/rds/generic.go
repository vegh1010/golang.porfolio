package db_rds

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/pg.v5"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var CONFIG_FILE = "../../config.json"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func DBConnect(databaseName string) (*pg.DB) {
	fmt.Println("Database connection established")

	configs := GetConfig()
	for key, value := range configs {
		fmt.Println(key, ":", value)
	}

	database, exist := configs["PGSQL_DATABASE"]
	if !exist {
		database = "postgres"
	}
	flag.Parse()

	dbusername := configs["PGSQL_USER"]
	dbpassword := configs["PGSQL_PASS"]
	dbhost := configs["PGSQL_HOST"]
	dbport := configs["PGSQL_PORT"]
	dbprefix := configs["PGSQL_DB_PREFIX"]
	addr := fmt.Sprint(dbhost, ":", dbport)
	prefixDatabaseName := fmt.Sprint(dbprefix, "_", databaseName)

	fmt.Println("Address :", addr)
	fmt.Println("Database :", prefixDatabaseName)

	db := pg.Connect(&pg.Options{
		User:     dbusername,
		Password: dbpassword,
		Addr:     addr,
		Database: database,
	})
	return db
}

func GetDatabaseName(database string) (string) {
	configs := GetConfig()

	return configs["PGSQL_DB_PREFIX"] + "_" + database
}

func GetValueColumns(columns []string) (string) {
	values := []string{}
	for i := 0; i < len(columns); i++ {
		values = append(values, fmt.Sprint("$", i+1))
	}
	return strings.Join(values, ",")
}

func GetConfig() (map[string]string) {
	var configs = make(map[string]string)
	data, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Println("Failed to read config file: ", CONFIG_FILE)
		os.Exit(1)
	} else {
		//parse json string into object
		err = json.Unmarshal(data, &configs)
		Check(err)
	}
	return configs
}
