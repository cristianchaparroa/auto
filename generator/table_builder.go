package generator

import (
	"github.com/cristianchaparroa/auto/connection"
)

// TableBuilder creates an instance of Table generator
type TableBuilder struct {
}

// NewTableBuilder creates a ponter to TableBuilder
func NewTableBuilder() *TableBuilder {
	return &TableBuilder{}
}

// GetTableGenerator retrieves the table generator according with the driver
func (b *TableBuilder) GetTableGenerator(Driver string) TableGenerator {
	if Driver == connection.PostgresDriver {
		return NewPostgresTable()
	}

	if Driver == connection.OracleDriver {

	}

	if Driver == connection.MysqlDriver {

	}
	return nil
}
