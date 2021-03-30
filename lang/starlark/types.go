package starlark

import (
	"fmt"
	"log"

	"github.com/fader4/okdoc/syntax"
)

func fmtPrintln(in string) {
	fmt.Println(in)
}

type CompositeToken struct {
	Start, End *syntax.Token
	Name       string
	// ref to source data
	lex *syntax.Lexer
}

type Token struct {
	*syntax.Token
	*syntax.Lexer
}

func (a *CompositeToken) Token() *syntax.Token {
	return &syntax.Token{
		Symbol: a.Start.Symbol,

		Start: a.Start.Start,
		End:   a.End.End,
		Pos:   a.Start.Pos,
	}
}

func (t *CompositeToken) MustBytes() []byte {
	dat, err := t.Token().Bytes(t.lex.Data)
	if err != nil {
		log.Println("starlark.CompositeToken#MustBytes failed get data:", err)
		return []byte{}
	}
	return dat
}

func (t *CompositeToken) HumanString() (string, error) {
	return t.Token().HumanString(t.lex.Data)
}
