package middlewares

import (
    "fmt"

    "github.com/gin-gonic/gin"
)

var BeforeExample gin.HandlerFunc = func(c *gin.Context) {
    // logic of the middleware goes here
    c.Next()

    // Pass on to the next-in-chain
}
