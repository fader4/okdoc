package okdoc

// Token it is recognized character group. Used by the parser.
type token struct {
	// symbol understandable to parser
	symbol int
	// start and end of data entering the symbol
	start, end int
}

func (t token) Read(dat []byte) []byte {
	return dat[t.start:t.end]
}

type tokens []*token

func (t *tokens) release(in ...*token) {
	*t = append(*t, in...)
}

type node struct {
	tokens tokens
}
