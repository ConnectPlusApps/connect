package models

type Preference struct {
    ID               uint   `gorm:"primaryKey"`
    UserID           uint   `gorm:"uniqueIndex;not null"`
    MatchDistance    int    `gorm:"default:50"` // in kilometers
    MinAge           int    `gorm:"default:18"`
    MaxAge           int    `gorm:"default:99"`
    NotifyNewMatches bool   `gorm:"default:true"`
    NotifyMessages   bool   `gorm:"default:true"`
    ShowOnlineStatus bool   `gorm:"default:true"`
    ShowLastActive   bool   `gorm:"default:true"`
    ShowDistance     bool   `gorm:"default:true"`
}
