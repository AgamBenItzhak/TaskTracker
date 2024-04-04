package server

import (
	"net/http"

	"github.com/AgamBenItzhak/TaskTracker/config"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/routes"
	"github.com/AgamBenItzhak/TaskTracker/internal/server/services"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Server provides functions for the server
type Server struct {
	Router   *chi.Mux
	DB       *pgxpool.Pool
	Services *services.Services
}

// NewServer creates a new Server instance
func NewServer(db *pgxpool.Pool, token string) *Server {
	s := &Server{
		Router:   chi.NewRouter(),
		DB:       db,
		Services: services.NewServices(db, token),
	}

	routes.Routes(s.Router, s.Services)

	return s
}

// Run runs the server
func (s *Server) Run() {
	addr := config.ServerConfig.Server.Address + ":" + config.ServerConfig.Server.Port
	http.ListenAndServe(addr, s.Router)
}
