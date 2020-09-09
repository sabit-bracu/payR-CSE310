package handlers

import (
	"net/http"
	"payR/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCustomers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"customers": services.GetAllCustomers(),
		})
	}
}

func GetCustomerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
			return
		}
		customer, err := services.GetCustomerByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"customer": customer,
		})
	}
}
