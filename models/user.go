package models

import "time"

type User struct {
	CreatedAt time.Time `dynamo:"created_at,unixtime" json:"created_at,omitempty"`
	Email     string    `dynamo:"email" json:"email"`
	Name      string    `dynamo:"name" json:"name"`
	Username  string    `dynamo:"username" json:"username"`
	UserID    string    `dynamo:"user_id" json:"user_id"`
}
