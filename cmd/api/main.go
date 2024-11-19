package main

import (
	"github.com/lpernett/godotenv"
	"icu.imta.gsarbaj.social/internal/env"
	"icu.imta.gsarbaj.social/internal/store"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//log.Println(os.Getenv("TEST"))

	cfg := config{
		address: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", ":5432"),
		},
	}

	storeDB := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  storeDB,
	}

	mux := app.mount()

	log.Fatalln(app.run(mux))
}
