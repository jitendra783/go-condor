package middlewares
import (
	"log"
    "github.com/gin-gonic/gin"

)
//JWT tokens are used for authentication,
func RegisterMiddlewares() {
    mwEngine := middlewares.Resolve()

    //Register your middlewares here
    mwEngine.Attach(BeforeExample)
	log.Println(BeforeExample)
}
//Check if a user is logged in
func Authverify(c *gin.Context) {
    Auth := auth.Resolve()
    ok, err := Auth.Check(c)
    if ok {
        // user is logged in
    }
}
//Get logged-in user's ID
func SomeHandler(c *gin.Context) {
    Auth := auth.Resolve()
    userID, err := Auth.UserID(c)
}