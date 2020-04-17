package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	router "signin/recurrent_rides_api/router"

	"github.com/subosito/gotenv"
)

//
func Init() {
	gotenv.Load()
}

func main() {
	log.Println(os.Getenv("APP_ID"))
	log.Println(os.Getenv("APP_SECRET"))
	Init()
	fmt.Println("Starting the application...")
	// Init router
	r := router.SetupRouter()
	// Listen and Serve in 0.0.0.0:8080
	port := os.Getenv("PORT")
	if port == "" {

		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
	fmt.Println("Terminating the application...")
}
