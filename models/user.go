package models

import (
	u "bitslibrary/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	Id        uint      `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Address   string    `json:"address"`
	Password  string    `json:"password"`
	Token     string    `json:"token";sql:"-"`
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	temp := &User{}

	err := GetDB().Table("users").Where("email=?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please try."), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.Id <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: user.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = ""

	response := u.Message(true, "Account has been created")
	response["user"] = user
	return response
}

func Login(email, password string) map[string]interface{} {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again.")
	}

	user.Password = ""

	tk := &Token{UserId: user.Id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

func GetAllUser() []*User {
	user := make([]*User, 0)
	err := GetDB().Table("users").Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func GetUser(id string) *User {
	user := &User{}
	err := GetDB().Table("users").Where("id=?", id).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func (user *User) Update(id string) map[string]interface{} {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Table("users").Where("id=?", id).Update(user)
	t := time.Now()
	GetDB().Table("users").Where("id=?", id).Update("updated_at", t)

	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}
