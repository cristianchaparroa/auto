package meta

import (
	"bytes"
	"fmt"
)

// Field represent the meta data  relevant related with field
type Field struct {
	// Field Name
	Name string
	// Field Type ej: int, string, ...
	Typ string
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
