package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwager "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
	"icu.imta.gsarbaj.social/docs" // This is required to generate swagger docs
	"icu.imta.gsarbaj.social/internal/auth"
	"icu.imta.gsarbaj.social/internal/mailer"
	"icu.imta.gsarbaj.social/internal/store"
	"net/http"
	"time"
)

type application struct {
	config        config
	store         store.Storage
	logger        *zap.SugaredLogger
	mailer        mailer.Client
	authenticator auth.Authenticator
}

type config struct {
	address     string
	db          dbConfig
	env         string
	apiURL      string
	mail        mailConfig
	frontendURL string
	auth        authConfig
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	username string
	password string
}

type mailConfig struct {
	sendGrid  sendGridConfig
	mailTrap  mailTrapConfig
	fromEmail string
	exp       time.Duration
}

type sendGridConfig struct {
	apiKey string
}

type mailTrapConfig struct {
	apiKey string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.With(app.BasicAuthMiddleware()).Get("/health", app.healthHandler)

		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.address)
		r.Get("/swagger/*", httpSwager.Handler(httpSwager.URL(docsURL)))

		//Posts

		r.Route("/posts", func(r chi.Router) {
			r.Use(app.AuthTokenMiddleware)
			r.Post("/", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {

				r.Use(app.postContextMiddleware)

				r.Get("/", app.getPostHandler)
				r.Delete("/", app.checkPostOwnership("admin", app.deletePostHandler))
				r.Patch("/", app.checkPostOwnership("moderator", app.updatePostHandler))

			})
		})

		//users

		r.Route("/users", func(r chi.Router) {

			r.Put("/activate/{token}", app.activateUserHandler)

			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)

				r.Get("/", app.getUserHandler)

				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)
				r.Get("/feed", app.getUserFeedHandler)
			})
		})

		// Public

		r.Route("/authentication", func(r chi.Router) {
			r.Post("/user", app.RegisterUserHandler)
			r.Post("/token", app.createTokenHandler)
		})
	})

	//auth

	return r
}

func (app *application) run(mux http.Handler) error {

	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

	srv := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	app.logger.Infow("Server has started", "addr", app.config.address, "env", app.config.env)

	return srv.ListenAndServe()
}
