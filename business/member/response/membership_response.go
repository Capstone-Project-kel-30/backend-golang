package response

import "gym-app/business/member/entity"

type MemberResponse struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	Tier        string `json:"tier"`
	ExpiredTire int    `json:"expired_tire"`
	UserID      int    `json:"user_id"`
}

func NewMemberResponse(membership entity.Member) MemberResponse {
	return MemberResponse{
		ID:          membership.ID,
		Status:      membership.Status,
		Tier:        membership.Tier,
		ExpiredTire: membership.ExpiredTire,
		UserID:      membership.UserID,
	}
}
