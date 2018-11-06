package schema

import (
	"database/sql"

	"github.com/cristianchaparroa/auto/meta"
)

// PartialResult is the first result of table created without relations
type PartialResult struct {
	// Res is the result after execute the sql generated
	Res sql.Result
	// Errors is in case of some error when is called the createTable process
	Error error

	// Model is the model processing
	Model *meta.ModelStruct

	// SqlExecuted is the sql generated without relations
	SqlExecuted string
	// Reltations are the fields that contains relations with other models
	// These relations must be proccesed after all tables witout dependeces are creted
	Relations []*meta.Field
}
