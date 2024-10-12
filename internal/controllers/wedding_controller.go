package controllers

import (
	"github.com/gin-gonic/gin"
	"try-golang/internal/services"
)

type WeddingController struct {
	service *services.WeddingService
}

func NewWeddingController(service *services.WeddingService) *WeddingController {
	return &WeddingController{service: service}
}
func (ctrl *WeddingController) CreateWedding(c *gin.Context) {
	response := ctrl.service.CreateWedding()
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
