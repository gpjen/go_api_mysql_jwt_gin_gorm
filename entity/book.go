package entity

import "time"

type Book struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title"`
	Description string    `gorm:"text" json:"description"`
	UserId      uint64    `gorm:"not null" json:"-"`
	User        User      `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}
