package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type WeddingController struct {
}

func CreateWedding(c *gin.Context) {
	c.JSON(200, "create")
}
func GetWeddings(c *gin.Context) {
	c.JSON(200, "get all")
}
func GetWeddingById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, fmt.Sprintf("get: %s", id))
}
func UpdateWedding(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, fmt.Sprintf("update: %s", id))
}
func DeleteWedding(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, fmt.Sprintf("delete: %s", id))
}
