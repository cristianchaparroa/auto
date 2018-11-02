package parser

import (
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

	if sc > 2 {
		tagsLine := components[2]
		// Todo: replace it for TagParser
		tags := strings.Split(tagsLine, ",")
		mf.Tags = tags
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
