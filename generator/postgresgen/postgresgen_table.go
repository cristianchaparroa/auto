package postgresgen

import (
	"bytes"
	"fmt"

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
func (g *PostgresTable) Generate(m *meta.ModelStruct) (string, error) {
	sql := fmt.Sprintf("CREATE TABLE %v ();", m.ModelName)

	var buffer bytes.Buffer
	buffer.WriteString(sql)

	fs := m.Fields

	cg := NewPostgresColumn()

	for _, f := range fs {
		sqlField, _ := cg.Create(m.ModelName, f)
		buffer.WriteString("\n\t" + sqlField)
	}

	return buffer.String(), nil
}
