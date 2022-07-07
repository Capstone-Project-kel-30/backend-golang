package booking

import (
	"errors"
	"strconv"

	"github.com/mashbens/cps/business/booking/entity"
	classService "github.com/mashbens/cps/business/class"
	userService "github.com/mashbens/cps/business/user"
)

type BookingRepo interface {
	InsertBooking(booking entity.Booking) (entity.Booking, error)
	GetSchedule(userID int) (entity.Booking, error)
	// FindBookingByID(bookingID int) (entity.Booking, error)
}

type BookingService interface {
	InsertBooking(booking entity.Booking) (*entity.Booking, error)
	GetSchedule(userID int) (*entity.Booking, error)
	// FindBookingByID(bookingID int) (*entity.Booking, error)
}

type bookingService struct {
	bookingRepo  BookingRepo
	userService  userService.UserService
	classService classService.ClassService
}

func NewBookingService(
	bookingRepo BookingRepo,
	userService userService.UserService,
	classService classService.ClassService,
) BookingService {
	return &bookingService{
		bookingRepo:  bookingRepo,
		userService:  userService,
		classService: classService,
	}
}

func (c *bookingService) InsertBooking(booking entity.Booking) (*entity.Booking, error) {
	// find user
	user, err := c.userService.FindUserByID(strconv.Itoa(booking.UserID))
	if err != nil {
		return nil, errors.New("User not found")
	}
	_ = user

	class, err := c.classService.FindClassByID(strconv.Itoa(booking.ClassID))
	if err != nil {
		return nil, errors.New("Class not found")
	}
	_ = class

	var tempStatus entity.Booking
	if class.UserBooked == class.Capacity-1 {
		status := "Full Booked"
		tempStatus.Class.Status = status
	} else {
		status := "Available"
		tempStatus.Class.Status = status
	}

	booking.Class.Status = tempStatus.Class.Status

	err = c.classService.UpdateClassStatus(strconv.Itoa(class.ID), booking.Class.Status)

	upt := c.classService.UpdateUserBooked(strconv.Itoa(booking.ClassID))
	_ = upt

	book, err := c.bookingRepo.InsertBooking(booking)
	if err != nil {
		return nil, err
	}

	book.Class.UserBooked++
	book.Class = ClassToBoo(*class)
	book.User = UserToBoo(*user)

	return &book, nil

}

func (c *bookingService) GetSchedule(userID int) (*entity.Booking, error) {
	booking, err := c.bookingRepo.GetSchedule(userID)
	if err != nil {
		return nil, err
	}

	user, err := c.userService.FindUserByID(strconv.Itoa(booking.UserID))
	if err != nil {
		return nil, err
	}

	class, err := c.classService.FindClassByID(strconv.Itoa(booking.ClassID))
	if err != nil {
		return nil, err
	}

	booking.User = UserToBoo(*user)
	booking.Class = ClassToBoo(*class)

	return &booking, nil
}
