package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler(c *gin.Context) {

	newvariable := c.Param("new") // get the parameter normal :param after url itself
	c.JOSN(http.StatusOK, gin.H{
		"message": newvariable,
	})
	log.Println("dhfjd")
	c.Abort()
	return
}

func Auth(c *gin.Context) {
	authToken := c.Request.Header["Authorization"] // get the 'authorization' header
	c.JOSN(http.StatusOK, gin.H{
		"message": authToken,
	})
	log.Println("dhfjd")
	c.AbortWithStatus(http.StatusBadRequest)
	return
}

func QueryHandler(c *gin.Context) {
	name := c.Query("name") // get the parameter from query parameters
	if name == "" {
		name = c.DefaultQuery("name", "guest") //use guest as default if name is missing

	}
	c.JOSN(http.StatusOK, gin.H{
		"message": name,
	})
	log.Println("dhfjd")
	c.Abort()
	return
}
