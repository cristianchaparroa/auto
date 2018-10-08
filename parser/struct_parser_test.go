package parser

import (
	"testing"
)

func TestNewModelStructParser(t *testing.T) {
	p := NewModelStructParser()

	if p == nil {
		t.Errorf("Expected a pointer to ModelStructParser but get a nil")
	}
}

func TestStructParserParseFile(t *testing.T) {

	fileTest := `package models

type AuthorTest struct {
  	   Id        string
  	   Name      string
  	   Books     []BookTest
  	   Editorial EditorialTest
}

// EditorialTest is another model
type EditorialTest struct {
    	 Id     string
    	 Name   string
    	 Autors []AuthorTest
}`
	p := NewModelStructParser()
	structs, err := p.ParseFile(fileTest)

	if err != nil {
		t.Error(err)
	}

	expectedStructs := 2
	resultStructs := len(structs)

	if resultStructs != expectedStructs {
		t.Errorf("Expected 2 structs but get:%v", resultStructs)
	}

}

func TestStructParserParse(t *testing.T) {

	structStr := `type EditorialTest struct {
      	 Id     string
      	 Name   string
      	 Autors []AuthorTest
  }`

	p := NewModelStructParser()
	m, err := p.Parse(structStr)

	if err != nil {
		t.Errorf("It generates an error to parse the string: %v", err)
	}

	if m == nil {
		t.Error("Expected a model but get a nil")
	}

}

func TestStructParserGetModels(t *testing.T) {
	fileTest := `package models

type AuthorTest struct {
			 Id        string
			 Name      string
			 Books     []BookTest
			 Editorial EditorialTest
}

// EditorialTest is another model
type EditorialTest struct {
			 Id     string
			 Name   string
			 Autors []AuthorTest
}`
	p := NewModelStructParser()
	ms, err := p.GetModels(fileTest)

	if err != nil {
		t.Error(err)
	}

	expectedSize := 2
	size := len(ms)

	if expectedSize != size {
		t.Errorf("Expected 2 model structs but get:%v", size)
	}
}

func TestStructParserGetModelsNOK(t *testing.T) {
	fileTest := `
	type AuthorTest struct {
			 Id        string
			 Name      string
			 Books     []BookTest
			 Editorial EditorialTest
	}

	// EditorialTest is another model
	type EditorialTest struct {
			 Id     string
			 Name   string
			 Autors []AuthorTest
	}`

	p := NewModelStructParser()
	_, err := p.GetModels(fileTest)

	if err == nil {
		t.Error("Expected an error but get nil")
	}

	if err.Error() != PackageParserEmptyLine {
		t.Errorf("Expected the error :%v , but get :%v", PackageParserEmptyLine, err.Error())
	}

}
