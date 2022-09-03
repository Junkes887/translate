package main

import (
	"log"
	"net/http"

	usecases "github.com/Junkes887/translate/usecases"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	http.HandleFunc("/search", usecases.GetTranslateAndSearch)
	http.HandleFunc("/translate", usecases.GetTranslate)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
