package main

import (
	"net/http"
	"log"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if len(port) < 10 {
		port = "3000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world, power by Golang"))
	})

	log.Printf("Listen on portt %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
