package generator

import (
	"testing"

	"github.com/cristianchaparroa/auto/meta"
	"github.com/cristianchaparroa/auto/parser"
)

func TestNewTypeColumnBuilder(t *testing.T) {
	b := NewTypeColumnBuilder()
	if b == nil {
		t.Error("Expected a pointer to NewTypeColumnBuilder but get nil")
	}
}

func TestGetTypeOK(t *testing.T) {

	b := NewTypeColumnBuilder()

	var test = []struct {
		Model        *meta.Field
		TypeExpected TypeColumn
	}{
		{&meta.Field{Typ: &meta.TypeField{Name: parser.TypeFieldString}}, TypePosgresVarchar},
		{&meta.Field{Typ: &meta.TypeField{Name: parser.TypeFieldInteger}}, TypePostgresInt},
		{&meta.Field{Typ: &meta.TypeField{Name: parser.TypeFieldBool}}, TypePostgresBoolean},
		{&meta.Field{Typ: &meta.TypeField{Name: parser.TypeFieldBigFloat}}, TypePostgresBigInt},
		{&meta.Field{Typ: &meta.TypeField{Name: parser.TypeFieldArray}}, TypePostgresRelation},
	}

	for _, tc := range test {
		tp, err := b.GetType(tc.Model)
		if err != nil {
			t.Error(err)
		}

		if tp != tc.TypeExpected {
			t.Errorf("Expected the type: %v but get :%v", tc.TypeExpected, tp)
		}
	}
}

func TestGetTypeNOK(t *testing.T) {
	b := NewTypeColumnBuilder()

	var test = []struct {
		Model         *meta.Field
		ErrorExpected string
	}{
		{nil, TypeColumnFieldNil},
		{&meta.Field{Typ: &meta.TypeField{Name: "TypeNotSupported"}}, TypeColumnNotSupported},
	}

	for _, tc := range test {
		_, err := b.GetType(tc.Model)

		if err.Error() != tc.ErrorExpected {
			t.Errorf("Expected the error:%v but get :%v", tc.ErrorExpected, err.Error())
		}
	}
}
