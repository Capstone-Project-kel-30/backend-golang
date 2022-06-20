package dto

type RegisterMemberRequest struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	Tier        string `json:"tier"`
	ExpiredTire int    `json:"expired_tire"`
	UserID      int    `json:"user_id"`
}
