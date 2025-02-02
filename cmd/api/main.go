package main

import (
	"github.com/lpernett/godotenv"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"icu.imta.gsarbaj.social/internal/auth"
	"icu.imta.gsarbaj.social/internal/db"
	"icu.imta.gsarbaj.social/internal/env"
	"icu.imta.gsarbaj.social/internal/mailer"
	"icu.imta.gsarbaj.social/internal/store"
	"icu.imta.gsarbaj.social/internal/store/cache"
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
	//log.Printf(env.GetString("MAIL_TRAP_API_KEY", ""))

	cfg := config{
		address:     env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("API_URL", "http://localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:3000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://gsarbaj:!Genryh38312290966@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pwd:     env.GetString("REDIS_PASSWORD", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", true),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
			mailTrap: mailTrapConfig{
				apiKey: env.GetString("MAIL_TRAP_API_KEY", "7ded1226b6c39fa20ea7ccf793645ec6"),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				username: env.GetString("AUTH_BASIC_USER", "admin"),
				password: env.GetString("AUTH_BASIC_PASSWORD", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "secret"),
				exp:    time.Hour * 24 * 7,
				iss:    "social",
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

	// Cache
	var rdb *redis.Client
	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pwd, cfg.redisCfg.db)
		logger.Info("Redis connection established")
	}

	storeDB := store.NewStorage(dbConnection)
	cacheStorage := cache.NewRedisStorage(rdb)

	//mailer := mailer.NewSendGrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)
	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.fromEmail)
	if err != nil {
		logger.Fatal(err)
	}

	jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	app := &application{
		config:        cfg,
		store:         storeDB,
		cacheStorage:  cacheStorage,
		logger:        logger,
		mailer:        mailtrap,
		authenticator: jwtAuthenticator,
	}

	mux := app.mount()

	logger.Fatalln("MUX", app.run(mux))
}
