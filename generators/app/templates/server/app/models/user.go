package models

import (
	"crypto/sha1" // if you need/want
	"gorm.io/gorm"
	"time"
)

type User struct { // example user fields
	gorm.Model
	Name              string
	IsAdmin           bool
	EncryptedPassword []byte
	Salt              []byte
	Password          string `sql:"-"`
}

func (u *User) CheckPassword(password string) bool {
	return string(u.encrypt(password, u.Salt)) == string(u.EncryptedPassword)
}

func (u *User) SetPassword(password string) {
	u.Salt = []byte(time.Now().String())
	u.EncryptedPassword = u.encrypt(password, u.Salt)
}

func (u *User) encrypt(password string, salt []byte) []byte {
	sha1 := sha1.New()
	sha1.Write([]byte(password))
	sha1.Write(salt)
	return sha1.Sum(nil)
}
