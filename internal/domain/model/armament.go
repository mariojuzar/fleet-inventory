package model

type Armament struct {
	BaseModel
	Id   int    `db:"id"`
	Name string `db:"name"`
}
