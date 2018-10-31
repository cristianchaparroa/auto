package meta

import (
	"bytes"
	"fmt"
)

// TypeTag represent the type of tag
type TypeTag string

func (tg *TypeTag) String() string {
	return fmt.Sprintf("%s", *tg)
}

// Tag represents the tag inside of sql tag
type Tag struct {
	Value string
	Typ   *TypeTag
}

func (t *Tag) String() string {
	return fmt.Sprintf("Tag[value:%s, type:%v]", t.Value, t.Typ)
}

// TypeField contains the data type of field
type TypeField struct {

	// Name should be string, int, array or the name of entity
	Name string

	// EntityRelated to Array
	EntityRelated string
}

func (f TypeField) String() string {
	return fmt.Sprintf("meta.TypeField(Name:%v, EntityRelated:%v)", f.Name, f.EntityRelated)
}

// Field represent the meta data  relevant related with field
type Field struct {
	// Field Name
	Name string
	// Field Type ej: int, string, ...
	Typ *TypeField
	// Tags annotated in the field
	Tags []string
}

func (f Field) String() string {
	return fmt.Sprintf("meta.Field(Name:%v, Type:%v ,Tags:%v)", f.Name, f.Typ, f.Tags)
}

//ModelStruct represent the meta data extracted from model
type ModelStruct struct {
	// This is the model name or struct name.
	ModelName string
	// This is the path in wich was readed the struct model
	ModelPath string
	// This is the package name in which was read the struct model
	PackageName string
	// This is the number of fields in the struct model
	NumberOfFields string

	Fields []*Field
}

func (m ModelStruct) String() string {

	var buffer bytes.Buffer

	for _, f := range m.Fields {
		buffer.WriteString(fmt.Sprintf("\t\t%v\n", f))
	}

	return fmt.Sprintf("\nmeta.ModelStruct(ModelName:%v, ModelPath:%v, PackageName:%v, \n \tFields:[\n%v\t]\n)\n",
		m.ModelName, m.ModelPath, m.PackageName, buffer.String())
}
