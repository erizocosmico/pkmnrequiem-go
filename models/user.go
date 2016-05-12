package models

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

type User struct {
	Badges   int    `json:"badges"`
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserStore struct {
	*mgo.Database
}

func (s UserStore) ByToken(token string) (*User, error) {
	return nil, nil
}

func NewUser(email, password, username string) *User {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := &User{
		Badges:   0,
		ID:       uuid.NewV4().String(),
		Email:    email,
		Password: string(encrypted),
		Username: username,
	}
	return user
}
