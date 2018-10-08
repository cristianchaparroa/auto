package parser

const (

	// PackageParserEmptyFile  error
	PackageParserEmptyFile string = "Error trying to retrieve the package name from empty file"

	//PackageParserEmptyLine error
	PackageParserEmptyLine string = "Error trying to retrieve the package name, the first line is empty"

	//PacakgeParserMalformedTwoTokensExpected error
	PacakgeParserMalformedTwoTokensExpected string = "Malformed definition of package statement, expected two tokens but get:%v (%v)"

	//PackageParserMalformedStmtPackage error
	PackageParserMalformedStmtPackage string = `Malformed defintion of pacakge statement, expected reserved word "package" but get:%v`

	//StructParserMalformedPatterNotFound error
	StructParserMalformedPatterNotFound string = "Malformed definition of struct statment. Struct pattern not found to be parsed"
	//StructParserMalformedTypeNotFound error
	StructParserMalformedTypeNotFound string = "Malformed definition of struct statement, It's not contains the reserved word type"

	// StructParserMalformedStructNotFound error
	StructParserMalformedStructNotFound string = "Malformed definition of struct statement, It's not contains the reserved word struct "

	// StructParserMalformedNameNotfound error
	StructParserMalformedNameNotfound string = "Malformed definition of struct statement, It's not contains the name of struct"

	// ValidatorValidateBodyError error
	ValidatorValidateBodyError string = "The body struct contains characters not allowed:%v"
)
