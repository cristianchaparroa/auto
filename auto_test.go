package auto

import (
	"fmt"
	"os"
	"testing"
)

func getCurentDirectory() (string, error) {
	currDir, err := os.Getwd()

	if err != nil {
		return "", err
	}
	return currDir, nil
}

func getModelDirectory() string {
	td, err := getCurentDirectory()

	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("%s/testdata/models", td)

	return path
}

func TestNewGenerator(t *testing.T) {
	g := NewGenerator()

	if g == nil {
		t.Error("Expected a pointer to Generator but get nil")
	}
}

func TestScan(t *testing.T) {
	path := getModelDirectory()

	g := NewGenerator()
	fs, err := g.Scan(path)

	if err != nil {
		t.Error(err)
	}

	if len(fs) != 3 {
		t.Errorf("Expected 3 files to scan but get :%v", len(fs))
	}

}

func TestReadFile(t *testing.T) {
	path := getModelDirectory()

	authorFile := fmt.Sprintf("%s/author.go", path)
	g := NewGenerator()
	bs, err := g.ReadFile(authorFile)

	if err != nil {
		t.Error(err)
	}

	if len(bs) == 0 {
		t.Error("Expected bytes from file but get empty file")
	}

}

func TestReadAll(t *testing.T) {
	path := getModelDirectory()

	g := NewGenerator()
	fs, err := g.Scan(path)

	if err != nil {
		t.Error(err)
	}

	bs, err := g.ReadAll(fs)

	if err != nil {
		t.Error(err)
	}

	if len(bs) != 3 {
		t.Errorf("Expected 3 files read but get %v", len(bs))
	}
}

func TestParse(t *testing.T) {
	path := getModelDirectory()

	g := NewGenerator()
	fs, err := g.Scan(path)

	if err != nil {
		t.Error(err)
	}

	bs, err := g.ReadAll(fs)

	if err != nil {
		t.Error(err)
	}

	if len(bs) != 3 {
		t.Errorf("Expected 3 files read but get %v", len(bs))
	}
	ms, err := g.ParseAll(bs)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(ms)
}

func TestParseAll(t *testing.T) {

}
