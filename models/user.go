package models

import "time"

type User struct {
	CreatedAt time.Time `dynamo:"created_at"`
	Email     string    `dynamo:"email"`
	Name      string    `dynamo:"name"`
	Username  string    `dynamo:"username"`
	UserID    string    `dynamo:"user_id"`
}
