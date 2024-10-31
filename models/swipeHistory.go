// models/swipeHistory.go
package models

import "gorm.io/gorm"

type SwipeHistory struct {
gorm.Model
UserID    uint
TargetUserID uint
Liked        bool
}