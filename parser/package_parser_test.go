package parser

import (
	"fmt"
	"testing"
)

func TestPackageParserNew(t *testing.T) {
	p := NewModelPackageParser()

	if p == nil {
		t.Error("Expected a pointer ModelPackageParser but get nil")
	}
}

func TestPacakgeParseOK(t *testing.T) {
	var test = []struct {
		fileStr      string
		NameExpected string
	}{
		{"package info \n type FakeStruct struct {}", "info"},
		{"package         tesla \n type FakeStruct struct {}", "tesla"},
	}

	p := NewModelPackageParser()

	for _, tc := range test {
		name, err := p.Parse(tc.fileStr)

		if err != nil {
			t.Error(err)
		}
		if name != tc.NameExpected {
			t.Errorf("expected %v but get:%v", tc.NameExpected, name)
		}
	}
}

func TestPackageParseNOK(t *testing.T) {
	var test = []struct {
		fileStr       string
		ErrorExpected string
	}{
		{"", PackageParserEmptyFile},
		{"\n type FakeStruct struct {}", PackageParserEmptyLine},
		{"package \n type FakeStruct struct {}", fmt.Sprintf(PacakgeParserMalformedTwoTokensExpected, 1, "package")},
		{"pack info \n type FakeStruct struct {}", fmt.Sprintf(PackageParserMalformedStmtPackage, "pack")},
	}

	p := NewModelPackageParser()

	for _, tc := range test {
		_, err := p.Parse(tc.fileStr)

		if err.Error() != tc.ErrorExpected {
			t.Errorf("Expected the error:%v but get:%v", tc.ErrorExpected, err.Error())
		}
	}

}
