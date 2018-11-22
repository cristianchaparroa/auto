package generator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cristianchaparroa/auto/meta"
)

// PostgresTable generates a query for create a table in postgres.
type PostgresTable struct {
}

// NewPostgresTable returns a pointer to PostgresTable
func NewPostgresTable() *PostgresTable {
	return &PostgresTable{}
}

// Generate the sql query to create a table for postgres
func (g *PostgresTable) Generate(m *meta.ModelStruct) (ITableResult, error) {

	result := &TableResult{}

	tableName := strings.ToUpper(m.ModelName)
	sql := fmt.Sprintf("CREATE TABLE %v ();", tableName)

	var buffer bytes.Buffer
	buffer.WriteString(sql)

	fs := m.Fields

	cg := NewPostgresColumn()

	relations := make([]*meta.Field, 0)

	for _, f := range fs {

		if f.IsRelation {
			relations = append(relations, f)
			continue
		}

		sqlField, _ := cg.Create(m.ModelName, f)
		buffer.WriteString(fmt.Sprintf("\n%s;", sqlField))
	}

	result.Relations = relations
	result.SQLResult = buffer.String()
	return result, nil
}

// CreateRelation generates a relation between to models
func (g *PostgresTable) CreateRelation(model *meta.ModelStruct, field *meta.Field) (string, error) {

	r := field.Relation

	if r.Typ == meta.OneToOne {
		return g.CreateOneToOne(model, field)
	}

	return "", nil
}

// CreateOneToOne generate the sql that generates the field and respective constraint
func (g *PostgresTable) CreateOneToOne(model *meta.ModelStruct, field *meta.Field) (string, error) {

	pc := NewPostgresColumn()

	parentTable := model.ModelName
	childTable := field.Relation.To

	pkParentTable := model.GetPrimaryKey()
	pkChildTable := field.Relation.PKRef

	constraintName := ""

	sqlField, err := pc.Create(parentTable, field)

	if err != nil {
		return "", err
	}

	sql := `ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%) REFERENCES %s (%s)`
	sqlConstraint := fmt.Sprintf(sql, childTable, constraintName, pkChildTable, parentTable, pkParentTable)

	return fmt.Sprintf("%s \n %s", sqlField, sqlConstraint), nil
}
