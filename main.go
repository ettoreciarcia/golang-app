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
	PORT := getEnv("PORT", "3000")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func logging() {

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

	// Loop over header names
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	resp := make(map[string]string)
	resp["hostname"] = getHostname()
	resp["version"] = getEnv("VERSION", "1")
	resp["value"] = strconv.Itoa(getRandomValue(1, 100))
	resp["remoteAddr"] = r.RemoteAddr
	jsonResp, err := json.MarshalIndent(resp, "", "    ")

	//log request on a file
	f, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	//File as standard output
	log.SetOutput(f)
	log.Printf(
		"%s\t\t%s\t\t%s\t",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
	)

	w.Write(jsonResp)

}
