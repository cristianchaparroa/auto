package testmodel

type EditorialTest struct {
	Id     string `sql:"pk"`
	Name   string
	Autors []AuthorTest `sql:"manyToOne"`
}

type Article struct {
	ISBN    string `sql:pk`
	Name    string
	Summary string `sql:len=300`
	Content string
}
