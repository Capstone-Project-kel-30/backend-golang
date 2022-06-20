package member

import (
	"gym-app/business/member/dto"
	"gym-app/business/member/entity"
	_member "gym-app/business/member/response"

	"github.com/mashingan/smapping"
)

type MemberRepository interface {
	InsertMember(member entity.Member) (entity.Member, error)
	FindByMemberID(memberID string) (entity.Member, error)
}

type MemberService interface {
	CreateMember(registerRequest dto.RegisterMemberRequest) (*_member.MemberResponse, error)
	FindMemberByMemberID(memberID string) (*_member.MemberResponse, error)
}

type memberService struct {
	memberRepo MemberRepository
}

func NewMemberService(membershipRepo MemberRepository) MemberService {
	return &memberService{
		memberRepo: membershipRepo,
	}
}

func (c *memberService) FindMemberByMemberID(memberID string) (*_member.MemberResponse, error) {
	membership, err := c.memberRepo.FindByMemberID(memberID)
	if err != nil {
		return nil, err
	}

	res := _member.NewMemberResponse(membership)
	return &res, nil
}

func (c *memberService) CreateMember(registerRequest dto.RegisterMemberRequest) (*_member.MemberResponse, error) {
	member := entity.Member{}
	err := smapping.FillStruct(&member, smapping.MapFields(&registerRequest))

	if err != nil {
		return nil, err
	}

	member, err = c.memberRepo.InsertMember(member)
	if err != nil {
		return nil, err
	}

	res := _member.NewMemberResponse(member)
	return &res, nil
}
