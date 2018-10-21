package postgresgen

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
	fmt.Println(sql)
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
	fmt.Println(sql)

}
