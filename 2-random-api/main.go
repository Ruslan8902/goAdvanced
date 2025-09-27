package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func getRandomN(w http.ResponseWriter, req *http.Request) {
	fmt.Println(rand.Intn(6) + 1)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", getRandomN)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
