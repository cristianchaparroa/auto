package schema

import (
	"fmt"

	"github.com/cristianchaparroa/auto/connection"
)

// PostgresManager implementation for Postgres schema
type PostgresManager struct {
	*DatabaseManager
}

//NewPostgresManager returns a manager for Postgres
func NewPostgresManager(c connection.Config) *PostgresManager {
	conn := GetConnection(c)

	pm := &PostgresManager{}
	pm.Connection = conn
	pm.Config = c

	return pm
}

// Clean ereses the Postgres schema
func (m *PostgresManager) Clean() error {

	schema := m.Config.Schema
	conn := m.Connection

	sqlDropCreate := fmt.Sprintf(`DROP SCHEMA IF EXIST %s CASCADE; CREATE SCHEMA %s;`, schema, schema)

	sqlPermissions := fmt.Sprintf(`GRANT ALL ON SCHEMA %s TO postgres; GRANT ALL ON SCHEMA %s TO %s;`, schema, schema, schema)

	_, err := conn.Exec(sqlDropCreate)

	if err != nil {
		return err
	}

	_, err = conn.Exec(sqlPermissions)

	if err != nil {
		return err
	}

	return nil
}
