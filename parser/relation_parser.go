package parser

import (
	"errors"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

// RelationParser defines the methods to parse the relation tag
type RelationParser interface {
	Parse(tag, fieldName string) (*meta.Relation, error)
}

// TagRelationParser implements the methods to parse the relation tag
type TagRelationParser struct {
}

// NewTagRelationParser returns a pointer to TagRelationParser
func NewTagRelationParser() *TagRelationParser {
	return &TagRelationParser{}
}

// Parse builds a relation
func (p *TagRelationParser) Parse(val, fieldName string) (*meta.Relation, error) {
	val = strings.Replace(val, "(", "", -1)
	val = strings.Replace(val, ")", "", -1)

	attributes := strings.Split(val, ";")

	rel := &meta.Relation{}

	for _, att := range attributes {
		att = strings.Replace(att, " ", "", -1)

		parts := strings.Split(att, ":")

		attType := parts[0]
		attVal := parts[1]

		if "type" == attType {
			typ, err := p.GetRelationType(attVal)

			if err != nil {
				return nil, err
			}

			rel.Typ = typ
		}

		if "to" == attType {
			rel.To = attVal
		}

		if "name" == attType {
			rel.Name = attVal
		}

		if "pkref" == attType {
			rel.PKRef = attVal
		}

		if "fk" == attType {
			rel.FK = attVal
		}
		if len(rel.To) == 0 {
			rel.To = fieldName
		}
	}

	return rel, nil
}

// GetRelationType retrieves the type of relation
func (p *TagRelationParser) GetRelationType(typ string) (meta.RelationType, error) {

	if string(meta.OneToOne) == typ {
		return meta.OneToOne, nil
	}

	if string(meta.OneToMany) == typ {
		return meta.OneToMany, nil
	}

	if string(meta.ManyToOne) == typ {
		return meta.ManyToOne, nil
	}

	if string(meta.ManyToMany) == typ {
		return meta.ManyToMany, nil
	}

	return "", errors.New("Unexpected relation type")
}
