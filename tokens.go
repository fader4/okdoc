package okdoc

// Token it is recognized character group. Used by the parser.
type token struct {
	// symbol understandable to parser
	symbol int
	// start and end of data entering the symbol
	start, end int

	// file position - file line
	// 1 - line number
	// 2 - number of characters from the beginning of the line
	// 3 - number of white spaces from the beginning of the line
	pos [3]int

	skipParser bool
}

func (t token) Read(lex *lexer) []byte {
	return lex.data[t.start:t.end]
}

type tokens []*token

func (t *tokens) release(in ...*token) int {
	*t = append(*t, in...)
	return len(*t) - 1
}

type node struct {
	tokens tokens
}

type tokenWithData struct {
	*token
	lex *lexer
}

func (t *tokenWithData) String() string {
	return string(t.Read(t.lex))
}
