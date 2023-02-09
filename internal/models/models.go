package models

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
