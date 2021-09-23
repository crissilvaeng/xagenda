package people

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

// PeopleHandler handles requests to the /people endpoint.
type PeopleHandler struct {
	// DB is the database connection.
	DB *gorm.DB
}

// Create adds a new person to the database.
// Returns 201 with the person if successful.
// Returns 400 if the request body is not valid JSON.
// Returns 500 if the person could not be created.
func (h *PeopleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var person Person
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

// GetAll retrieves all people from the database, except if a "q" query parameter was provided, then search for people with name match.
// Returns 200 with the people if successful.
// Returns 404 if no people were found.
// Returns 500 if the people could not be retrieved.
func (h *PeopleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	var people []Person
	if q == "" {
		if err := h.DB.Find(&people).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := h.DB.Where("name LIKE ?", "%"+q+"%").Find(&people).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(people); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Get retrieves a person from the database.
// Returns 200 with the person if successful.
// Returns 404 if no person was found.
// Returns 500 if the person could not be retrieved.
func (h *PeopleHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var person Person
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

// Update updates a person in the database.
// Returns 200 with the person if successful.
// Returns 400 if the request body is not valid JSON.
// Returns 404 if no person was found.
// Returns 500 if the person could not be updated.
func (h *PeopleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var person Person
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

// Delete removes a person from the database.
// Returns 200 if successful.
// Returns 404 if no person was found.
// Returns 500 if the person could not be removed.
func (h *PeopleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var person Person
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

// NewPeopleHandler returns a new PeopleHandler.
// Englobes the request to the database.
func NewPeopleHandler(db *gorm.DB) *PeopleHandler {
	return &PeopleHandler{DB: db}
}
