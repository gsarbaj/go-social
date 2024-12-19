package main

import (
	"github.com/lpernett/godotenv"
	"icu.imta.gsarbaj.social/internal/db"
	"icu.imta.gsarbaj.social/internal/env"
	"icu.imta.gsarbaj.social/internal/store"
	"log"
)

const version = "0.0.1"

//	@title			Imta Example API
//	@description	API for Imta golang code studies
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//log.Println(os.Getenv("TEST"))

	cfg := config{
		address: env.GetString("ADDR", ":8080"),
		apiURL:  env.GetString("API_URL", "http://localhost:8080"),
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
		log.Println("DB ERROR", err.Error())
	}

	defer dbConnection.Close()
	log.Println("Database connection established")

	storeDB := store.NewStorage(dbConnection)

	app := &application{
		config: cfg,
		store:  storeDB,
	}

	mux := app.mount()

	log.Fatalln("MUX", app.run(mux))
}
