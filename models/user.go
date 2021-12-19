package models

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) Create(hash string) error {
	cmd := fmt.Sprintf("INSERT INTO %s (name, email, password_hash, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", tableNameUsers)
	_, err := DbConnection.Exec(cmd, u.Name, u.Email, hash, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}
	return err
}

func (u *User) GenerateHash() string {
	b := []byte(u.Password)
	hashed, err := bcrypt.GenerateFromPassword(b, 12)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed)
}
