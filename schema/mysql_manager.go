package schema

import "github.com/cristianchaparroa/auto/connection"

// MysqlManager implementation for Mysql schema
type MysqlManager struct {
	*DatabaseManager
}

// NewMysqlManager returns a pointer for MysqlManager
func NewMysqlManager(c connection.Config) *MysqlManager {
	conn := GetConnection(c)

	pm := &MysqlManager{}
	pm.Connection = conn
	pm.Config = c

	return pm
}

// Clean erease the Mysql schema
func (m *MysqlManager) Clean() error {
	return nil
}
