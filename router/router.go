package router

import (
	"signin/recurrent_rides_api/controller"

	"github.com/gorilla/mux"
)

//router
func SetupRouter() *mux.Router {

	router := mux.NewRouter()
	// routes
	router.HandleFunc("/login", controller.CreateEndpoint).Methods("POST")
	return router
}
