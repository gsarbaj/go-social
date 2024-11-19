package main

import "log"

func main() {

	cfg := config{
		address: ":8080",
	}

	app := &application{
		config: cfg,
	}

	log.Fatalln(app.run())

}
