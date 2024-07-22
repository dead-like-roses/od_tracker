package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to DB")

	return db
}

func RunMigrations(db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS devices(
                 device_id character(36) NOT NULL PRIMARY KEY UNIQUE,
                 name text NOT NULL
                 );`
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `CREATE TABLE IF NOT EXISTS activity(
			 device_id character(36) NOT NULL references devices(device_id),
			 posted_at timestamp with time zone NOT NULL,
			 app text NOT NULL
			 );`
	_, err = db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
