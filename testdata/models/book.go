package testmodel

import "time"

type BookTest struct {
	Id    string `sql:"pk"`
	Title string `sql:"nullable=false,name=titulo"`
	Date  time.Time
}
