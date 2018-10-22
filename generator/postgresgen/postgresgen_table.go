package postgresgen

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
func (g *PostgresTable) Generate(m *meta.ModelStruct) (string, error) {

	tableName := strings.ToUpper(m.ModelName)
	sql := fmt.Sprintf("CREATE TABLE %v ();", tableName)

	var buffer bytes.Buffer
	buffer.WriteString(sql)

	fs := m.Fields

	cg := NewPostgresColumn()

	for _, f := range fs {
		sqlField, _ := cg.Create(m.ModelName, f)
		buffer.WriteString(fmt.Sprintf("\n%s;", sqlField))
	}

	return buffer.String(), nil
}
