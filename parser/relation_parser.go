package parser

import "github.com/cristianchaparroa/auto/meta"

// RelationParser defines the methods to parse the relation tag
type RelationParser interface {
	Parse(tag string) (*meta.Relation, error)
}

// TagRelationParser implements the methods to parse the relation tag
type TagRelationParser struct {
}

// Parse builds a relation
func (p *TagRelationParser) Parse(tag string) (*meta.Relation, error) {
	return nil, nil
}
