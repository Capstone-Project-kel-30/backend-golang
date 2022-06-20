package member

import (
	"gym-app/business/member"
	"gym-app/business/member/entity"

	"gorm.io/gorm"
)

type MemberPostgresRepo struct {
	db *gorm.DB
}

func NewMemberPostgresRepo(db *gorm.DB) member.MemberRepository {
	return &MemberPostgresRepo{
		db: db,
	}
}

func (c *MemberPostgresRepo) InsertMember(member entity.Member) (entity.Member, error) {
	c.db.Save(&member)
	return member, nil
}

func (c *MemberPostgresRepo) FindByMemberID(memberID string) (entity.Member, error) {
	var member entity.Member
	res := c.db.Where("member_id = ?", memberID).Take(&member)
	if res.Error != nil {
		return member, res.Error
	}
	return member, nil
}
