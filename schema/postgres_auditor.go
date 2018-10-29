package schema

// PostgresAuditor implemets the auditor for Postgres
type PostgresAuditor struct{}

// Create the auditor for Postgres
func (a *PostgresAuditor) Create() (string, error) {
	return "", nil
}
