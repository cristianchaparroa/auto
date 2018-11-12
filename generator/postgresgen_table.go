package generator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

// PostgresTable generates a query for create a table in postgres.
type PostgresTable struct {
}

// NewPostgresTable returns a pointer to PostgresTable
func NewPostgresTable() *PostgresTable {
	return &PostgresTable{}
}

// Generate the sql query to create a table for postgres
func (g *PostgresTable) Generate(m *meta.ModelStruct) (ITableResult, error) {

	result := &TableResult{}

	tableName := strings.ToUpper(m.ModelName)
	sql := fmt.Sprintf("CREATE TABLE %v ();", tableName)

	var buffer bytes.Buffer
	buffer.WriteString(sql)

	fs := m.Fields

	cg := NewPostgresColumn()

	relations := make([]*meta.Field, 0)

	for _, f := range fs {

		if f.IsRelation {
			relations = append(relations, f)
			continue
		}

		sqlField, _ := cg.Create(m.ModelName, f)
		buffer.WriteString(fmt.Sprintf("\n%s;", sqlField))
	}

	result.Relations = relations
	result.SQLResult = buffer.String()
	return result, nil
}

// CreateRelation generates a relation between to models
func (g *PostgresTable) CreateRelation(field *meta.Field) (string, error) {
	return "", nil
}
