package models

import "github.com/google/uuid"

type User struct {
	Id       int
	PublicId uuid.UUID
	Login    string
	Password string
}
