package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"try-golang/internal/controllers"
	"try-golang/internal/services"
)

func main() {
	router := gin.Default()

	// 서비스 생성
	weddingService := services.NewWeddingService()

	// 컨트롤러 생성 및 서비스 주입
	weddingController := controllers.NewWeddingController(weddingService)
	router.POST("/wedding", weddingController.CreateWedding)
	router.GET("/wedding", weddingController.GetWeddings)
	router.GET("/wedding/:id", weddingController.GetWeddingById)
	router.PUT("/wedding/:id", weddingController.UpdateWedding)
	router.DELETE("/wedding/:id", weddingController.DeleteWedding)

	router.Run(":8999")
	fmt.Println("Server is running on port 8999")

}
