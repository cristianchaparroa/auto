package parser

import (
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

const (

	// TypeFieldString is the name of string type
	TypeFieldString string = "string"

	// TypeFieldInteger is the name of int type
	TypeFieldInteger string = "int"

	// TypeFieldArray is the name of array type for any field represented
	// by symbole []interface{}
	TypeFieldArray string = "array"

	// TypeFieldFloat64 is the name of float64
	TypeFieldFloat64 string = "float64"

	// TypeFieldBool is the name of bool type
	TypeFieldBool string = "bool"
	// TypeFieldBigFloat is the name of big.Float from math/big
	// Should be declared in the way big.Float in other case it cant be able to
	// recognize the type
	TypeFieldBigFloat string = "big.Float"

	// TypeFieldTime is the name of time from time package  to represent dates
	TypeFieldTime string = " time.Time"
)

// TypeFieldBuilder return the type of model
type TypeFieldBuilder struct {
}

// NewTypeFieldBuilder returns pointer to TypeFieldBuilder
func NewTypeFieldBuilder() *TypeFieldBuilder {
	return &TypeFieldBuilder{}
}

// GetType returns the type of field
func (b *TypeFieldBuilder) GetType(typeField string) (*meta.TypeField, error) {

	isArrayType := strings.Contains(typeField, "[]")

	if isArrayType {
		parts := strings.Split(typeField, "[]")
		entity := parts[1]

		return &meta.TypeField{Name: TypeFieldArray, EntityRelated: entity}, nil
	}

	return &meta.TypeField{Name: typeField}, nil
}
