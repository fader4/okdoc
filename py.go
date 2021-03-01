package okdoc

import (
	"fmt"

	"github.com/pkg/errors"
)

func PyParse(dat []byte) (*ReportFile, error) {
	lex, err := pyLex(dat)
	if err != nil {
		return nil, errors.Wrap(err, "tokenization error")
	}
	lexForParser := &pyLexerTokenIter{lex: lex, onlyTokens: []string{
		"bracket",
		"string",
		"comment",
		"keyword",
		"op_and_punct",
		"ident",
	}}
	out := pyParse(lexForParser)
	if out != 0 {
		return nil, fmt.Errorf("parser error: %v", lexForParser.errors)
	}
	return lexForParser.res, nil
}

var _ reporter = (*pyLexerTokenIter)(nil)

type pyLexerTokenIter struct {
	// associated tokenizer
	lex *lexer
	// iterator of tokens
	iter int
	// filter tokens interesting to the parser by labels
	onlyTokens []string
	errors     []string

	res *ReportFile
}

func (l *pyLexerTokenIter) SetReport(r *ReportFile) {
	l.res = r
}

type reporter interface {
	SetReport(*ReportFile)
}

func (l *pyLexerTokenIter) Lex(out *pySymType) (symbol int) {

skipToken:
	if !l.hasNext() {
		// EOF
		return 0
	}

	token := l.next()

	// if sets filter and not matched - skip token for current parser
	if len(l.onlyTokens) != 0 && !token.matchAtLeastOneLabels(l.onlyTokens...) {
		goto skipToken
	}

	out.token = &tokenWithLexer{token: token, lex: l.lex}

	if pyDebug == 1 {
		firstChar := token.pos[1] == token.pos[2]

		lineInfo := fmt.Sprintf("\t>%d:\t", token.pos[1])
		if firstChar {
			lineInfo = fmt.Sprintf("L%dT%d:", token.pos[0], token.pos[2])
		}
		charInfo := fmt.Sprintf("%q", string(token.Read(l.lex)))
		fmt.Println(
			"[PARSER]",
			lineInfo,
			charInfo,
		)
	}
	return token.symbol
}

func (l *pyLexerTokenIter) next() *token {
	if !l.hasNext() {
		return nil
	}
	token := l.lex.tokens[l.iter]
	l.iter++
	return token
}

func (l *pyLexerTokenIter) hasNext() bool {
	if l.iter == len(l.lex.tokens) {
		return false
	}
	return true
}

func (l *pyLexerTokenIter) current() *token {
	return l.lex.tokens[l.iter]
}

func (l *pyLexerTokenIter) Error(msg string) {
	token := l.current()
	msg = fmt.Sprintf(
		"parser error #%d: err=%q before %q",
		l.iter,
		msg,
		string(token.Read(l.lex)),
	)
	l.errors = append(l.errors, msg)
}
