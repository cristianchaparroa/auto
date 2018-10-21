package generator

import "github.com/cristianchaparroa/auto/generator/postgresgen"

const (
	// MysqlDriver is the driver used to build tables for Mysql
	MysqlDriver string = "mysql"
	// OracleDriver is the driver used to build tables for Oracle
	OracleDriver string = "oracle"
	// PostgresDriver is the driver used to build tables for Postgres
	PostgresDriver string = "postgres"
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
	if b.Driver == PostgresDriver {
		return postgresgen.NewPostgresTable()
	}

	if b.Driver == OracleDriver {

	}

	if b.Driver == MysqlDriver {

	}
	return nil
}
