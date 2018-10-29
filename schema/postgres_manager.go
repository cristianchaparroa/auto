package schema

import (
	"fmt"

	"github.com/cristianchaparroa/auto/connection"
	"github.com/cristianchaparroa/auto/meta"
)

// PostgresManager implementation for Postgres schema
type PostgresManager struct {
	*DatabaseManager
}

//NewPostgresManager returns a manager for Postgres
func NewPostgresManager(c *connection.Config) *PostgresManager {
	conn := GetConnection(c)

	fmt.Println(conn)
	pm := &PostgresManager{&DatabaseManager{Config: c, Connection: conn}}
	return pm
}

// Clean ereses the Postgres schema
func (m *PostgresManager) Clean() error {

	schema := m.Config.Schema
	conn := m.Connection

	sqlDropCreate := fmt.Sprintf(`DROP SCHEMA IF EXISTS %s CASCADE;
															  CREATE SCHEMA %s;`, schema, schema)

	_, err := conn.Exec(sqlDropCreate)

	if err != nil {
		return err
	}

	return nil
}

// Execute generates the whole changes in the schema
func (m *PostgresManager) Execute(ms []*meta.ModelStruct) error {

	fmt.Println("Executing...")
	err := m.Clean()

	if err != nil {
		return err
	}

	m.CreateTables(ms)

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
