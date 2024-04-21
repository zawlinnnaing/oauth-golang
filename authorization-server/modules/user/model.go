package user

import "time"

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null;default:null" json:"email"`
	FirstName string    `gorm:"not null;default:null" json:"first_name"`
	LastName  string    `gorm:"not null;default:null" json:"last_name"`
	Password  string    `gorm:"not null;default:null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
