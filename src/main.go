package main

import (
	"go-api-demo/infrastructure/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("start server...")
	r := gin.Default()
	r.Use(middleware.Logger())

	// /book/v1/メソッド名
	v1 := r.Group("/book/v1")
	{
		v1.GET("/GetBookList", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "GetBookList success",
			})
		})
		v1.POST("/CreateBook", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "CreateBook success",
			})
		})
		v1.PUT("/UpdateBook", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "UpdateBook success",
			})
		})
		v1.DELETE("/DeleteBook", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "DeleteBook success",
			})
		})
	}

	log.Fatal(r.Run())
}
