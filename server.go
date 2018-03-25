package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// NOTE: Ideally there should be api prefix like "/api/v1"
	http.HandleFunc("/", helloWorld)

	// NOTE: We can accept host/port from flags as well
	address := "0.0.0.0:3333"
	log.Println("Server running on", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

// Message struct to be displayed
type Message struct {
	Text string `json:"text"`
}

// helloWorld endpoint
func helloWorld(res http.ResponseWriter, req *http.Request) {

	message := &Message{
		Text: "hello world",
	}

	renderJSON(res, http.StatusOK, message)
}

// RenderJSON as rest response
func renderJSON(w http.ResponseWriter, status int, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// NOTE: Uncomment to display api response
	// log.Println("Render response", obj)
	return json.NewEncoder(w).Encode(obj)
}
