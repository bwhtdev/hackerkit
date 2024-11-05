package types

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
  Name        string    `json:"name"`
  Avatar      string    `json:"avatar"`
  Description string    `json:"description"`
	Token       string    `json:"token"`
}

type CreateUserRequest struct {
	Username  string `json:"username"`
  Name      string `json:"name"`
  Avatar    string `json:"avatar"`
	Password  string `json:"password"`
}

type UpdateUserRequest struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
  Name        string    `json:"name"`
  Avatar      string    `json:"avatar"`
  Description string    `json:"description"`
}

type User struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
  Name              string    `json:"name"`
  Avatar            string    `json:"avatar"`
  Description       string    `json:"description"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewUser(username, name, avatar, password string) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:          username,
    Name:              name,
    Avatar:            avatar,
    Description:       "",
    EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
