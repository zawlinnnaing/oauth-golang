package client_app

import (
	"time"

	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/user"
)

type ClientApp struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	RedirectURI string    `gorm:"not null" json:"redirect_uri"`
	UserID      string    `gorm:"not null" json:"user_id"`
	User        user.User `json:"user"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
