package main

import (
	"log"
	"net/http"
)

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//_, err := w.Write([]byte(`"status": "ok"`))
	//if err != nil {
	//	return
	//}

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Println(err.Error())
		}
	}

}
