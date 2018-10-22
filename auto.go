package auto

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cristianchaparroa/auto/connection"
	"github.com/cristianchaparroa/auto/meta"
	"github.com/cristianchaparroa/auto/parser"
	"github.com/cristianchaparroa/auto/scanner"
	"github.com/cristianchaparroa/auto/schema"
)

// Auto defines the interface to Auto tool
type Auto interface {
	// Generate tables from models
	Generate(path string) error

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
func (g *Generator) Generate(path, driver, host, user, pass, database string, port int) error {

	// 1. Scan the paths
	// 2. Read the files
	// 3. parse the files
	// 4. get a database managr
	// 5. connect with respective database
	// 6. execute all changes scanned
	paths, err := g.Scan(path)

	if err != nil {
		log.Println(err)
		return err
	}

	bs, err := g.ReadAll(paths)

	if err != nil {
		panic(err)
	}

	_, err = g.ParseAll(bs)

	if err != nil {
		panic(err)
	}

	c := connection.Config{Driver: driver, Host: host, Port: port, User: user, Password: pass}
	dm := schema.NewDatabaseManager(c)

	return dm.Execute()
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
		fmt.Print(err)
	}
	return b, err
}

// ReadAll  files in the models directory
func (g *Generator) ReadAll(paths []string) ([][]byte, error) {

	files := make([][]byte, 0)

	for _, p := range paths {

		bs, err := g.ReadFile(p)
		if err != nil {
			log.Printf("Is not possible to read the file:%v, get the error:%v ", p, err)
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
