package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"try-golang/internal/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"try-golang/internal/handlers"
	"try-golang/internal/models"
	"try-golang/internal/services"
)

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("weddings.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 테이블 자동 생성
	db.AutoMigrate(&models.Wedding{}, &models.People{})

	//레포지토리 생성
	weddingRepository := repository.NewWeddingRepository(db)
	peopleRepository := repository.NewPeopleRepository(db)

	// 서비스 생성
	weddingService := services.NewWeddingService(weddingRepository, peopleRepository)

	// 컨트롤러 생성 및 서비스 주입
	router.POST("/wedding", handlers.CreateWedding(weddingService))
	router.GET("/wedding", handlers.GetWeddings(weddingService))
	router.GET("/wedding/:id", handlers.GetWeddingById(weddingService))
	router.PUT("/wedding/:id", handlers.UpdateWedding(weddingService))
	router.DELETE("/wedding/:id", handlers.DeleteWedding(weddingService))
	router.POST("/wedding/:id/people", handlers.CreatePeople(weddingService))
	// 서버 시작 전에 로그 출력
	fmt.Println("Server is starting on port 8999")

	// 서버 실행
	if err := router.Run(":8999"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
