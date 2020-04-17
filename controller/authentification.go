package controller

import (
	"fmt"
	"net/http"
)

//endpoint creation
// func CreateEndpoint(w http.ResponseWriter, r *http.Request) {

// 	response, err := http.Post(os.Getenv("URL"), "application/json", r.Body)
// 	if err != nil {
// 		fmt.Printf("The HTTP request failed with error %s\n", err)
// 	} else {
// 		data, _ := ioutil.ReadAll(response.Body)
// 		w.WriteHeader(response.StatusCode)
// 		w.Write(data)
// 		fmt.Fprintf(w, "Hello world")
// 		return
// 	}

// }
func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
