package main

import (
	"github.com/lpernett/godotenv"
	"go.uber.org/zap"
	"icu.imta.gsarbaj.social/internal/db"
	"icu.imta.gsarbaj.social/internal/env"
	"icu.imta.gsarbaj.social/internal/mailer"
	"icu.imta.gsarbaj.social/internal/store"
	"log"
	"time"
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
		address:     env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("API_URL", "http://localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:4000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://gsarbaj:!Genryh38312290966@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database
	dbConnection, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		//log.Println("DB ERROR", err.Error())
		logger.Fatal(err)
	}

	defer dbConnection.Close()
	logger.Info("Database connection established")

	storeDB := store.NewStorage(dbConnection)

	mailer := mailer.NewSendGrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)

	app := &application{
		config: cfg,
		store:  storeDB,
		logger: logger,
		mailer: mailer,
	}

	mux := app.mount()

	logger.Fatalln("MUX", app.run(mux))
}
