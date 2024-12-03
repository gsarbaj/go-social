package main

import (
	"icu.imta.gsarbaj.social/internal/db"
	"icu.imta.gsarbaj.social/internal/env"
	"icu.imta.gsarbaj.social/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://gsarbaj:!Genryh38312290966@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	str := store.NewStorage(conn)

	db.Seed(str)
}
