package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Mysql is the concret connection to database
type Mysql struct {
	Config Config
	Db     *sql.DB
}

// NewMysqlConnection generates a Mysql pointer
func NewMysqlConnection(c Config) *Mysql {
	return &Mysql{Config: c}
}

// Open generates a connection with Mysql database
func (m *Mysql) Open() (*sql.DB, error) {
	c := m.Config
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := sql.Open(MysqlDriver, source)

	if err != nil {
		return nil, err
	}

	m.Db = db
	return db, nil
}

// Close the current connection with Mysql database
func (m *Mysql) Close() error {
	return m.Db.Close()
}
