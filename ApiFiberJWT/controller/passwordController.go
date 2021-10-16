package controller

import (
	"math/rand"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"github.com/guifgr/go-learn/ApiFiberJWT/database"
	"github.com/guifgr/go-learn/ApiFiberJWT/models"
	"golang.org/x/crypto/bcrypt"
)

func Forgot(c *fiber.Ctx) error{
	data := ParseData(c)

	email := data["email"]
	token := RandStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: email,
		Token: token,
	}

	database.DB.Create(&passwordReset)

	from := "admin@example.com"

	to := []string{
		email,
	}

	url := "http://localhost:5000/reset/" + token

	message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")

	err := smtp.SendMail("localhost:1025", nil, from, to, message)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check your email box",
	})

}

func Reset(c *fiber.Ctx) error{
	data := ParseData(c)

	if data["password"] !=  data["passwordConfirm"]{
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Password not match",
		})
	}

	passwordReset := models.PasswordReset{}

	database.DB.Where("token = ?", data["token"]).Last(&passwordReset)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "sucess",
	})
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}