package schema

import "github.com/cristianchaparroa/auto/connection"

// OracleManager implementation for Oracle schema
type OracleManager struct {
	*DatabaseManager
}

// NewOracleManager returns a pointer for OracleManager
func NewOracleManager(c *connection.Config) *OracleManager {
	conn := GetConnection(c)

	pm := &OracleManager{}
	pm.Connection = conn
	pm.Config = c

	return pm
}

// Clean erase the Oracle schema
func (m *OracleManager) Clean() error {
	return nil
}
