package member

import (
	"github.com/mashbens/cps/business/member/entity"
)

type Membership struct {
	ID            int    `gorm:"primary_key:auto_increment" json:"-"`
	Type          string `gorm:"type:varchar(100)" `
	Price         int
	Super_adminID int
	Super_admin   entity.SuperAdmin `gorm:"foreignkey:Super_adminID" json:"-"`
	Duration      int
	Description   string `gorm:"type:varchar(100)" json:"-"`
}

func (m *Membership) toService() entity.Membership {
	return entity.Membership{
		ID:            int(m.ID),
		Type:          m.Type,
		Price:         m.Price,
		Duration:      m.Duration,
		Super_adminID: m.Super_adminID,
		Super_admin:   m.Super_admin,
		Description:   m.Description,
	}
}

func fromService(member entity.Membership) Membership {
	return Membership{
		ID:            int(member.ID),
		Type:          member.Type,
		Price:         member.Price,
		Duration:      member.Duration,
		Super_adminID: member.Super_adminID,
		Super_admin:   member.Super_admin,
		Description:   member.Description,
	}
}

func toServiceList(data []Membership) []entity.Membership {
	a := []entity.Membership{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
