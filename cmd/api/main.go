package main

import (
	"github.com/lpernett/godotenv"
	"icu.imta.gsarbaj.social/internal/db"
	"icu.imta.gsarbaj.social/internal/env"
	"icu.imta.gsarbaj.social/internal/store"
	"log"
)

const version = "0.0.1"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//log.Println(os.Getenv("TEST"))

	cfg := config{
		address: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://gsarbaj:!Genryh38312290966@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	dbConnection, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err.Error())
	}

	defer dbConnection.Close()
	log.Println("Database connection established")

	storeDB := store.NewStorage(dbConnection)

	app := &application{
		config: cfg,
		store:  storeDB,
	}

	mux := app.mount()

	log.Fatalln(app.run(mux))
}
