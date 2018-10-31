package parser

import (
	"github.com/cristianchaparroa/auto/meta"
)

const (
	// TypeTagName represents a tag that indicate what should be the name of column
	// in database
	TypeTagName meta.TypeTag = "name"

	// TypeTagType force to stablish a specific type for column
	TypeTagType meta.TypeTag = "type"

	// TypeTagPrimaryKey indicates that column should be the primary key
	TypeTagPrimaryKey meta.TypeTag = "pk"

	// TypeTagUnique indicates that colum data must be unique in database
	TypeTagUnique meta.TypeTag = "unique"

	// TypeTagNullable represents that data nullabe should be stored in column
	TypeTagNullable meta.TypeTag = "nullable"

	// TypeTagLength indicates the length of column
	TypeTagLength meta.TypeTag = "length"
)

// TypeTagBuilder is in charge to build the type of tag
type TypeTagBuilder struct {
}

// NewTypeTagBuilder returns  a pointer to TypeTagBuilder
func NewTypeTagBuilder() *TypeTagBuilder {
	return &TypeTagBuilder{}
}

// GetType returns the type of tag
func (b *TypeTagBuilder) GetType(tag string) *meta.TypeTag {

	typ := meta.TypeTag(tag)

	if meta.TypeTag(tag) == TypeTagName {
		return &typ
	}

	if meta.TypeTag(tag) == TypeTagType {
		return &typ
	}

	if meta.TypeTag(tag) == TypeTagPrimaryKey {
		return &typ
	}

	if meta.TypeTag(tag) == TypeTagUnique {
		return &typ
	}

	if meta.TypeTag(tag) == TypeTagNullable {
		return &typ
	}

	if meta.TypeTag(tag) == TypeTagLength {
		return &typ
	}

	return nil
}
