package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/nosleep", NoSleepHandler)

	fmt.Println("Server starting at", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Couldn't Start Server")
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(6)
	time.Sleep(time.Second * time.Duration(n))
	fmt.Fprintf(w, "Reply after %d seconds", n)
}

func NoSleepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reply immediately")
}
