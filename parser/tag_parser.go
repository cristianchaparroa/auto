package parser

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
