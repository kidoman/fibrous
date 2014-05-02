package main

import "fmt"

type User struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func UserKey(id int) string {
	return fmt.Sprintf("user:%d", id)
}

func (u *User) Key() string {
	return UserKey(u.ID)
}
