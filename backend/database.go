package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// LoadDatabase opens the database connection
func LoadDatabase() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			viper.GetString("db_user"),
			viper.GetString("db_pass"),
			viper.GetString("db_host"),
			viper.GetString("db_port"),
			viper.GetString("db_name"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}
	return db, nil
}
