package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var v = viper.New()

func init() {
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath("../config")
}

func GetConnectionDB() *sql.DB {
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		v.GetString("DB_USER"),
		v.GetString("DB_PASS"),
		v.GetString("DB_HOST"),
		v.GetString("DB_PORT"),
		v.GetString("DB_NAME"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("cannot connect to database: %w", err))
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
