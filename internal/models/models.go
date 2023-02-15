package models

import (
	"fmt"
	"time"
)

type City struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Carriers struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type User struct {
	Id       int    `db:"id" json:"id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

type MainRoute struct {
	Id        int
	BusId     int
	CarrierId int
	Name      string
}

type MidRoute struct {
	Id       int
	FromCity int
	ToCity   int
	FromTime time.Time
	ToTime   time.Time
	Price    float32
	Rating   float32
}

type SearchRoute struct {
	MD          MidRoute
	CarrierName string
}

type SerchParams struct {
	FromCity int
	ToCity   int
	Date     time.Time
}

type Token struct {
	Token string `json:"token"`
}

func (u *User) Validate() error {
	switch {
	case u.Login == "":
		return fmt.Errorf("Login must be field")

	case u.Password == "":
		return fmt.Errorf("Password must be field")

	case len(u.Password) < 8:
		return fmt.Errorf("Password must be at least 8 characters")
	}
	return nil
}
