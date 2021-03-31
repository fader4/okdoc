package starlark

import (
	"fmt"

	"github.com/fader4/okdoc/syntax"
	"github.com/pkg/errors"
)

type Token struct {
	*syntax.CompositeToken
}

func Parse(dat []byte) ([]*Token, error) {
	lex, err := newPreprocessing([]byte(dat))
	if err != nil {
		return nil, err
	}

	iter := syntax.NewLexIter(lex, "def", "comment", "return", "load", "module")

	var tokens = []*Token{}
	var startToken *syntax.Token

	for {
		token := iter.Next()
		if token == nil {
			break
		}

		switch token.Symbol {
		case commentMultiline, def, load, module:
			startToken = token

		case returnKeyword:
			newToken := &syntax.CompositeToken{
				Start:  token,
				End:    token,
				Lex:    lex,
				Symbol: returnKeyword,
			}
			tokens = append(tokens, &Token{newToken})
		case commentInline:
			newToken := &syntax.CompositeToken{
				Start:  token,
				End:    token,
				Lex:    lex,
				Symbol: commentInline,
			}
			tokens = append(tokens, &Token{newToken})
		case endModule:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: module,
			}
			tokens = append(tokens, &Token{newToken})
		case endCommentMultiline:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: commentMultiline,
			}
			tokens = append(tokens, &Token{newToken})
		case endLoad:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: load,
			}
			tokens = append(tokens, &Token{newToken})
		case endDef:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: def,
			}
			tokens = append(tokens, &Token{newToken})
		}
	}

	return tokens, nil
}

func (t *Token) Parse() error {
	lex, err := newTokenizer(t.MustBytes())
	if err != nil {
		return errors.Wrap(err, "tokenization error")
	}
	interestedTokens := []string{
		"bracket",
		"literal",
		"ident",
		"op_and_punct",
		"dict",

		"comment",
		"def",
		"load",
		"module",
		"return",
	}
	lexer := &starlarkLex{
		LexIter: syntax.NewLexIter(lex, interestedTokens...),
		debug:   true,
		// refToken: t,
	}
	out := starlarkParse(lexer)
	if out != 0 {
		return fmt.Errorf("parser error: %v", lexer.Errors)
	}
	return nil
}

type starlarkLex struct {
	*syntax.LexIter
	debug bool
}

func (l *starlarkLex) Lex(out *starlarkSymType) (symbol int) {
	token := l.Next()
	if token == nil {
		return 0
	}
	out.token = &syntax.TokenWithData{token, l.LexIter.Lex()}

	if l.debug {
		tokenBytesm, err := token.HumanString(l.LexIter.Lex().Data)
		if err != nil {
			panic("failed render token")
		}
		fmt.Println(string(tokenBytesm))
	}
	return token.Symbol
}
