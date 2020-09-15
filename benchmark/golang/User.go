package main

import (
	"github.com/watercraft/validator"
)

type User struct {
	Name     string `json:"name" validate:"nonzero"`
	Username string `json:"username" validate:"nonzero"`
}

func (s User) Validate() error {

	return validator.Validate(s)
}
