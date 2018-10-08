package scanner

import (
	"fmt"
	"os"
	"strings"
)

//PackageScanner defines de functions to scan the models
type PackageScanner interface {
	// Scan retrieves the files to parse in the model directory
	Scan(directoryPath string) ([]string, error)
}

// Scanner is the implementation of PackageScanner
type Scanner struct {
}

// NewPackageScanner returns a pointer of scanner
func NewPackageScanner() *Scanner {
	return &Scanner{}
}

// Scan read the models in foleder to scan
func (s *Scanner) Scan(directoryPath string) ([]string, error) {

	file, err := os.Open(directoryPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	list, err := file.Readdirnames(0)

	if err != nil {
		return nil, err
	}

	models := make([]string, 0)

	for _, model := range list {

		if s.IsModelToScan(model) {
			models = append(models, fmt.Sprintf("%s/%s", directoryPath, model))
		}
	}

	return models, nil
}

// IsModelToScan verifies if is a go extension
func (s *Scanner) IsModelToScan(modelName string) bool {
	return strings.HasSuffix(modelName, ".go")
}
