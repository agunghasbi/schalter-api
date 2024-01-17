package models

import (
	"github.com/dgrijalva/jwt-go"
	u "github.com/agunghasbi/schalter-api/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"os"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	BirthDate string `json:"birth_date"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is not valid"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password minimum 6 characters."), false
	}

	// Email must be unique
	temp := &User{}

	// Check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please try again"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email already in use by another user"), false
	}

	return u.Message(true, "Requirement passed"), true
}

func (user *User) Create() (map[string]interface{}) {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error")
	}

	// TODO Create new JWT Token for the newly registered user
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = ""  // Delete Password

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}

func Login(email string, password string) (map[string]interface{}) {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error;
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "User not found")
		}
		return u.Message(false, "Connection error. Please Retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	user.Password = "" // Delete Password
	
	// TODO Create JWT token
	tk := &Token{UserId:user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp



}