package parser

import (
	"errors"
	"regexp"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

// StructParser defines the parser to get struct models from file
type StructParser interface {

	// ParseFile retrieves all structs(array of strings) in a file
	ParseFile(file string) ([]string, error)

	// GetModels retrieves all metadata of models scaned
	GetModels(file string) ([]*meta.ModelStruct, error)

	// Parse retrieve a metda of from struct represented by string
	Parse(structStr string) (*meta.ModelStruct, error)

	// Validate if is  well formed the struct
	Validate(structStr string) (bool, error)

	// ParseName retrieve the name of model
	ParseName(structStr string) string

	// ParsePackage Retrieves the package name in which is the model
	ParsePackage(file string) (string, error)
}

// ModelStructParser parses a file wich contains a posible representation entities
type ModelStructParser struct {
}

// NewModelStructParser creates a pointer to ModelStructParser
func NewModelStructParser() *ModelStructParser {
	return &ModelStructParser{}
}

// ParseFile find structures in a file. It retrieve a array of structs
// whitout parse.
func (p *ModelStructParser) ParseFile(file string) ([]string, error) {

	r, _ := regexp.Compile(`(?m)^type\s*[a-zA-Z]{0,1}[a-zA-Z0-9]*\s*struct{0,1}\s*{(.|\n)*?}`)

	structs := r.FindAllString(file, -1)
	return structs, nil
}

// GetModels retreive models structs from file to parse
func (p *ModelStructParser) GetModels(file string) ([]*meta.ModelStruct, error) {
	structs, _ := p.ParseFile(file)

	ms := make([]*meta.ModelStruct, 0)

	packageName, err := p.ParsePacakge(file)

	if err != nil {
		return nil, err
	}

	for _, s := range structs {
		m, _ := p.Parse(s)

		if m != nil {
			m.PackageName = packageName
			ms = append(ms, m)
		}

	}
	return ms, nil
}

// Parse retrieves a model struct from string representation of entity model
func (p *ModelStructParser) Parse(structStr string) (*meta.ModelStruct, error) {
	model := &meta.ModelStruct{}

	// validate the statement definition
	r := regexp.MustCompile(`(?m)^type\s*[a-zA-Z]{0,1}[a-zA-Z0-9]*\s*struct{0,1}`)
	isValid := r.MatchString(structStr)

	if !isValid {
		return nil, errors.New(StructParserMalformedPatterNotFound)
	}

	// parse name
	name, err := p.ParseName(structStr)

	if err != nil {
		return nil, err
	}

	model.ModelName = name
	// parse body
	pattern := `\s*{(.|\n)*?}`
	r = regexp.MustCompile(pattern)

	content := r.FindString(structStr)
	content = strings.Replace(content, "{", "", -1)
	content = strings.Replace(content, "}", "", -1)

	fp := NewModelFieldParser()
	fields := fp.ParseFields(content)
	model.Fields = fields

	return model, nil
}

// ParseName retrieves the name of struct from parsed struct representation
func (p *ModelStructParser) ParseName(structStr string) (string, error) {
	r := regexp.MustCompile(`(?m)^type\s*[a-zA-Z]{0,1}[a-zA-Z0-9]*\s*struct{0,1}`)
	structStmLine := r.FindString(structStr)

	structStmLine = strings.Replace(structStmLine, "type", "", -1)
	structStmLine = strings.Replace(structStmLine, "struct", "", -1)
	name := FilterSpaces(structStmLine)

	r = regexp.MustCompile(`^[a-zA-Z].([a-zA-Z0-9]+)`)
	isValidName := r.MatchString(name)

	if !isValidName {
		return "", errors.New(StructParserMalformedNameNotfound)
	}

	return name, nil
}

// ParsePacakge retrieves the package name of model entities
func (p *ModelStructParser) ParsePacakge(fileStr string) (string, error) {
	pp := NewModelPackageParser()
	return pp.Parse(fileStr)
}
