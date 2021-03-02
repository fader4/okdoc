package meta

import (
	"fmt"

	"github.com/fader4/okdoc/syntax"
	"github.com/pkg/errors"
)

func Parse(dat []byte) error {
	lex, err := newTokenizer(dat)
	if err != nil {
		return errors.Wrap(err, "tokenization error")
	}
	lexForParser := &lexerTokenIter{lex: lex, onlyTokens: []string{
		"bracket",
		"string",
		"op_and_punct",
		"ident",
	}}
	out := metaParse(lexForParser)
	if out != 0 {
		return fmt.Errorf("parser error: %v", lexForParser.errors)
	}

	return nil
}

type lexerTokenIter struct {
	// associated tokenizer
	lex *syntax.Lexer
	// iterator of tokens
	iter int
	// filter tokens interesting to the parser by labels
	onlyTokens []string
	errors     []string
}

func (l *lexerTokenIter) Lex(out *metaSymType) (symbol int) {
skipToken:
	if !l.hasNext() {
		// EOF
		return 0
	}

	token := l.next()

	fmt.Println("call Lex", token.Labels)

	// if sets filter and not matched - skip token for current parser
	if len(l.onlyTokens) != 0 && !token.MatchAtLeastOneLabels(l.onlyTokens...) {
		goto skipToken
	}

	out.token = token

	if metaDebug == 1 {
		firstChar := token.Pos[1] == token.Pos[2]

		lineInfo := fmt.Sprintf("\t>%d:\t", token.Pos[1])
		if firstChar {
			lineInfo = fmt.Sprintf("L%dSP%d:", token.Pos[0], token.Pos[2])
		}

		tokenBytes, _ := l.lex.Data.Render(token)
		charInfo := fmt.Sprintf("%q", string(tokenBytes))
		fmt.Println(
			"[PARSER]",
			lineInfo,
			charInfo,
		)
	}
	return token.Symbol
}

func (l *lexerTokenIter) next() *syntax.Token {
	if !l.hasNext() {
		return nil
	}
	token := l.lex.Tokens[l.iter]
	l.iter++
	return token
}

func (l *lexerTokenIter) hasNext() bool {
	if l.iter == len(l.lex.Tokens) {
		return false
	}
	return true
}

func (l *lexerTokenIter) current() *syntax.Token {
	return l.lex.Tokens[l.iter]
}

func (l *lexerTokenIter) Error(msg string) {
	token := l.current()
	tokenBytes, _ := l.lex.Data.Render(token)
	msg = fmt.Sprintf(
		"parser error #%d: err=%q before %q",
		l.iter,
		msg,
		string(tokenBytes),
	)
	l.errors = append(l.errors, msg)
}
