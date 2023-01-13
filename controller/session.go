package controllers
import (
	"log"
	"github.com/gin-gonic/gin"

)
func Sessions(c *gin.Context) {
    Session.Set("key", "value", c)
    Session.Get("key", c)
    Session.Has("key", c)
    Session.Delete("key", c)
    Session.Pull("key", c)
    Session.Clear(c)
	log.Println(c)

}
func RedirectSession(c *gin.Context){
    c.Redirect(http.StatusMovedPermanently, "https://google.com")
}
func JwtTokenCreation (c *gin.Context) {
    payload := map[string]string{"userId": "123"}
	//Generate JWT token
	token, err := JWT.CreateToken(payload)

//Extract authorization token from header
	token, error = JWT.ExtractToken(c)


	//Decode a token
		// first let's extract the token from the request header
	token, error = JWT.ExtractToken(c)
	
		// decode
	payload, error = JWT.DecodeToken(token)
	
	// Generate JWT refresh token#
		payload = map[string]string{"userId": "123"}
		token, error = JWT.CreateRefreshToken(payload)
	

	// Validate a token#

		// first let's extract the token from the request header
		token, error := JWT.ExtractToken(c)
	
		// validate
		ok, err := JWT.ValidateToken(token)

}
