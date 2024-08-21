package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Sub string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:now()"`
}