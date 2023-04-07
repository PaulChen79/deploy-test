package gorm

import (
	"deploy-test/domain"
	"time"
)

type Todo struct {
	ID        uint       `gorm:"auto_increment primary_key"`
	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"default:null"`
	Title     string     `gorm:"type:varchar(255);not null"`
	Content   string     `gorm:"type:varchar(255);not null"`
	IsDone    *bool      `gorm:"type:int;not null;default:0"`
}

func TodoModelToDomain(input *Todo) *domain.Todo {
	return &domain.Todo{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		DeletedAt: input.DeletedAt,
		Title:     input.Title,
		Content:   input.Content,
		IsDone:    input.IsDone,
	}
}

func TodoDomainToModel(input *domain.Todo) *Todo {
	return &Todo{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		DeletedAt: input.DeletedAt,
		Title:     input.Title,
		Content:   input.Content,
		IsDone:    input.IsDone,
	}
}
