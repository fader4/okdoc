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

skipToken:
	if !l.hasNext() {
		// EOF
		return 0
	}

	token := l.next()
	if token.skipParser {
		goto skipToken
	}

	out.token = &tokenWithData{token: token, lex: l.lex}

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
