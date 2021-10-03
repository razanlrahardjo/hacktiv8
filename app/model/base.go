package model

import (
	"gorm.io/gorm"
	"time"
)

// Base model
type Base struct {
	ID        int           `json:"id,omitempty" gorm:"primaryKey;unique;auto_increment"`
	CreatedAt time.Time     `json:"created_at,omitempty" gorm:"type:timestamp" format:"date-time" swaggertype:"string"`
	UpdatedAt time.Time     `json:"updated_at,omitempty" gorm:"type:timestamp" format:"date-time" swaggertype:"string"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}

// BeforeCreate Data
func (b *Base) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

// BeforeUpdate Data
func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now()
	b.UpdatedAt = now
	return nil
}
