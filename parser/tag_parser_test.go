package parser

import (
	"testing"
)

func TestNewModeTagParser(t *testing.T) {
	p := NewModelTagParser()

	if p == nil {
		t.Errorf("Expected a pointer to ModelTagParser but get nil")
	}
}

func TestTagParserParse(t *testing.T) {

	var test = []struct {
		TagsLine   string
		TagsNumber int
	}{
		{`sql:"pk"`, 1},
		{`sql:"nullable=false,name=titulo", json:"title"`, 2},
		{`sql:"nullable=false,name=titulo,unique=true"`, 3},
		{`sql:"pk,name=title,length=80" json:"name"`, 3},
		{`sql:"type=bigint"`, 1},
		{`json:"name"`, 0},
		{`sql:"rel=(type:11; to:PostDetail)"`, 1},
	}

	p := NewModelTagParser()

	for _, tc := range test {
		ts := p.Parse(tc.TagsLine)
		if len(ts) != tc.TagsNumber {
			t.Errorf("Expected: %d tags(%s) but get:%d ", tc.TagsNumber, tc.TagsLine, len(ts))
		}
	}

}
