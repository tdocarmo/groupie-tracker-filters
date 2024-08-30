package main

import (
	"groupie-tracker-filters/routes"
	"log"
	"net/http"
)

func main() {
	// Serve static files (CSS and JS)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./client/style"))))

	// Initialize routes for error pages
	routes.InitErrorRoutes()

	// Initialize all application routes
	routes.Init()

	// Start the server on port 3001
	log.Fatal(http.ListenAndServe(":3001", nil))
}
