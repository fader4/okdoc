package okdoc

import (
	"fmt"

	"github.com/pkg/errors"
)

func PyParse(dat []byte) error {
	lex, err := pyLex(dat)
	if err != nil {
		return errors.Wrap(err, "tokenization error")
	}
	lexForParser := &pyLexerTokenIter{lex: lex}
	out := pyParse(lexForParser)
	if out != 0 {
		return fmt.Errorf("parser error: %v", lexForParser.errors)
	}
	return nil
}

type pyLexerTokenIter struct {
	lex    *lexer
	iter   int
	errors []string
}

func (l *pyLexerTokenIter) Lex(out *pySymType) (symbol int) {
	if !l.hasNext() {
		// EOF
		return 0
	}

	currentIdx := l.iter
	token := l.next()
	if pyDebug == 1 {
		fmt.Printf(
			"[PARSER] next token #%d: %q\n",
			currentIdx+1,
			string(token.Read(l.lex.data)),
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
		"parser error #%d: err=%q for %q",
		l.iter,
		msg,
		string(token.Read(l.lex.data)),
	)
	l.errors = append(l.errors, msg)
}
