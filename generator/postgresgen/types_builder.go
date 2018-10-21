package postgresgen

import (
	"errors"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
	"github.com/cristianchaparroa/auto/parser"
	"github.com/fatih/camelcase"
)

const (
	// TypePosgresVarchar is the keyword Varchar
	TypePosgresVarchar TypeColumn = "varchar"

	// TypePostgresNumeric is the keyword Numeric
	TypePostgresNumeric TypeColumn = "numeric"

	// TypePostgresInt reprsent a integer sql  type supported by postgres
	TypePostgresInt TypeColumn = "int"

	// TypePostgresBoolean represents a SQl Boolean type
	TypePostgresBoolean TypeColumn = "boolean"

	// TypePostgresBigInt represents a SQl Bigint type
	TypePostgresBigInt TypeColumn = "bigint"

	// TypePostgresTimestamp  represents a SQL timestamp type
	TypePostgresTimestamp TypeColumn = "timestamp"

	// TypePostgresTime represents a SQL time type
	TypePostgresTime TypeColumn = "time"

	// TypePostgresDate represents a SQL date type
	TypePostgresDate TypeColumn = "date"

	// TypePostgresRelation represents a relation between models
	TypePostgresRelation TypeColumn = "relation"
)

// TypeColumn is the type of column generated from meta field
type TypeColumn string

// TypeColumnBuilder geneates the properly type for a column according
// with meta data  of field in a model
type TypeColumnBuilder struct {
}

// NewTypeColumnBuilder generates a pointer to TypeColumnBuilder
func NewTypeColumnBuilder() *TypeColumnBuilder {
	return &TypeColumnBuilder{}
}

// GetType returns the type of column  related to meta field
func (b *TypeColumnBuilder) GetType(f *meta.Field) (TypeColumn, error) {

	if f == nil {
		return "", errors.New(TypeColumnFieldNil)
	}

	ft := f.Typ

	if ft.Name == parser.TypeFieldString {
		return TypePosgresVarchar, nil
	}

	if ft.Name == parser.TypeFieldInteger {
		return TypePostgresInt, nil
	}

	if ft.Name == parser.TypeFieldBool {
		return TypePostgresBoolean, nil
	}

	if ft.Name == parser.TypeFieldBigFloat {
		return TypePostgresBigInt, nil
	}

	if ft.Name == parser.TypeFieldArray {
		return TypePostgresRelation, nil
	}

	return "", errors.New(TypeColumnNotSupported)
}

// GetName retrieves the name of column
// It'll split the  field name in camel case and return it in upper and between
// camel case it'll put a underscore
func (b *TypeColumnBuilder) GetName(fieldName string) string {
	name := camelcase.Split(fieldName)
	return strings.ToUpper(strings.Join(name, "_"))

}
