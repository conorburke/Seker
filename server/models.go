package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	firstName string
	lastName string
	email string
	street1 string
	street2 string
	city string
	state string
	zipcode string
	password string
}

type Tool struct {
	gorm.Model
	Owner User
	Name string
}