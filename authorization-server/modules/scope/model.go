package scope

import (
	"time"
)

type Scope struct {
	Name        string    `gorm:"primaryKey" json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
