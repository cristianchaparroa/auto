package schema

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/cristianchaparroa/auto/connection"
	"github.com/cristianchaparroa/auto/meta"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestParalleManagerCreateTables(t *testing.T) {
	defer timeTrack(time.Now(), "TestManagerCreateTables")

	// 1. mocking the dependencies for connections
	config := &connection.Config{Driver: connection.PostgresDriver}
	db, mock, err := sqlmock.New()
	pm := NewParallelManager(config)
	pm.Connection = db

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 2. generating the test cases
	models := make([]*meta.ModelStruct, 0)

	for i := 0; i < 10000; i++ {
		name := fmt.Sprintf("Ebook%v", i)
		name = strings.ToUpper(name)

		// mock any call to sql.DB.Exec
		mock.ExpectExec(`\s*CREATE(.|\n)*?;`).WillReturnResult(sqlmock.NewResult(0, 0))
		e := &meta.ModelStruct{ModelName: name}
		models = append(models, e)
	}

	// 3. excuting the test
	results, errors := pm.CreateTables(models)

	expectedSize := len(models)
	resultSize := len(results)

	// 4. Comparing results
	if resultSize != expectedSize {
		t.Errorf("Expected a result of size :%v, but get:%v", expectedSize, resultSize)
	}

	if len(errors) > 0 {
		t.Errorf("Expected Zero(0) errors but get:%d", len(errors))
	}

	for _, err := range errors {
		fmt.Println(err)
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
