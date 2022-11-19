package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	PORT := getEnv("PORT", "8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

// Check to verify if a env variables exists, if it doesn't it return a default value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getRandomValue(rangeLower int, rangeUpper int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(rangeUpper-rangeLower) + rangeLower
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hostname
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := make(map[string]string)
	resp["message"] = "Status Created"
	resp["hostname"] = getHostname()
	resp["version"] = getEnv("VERSION", "1")
	resp["value"] = strconv.Itoa(getRandomValue(1, 100))
	jsonResp, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}
