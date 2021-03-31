package starlark

import (
	"fmt"

	"github.com/fader4/okdoc/syntax"
)

func Parse(dat []byte) ([]*syntax.CompositeToken, error) {
	lex, err := newPreprocessing([]byte(dat))
	if err != nil {
		return nil, err
	}

	lexForParser := &starlarkLex{
		LexIter: syntax.NewLexIter(lex, "def", "comment", "return", "load", "module"),
		debug:   true,
	}

	var tokens = []*syntax.CompositeToken{}
	var startToken *syntax.Token
	container := &starlarkSymType{}
	for {
		symbol := lexForParser.Lex(container)
		if symbol == 0 {
			break
		}

		switch symbol {
		case returnKeyword:
			newToken := &syntax.CompositeToken{
				Start:  container.token.Token,
				End:    container.token.Token,
				Lex:    lex,
				Symbol: returnKeyword,
			}
			tokens = append(tokens, newToken)
		case commentInline:
			newToken := &syntax.CompositeToken{
				Start:  container.token.Token,
				End:    container.token.Token,
				Lex:    lex,
				Symbol: commentInline,
			}
			tokens = append(tokens, newToken)
		case commentMultiline, def, load, module:
			startToken = container.token.Token
		case endModule:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    container.token.Token,
				Lex:    lex,
				Symbol: module,
			}
			tokens = append(tokens, newToken)
		case endCommentMultiline:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    container.token.Token,
				Lex:    lex,
				Symbol: commentMultiline,
			}
			tokens = append(tokens, newToken)
		case endLoad:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    container.token.Token,
				Lex:    lex,
				Symbol: load,
			}
			tokens = append(tokens, newToken)
		case endDef:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    container.token.Token,
				Lex:    lex,
				Symbol: def,
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
