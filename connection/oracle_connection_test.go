package connection

import "testing"

func TestOracleOpen(t *testing.T) {
	c := Config{}
	oc := NewOracleConnection(c)

	_, err := oc.Open()

	if err != nil {
		t.Error(err)
	}
}
