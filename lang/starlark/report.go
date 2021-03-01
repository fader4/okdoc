package starlark

type Report struct {
	Files []*ReportFile
}
type ReportFile struct {
	Filename  string
	Imports   []*ReportFile_Import
	Exports   []*ReportFile_Export
	Functions []*ReportFile_Function

	Comments []*ReportFile_Comment
	Returns  []*ReportFile_Return
}

type ReportFile_Comment struct {
	Value     string
	MultiLine bool

	Pos pos
}

type ReportFile_Import struct {
	// vendor name
	Name string
	// extracts one or more values from imported module
	Args []ArgumentNode

	Pos pos
}

type ReportFile_Export struct {
	// module name
	Name string
	// exporting one or more values from current module
	Args []ArgumentNode

	Pos pos
}

type ReportFile_Function struct {
	Name    string
	Comment *ReportFile_Comment

	Pos          pos
	PosOfReturns []pos
}

type ReportFile_Return struct {
	Pos pos
}
