package people

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/crissilvaeng/xagenda/pkg/people"
	"gorm.io/gorm"
)

type PeopleHandler struct {
	DB *gorm.DB
}

func (h *PeopleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var person people.PersonInfo
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.DB.Create(&person).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PeopleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var people []people.PersonInfo
	if err := h.DB.Find(&people).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(people); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PeopleHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var person people.PersonInfo
	if err := h.DB.First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PeopleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var person people.PersonInfo
	if err := h.DB.First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.DB.Save(&person).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PeopleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var person people.PersonInfo
	if err := h.DB.First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.DB.Delete(&person).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *PeopleHandler) Search(w http.ResponseWriter, r *http.Request) {}

func NewPeopleHandler(db *gorm.DB) *PeopleHandler {
	return &PeopleHandler{DB: db}
}
