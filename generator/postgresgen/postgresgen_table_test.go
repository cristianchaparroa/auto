package postgresgen

import (
	"fmt"
	"testing"

	"github.com/cristianchaparroa/auto/meta"
	"github.com/cristianchaparroa/auto/parser"
)

func TestNewPostgresTable(t *testing.T) {
	pt := NewPostgresTable()
	if pt == nil {
		t.Error("Expected a pointer to PostgresTable but get nil")
	}
}

func TestPostgresTableGenerate(t *testing.T) {
	pt := NewPostgresTable()

	m := &meta.ModelStruct{ModelName: "user"}

	sql, err := pt.Generate(m)

	if err != nil {
		t.Error(err)
	}

	sqlExpected := `CREATE TABLE user ();`

	if sql != sqlExpected {
		t.Errorf("Expected the following sql:%v but get:%v", sqlExpected, sql)
	}
}

func TestPostgresTableGenerateWithFields(t *testing.T) {
	pt := NewPostgresTable()

	m := &meta.ModelStruct{ModelName: "user"}

	typeString := &meta.TypeField{Name: parser.TypeFieldString}
	fName := &meta.Field{Name: "name", Typ: typeString}
	fEmail := &meta.Field{Name: "email", Typ: typeString}

	fs := make([]*meta.Field, 0)
	fs = append(fs, fName)
	fs = append(fs, fEmail)

	m.Fields = fs

	sql, err := pt.Generate(m)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(sql)
}
