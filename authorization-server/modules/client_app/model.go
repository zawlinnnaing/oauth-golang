package client_app

import "time"

type ClientApp struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	RedirectURI *string   `json:"redirect_uri"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
