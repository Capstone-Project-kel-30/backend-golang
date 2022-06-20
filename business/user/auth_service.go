package user

import (
	"errors"
	"fmt"
	"log"

	// "gym-app/config"

	// "github.com/mailjet/mailjet-apiv3-go"
	"github.com/pquerna/otp/totp"
	"github.com/xlzd/gotp"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) error
	GenerateTOTP(email string) string
	SendOTPtoEmail(otp string, name string, email string) error
	SendEmailVerification(otp string, email string) error
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

	// config := config.GetConfig()

	// publicKey := config.Mailjet.PublicKey
	// privateKey := config.Mailjet.PrivateKey

	// mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)

	// htmlpart := fmt.Sprintf(`<H3>Hai %s, Terimakasih sudah daftarin dirimu di GYM30. Dibawah ini adalah kode verifikasi kamu<H3/> <h1>%s</h1>`, name, otp)
	// messagesInfo := []mailjet.InfoMessagesV31{
	// 	{
	// 		From: &mailjet.RecipientV31{
	// 			Email: "bens.sky69@gmail.com",
	// 			Name:  "GYM30",
	// 		},
	// 		To: &mailjet.RecipientsV31{
	// 			mailjet.RecipientV31{
	// 				Email: email,
	// 				Name:  name,
	// 			},
	// 		},
	// 		Subject:  "OTP Key GYM30",
	// 		TextPart: otp,
	// 		HTMLPart: htmlpart,
	// 		CustomID: "AppGettingStartedTest",
	// 	},
	// }
	// messages := mailjet.MessagesV31{Info: messagesInfo}
	// res, err := mailjetClient.SendMailV31(&messages)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// log.Println(res)
	return nil
}

func (c *authService) SendEmailVerification(otp string, email string) error {
	// config := config.GetConfig()

	// publicKey := config.Mailjet.PublicKey
	// privateKey := config.Mailjet.PrivateKey

	// mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)
	// htmlpart := fmt.Sprintf(`<H3> Dibawah ini adalah kode verifikasi Forgot Password kamu<H3/><h1>%s</h1>`, otp)
	// messagesInfo := []mailjet.InfoMessagesV31{
	// 	{
	// 		From: &mailjet.RecipientV31{
	// 			Email: "bens.sky69@gmail.com",
	// 			Name:  "GYM30",
	// 		},
	// 		To: &mailjet.RecipientsV31{
	// 			mailjet.RecipientV31{
	// 				Email: email,
	// 				Name:  email,
	// 			},
	// 		},
	// 		Subject:  "Forgot Password OTP Key GYM30",
	// 		TextPart: otp,
	// 		HTMLPart: htmlpart,
	// 		CustomID: "AppGettingStartedTest",
	// 	},
	// }
	// messages := mailjet.MessagesV31{Info: messagesInfo}
	// res, err := mailjetClient.SendMailV31(&messages)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// log.Println(res)
	return nil
}
