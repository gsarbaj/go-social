package main

import (
	"icu.imta.gsarbaj.social/internal/store"
	"net/http"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	//pagination, filters, sort

	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(42), fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if errr := app.jsonResponse(w, http.StatusOK, feed); errr != nil {
		app.internalServerError(w, r, errr)
	}

}
