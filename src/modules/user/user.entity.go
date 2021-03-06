package user

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username     string    `gorm:"unique;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string `gorm:"unique;not null"`
	PasswordHash string
	UserStatus   int
	UserRole     string
}