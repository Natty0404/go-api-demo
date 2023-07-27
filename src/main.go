package main

import (
	"encoding/json"
	"fmt"
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

	mux.Handle("/book/v1/", http.StripPrefix("/book/v1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bookID := r.URL.Query().Get("book_id")
			book, err := goDemoAPI.GetBook(r.Context(), bookID)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
		case http.MethodPost:
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("CreateBook success")
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("UpdateBook success")
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("DeleteBook success")
		default:
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

	// /book/v1/メソッド名
	// v1 := r.Group("/book/v1")
	// {
	// 	v1.GET("/GetBookList", func(context *gin.Context) {
	// 		context.JSON(200, gin.H{
	// 			"message": "GetBookList success",
	// 		})
	// 	})
	// 	v1.POST("/CreateBook", func(context *gin.Context) {
	// 		context.JSON(200, gin.H{
	// 			"message": "CreateBook success",
	// 		})
	// 	})
	// 	v1.PUT("/UpdateBook", func(context *gin.Context) {
	// 		context.JSON(200, gin.H{
	// 			"message": "UpdateBook success",
	// 		})
	// 	})
	// 	v1.DELETE("/DeleteBook", func(context *gin.Context) {
	// 		context.JSON(200, gin.H{
	// 			"message": "DeleteBook success",
	// 		})
	// 	})
	// }
	// log.Fatal(r.Run())

	if err := run(); err != nil {
		log.Fatalf("err: %v", err)
	}
}
