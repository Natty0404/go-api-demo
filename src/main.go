package main

import (
	"encoding/json"
	"fmt"
	"go-api-demo/handlers"
	"go-api-demo/infrastructure/middleware"
	"go-api-demo/interfaceadapter/controller"
	"go-api-demo/interfaceadapter/gateway"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}

func run() error {

	mux := http.NewServeMux()

	goDemoAPI := controller.NewGoApiDemo(
		gateway.NewBook(),
	)

	bookHandlers := map[string]func(http.ResponseWriter, *http.Request){
		http.MethodGet: handlers.GetBookHandler(goDemoAPI),
		// http.MethodPost: handlers.PostBookHandler(goDemoAPI),
		// http.MethodPut: handlers.PutBookHandler(goDemoAPI),
		// http.MethodDelete: handlers.DeleteBookHandler(goDemoAPI),
	}

	mux.Handle("/book/v1/", http.StripPrefix("/book/v1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if handler, ok := bookHandlers[r.Method]; ok {
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Not Found")
		}
	})))

	mux.Handle("/healthcheck", http.HandlerFunc(healthCheck))

	return nil
}

func main() {
	log.Println("start server...")
	r := gin.Default()
	r.Use(middleware.Logger())

	if err := run(); err != nil {
		log.Fatalf("err: %v", err)
	}
}
