package postgresgen

import (
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

	m := &meta.ModelStruct{ModelName: "users"}

	sql, err := pt.Generate(m)

	if err != nil {
		t.Error(err)
	}

	sqlExpected := `CREATE TABLE USERS ();`

	if sql != sqlExpected {
		t.Errorf("Expected the following sql:%v but get:%v", sqlExpected, sql)
	}
}

func TestPostgresTableGenerateWithFields(t *testing.T) {
	pt := NewPostgresTable()

	m := &meta.ModelStruct{ModelName: "users"}

	typeString := &meta.TypeField{Name: parser.TypeFieldString}
	typeInt := &meta.TypeField{Name: parser.TypeFieldInteger}

	fName := &meta.Field{Name: "name", Typ: typeString}
	fEmail := &meta.Field{Name: "email", Typ: typeString}
	fCL := &meta.Field{Name: "likes", Typ: typeInt}
	fAnother := &meta.Field{Name: "camelCaseName", Typ: typeInt}

	fs := make([]*meta.Field, 0)
	fs = append(fs, fName)
	fs = append(fs, fEmail)
	fs = append(fs, fCL)
	fs = append(fs, fAnother)

	m.Fields = fs

	_, err := pt.Generate(m)

	if err != nil {
		t.Error(err)
	}
	//fmt.Println(sql)
}
