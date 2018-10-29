package connection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	cb := NewBuilder()

	if cb == nil {
		t.Error("Expected an instance of Builder but get nil")
	}
}

func TestBuilderGetConnectionNOK(t *testing.T) {
	cb := NewBuilder()

	c := &Config{Driver: "somethingwrong"}
	_, err := cb.GetConnection(c)

	expectedErr := fmt.Sprintf("%s :%s", ConnectionNotSupportedError, c.Driver)

	if err.Error() != expectedErr {
		t.Errorf("Expected %v but get %v", expectedErr, err.Error())
	}
}

func TestBuilderGetConnection(t *testing.T) {
	cb := NewBuilder()

	var test = []struct {
		Driver string
		Class  string
	}{
		{PostgresDriver, "*connection.Postgres"},
		{MysqlDriver, "*connection.Mysql"},
		{OracleDriver, "*connection.Oracle"},
	}

	for _, tc := range test {
		c := &Config{Driver: tc.Driver}
		conn, err := cb.GetConnection(c)

		if err != nil {
			t.Error(err)
		}

		if fmt.Sprintf("%v", reflect.TypeOf(conn)) != tc.Class {
			t.Errorf("Expected %v but get: %v", tc.Class, reflect.TypeOf(conn))
		}
	}
}
