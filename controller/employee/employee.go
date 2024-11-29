package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all employees
func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Retrieved all employees",
	})
}

// GET employee by ID
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

// PUT method handler
func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated employee",
	})
}

// POST method handler by ID
func PostEmployee(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Created or updated employee",
		"id":      id,
	})
}

// DELETE method handler
func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted employee",
	})
}
