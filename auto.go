package auto

import (
	"io/ioutil"

	"github.com/cristianchaparroa/auto/config"
	"github.com/cristianchaparroa/auto/meta"
	"github.com/cristianchaparroa/auto/parser"
	"github.com/cristianchaparroa/auto/scanner"
	"github.com/cristianchaparroa/auto/schema"
)

// Auto defines the interface to Auto tool
type Auto interface {
	// Generate tables from models
	Generate(c *config.Config) error

	// Scan models  directory and retrieve the file paths to parse files
	Scan(path string) ([]string, error)

	// Parse a file and generate meta models for structus on that file
	Parse(file []byte) ([]*meta.ModelStruct, error)

	// Parse files and generate meta models for structus in all model directory
	ParseAll(files [][]byte) ([]*meta.ModelStruct, error)

	// ReadFile read a file that contains models and retrive the information on that
	ReadFile(path string) ([]byte, error)

	// ReadAll files to parsed
	ReadAll(paths []string) ([][]byte, error)
}

// Generator  integratees all the flows for generates the tables from models
type Generator struct {
}

// NewGenerator generate a pointer to Generator
func NewGenerator() *Generator {

	return &Generator{}
}

// Generate tables from models
func (g *Generator) Generate(c *config.Config) error {
	// 1. Scan the paths
	// 2. Read the files
	// 3. parse the files
	// 4. get a database managr
	// 5. connect with respective database
	// 6. execute all changes scanned
	paths, err := g.Scan(c.PathModels)

	if err != nil {
		panic(err)
	}

	bs, err := g.ReadAll(paths)

	if err != nil {
		panic(err)
	}

	ms, err := g.ParseAll(bs)

	if err != nil {
		panic(err)
	}

	sb := schema.NewManagerBuilder(c.Driver)

	dbc := config.NewDatabaseConfig(c)

	dm := sb.GetManager(dbc)

	if dm == nil {
		panic("It's not possible to retrive the schema manager")
	}
	err = dm.Execute(ms)

	if err != nil {
		panic(err)
	}
	return nil
}

// Scan models  directory and retrieve the file paths to parse files
func (g *Generator) Scan(path string) ([]string, error) {
	s := scanner.NewPackageScanner()
	return s.Scan(path)
}

// ReadFile read a file in wich is supposed exist an model to be parsed
func (g *Generator) ReadFile(path string) ([]byte, error) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// ReadAll  files in the models directory
func (g *Generator) ReadAll(paths []string) ([][]byte, error) {

	files := make([][]byte, 0)

	for _, p := range paths {

		bs, err := g.ReadFile(p)
		if err != nil {
			return nil, err
		} else {
			files = append(files, bs)
		}

	}
	return files, nil
}

// Parse a files
func (g *Generator) Parse(file []byte) ([]*meta.ModelStruct, error) {

	mp := parser.NewModelStructParser()
	models, err := mp.GetModels(string(file))

	if err != nil {
		return nil, err
	}
	return models, nil
}

// ParseAll  files in model directory
func (g *Generator) ParseAll(files [][]byte) ([]*meta.ModelStruct, error) {
	models := make([]*meta.ModelStruct, 0)

	for _, f := range files {
		ms, err := g.Parse(f)

		if err != nil {
			return nil, err
		}
		models = append(models, ms...)

	}
	return models, nil
}
