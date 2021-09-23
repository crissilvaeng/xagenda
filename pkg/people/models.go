package people

import "time"

// Person is a struct that holds base info about and individual.
type Person struct {
	// ID is the primary key for the table.
	// As default, GORM uses a unsigned int for ID, but for real world environments, consider use more robust solution for ID.
	// For example, Ticket Server, UUID, Twitter Snowflake, etc.
	ID uint `json:"id" gorm:"primary_key"`
	// Name is used to perform search in Person model using simple LIKE statement.
	// This strategy is not optimized for performance, but for simplicity.
	// Please, refer other full_search, fuzzy_search, etc. strategies for more information.
	// PostreSQL (pg_trgm extension): https://www.postgresql.org/docs/9.6/pgtrgm.html
	// PostgreSQL (fuzzystrmatch extension): https://www.postgresql.org/docs/9.1/fuzzystrmatch.html
	// PostgresSQL (tsvector and tsquery): https://www.postgresql.org/docs/10/datatype-textsearch.html
	Name string `json:"name"`
	// Emails is a slice of emails of the person.
	// Like Google People API: https://developers.google.com/people/api/rest/v1/people/get
	Emails []string `json:"emails"`
	// Phone is a slice of phones of the person.
	// Like Google People API: https://developers.google.com/people/api/rest/v1/people/get
	Phones []string `json:"phones"`
	// CreatedAt is the time when the person was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time when the person was updated.
	UpdatedAt time.Time `json:"updated_at"`
}
