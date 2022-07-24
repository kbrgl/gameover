package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `json:"username"`
	Email          string `json:"email"`
	PasswordDigest string `json:"password_digest"`
}
