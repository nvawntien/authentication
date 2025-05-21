package model

import "time"

type User struct {
	UserId    string    `db:"user_id, omitempty"`
	FullName  string    `db:"full_name, omitempty"`
	Email     string    `db:"email, omitempty"`
	Password  string    `db:"password, omitempty"`
	Role      string    `db:"role, omitempty"`
	CreatedAt time.Time `db:"created_at, omitempty"`
	UpdateAt  time.Time `db:"updated_at, omitempty"`
	Token     string
}
