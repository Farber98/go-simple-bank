package api

import (
	db "go-simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}

	// Add currency validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	// Add routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)

	server.router = router
	return server
}

// Start runs the HTTP server on specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// errorResponse returns an error as gin.H (JSON)
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
