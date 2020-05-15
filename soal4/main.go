package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// cut string by the given length
func cut(str string, length int) int {
	a, _ := strconv.Atoi(str[0:length])
	return a
}

// check is the given number + 1 is equal to next number
func check(str string, n1, length int) bool {
	n2 := cut(str, length)
	if n1+1 == n2 {
		return true
	}
	return false
}

// check each digit
func loop(providedStringInput string, cutLengthInput int) ([]int, bool) {
	cutLength := cutLengthInput
	providedString := providedStringInput
	var n int
	arr := []int{}
	isCorrect := true
	totalFalse := 0
	for true {
		if len(providedString) <= cutLength {
			nx, _ := strconv.Atoi(providedString)
			arr = append(arr, nx)
			break
		}
		n = cut(providedString, cutLength)
		providedString = strings.Replace(providedString, strconv.Itoa(n), "", 1)
		arr = append(arr, n)
		if !check(providedString, n, cutLength) {
			if totalFalse >= 2 {
				isCorrect = false
				break
			}
			if check(providedString, n+1, cutLength) {
				continue
			} else if check(providedString, n, cutLength+1) {
				cutLength++
			} else {
				totalFalse++
			}
		}
	}
	return arr, isCorrect
}

// find first mis matched number
func getMissingNumber(arr []int) int {
	a := arr[:1][0]
	z := arr[len(arr)-1:][0]
	index := 0
	var missingNumber int
	for i := a; i < z; i++ {
		if i != arr[index] {
			missingNumber = i
			break
		}
		index++
	}
	return missingNumber
}

func findMissingNumber(str string) int {
	index := 1
	var missingNumber int
	for true {
		arr, isCorrect := loop(str, index)
		if isCorrect {
			missingNumber = getMissingNumber(arr)
			break
		}
		index++
	}
	return missingNumber
}

// OutputJSON format output json
type OutputJSON struct {
	MissingNumber int `json:"missing_number"`
}

func findMissingNumberController(w http.ResponseWriter, r *http.Request) {
	input, _ := r.URL.Query()["input"]
	number := OutputJSON{
		MissingNumber: findMissingNumber(input[0]),
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
