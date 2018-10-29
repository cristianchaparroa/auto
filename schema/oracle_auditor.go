package schema

// OracleAuditor implemets the auditor for Oracle
type OracleAuditor struct{}

// Create the auditor for Oracle
func (a *OracleAuditor) Create() (string, error) {
	return "", nil
}
