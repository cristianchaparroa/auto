package testmodel

type AuthorTest struct {
	Id        string        `sql:"pk"`
	Name      string        `sql:"len=50"`
	Books     []BookTest    `sql:"manyToMany"`
	Editorial EditorialTest `sql:"oneToMany"`
}
