package schema

// MysqlAuditor implemets the auditor for Mysql
type MysqlAuditor struct{}

// Create the auditor for Mysql
func (a *MysqlAuditor) Create() (string, error) {
	return "", nil
}

func (a *MysqlAuditor) ExistsAuditor() bool {
	return false
}
