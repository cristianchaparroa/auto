package schema

import (
	"database/sql"

	"github.com/cristianchaparroa/auto/connection"
	"github.com/cristianchaparroa/auto/generator"
	"github.com/cristianchaparroa/auto/meta"
)

// Manager is in charge of abstract the sql generation assets such as
// Tables, Sequences, Foreign Keys and Indexes.
type Manager interface {

	// Execute creates, updates, delete all changes in models scanned
	Execute() error

	// Clean the schema
	Clean() error

	// CreateTable table in the current schema
	CreateTable(ms *meta.ModelStruct) (sql.Result, error)

	// CreateTables creates multiples tables
	CreateTables(ms []*meta.ModelStruct) ([]sql.Result, error)
}

// ManagerBuilder handle the schema Manager
type ManagerBuilder struct {
}

// NewManagerBuilder create a pointer to ManagerBuilder
func NewManagerBuilder(driver string) *ManagerBuilder {
	return &ManagerBuilder{}
}

// GetManager gets a manager for driver selected
func (m *ManagerBuilder) GetManager(c connection.Config) Manager {

	if c.Driver == connection.PostgresDriver {
		return NewPostgresManager(c)
	}

	if c.Driver == connection.OracleDriver {

	}

	if c.Driver == connection.MysqlDriver {

	}

	return nil
}

// DatabaseManager is concret manager for database
type DatabaseManager struct {
	Config     connection.Config
	Connection *sql.DB
}

// NewDatabaseManager creates a pointer to DatabaseManager
func GetConnection(c connection.Config) *sql.DB {

	cb := connection.NewBuilder()

	con, err := cb.GetConnection(c)

	if err != nil {
		panic(err)
	}

	connection, err := con.Open()

	if err != nil {
		panic(err)
	}
	return connection
}

// CreateTable create a table in database
func (m *DatabaseManager) CreateTable(ms *meta.ModelStruct) (sql.Result, error) {

	c := m.Connection

	tb := generator.NewTableBuilder()
	tg := tb.GetTableGenerator(m.Config.Driver)
	sql, err := tg.Generate(ms)

	if err != nil {
		return nil, err
	}

	res, err := c.Exec(sql)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateTables creates in database multiples tables
func (m *DatabaseManager) CreateTables(ms []*meta.ModelStruct) ([]sql.Result, error) {

	rs := make([]sql.Result, 0)
	for _, model := range ms {
		r, err := m.CreateTable(model)

		if err != nil {
			return nil, err
		}

		rs = append(rs, r)
	}

	return rs, nil
}

// Execute generates the whole changes in the schema
func (m *DatabaseManager) Execute() error {

	// defines here how to do the things
	// drop-create:
	//    1. Clean SCHEMA
	// 		2. Generate all queries to create
	// Update strategy:
	// 	  1. First verify if table exist
	//	  2. If table doesn't exist create it and all
	// 	  3. If table exist check the modifications in fields
	return nil
}
