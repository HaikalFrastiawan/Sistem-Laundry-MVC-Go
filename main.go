package main

import (
	"net/http"
	"log"
	"Sistem-Laundry/config"
	"Sistem-Laundry/controllers/homecontroller"

)

func main() {
	
	//Connect DB
	config.ConnectDb()

	//routing
	http.HandleFunc("/", homecontroller.Index)


	log.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
