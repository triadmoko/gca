package model

import "time"

type User struct {
	ID        uint32
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Phone     uint64
	Address   string
	Role      int
	Image     string
	Password  string
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}
