package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

// FieldParser defines the methods to parse a fields in entity model
type FieldParser interface {

	// Parse retrieve the representation of one field of entiy model
	Parse(fieldStr string) (*meta.Field, error)

	// ParseFields retrieves an array of fields from content struct
	ParseFields(structContent string) []*meta.Field

	// IsRelation verifies is the current field is an relation statement
	IsRelation(tags []*meta.Tag) bool
}

// ModelFieldParser geneates a meta field  throug string representation of entity model
type ModelFieldParser struct {
}

// NewModelFieldParser retrieves a pointer to NewModelFieldParser struct
func NewModelFieldParser() *ModelFieldParser {
	return &ModelFieldParser{}
}

// Parse generates meta field
func (p *ModelFieldParser) Parse(fieldStr string) (*meta.Field, error) {

	filtered := p.FilterComments(fieldStr)
	components := strings.Fields(filtered)
	sc := len(components)

	if sc == 0 {
		return nil, nil
	}

	mf := &meta.Field{}

	if sc > 0 {
		name := components[0]
		mf.Name = name
	}

	if sc > 1 {
		tb := NewTypeFieldBuilder()
		typ := components[1]
		t, err := tb.GetType(typ)

		if err == nil {
			mf.Typ = t
		}
	}

	tags := p.ParseTags(filtered)

	if len(tags) > 0 {
		mf.Tags = tags
		mf.IsRelation = p.IsRelation(tags)
	}

	if mf.IsRelation {
		rel, err := p.ParseRelation(mf.Name, tags)

		if err != nil {
			return nil, err
		}

		mf.Relation = rel
	}

	return mf, nil
}

// FilterComments filter the comments above of line definition of field
func (p *ModelFieldParser) FilterComments(fieldStr string) string {
	pattern := `(?m)^[a-zA-Z](.*|\n)*`

	r, _ := regexp.Compile(pattern)

	filtered := r.FindString(fieldStr)

	if len(filtered) > 0 {
		return filtered
	}
	return fieldStr
}

// IsCommentedLine verifies if is a commented line to skeep it.
func (p *ModelFieldParser) IsCommentedLine(fieldStr string) bool {
	pattern := `^(\s*\/\/.*)`

	isCommented, _ := regexp.MatchString(pattern, fieldStr)

	return isCommented
}

// ParseFields  it parses the multiples fields in content body of struct
func (p *ModelFieldParser) ParseFields(structContent string) []*meta.Field {
	filtered := p.FilterComments(structContent)

	lines := strings.Split(filtered, "\n")

	fs := make([]*meta.Field, 0)

	for _, l := range lines {

		if p.IsCommentedLine(l) {
			continue
		}

		l = strings.TrimSpace(l)

		if len(l) > 1 {
			mf, _ := p.Parse(l)
			fs = append(fs, mf)
		}
	}
	return fs
}

// IsRelation verifies if the field is a relation statement
func (p *ModelFieldParser) IsRelation(tags []*meta.Tag) bool {

	for _, t := range tags {
		if fmt.Sprintf("%v", t.Typ) == string(TypeTagRelation) {
			return true
		}
	}

	return false
}

// ParseTags retrieves the tags in a field statement
func (p *ModelFieldParser) ParseTags(fieldStr string) []*meta.Tag {
	tp := NewModelTagParser()
	tagsLine := tp.ExtractTagStatement(fieldStr)

	tags := tp.Parse(tagsLine)
	return tags
}

// ParseRelation retrieves the relation statement according with field definition
func (p *ModelFieldParser) ParseRelation(fieldName string, tags []*meta.Tag) (*meta.Relation, error) {

	relStmt := ""
	for _, t := range tags {
		if fmt.Sprintf("%v", t.Typ) == string(TypeTagRelation) {
			relStmt = t.Value
		}
	}

	rp := NewTagRelationParser()
	rel, err := rp.Parse(relStmt, fieldName)

	if err != nil {
		return nil, err
	}
	return rel, nil
}
