package domain

import "time"

type Todo struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	IsDone    *bool      `json:"isDone"`
}
