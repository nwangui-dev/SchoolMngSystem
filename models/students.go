package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FirstName string `gorm:"type:varchar(100);not null" json:"first_name" validate:"required"`
	LastName  string `gorm:"type:varchar(100);not null" json:"last_name" validate:"required"`
	Email string `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
	Grade string `gorm:"type:varchar(10);not null" json:"grade" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	
}
