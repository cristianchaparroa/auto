package generator

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

	m := &meta.ModelStruct{ModelName: "users"}

	result, err := pt.Generate(m)

	if err != nil {
		t.Error(err)
	}

	sqlExpected := `CREATE TABLE USERS ();`

	if result.GetSQLResult() != sqlExpected {
		t.Errorf("Expected the following sql:%v but get:%v", sqlExpected, result.GetSQLResult())
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

func TestPostgresTableCreateOneToOne(t *testing.T) {
	m := &meta.ModelStruct{}
	m.ModelName = "PostDetail"

	typeInt := &meta.TypeField{Name: parser.TypeFieldInteger}

	pk := &meta.Field{Name: "Id", Typ: typeInt}
	fs := make([]*meta.Field, 0)
	fs = append(fs, pk)

	f := &meta.Field{Name: "Post"}
	f.IsRelation = true

	r := &meta.Relation{}
	r.To = "Post"
	r.Name = "post_id"
	r.PKRef = "Id" // This the Post.Id
	r.FK = "fk_post_id_post_detail_id"
	f.Relation = r

	fs = append(fs, f)

	m.Fields = fs

	pg := NewPostgresTable()
	sql, err := pg.CreateOneToOne(m, f)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(sql)
}
