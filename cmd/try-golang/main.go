package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"try-golang/internal/controllers"
)

func main() {
	router := gin.Default()
	router.POST("/wedding", controllers.CreateWedding)
	router.GET("/wedding", controllers.GetWeddings)
	router.GET("/wedding/:id", controllers.GetWeddingById)
	router.PUT("/wedding/:id", controllers.UpdateWedding)
	router.DELETE("/wedding/:id", controllers.DeleteWedding)

	router.Run(":8999")
	fmt.Println("Server is running on port 8999")

}
