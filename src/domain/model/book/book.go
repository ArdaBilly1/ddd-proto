package book

import "time"

type Book struct {
	ID        int       `json:"book" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:50"`
	Type      string    `json:"type" gorm:"size:30"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
