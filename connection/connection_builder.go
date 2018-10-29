package connection

import (
	"database/sql"
	"fmt"
)

// Connection defines the methods to generate a connection to database
type Connection interface {
	Open() (*sql.DB, error)
	Close() error
}

// Builder retrieves spcific connection according with
// driver configuration
type Builder struct {
}

// NewBuilder retrieves a pointer to  ConnectionBuilder
func NewBuilder() *Builder {
	return &Builder{}
}

// GetConnection retrieves specific connection according with driver
func (b *Builder) GetConnection(c *Config) (Connection, error) {

	if c.Driver == PostgresDriver {
		pc := NewPostgresConnection(c)
		return pc, nil
	}

	if c.Driver == OracleDriver {
		oc := NewOracleConnection(c)
		return oc, nil
	}

	if c.Driver == MysqlDriver {
		mc := NewMysqlConnection(c)
		return mc, nil
	}

	return nil, fmt.Errorf("%s :%s", ConnectionNotSupportedError, c.Driver)
}
