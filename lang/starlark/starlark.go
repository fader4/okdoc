package starlark

import (
	"fmt"

	"github.com/fader4/okdoc/syntax"
)

func Parse(dat []byte) ([]*CompositeToken, error) {
	lex, err := newPreprocessing([]byte(dat))
	if err != nil {
		return nil, err
	}

	lexForParser := &starlarkLex{
		LexIter: syntax.NewLexIter(lex, "def", "comment", "return", "load", "module"),
		debug:   true,
	}

	var tokens = []*CompositeToken{}
	var startToken *syntax.Token
	container := &starlarkSymType{}
	for {
		symbol := lexForParser.Lex(container)
		if symbol == 0 {
			break
		}

		switch symbol {
		case returnKeyword:
			newToken := &CompositeToken{
				Start: container.token.Token,
				End:   container.token.Token,
				lex:   lex,
				Name:  "return",
			}
			tokens = append(tokens, newToken)
		case commentInline:
			newToken := &CompositeToken{
				Start: container.token.Token,
				End:   container.token.Token,
				lex:   lex,
				Name:  "inline comment",
			}
			tokens = append(tokens, newToken)
		case beginMultilineComment, beginDef, beginLoad, beginModule:
			startToken = container.token.Token
		case endModule:
			newToken := &CompositeToken{
				Start: startToken,
				End:   container.token.Token,
				lex:   lex,
				Name:  "module",
			}
			tokens = append(tokens, newToken)
		case endMultilineComment:
			newToken := &CompositeToken{
				Start: startToken,
				End:   container.token.Token,
				lex:   lex,
				Name:  "multiline comment",
			}
			tokens = append(tokens, newToken)
		case endLoad:
			newToken := &CompositeToken{
				Start: startToken,
				End:   container.token.Token,
				lex:   lex,
				Name:  "load",
			}
			tokens = append(tokens, newToken)
		case endDef:
			newToken := &CompositeToken{
				Start: startToken,
				End:   container.token.Token,
				lex:   lex,
				Name:  "def",
			}
			tokens = append(tokens, newToken)
		}
	}

	return tokens, nil
}

// func Parse(dat []byte) error {
// 	lex, err := newTokenizer(dat)
// 	if err != nil {
// 		return errors.Wrap(err, "tokenization error")
// 	}
// 	interestedTokens := []string{
// 		"bracket",
// 		"op_and_punct",
// 		"ident",
// 		"literal",
// 		"comment",
// 	}
// 	lexer := &starlarkLex{
// 		LexIter: syntax.NewLexIter(lex, interestedTokens...),
// 		debug:   true,
// 	}
// 	out := starlarkParse(lexer)
// 	if out != 0 {
// 		return fmt.Errorf("parser error: %v", lexer.Errors)
// 	}
// 	return nil
// }

type starlarkLex struct {
	*syntax.LexIter
	debug bool
}

func (l *starlarkLex) Lex(out *starlarkSymType) (symbol int) {
	token := l.Next()
	if token == nil {
		return 0
	}
	out.token = &Token{token, l.LexIter.Lex()}

	if l.debug {
		tokenBytesm, err := token.HumanString(l.LexIter.Lex().Data)
		if err != nil {
			panic("failed render token")
		}
		fmt.Println(string(tokenBytesm))
	}
	return token.Symbol
}
