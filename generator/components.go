package generator

import (
	"github.com/cristianchaparroa/auto/meta"
)

type ITableResult interface {
	GetModel() *meta.ModelStruct
	GetRelations() []*meta.Field
	GetSQLResult() string
}

type TableResult struct {
	Model     *meta.ModelStruct
	Relations []*meta.Field
	SqlResult string
}

func (r *TableResult) GetModel() *meta.ModelStruct {
	return r.GetModel()
}

func (r *TableResult) GetRelations() []*meta.Field {
	return r.Relations
}

func (r *TableResult) GetSQLResult() string {
	return r.SqlResult
}

// TableGenerator interface defines the methods that allow you to create sql for tables
type TableGenerator interface {
	// Generate the sql query to create a table with fields
	Generate(*meta.ModelStruct) (ITableResult, error)
}

// ColumnGenerator interface defines the methods for create or modified  a column
// of model
type ColumnGenerator interface {
	// Create returns the sql that allow to create a colum based on meta field data
	Create(tableName string, f *meta.Field) (string, error)

	// ChangeType return the sql that modifies the original column definition based
	// on meta field data
	ChangeType(tableName string, f *meta.Field) (string, error)
}

// ConstraintGenerator inteface defines the way to create, moified or delete a
// constraint related to a column
type ConstraintGenerator interface {
	// Create returns the sql that allow to create a constraint based on meta field data
	Create(*meta.Field) (string, error)
}
