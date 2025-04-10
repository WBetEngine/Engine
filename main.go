
package main

import (
	"log"
	"net/http"

	"whitelabel/internal/customer/config"
	"whitelabel/internal/customer/routers"
)

func main() {
	// Inisialisasi aplikasi customer
	log.Println("Menginisialisasi aplikasi...")
	config.Init()
	
	// Set up file server for static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup customer routes
	routers.SetupRoutes()

	// Start the server
	log.Println("Server starting on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
