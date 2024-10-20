package handlers

import (
	"github.com/google/uuid"
	"net/http"
	"try-golang/internal/services"
	"try-golang/pkg/dto"

	"github.com/gin-gonic/gin"
)

func CreateWedding(service *services.WeddingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var weddingDto dto.WeddingDto
		if err := c.ShouldBindJSON(&weddingDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		message, err := service.CreateWedding(weddingDto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, message)
	}
}

func GetWeddings(service *services.WeddingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := service.GetWeddings()
		c.JSON(http.StatusOK, response)
	}
}

func GetWeddingById(service *services.WeddingService) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		weddingID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is not UUID"})
			return
		}
		response, err := service.GetWeddingById(weddingID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Wedding not found"})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func UpdateWedding(service *services.WeddingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		weddingID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is not UUID"})
			return
		}
		var weddingDto dto.WeddingDto
		if err := c.ShouldBindJSON(&weddingDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := service.UpdateWedding(weddingID, weddingDto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func DeleteWedding(service *services.WeddingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		weddingID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is not UUID"})
			return
		}
		err = service.DeleteWedding(weddingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

func CreatePeople(service *services.WeddingService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		weddingID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is not UUID"})
			return
		}
		var peopleDto dto.PeopleDto
		if err := c.ShouldBindJSON(&peopleDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := service.CreatePeople(weddingID, peopleDto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
