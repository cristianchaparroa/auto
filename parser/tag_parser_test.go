package parser

import "testing"

func TestNewModeTagParser(t *testing.T) {
	p := NewModelTagParser()

	if p == nil {
		t.Errorf("Expected a pointer to ModelTagParser but get nil")
	}
}
