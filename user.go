package main

import (
	"fmt"
)

type User struct {
	ID                     int    `json:"id"`
	lastName               string `json:"lastName"`
	firstName              string `json:"firstName"`
	identificationDocument string `json:"identificationDocument"`
	email                  string `json:"email"`
	passwd                 string `json:"passwd"`
}

func (u *User) showUserInfo() {
	fmt.Println(u.lastName+","+u.firstName+","+u.identificationDocument+","+u.email+","+u.passwd)
}