package scanner

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func getCurentDirectory() (string, error) {
	currDir, err := os.Getwd()

	if err != nil {
		return "", err
	}
	return currDir, nil
}

func getTestDataDirectory() (string, error) {
	currDir, err := getCurentDirectory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", currDir, "testdata"), nil
}

func TestIntegrationScanFindModels(t *testing.T) {
	testDataDir, err := getTestDataDirectory()
	if err != nil {
		t.Error(err)
	}
	testDataDir = strings.Replace(testDataDir, "/scanner", "", -1)

	s := NewPackageScanner()

	models, err := s.Scan(fmt.Sprintf("%s/models", testDataDir))

	if err != nil {
		t.Error(err)
	}

	if len(models) == 0 {
		t.Error("Expected models to scan, but is not possible find nothing")
	}

	for _, model := range models {
		fmt.Println(model)
	}
}

func TestIntegrationScanModelsNotFound(t *testing.T) {
	path := "/var/models/notmodels"

	s := NewPackageScanner()
	_, err := s.Scan(path)

	if err == nil {
		t.Error("Expected an error but get nothing")
	}
}
