package api

import (
	"github.com/crissilvaeng/xagenda/internal/pkg/support"
	"github.com/crissilvaeng/xagenda/pkg/people"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Service represents the service with fields to be injected into the handlers.
type Service struct {
	// Router is the router to be used by the handlers.
	Router *mux.Router
	// DB is the database to be used by the handlers.
	DB *gorm.DB
	// Logger is the logger to be used by the handlers.
	Logger *support.Logger
}

func (srv *Service) setRouters() {
	p := people.NewPeopleHandler(srv.DB)
	srv.Get("/people", p.GetAll)
	srv.Get("/people/{id}", p.Get)
	srv.Post("/people", p.Create)
	srv.Put("/people/{id}", p.Update)
	srv.Delete("/people/{id}", p.Delete)
}

// NewService returns a new service instance. It expects the service to be used in a HTTP server.
func NewService(r *mux.Router, db *gorm.DB, log *support.Logger) *Service {
	srv := &Service{
		Router: r,
		DB:     db,
		Logger: log,
	}
	srv.setRouters()
	return srv
}
