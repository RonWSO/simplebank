package api

import (
	db "backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

// define the struct of the Http Server
type Server struct {
	store  *db.Store
	Router *gin.Engine
}

// Creates a new instance server
func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	server.Router = router
	return server
}

// Start runs the HTTP Server on a specific address.
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
