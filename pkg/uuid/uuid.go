package uuid

import uuid "github.com/satori/go.uuid"

func NewV4() string {
	return uuid.NewV4().String()
}

func Nil() string {
	return uuid.Nil.String()
}
