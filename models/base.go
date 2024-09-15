package models

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
