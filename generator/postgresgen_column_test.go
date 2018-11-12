package generator

import (
	"fmt"
	"testing"

	"github.com/cristianchaparroa/auto/meta"
	"github.com/cristianchaparroa/auto/parser"
)

func TestPostgresgenColumnCreate(t *testing.T) {

	g := NewPostgresColumn()

	m := &meta.Field{Name: "Name", Typ: &meta.TypeField{Name: parser.TypeFieldString}}

	sql, err := g.Create("user", m)

	if err != nil {
		t.Error(err)
	}

	sqlExpected := `ALTER TABLE USER ADD COLUMN NAME VARCHAR`

	if sqlExpected != sql {
		t.Errorf("Expected %v but get:%v", sqlExpected, sql)
	}
}

func TestPosgresgenColumnCreateNOK(t *testing.T) {
	g := NewPostgresColumn()

	_, err := g.Create("user", nil)

	if err == nil {
		t.Error("Expected an error but get nil")
	}

}

func TestPostgresgenColumnChangeType(t *testing.T) {
	g := NewPostgresColumn()

	m := &meta.Field{Name: "Name", Typ: &meta.TypeField{Name: parser.TypeFieldInteger}}
	sql, err := g.ChangeType("user", m)

	if err != nil {
		t.Error(err)
	}

	sqlExpected := `ALTER TABLE USER ALTER COLUMN NAME TYPE INT`

	if sqlExpected != sql {
		t.Errorf("Expected %v but get:%v", sqlExpected, sql)
	}
}

func TestPostgresgenColumnChangeTypeNOK(t *testing.T) {
	g := NewPostgresColumn()

	_, err := g.ChangeType("user", nil)

	if err == nil {
		t.Error("Expected an error but get nil")
	}
}

func TestPostgresTableGenerateAttributes(t *testing.T) {
	pkField := &meta.Field{Name: "Title", Typ: &meta.TypeField{Name: parser.TypeFieldString}}
	pkField.IsPrimaryKey = true

	uniqueField := &meta.Field{Name: "Title", Typ: &meta.TypeField{Name: parser.TypeFieldString}}
	uniqueField.IsUnique = true

	notNullField := &meta.Field{Name: "Title", Typ: &meta.TypeField{Name: parser.TypeFieldString}}
	notNullField.IsNotNull = true

	sqlExpected := `ALTER TABLE USER ADD COLUMN TITLE VARCHAR`

	var test = []struct {
		Field       *meta.Field
		SQLExpected string
	}{
		{pkField, fmt.Sprintf("%s PRIMARY KEY", sqlExpected)},
		{notNullField, fmt.Sprintf("%s NOT NULL", sqlExpected)},
		{uniqueField, fmt.Sprintf("%s UNIQUE", sqlExpected)},
	}

	g := NewPostgresColumn()
	for _, tc := range test {
		sqlResult, err := g.Create("USER", tc.Field)

		if err != nil {
			t.Error(err)
		}
		if sqlResult != tc.SQLExpected {
			t.Errorf("Expected:%s, but get:%s", tc.SQLExpected, sqlResult)
		}
	}
}
