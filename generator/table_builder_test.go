package generator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cristianchaparroa/auto/connection"
)

func TestNewTableBuilder(t *testing.T) {
	tb := NewTableBuilder()

	// TODO: implmenent de appropiate test case when it'll implemented the
	// gnerators for Oracle and Mysql
	var test = []struct {
		Driver string
		Class  string
	}{
		{connection.MysqlDriver, "<nil>"},
		{connection.OracleDriver, "<nil>"},
		{connection.PostgresDriver, "*postgresgen.PostgresTable"},
	}

	for _, tc := range test {
		g := tb.GetTableGenerator(tc.Driver)

		class := fmt.Sprintf("%v", reflect.TypeOf(g))

		if class != tc.Class {
			t.Errorf("Expected %s but get:%s", tc.Class, class)
		}
	}

}
