package schema

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/cristianchaparroa/auto/connection"
	"github.com/cristianchaparroa/auto/generator"
	"github.com/cristianchaparroa/auto/meta"
	log "github.com/sirupsen/logrus"
)

// ParallelManager is an implementation that process all in parallel
type ParallelManager struct {
	Config     *connection.Config
	Connection *sql.DB
}

// NewParallelManager generates a pointer to ParallelManager
func NewParallelManager(c *connection.Config) *ParallelManager {
	return &ParallelManager{Config: c}
}

// CreateTable insert a table into database
func (m *ParallelManager) CreateTable(ms *meta.ModelStruct, queue chan *PartialResult) {
	pr := &PartialResult{}

	log.Info(fmt.Sprintf("auto:processing the model: %s \n", ms.ModelName))

	tb := generator.NewTableBuilder()
	tg := tb.GetTableGenerator(m.Config.Driver)

	tableResult, err := tg.Generate(ms)
	log.Info(fmt.Sprintf("auto:The sql generated is:%s \n", tableResult.SqlResult))

	if err != nil {
		pr.Error = err
		queue <- pr
		return
	}

	c := m.Connection
	res, err := c.Exec(tableResult.SqlResult)

	if err != nil {
		pr.Error = err
		queue <- pr
		return
	}

	pr.Model = ms
	pr.SqlExecuted = tableResult.SqlResult
	pr.Res = res
	pr.Relations = tableResult.Relations
	log.Info(fmt.Sprintf("auto:processing the model: %s was finnished \n\n", ms.ModelName))

	queue <- pr
}

// CreateTables create tables an returns all errors processing all model structs
func (m *ParallelManager) CreateTables(models []*meta.ModelStruct) ([]*PartialResult, []error) {

	queue := make(chan *PartialResult, 1)

	var wg sync.WaitGroup
	wg.Add(len(models))

	for _, model := range models {
		go m.CreateTable(model, queue)
	}

	results := make([]*PartialResult, 0)

	go func() {
		for res := range queue {
			results = append(results, res)
			wg.Done()
		}
	}()

	wg.Wait()

	errors := make([]error, 0)

	for _, resp := range results {
		if resp.Error != nil {
			errors = append(errors, resp.Error)
		}
	}

	return results, errors
}
