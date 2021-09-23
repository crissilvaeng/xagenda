package people

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateBadRequest(t *testing.T) {
	var conn *sql.DB

	conn, _, err := sqlmock.New()
	require.NoError(t, err)
	defer conn.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	require.NoError(t, err)

	handler := NewPeopleHandler(db)

	req := httptest.NewRequest(http.MethodGet, "/person", nil)
	w := httptest.NewRecorder()
	handler.Create(w, req)

	resp := w.Result()
	if _, err = ioutil.ReadAll(resp.Body); err != nil {
		t.Error(err)
	}

	if http.StatusBadRequest != resp.StatusCode {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestCreateInternalServerError(t *testing.T) {
	var conn *sql.DB

	conn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer conn.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	require.NoError(t, err)

	handler := NewPeopleHandler(db)

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "people" ("name","email","phone") VALUES ($1,$2,$3) RETURNING "id"`).WithArgs("Joe Doe", "john@doe.com", "123456789").WillReturnError(sql.ErrConnDone)
	mock.ExpectCommit()

	req := httptest.NewRequest(http.MethodGet, "/person",
		strings.NewReader(`{"name": "Joe Doe", "email": "john@doe.com", "phone": "123456789"}`))
	w := httptest.NewRecorder()
	handler.Create(w, req)

	resp := w.Result()
	if _, err = ioutil.ReadAll(resp.Body); err != nil {
		t.Error(err)
	}

	if http.StatusInternalServerError != resp.StatusCode {
		t.Errorf("expected %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestCreateOk(t *testing.T) {
	var conn *sql.DB

	conn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer conn.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	require.NoError(t, err)

	handler := NewPeopleHandler(db)

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "people" ("name","email","phone") VALUES ($1,$2,$3) RETURNING "id"`).WithArgs("Joe Doe", "john@doe.com", "123456789").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	req := httptest.NewRequest(http.MethodGet, "/person",
		strings.NewReader(`{"name": "Joe Doe", "email": "john@doe.com", "phone": "123456789"}`))
	w := httptest.NewRecorder()
	handler.Create(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if strings.TrimSpace(string(body)) != `{"id":1, "name": "Joe Doe", "email": "john@doe.com", "phone": "123456789"}` {
		t.Errorf("expected %s, got %s", `{"id":1, "name": "Joe Doe", "email": "john@doe.com", "phone": "123456789"}`, string(body))
	}

	if http.StatusCreated != resp.StatusCode {
		t.Errorf("expected %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}
