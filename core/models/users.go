package models

import "time"

// User represents a users table in the postgresql database.
type User struct {
	Id           int64      `json:"id" db:"id"`
	Email        string     `json:"email" db:"email"`
	FirstName    string     `json:"firstname" db:"firstname"`
	LastName     string     `json:"lastname" db:"lastname"`
	ParentUserId *int64     `json:"parent_user_id,omitempty" db:"parent_user_id"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"`
	DeletedAt    *time.Time `json:"updated_at,omitempty" db:"deleted_at"`
	MergedAt     *time.Time `json:"merged_at,omitempty" db:"merged_at"`
}
