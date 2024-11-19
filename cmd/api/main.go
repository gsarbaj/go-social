package main

import (
	"github.com/lpernett/godotenv"
	"icu.imta.gsarbaj.social/internal/env"
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
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatalln(app.run(mux))
}
