package starlark

// Token it is recognized character group. Used by the parser.
type token struct {
	// symbol understandable to parser
	symbol int
	// start and end of data entering the symbol
	start, end int

	// name of the group
	labels map[string]bool

	// file position
	pos pos
}

func (t *token) addLabel(labels ...string) {
	if t.labels == nil {
		t.labels = map[string]bool{}
	}
	for _, label := range labels {
		t.labels[label] = true
	}
}

func (t *token) matchAtLeastOneLabels(labels ...string) bool {
	for _, label := range labels {
		if t.labels[label] {
			return true
		}
	}
	return false
}

func (t token) Read(lex *lexer) []byte {
	return lex.data[t.start:t.end]
}

type tokens []*token

func (t *tokens) release(in ...*token) int {
	*t = append(*t, in...)
	return len(*t) - 1
}

type tokenWithLexer struct {
	*token
	lex *lexer
}

func (t *tokenWithLexer) Pos() pos {
	return t.pos
}

func (t *tokenWithLexer) String() string {
	return string(t.Read(t.lex))
}
