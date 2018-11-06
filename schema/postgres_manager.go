package schema

import (
	"bytes"
	"errors"
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

	resps, errs := m.CreateTables(ms)

	fmt.Println("--------responses------")
	fmt.Println(len(resps))
	var buffer bytes.Buffer

	for _, err := range errs {
		buffer.WriteString(err.Error() + "\n")
	}

	finalErr := buffer.String()

	if len(finalErr) > 0 {
		return errors.New(finalErr)
	}

	fmt.Println("--------relations------")

	fs := make([]*meta.Field, 0)
	for _, r := range resps {
		fs = append(fs, r.Relations...)
	}
	fmt.Println(len(fs))
	for _, rel := range fs {
		fmt.Println(rel)
	}

	// from here creates the relations

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
