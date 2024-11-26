package api

import (
	db "api/db/sqlc"

	"github.com/gin-gonic/gin"
)

// define the struct of the Http Server
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates a new instance server
func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	server.router = router
	return server
}

// Start runs the HTTP Server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
