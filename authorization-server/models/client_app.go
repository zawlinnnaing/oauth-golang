package models

import "time"

type ClientApp struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	RedirectURI string    `gorm:"not null" json:"redirect_uri"`
	UserID      string    `gorm:"not null" json:"user_id"`
	User        User      `json:"user"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
