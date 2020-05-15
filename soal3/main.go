package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"test.pic/soal3/helper"
)

// OutputJSON format output json
type OutputJSON struct {
	MissingNumber int `json:"missing_number"`
}

func findMissingNumberController(w http.ResponseWriter, r *http.Request) {
	input, _ := r.URL.Query()["input"]
	number := OutputJSON{
		MissingNumber: helper.FindMissingNumber(input[0]),
	}
	json.NewEncoder(w).Encode(number)
}

func handleRequest() {
	http.HandleFunc("/find_missing_number", findMissingNumberController)
	port := "8080"
	fmt.Println("Server running on port:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	handleRequest()
}
