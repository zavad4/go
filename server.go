package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//defining structure for response
type date struct {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/time", handleTime)
	fmt.Printf("Listening at port 8795\n")
	http.ListenAndServe(":8795", nil)
}

func handleTime(w http.ResponseWriter, req *http.Request) {
	var currentTime = time.Now().Format(time.RFC3339)
	var response = date{currentTime}
	var jsonResponse, error = json.Marshal(&response)
	if error == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}

//comment for commit Zavodovska