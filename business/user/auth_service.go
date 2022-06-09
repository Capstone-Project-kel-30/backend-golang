package user

import (
	"errors"
	"fmt"
	"log"

	"github.com/pquerna/otp/totp"
	"github.com/xlzd/gotp"

	// "gopkg.in/gomail.v2"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) error
	// CreateUser(user dto.) entity.User
	GenerateTOTP(email string) string
	SendOTPtoEmail(otp string, name string, email string) error
}

type authService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (c *authService) VerifyCredential(email string, password string) error {
	user, err := c.userRepo.FindByEmail(email)
	if err != nil {
		println(err.Error())
		return err
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}
	return nil
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (c *authService) GenerateTOTP(email string) string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "golang",
		AccountName: email,
		Period:      180,
		SecretSize:  20,
	})
	if err != nil {
		log.Println(err)
	}

	totp := gotp.NewDefaultTOTP(key.Secret()).Now()
	fmt.Println("TOTP:", totp)

	return totp
}

func (c *authService) SendOTPtoEmail(otp string, name string, email string) error {

	// Setup mail server
	// less secure apps needed but not working

	// mail := gomail.NewMessage()
	// mail.SetAddressHeader("From", "email", "GYM30")
	// mail.SetAddressHeader("To", email, name)
	// mail.SetHeader("Subject", "OTP")
	// mail.SetBody("text/plain", otp)

	// Send := gomail.NewDialer("smtp.gmail.com", 587, "email", "pw")

	// if err := Send.DialAndSend(mail); err != nil {
	// 	panic(err)
	// }
	return nil
}
