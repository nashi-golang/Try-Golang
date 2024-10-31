package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"try-golang/internal/handlers"
	"try-golang/internal/models"
	"try-golang/internal/repository"
	"try-golang/internal/services"
)

// @title Try-Gorang
// @version 1.0
// @description Try Try Try
// @host localhost:8999
// @BasePath /

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("weddings.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 테이블 자동 생성
	db.AutoMigrate(&models.People{}, &models.Wedding{}, &models.Photo{})

	//레포지토리 생성
	weddingRepository := repository.NewWeddingRepository(db)
	peopleRepository := repository.NewPeopleRepository(db)
	photoRepository := repository.NewPhotoRepository(db)

	// 서비스 생성
	weddingService := services.NewWeddingService(weddingRepository, peopleRepository, photoRepository)

	// 컨트롤러 생성 및 서비스 주입
	api := router.Group("/wedding")
	handlers.RegisterRoutes(api, weddingService)

	// 서버 시작 전에 로그 출력
	fmt.Println("Server is starting on port 8999")

	// 서버 실행
	if err := router.Run(":8999"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
