package api

import (
	"github.com/crissilvaeng/xagenda/internal/pkg/people"
	"github.com/crissilvaeng/xagenda/internal/pkg/support"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Service struct {
	Router *mux.Router
	DB     *gorm.DB
	Logger *support.Logger
}

func (srv *Service) setRouters() {
	p := people.NewPeopleHandler(srv.DB)
	srv.Get("/people", p.GetAll)
	srv.Get("/people/{id}", p.Get)
	srv.Post("/people", p.Create)
	srv.Put("/people/{id}", p.Update)
	srv.Delete("/people/{id}", p.Delete)
	srv.Query("/people", p.Search)
}

func NewService(r *mux.Router, db *gorm.DB, log *support.Logger) *Service {
	srv := &Service{
		Router: r,
		DB:     db,
		Logger: log,
	}
	srv.setRouters()
	return srv
}
