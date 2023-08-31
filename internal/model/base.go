package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UUID      string         `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
