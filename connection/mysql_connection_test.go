package connection

import "testing"

func TestMysqlOpen(t *testing.T) {
	c := Config{}
	mc := NewMysqlConnection(c)

	_, err := mc.Open()

	if err != nil {
		t.Error(err)
	}
}
