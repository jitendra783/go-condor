package controllers
import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

)
type  User Struct {
   Name string	`json:"name"`

}
func GetUser(c *gin.Context){
	c.Param("id")
	log.Println("njkdkjf")

	c.HTML(http.StatusOK, "index.html", gin.H{
        "message": "this is a message",
    })



}
func AllGetUser(c *gin.Context){
	log.Println("njkdkjf")
	c.JSON(http.StatusOK, gin.H{
        "message": "this is a message",
    })

}

func CreateUser(c *gin.Context){
	var user User
    err := c.ShouldBindJSON(&user)
    if err != nil {
        // validation failed
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // validation passed, the variable user is filled with the request data
    fmt.Println(user)
}
	log.Println("njkdkjf")

}
func UpdateUser(c *gin.Context){
	c.Param("id")
	
}
func DeleteUser(c *gin.Context){

		c.Param("id")
		c.XML(http.StatusOK, gin.H{
			"message": "this is a message",
		})

}