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

	CreateTable(ms *meta.ModelStruct) (sql.Result, error)

	CreateTables(ms []*meta.ModelStruct) (sql.Result, error)
}

// DatabaseManager is concret manager for database
type DatabaseManager struct {
	Config     connection.Config
	Connection *sql.DB
}

// NewDatabaseManager creates a pointer to DatabaseManager
func NewDatabaseManager(c connection.Config) *DatabaseManager {

	cb := connection.NewBuilder()

	con, err := cb.GetConnection(c)

	if err != nil {
		panic(err)
	}

	connection, err := con.Open()

	if err != nil {
		panic(err)
	}

	dm := &DatabaseManager{Config: c, Connection: connection}
	return dm
}

// CreateTable create a table in database
func (m *DatabaseManager) CreateTable(ms *meta.ModelStruct) (sql.Result, error) {
	c := m.Connection

	tb := generator.NewTableBuilder(m.Config.Driver)
	tg := tb.GetTableGenerator()
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
func (m *DatabaseManager) CreateTables(ms []*meta.ModelStruct) (sql.Result, error) {
	return nil, nil
}

func (m *DatabaseManager) Execute() error {
	return nil
}
