package schema

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cristianchaparroa/auto/connection"
)

func TestNewAuditorBuilder(t *testing.T) {
	ab := NewAuditorBuilder()

	if ab == nil {
		t.Error("Expected a new builder but get a nil pointer")
	}

}

func TestGetAuditor(t *testing.T) {
	ab := NewAuditorBuilder()

	var test = []struct {
		Driver string
		Class  string
	}{
		{connection.PostgresDriver, "*schema.PostgresAuditor"},
		{connection.MysqlDriver, "*schema.MysqlAuditor"},
		{connection.OracleDriver, "*schema.OracleAuditor"},
	}

	for _, tc := range test {
		a := ab.GetAuditor(tc.Driver)
		auditorType := fmt.Sprintf("%v", reflect.TypeOf(a))
		if auditorType != tc.Class {
			t.Errorf("Expected %s but get:%s", tc.Class, auditorType)
		}
	}
}

func TestGetAuditorNOK(t *testing.T) {
	ab := NewAuditorBuilder()

	auditor := ab.GetAuditor("AnotherDatabaseDriver")

	if auditor != nil {
		auditorType := fmt.Sprintf("%v", reflect.TypeOf(auditor))
		t.Errorf("Expected nil but get:%s", auditorType)
	}
}
