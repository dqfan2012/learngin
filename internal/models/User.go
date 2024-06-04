package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"primaryKey"`
	FirstName       string         `gorm:"size:255;not null"`
	LastName        string         `gorm:"size:255;not null"`
	Email           string         `gorm:"size:255;unique;not null"`
	Password        string         `gorm:"size:255;not null"`
	Role            string         `gorm:"type:VARCHAR(255);default:'reader';check:role IN ('reader', 'publisher')"`
	RememberToken   string         `gorm:"size:100"`
	CreatedAt       time.Time      `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
	EmailVerifiedAt *time.Time     `gorm:"type:timestamptz"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate sets the timestamps to UTC before inserting the record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = now
	return
}

// BeforeUpdate sets the UpdatedAt timestamp to UTC before updating the record
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return
}
