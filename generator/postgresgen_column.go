package generator

import (
	"fmt"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

// PostgresColumn allows to crete columns for a table in Postgres database
type PostgresColumn struct {
}

// NewPostgresColumn returns a pointer to PostgresColumn
func NewPostgresColumn() *PostgresColumn {
	return &PostgresColumn{}
}

// Create generates a sql to create column
func (g *PostgresColumn) Create(tableName string, f *meta.Field) (string, error) {
	tb := NewTypeColumnBuilder()
	dataType, err := tb.GetType(f)

	if err != nil {
		return "", err
	}

	dataTypeStr := strings.ToUpper(fmt.Sprintf(`%v`, dataType))
	tableName = strings.ToUpper(tableName)
	columnName := tb.GetName(f.Name)
	sql := fmt.Sprintf(`ALTER TABLE %v ADD COLUMN %v %v`, tableName, columnName, dataTypeStr)

	if f.IsPrimaryKey {
		sql = fmt.Sprintf("%s PRIMARY KEY", sql)
	}
	if f.IsNotNull {
		sql = fmt.Sprintf("%s NOT NULL", sql)
	}

	if f.IsUnique {
		sql = fmt.Sprintf("%s UNIQUE", sql)
	}

	return sql, nil
}

// ChangeType generates a sql to alter the type definition of column
func (g *PostgresColumn) ChangeType(tableName string, f *meta.Field) (string, error) {
	tb := NewTypeColumnBuilder()
	dataType, err := tb.GetType(f)

	if err != nil {
		return "", err
	}

	dataTypeStr := strings.ToUpper(fmt.Sprintf(`%v`, dataType))
	tableName = strings.ToUpper(tableName)
	columnName := tb.GetName(f.Name)

	sql := fmt.Sprintf(`ALTER TABLE %v ALTER COLUMN %v TYPE %v`, tableName, columnName, dataTypeStr)
	return sql, nil
}

// CreatePrimaryKey creates the query
func (g *PostgresColumn) CreatePrimaryKey(tableName, columnName string) string {
	sql := fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY (%s)", tableName, columnName)
	return sql
}
