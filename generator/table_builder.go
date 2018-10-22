package generator

import (
	"github.com/cristianchaparroa/auto/connection"
	"github.com/cristianchaparroa/auto/generator/postgresgen"
)

// TableBuilder creates an instance of Table generator
type TableBuilder struct {
	Driver string
}

// NewTableBuilder creates a ponter to TableBuilder
func NewTableBuilder(driver string) *TableBuilder {
	return &TableBuilder{Driver: driver}
}

// GetTableGenerator retrieves the table generator according with the driver
func (b *TableBuilder) GetTableGenerator() TableGenerator {
	if b.Driver == connection.PostgresDriver {
		return postgresgen.NewPostgresTable()
	}

	if b.Driver == connection.OracleDriver {

	}

	if b.Driver == connection.MysqlDriver {

	}
	return nil
}
