package api

import(
	"github.com/gin-gonic/gin"
)
tyoe Server struct {
	store db.Store
	router *gin.Engine


}
func NewServer(store db.Store) *Server{
	server := & Server{store: store}
	router:= gin.Default()

	router.Post("/accounts",serer.createAccount)
	router.Get("/accounts",server.createAccount)
	
	server.router =router
	return router

}
func (server.Server) Start(address string ) erroe 