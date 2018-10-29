package connection

import "testing"

func TestPostgresOpen(t *testing.T) {
	c := &Config{}
	pc := NewPostgresConnection(c)

	_, err := pc.Open()
	if err != nil {
		t.Error(err)
	}
}
