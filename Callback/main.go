package main

import (
	"fmt"
	"io"
	"net/http"
)

func Fetch(addr string) (string, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	return string(body), nil
}

func Sync() {
	for i := 0; i < 10; i++ {
		Fetch("http://localhost:8080")
	}
}

func Async() {
	for i := 0; i < 10; i++ {
		go Fetch("http://localhost:8080")
	}

	for {
	}
}

func main() {

	fmt.Println("Sync:")
	Sync()
	fmt.Println()
	fmt.Println("Async:")
	Async()

}
