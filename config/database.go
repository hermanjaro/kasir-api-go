package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

var DB *sql.DB

func ConnectDB() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables: %v", err)
	}

	dbURL := viper.GetString("DB_URL")
	if dbURL == "" {
		return fmt.Errorf("DB_URL environment variable is not set")
	}

	var err error
	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		return err
	}

	// Configure connection pool for Supabase pooler
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(2)
	DB.SetConnMaxLifetime(300)

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Successfully connected to PostgreSQL database")
	return nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
