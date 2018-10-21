package testmodel

import "time"

type BookTest struct {
	Id    string `sql:"pk"`
	Title string `sql:"notnull=true," `
	Date  time.Time
}
