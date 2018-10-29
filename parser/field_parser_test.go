package parser

import (
	"testing"
)

func TestFilterComments(t *testing.T) {
	test := `//This is a comment inside of definition of field
Editorial EditorialTest` + " `sql:\"oneToMany\"`"

	fp := NewModelFieldParser()

	result := fp.FilterComments(test)

	expected := "Editorial EditorialTest `sql:\"oneToMany\"`"

	if result != expected {
		t.Errorf("Expected:%v but get:%v", expected, result)
	}
}

func TestFieldParse(t *testing.T) {

	var tests = []struct {
		Line      string // line to test
		ShouldNil bool
	}{
		{"", true},
		{"Id string", false},
		{"Name      string        `sql:\"len=50\"`", false},
		{"Books     []BookTest    `sql:\"manyToMany\"` //Comment in line", false},
		{`//This is a comment inside of definition of field
Editorial EditorialTest` + " `sql:\"oneToMany\"`", false},
	}

	fp := NewModelFieldParser()

	for _, tc := range tests {
		mf, err := fp.Parse(tc.Line)

		if err != nil {
			t.Error(err)
		}

		if mf == nil && !tc.ShouldNil {
			t.Errorf("The field( %v ) Line parsed is null", tc.Line)
		}
	}
}

func Test_ParseFields(t *testing.T) {
	fp := NewModelFieldParser()

	bodyContent := "Id        string        `sql:\"pk\"` \n" +
		"Name      string        `sql:\"len=50\"` \n" +
		"Books  []BookTest        `sql:\"manyToMany\"` //another comment \n"
	fs := fp.ParseFields(bodyContent)

	rs := len(fs)

	if rs != 3 {
		// fmt.Println(fs)
		t.Errorf("Expected 3 but get:%v", rs)
	}
}
