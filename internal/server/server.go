package server

import (
	"github.com/AgamBenItzhak/TaskTracker/internal/db"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var ServerInstance *Server

// Server provides functions for the server
type Server struct {
	Router   *gin.Engine
	DB       db.PgxIface
	Services *services.Services
}

// NewServer creates a new Server instance
func NewServer() (*Server, error) {
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}

	router := gin.Default()
	services := services.NewServices(db.Pool)

	return &Server{
		Router:   router,
		DB:       db.Pool,
		Services: services,
	}, nil
}

func init() {
	server, err := NewServer()
	if err != nil {
		panic(err)
	}
	ServerInstance = server
}

// Run runs the server
func (s *Server) Run() {
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	s.Router.Run(addr)
}
