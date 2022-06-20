package entity

import (
	"gym-app/business/user/entity"

	"gorm.io/gorm"
)

type Member struct {
	ID          int    `gorm:"primary_key:auto_increment" json:"-"`
	Status      string `gorm:"type:varchar(100)" json:"-"`
	Tier        string `gorm:"type:varchar(100)" json:"-"`
	ExpiredTire int    `gorm:"type:varchar(100)" json:"-"`
	UserID      int
	User        entity.User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	*gorm.Model
}
