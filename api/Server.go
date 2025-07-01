package api

import (
	"fmt"

	db "github.com/PhilaniAntony/simplebank/db/sqlc"
	"github.com/PhilaniAntony/simplebank/token"
	"github.com/PhilaniAntony/simplebank/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	//Routes without authentication
	router.POST("/auth/users", server.createUser)
	router.POST("/auth/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.PUT("/accounts/:id", server.updateAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	authRoutes.POST("/entries", server.createEntry)
	authRoutes.GET("/entries/:id", server.getEntry)
	authRoutes.GET("/entries", server.listEntries)
	authRoutes.POST("/transfers", server.createTransfer)
	authRoutes.GET("/transfers/:id", server.getTransfer)
	authRoutes.GET("/transfers", server.listTransfers)
	authRoutes.GET("/users/:username", server.getUser)
	authRoutes.GET("/users", server.listUsers)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Message": err.Error()}
}
