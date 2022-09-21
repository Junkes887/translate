package main

import (
	"log"
	"net/http"

	usecases "github.com/Junkes887/translate/usecases"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load()
	router := httprouter.New()
	c := cors.AllowAll()
	handlerCors := c.Handler(router)
	router.GET("/search", usecases.GetTranslateAndSearch)
	router.GET("/translate", usecases.GetTranslate)
	log.Fatal(http.ListenAndServe(":8090", handlerCors))
}
