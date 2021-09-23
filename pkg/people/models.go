package people

import "time"

// PersonInfo is a struct that holds base info about and individual.
// As default, GORM uses a unsigned int for ID, but for real world environments, consider use more robust solution for ID. For example, Ticket Server, UUID, Twitter Snowflake, etc.
type PersonInfo struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}
