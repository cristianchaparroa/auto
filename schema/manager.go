package schema

import (
	"database/sql"
	"fmt"

	"github.com/cristianchaparroa/auto/connection"
	"github.com/cristianchaparroa/auto/generator"
	"github.com/cristianchaparroa/auto/meta"

	log "github.com/sirupsen/logrus"
)

// Manager is in charge of abstract the sql generation assets such as
// Tables, Sequences, Foreign Keys and Indexes.
type Manager interface {

	// Execute creates, updates, delete all changes in models scanned
	Execute(ms []*meta.ModelStruct) error

	// Clean the schema
	Clean() error

	// CreateTable table in the current schema
	CreateTable(ms *meta.ModelStruct) (*PartialResult, error)

	// CreateTables creates multiples tables
	CreateTables(ms []*meta.ModelStruct) ([]*PartialResult, []error)
}

// ManagerBuilder handle the schema Manager
type ManagerBuilder struct {
}

// NewManagerBuilder create a pointer to ManagerBuilder
func NewManagerBuilder(driver string) *ManagerBuilder {
	return &ManagerBuilder{}
}

// GetManager gets a manager for driver selected
func (m *ManagerBuilder) GetManager(c *connection.Config) Manager {

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
	Config     *connection.Config
	Connection *sql.DB
}

// GetConnection creates a pointer to DatabaseManager
func GetConnection(c *connection.Config) *sql.DB {

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
func (m *DatabaseManager) CreateTable(ms *meta.ModelStruct) (*PartialResult, error) {

	pr := &PartialResult{}

	log.Info(fmt.Sprintf("auto:processing the model: %s \n", ms.ModelName))
	c := m.Connection

	tb := generator.NewTableBuilder()
	tg := tb.GetTableGenerator(m.Config.Driver)

	tableResult, err := tg.Generate(ms)
	log.Info(fmt.Sprintf("\nauto:The sql generated is: \n %s \n", tableResult.SqlResult))

	if err != nil {
		return nil, err
	}

	res, err := c.Exec(tableResult.SqlResult)

	if err != nil {
		return nil, err
	}

	pr.Model = ms
	pr.SqlExecuted = tableResult.SqlResult
	pr.Res = res
	pr.Relations = tableResult.Relations
	log.Info(fmt.Sprintf("auto:processing the model: %s was finnished \n\n", ms.ModelName))
	return pr, nil
}

// CreateTables creates in database multiples tables
func (m *DatabaseManager) CreateTables(ms []*meta.ModelStruct) ([]*PartialResult, []error) {

	rs := make([]*PartialResult, 0)

	errors := make([]error, 0)

	for _, model := range ms {
		r, err := m.CreateTable(model)

		if err != nil {
			errors = append(errors, err)
			continue
		}

		rs = append(rs, r)
	}

	return rs, errors
}
