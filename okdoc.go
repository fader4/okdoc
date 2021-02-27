package okdoc

import "errors"

func Parse(in []byte) (*Content, error) {
	p := &Parser{lexer: newLexer(in)}
	e := yyNewParser().Parse(p)
	if e != 0 {
		return p.Content, errors.New("invalid format")
	}
	return p.Content, nil
}

type Parser struct {
	*lexer

	Content *Content
}
