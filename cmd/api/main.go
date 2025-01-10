package main

import "log"

func main() {
	// Start the API server
	// This will listen for incoming requests
	// and serve the appropriate response
	app := &application{
		config: config{
			addr: ":8000",
		},
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
