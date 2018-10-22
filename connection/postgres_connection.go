package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Postgres is the postgres connection implementation
type Postgres struct {
	Config Config
	Db     *sql.DB
}

// NewPostgresConnection generates a pointer to Postgres
func NewPostgresConnection(c Config) *Postgres {
	return &Postgres{Config: c}
}

// Open it the function that generates a connection with postgres
func (p *Postgres) Open() (*sql.DB, error) {
	c := p.Config
	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.Host, c.Port, c.User, c.Password, c.Database)

	db, err := sql.Open(PostgresDriver, source)

	if err != nil {
		return nil, err
	}

	p.Db = db
	return db, nil
}

// Close the current connection with postgres database
func (p *Postgres) Close() error {
	return p.Db.Close()
}
