package main

import (
	"fmt"
	"net/http"
	"os"
	router "signin/recurrent_rides_api/router"

	"github.com/gorilla/handlers"
	"github.com/subosito/gotenv"
)

//
func Init() {
	gotenv.Load()
}

func main() {
	Init()
	fmt.Println("Starting the application...")
	// Init router
	r := router.SetupRouter()
	// Listen and Serve in 0.0.0.0:8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"https://craftfoundry-signin-front.herokuapp.com/", "http://localhost:3000"}))(r))
	fmt.Println(err)
}
