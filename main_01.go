package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main01() {
	r := gin.Default()
	// Get method C=POST R=GET U=PUT D=DELETE

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD",
		})
	})
	r.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD",
		})
	})
	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD",
		})
	})
	r.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD",
		})
	})
	//-------------------------------------employee-----------------------------------------
	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD",
		})
	})
	//POST METHOD
	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD",
		})
	})
	//PUT METHOD
	r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD",
		})
	})
	//DELETE METHOD
	r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD",
		})
	})
	//---------------------------------customer-------------------------------------------
	//CUSTOMER

	r.GET("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "custome METHOD",
		})
	})

	r.PUT("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "customer METHOD",
		})
	})

	r.POST("/cusmoter", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "customer METHOD",
		})
	})

	r.DELETE("/cusmoter", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "customer METHOD",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
