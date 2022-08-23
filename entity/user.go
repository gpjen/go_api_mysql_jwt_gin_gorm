package entity

import "time"

type User struct {
	ID        uint64    `gorm:"primary_ket;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	Password  string    `gorm:"->;<-;not null;not null" json:"-"`
	Active    bool      `gorm:"default:true" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
