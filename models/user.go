package models

import (
	"database/sql"
	"gorm_example/database"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB, _ = database.Conn()
}

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	// DeletedAt    time.Time `gorm:"index"`
}

func (user *User) AddUser() {
	DB.Create(user)
}
