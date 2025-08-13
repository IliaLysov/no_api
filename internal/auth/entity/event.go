package entity

import "time"

type CreateEvent struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	OccurredAt time.Time   `json:"occurred_at"`
	Payload    interface{} `json:"payload"`
}

type UserCreated struct {
	Email string `json:"email"`
}

type UserLoggedIn struct {
	Email string `json:"email"`
	IP    string `json:"ip,omitempty"`
}

type UserProtected struct {
	ID string `json:"id"`
}
