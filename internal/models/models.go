package models

type City struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Carriers struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
