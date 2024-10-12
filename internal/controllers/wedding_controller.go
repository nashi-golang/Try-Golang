package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"try-golang/internal/services"
	"try-golang/pkg/dto"
)

type WeddingController struct {
	service *services.WeddingService
}

func NewWeddingController(service *services.WeddingService) *WeddingController {
	return &WeddingController{service: service}
}
func (ctrl *WeddingController) CreateWedding(c *gin.Context) {
	var weddingDto dto.WeddingDto
	if err := c.ShouldBindJSON(&weddingDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := ctrl.service.CreateWedding(weddingDto)
	c.JSON(200, response)
}
func (ctrl *WeddingController) GetWeddings(c *gin.Context) {
	response := ctrl.service.GetWeddings()
	c.JSON(200, response)
}
func (ctrl *WeddingController) GetWeddingById(c *gin.Context) {
	response := ctrl.service.GetWeddingById(c.Param("id"))
	c.JSON(200, response)
}
func (ctrl *WeddingController) UpdateWedding(c *gin.Context) {
	response := ctrl.service.UpdateWedding(c.Param("id"))
	c.JSON(200, response)
}
func (ctrl *WeddingController) DeleteWedding(c *gin.Context) {
	response := ctrl.service.DeleteWedding(c.Param("id"))
	c.JSON(200, response)
}
