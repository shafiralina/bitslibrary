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
	Id        uint      `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Address   string    `json:"address"`
	Password  string    `json:"password"`
	Token     string    `json:"token";sql:"-"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) Validate() (map[string]interface{}, bool) {
	user.Email = strings.ToLower(user.Email)

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

	temp2 := &User{}
	err2 := GetDB().Table("users").Where("mobile=?", user.Mobile).First(temp2).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please try."), false
	}
	if temp2.Mobile != "" {
		return u.Message(false, "Mobile already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (user *User) Create(conn *gorm.DB) map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	user.Email = strings.ToLower(user.Email)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	conn.Create(user)

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

func Login(conn *gorm.DB, mobile, email, password string) map[string]interface{} {
	user := &User{}

	if mobile == "" {
		email = strings.ToLower(email)
		err := conn.Table("users").Where("email = ?", email).First(user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return u.Message(false, "Email address not found")
			}
			return u.Message(false, "Connection error. Please retry.")
		}
	} else {
		err2 := conn.Table("users").Where("mobile = ?", mobile).First(user).Error
		if err2 != nil {
			if err2 == gorm.ErrRecordNotFound {
				return u.Message(false, "Mobile not found")
			}
			return u.Message(false, "Connection error. Please retry.")
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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

func GetAllUser(conn *gorm.DB) []*User {
	user := make([]*User, 0)
	err := conn.Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}

func GetUser(conn *gorm.DB, id string) *User {
	user := &User{}
	err := conn.Where("id=?", id).First(user).Error
	if err != nil {
		return nil
	}

	return user
}

func (user *User) Update(conn *gorm.DB, id string) map[string]interface{} {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	user1 := &User{}
	conn.Where("id=?", id).First(&user1)
	conn.Model(&user1).Updates(user)

	resp := u.Message(true, "success")
	resp["user"] = user

	return resp
}
