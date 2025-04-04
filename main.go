package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbDriver := os.Getenv("POSTGRES_DRIVER")
	dbSource := os.Getenv("POSTGRES_URL")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
