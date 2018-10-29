package schema

import (
	"github.com/cristianchaparroa/auto/connection"
)

// Auditor defines all functions to audit the schema
type Auditor interface {
	// Create returns the sql that create the trigger to audit the changes in schema
	Create() (string, error)

	// ExistsAuditor verifies if exist an auditor
	ExistsAuditor() bool
}

// AuditorBuilder creates an specific autor acording with the setup driver
type AuditorBuilder struct {
}

// NewAuditorBuilder returns a pointer to AuditorBuilder
func NewAuditorBuilder() *AuditorBuilder {
	return &AuditorBuilder{}
}

// GetAuditor returns the auditor according with the driver used
func (a *AuditorBuilder) GetAuditor(driver string) Auditor {

	if driver == connection.PostgresDriver {
		return &PostgresAuditor{}
	}

	if driver == connection.OracleDriver {
		return &OracleAuditor{}
	}

	if driver == connection.MysqlDriver {
		return &MysqlAuditor{}
	}

	return nil
}
