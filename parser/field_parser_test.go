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
		"Books  []BookTest        `sql:\"manyToMany\"` //another comment \n" +
		"// AnotherField string"
	fs := fp.ParseFields(bodyContent)

	rs := len(fs)

	if rs != 3 {
		t.Errorf("Expected 3 but get:%v", rs)
	}
}

func TestParseRelationFields(t *testing.T) {

	var test = []struct {
		FieldStr   string
		IsRelation bool
	}{
		{"Name      string", false},
		{"PostDetail PostDetail `sql:\"rel=(type:11; to:PostDetail)\"`", true},
		{"PostDetail PostDetail `sql:\"rel=(type:11)\"`", true},
		{"PostDetail PostDetail `sql:\"rel=(type:11; to:PostDetail)\", json:\"post_detail\"`", true},
		{"Post Post `sql:\"rel=(type:11; mappedBy:Post; name:post_id)\"`", true},
		{"Comments []Comments `sql:\"rel=(type:1*;to:Comments)\"`", true},
		{"Post Post `sql:\"rel=(type:*1;to:Post; name:post_id)\"`", true},
		{"Post Post `sql:\"rel=(type:*1; name:post_id)\"`", true},
		{"Tags []Tag `sql:\"rel=(type:**; to:Tag)\"`", true},
		{"Tags []Tag `sql:\"rel=(type:**)\"`", true},
		{"Posts []Post `sql:\"rel=(type:**; to:Post)\"`", true},
	}

	fp := NewModelFieldParser()

	for _, tc := range test {

		fs, err := fp.Parse(tc.FieldStr)

		if err != nil {
			t.Error(err)
		}

		if fs.IsRelation != tc.IsRelation {
			t.Errorf("Expected relation equals to:%v, but get:%v", tc.IsRelation, fs.IsRelation)
		}
	}

}
