package users

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	HashKey   string `json:"hash_key"`
	StudentId string `json:"student_id"`
	Email     string `json:"email"`
}

func (user *User) setId() error {
	if user.HashKey != "" {
		return errors.New("already has hash key")
	}
	user.HashKey = uuid.New().String()
	return nil
}

func (user User) getSchoolEmail() string {
	return fmt.Sprintf("%s@mail.hongik.ac.kr", user.StudentId)
}
