package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	//log.Printf("internal server error: %s path: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	//log.Printf("internal server error: %s path: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path)
	writeJSONError(w, http.StatusForbidden, "forbidden")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log.Printf("bad request error: %s path: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log.Printf("conflict error: %s path: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Errorf("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log.Printf("not found error: %s path: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log.Printf("not found error: %s path: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log.Printf("not found error: %s path: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}
