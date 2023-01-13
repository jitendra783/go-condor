package router

import (
	"log"
	"net/http"
	"golean/controllers"
	"go get github.com/gocondor/installer/gocondor"
)
func RegisterRoutes() {
	router := routing.Resolve()
	log.Println(router)
	user := router.Group("user")
	{
		user.Get("/r",controllers.)
	}
	todos :=router.Group("todos")
	{
		todos.Get("/", controllers)
	}

}
