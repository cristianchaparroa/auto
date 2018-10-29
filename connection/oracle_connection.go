package connection

import (
	"database/sql"
	//	_ "github.com/mattn/go-oci8"
)

// Oracle is the concret connection to database
type Oracle struct {
	Config *Config
	Db     *sql.DB
}

// NewOracleConnection generates a Oracle pointer
func NewOracleConnection(c *Config) *Oracle {
	return &Oracle{Config: c}
}

// Open generates a connection with Oracle database
func (o *Oracle) Open() (*sql.DB, error) {
	/*c := o.Config

	dsn := fmt.Sprintf("%s/%s//%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.Orcl)

	db, err := sql.Open("oci8", dsn)
	if err != nil {
		return nil, err
	}

	o.Db = db
	return db, nil*/

	return nil, nil
}

// Close the current connection with Oracle database
func (o *Oracle) Close() error {
	return o.Db.Close()
}
