package models

import (
    "time"
)

type Message struct {
    ID         uint      `gorm:"primaryKey"`
    SenderID   uint      `gorm:"not null"`
    ReceiverID uint      `gorm:"not null"`
    Content    string    `gorm:"type:text;not null"`
    IsRead     bool      `gorm:"default:false"`
    CreatedAt  time.Time `gorm:"autoCreateTime"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
