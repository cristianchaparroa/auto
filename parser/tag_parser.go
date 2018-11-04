package parser

import (
	"regexp"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

// TagParser prove the definition of methods to parse a tag in a field
type TagParser interface {
}

// ModelTagParser parses the tags in field
type ModelTagParser struct {
}

//NewModelTagParser creates a pointer for
func NewModelTagParser() *ModelTagParser {
	return &ModelTagParser{}
}

// Parse the tags in column
func (p *ModelTagParser) Parse(tagsLine string) []*meta.Tag {

	tagsLine = p.ExtractSQLContent(tagsLine)

	if len(tagsLine) == 0 {
		return nil
	}

	tagsLine = strings.Replace(tagsLine, "sql:", "", -1)
	tagsLine = strings.Replace(tagsLine, "\"", "", -1)
	tags := strings.Split(tagsLine, ",")

	ts := make([]*meta.Tag, 0)

	tb := NewTypeTagBuilder()

	for _, t := range tags {

		parts := strings.Split(t, "=")
		sz := len(parts)
		tag := &meta.Tag{}

		if sz > 0 && p.IsPrimaryKeyTag(parts[0]) {
			val := parts[0]
			tp := tb.GetType(val)
			tag.Value = "pk"
			tag.Typ = tp
			ts = append(ts, tag)
			continue
		}

		if sz > 0 {
			typ := parts[0]
			tp := tb.GetType(typ)
			tag.Typ = tp
		}

		if sz > 1 {
			val := parts[1]
			tag.Value = val
		}
		ts = append(ts, tag)
	}
	return ts
}

// ExtractSQLContent retrieves just the content inside on sql:" ... "
func (p *ModelTagParser) ExtractSQLContent(tagsLine string) string {

	pattern := `\s*sql:"(.|\n)*?"`

	r := regexp.MustCompile(pattern)

	sqlContent := r.FindString(tagsLine)

	return sqlContent
}

// IsPrimaryKeyTag verifies if value of tag is a primary key
func (p *ModelTagParser) IsPrimaryKeyTag(value string) bool {
	if "pk" == value {
		return true
	}
	return false
}

// ExtractTagStatement retrieves just the tags statement inside of fieldStr
func (p *ModelTagParser) ExtractTagStatement(fieldStr string) string {
	pattern := "\\s*`(.|\n)*?`"

	r := regexp.MustCompile(pattern)

	tagsStatement := r.FindString(fieldStr)

	return tagsStatement
}
