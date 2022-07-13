package resp

import (
	booking "github.com/mashbens/cps/business/booking/entity"
)

type BookResp struct {
	UserID     int             `json:"user_id"`
	ClassSlice []booking.Class `json:"class"`
}

func FromService(c booking.Booking) BookResp {
	return BookResp{
		UserID:     int(c.UserID),
		ClassSlice: c.ClassSlice,
	}
}
