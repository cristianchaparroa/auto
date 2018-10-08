package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// PackageParser define the methods to retrieve information related with model package
type PackageParser interface {
	// Retrive the name of the package in which is locate the models
	Parse(fileStr string) (string, error)
}

// ModelPackageParser retrieves information related with model package
type ModelPackageParser struct {
}

//NewModelPackageParser creates a pointer to NewModelPackageParser struct
func NewModelPackageParser() *ModelPackageParser {
	return &ModelPackageParser{}
}

// Parse generete the name of package if is posible to parse the file
// in other case retrieves an error related with the issue
func (p *ModelPackageParser) Parse(fileStr string) (string, error) {
	lines := strings.Split(fileStr, "\n")

	if len(lines) == 1 && lines[0] == fileStr {
		return "", errors.New(PackageParserEmptyFile)
	}

	packageLine := lines[0]

	if len(packageLine) == 0 {
		return "", errors.New(PackageParserEmptyLine)
	}

	packageLine = FilterSpaces(packageLine)
	patternSplit := `\s`

	r, _ := regexp.Compile(patternSplit)
	parts := r.Split(packageLine, -1)

	if len(parts) != 2 {
		return "", fmt.Errorf(PacakgeParserMalformedTwoTokensExpected, len(parts), packageLine)
	}

	packageStmt := parts[0]

	if PackageStatementToken != packageStmt {
		return "", fmt.Errorf(PackageParserMalformedStmtPackage, packageStmt)
	}

	packageName := parts[1]

	return packageName, nil
}
