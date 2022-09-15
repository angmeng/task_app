package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	dbname := os.Getenv("DATABASE_NAME")
	dbhost := os.Getenv("DATABASE_HOST")
	dbuser := os.Getenv("DATABASE_USER")
	dbpass := os.Getenv("DATABASE_PASS")

	connectionString := "host=" + dbhost + " user=" + dbuser + " dbname=" + dbname + " sslmode=disable" + " password=" + dbpass
	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	schemaMigrations := &migrate.FileMigrationSource{
		Dir: "migrations/postgresql",
	}

	n, err := migrate.Exec(dbConn.DB, "postgres", schemaMigrations, migrate.Up)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Printf("Applied %d migrations!\n", n)
}
