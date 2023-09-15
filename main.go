package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	cluster0()

	mux := chi.NewRouter()

	mux.Post("/api", CreateUser)

	fmt.Println("Server is running on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
