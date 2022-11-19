package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Junkes887/translate/database"
	usecases "github.com/Junkes887/translate/usecases"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	usecases := usecases.Client{
		DB: database.CreateConnectionRedis(),
	}

	godotenv.Load()
	router := httprouter.New()
	c := cors.AllowAll()
	handlerCors := c.Handler(router)
	router.GET("/search", usecases.GetTranslateAndSearch)
	fmt.Println("STARTED API IN: 8090")
	log.Fatal(http.ListenAndServe(":8090", handlerCors))
}
