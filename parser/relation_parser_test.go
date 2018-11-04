package parser

import (
	"testing"

	"github.com/cristianchaparroa/auto/meta"
)

func TestTagRelationParserParse(t *testing.T) {
	var test = []struct {
		RelationLine string
		FieldName    string
		ExpectedType meta.RelationType
	}{
		{"(type:11; to:PostDetail)", "Post", "11"},
		{"(type:1*; to:PostDetail)", "Post", "1*"},
		{"(type:*1; to:PostDetail)", "Post", "*1"},
		{"(type:**; to:PostDetail)", "Post", "**"},
	}

	rp := NewTagRelationParser()

	for _, tc := range test {
		rel, err := rp.Parse(tc.RelationLine, tc.FieldName)

		if err != nil {
			t.Error(err)
		}

		if rel.Typ != tc.ExpectedType {
			t.Errorf("Expected the following typ:%v but get:%v", tc.ExpectedType, rel.Typ)
		}
	}
}
