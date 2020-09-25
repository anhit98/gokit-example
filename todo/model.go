package todo

import (
	"time"
)

type Todo struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
