package syntax

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type CompositeToken struct {
	Symbol     int
	Start, End *Token

	// ref to source data
	Lex *Lexer
}

func (a *CompositeToken) Token() *Token {
	return &Token{
		Symbol: a.Symbol,
		Start:  a.Start.Start,
		End:    a.End.End,
		Pos:    a.Start.Pos,
	}
}

func (t *CompositeToken) TokenWithData() *TokenWithData {
	return &TokenWithData{
		Token: t.Token(),
		Lexer: t.Lex,
	}
}

func (t *CompositeToken) MustBytes() []byte {
	dat, err := t.Token().Bytes(t.Lex.Data)
	if err != nil {
		log.Println("starlark.CompositeToken#MustBytes failed get data:", err)
		return []byte{}
	}
	return dat
}

func (t *CompositeToken) HumanString() (string, error) {
	return t.Token().HumanString(t.Lex.Data)
}

type TokenWithData struct {
	*Token
	*Lexer
}

func (t *TokenWithData) MustBytes() []byte {
	dat, err := t.Token.Bytes(t.Lexer.Data)
	if err != nil {
		log.Println("annotation.Token#MustBytes failed get data:", err)
		return []byte{}
	}
	return dat
}

func (t *TokenWithData) Ident() Ident_ {
	return Ident(t.MustBytes())
}

func (t *TokenWithData) String() StringLiteral {
	return String(t.MustBytes())
}

func (t *TokenWithData) Int() IntLiteral {
	return Int(t.MustBytes())
}

func (t *TokenWithData) Null() Null_ {
	return Null_{}
}

func (t *TokenWithData) Float() FloatLiteral {
	return Float(t.MustBytes())
}

func (t *TokenWithData) Bool() BoolLiteral {
	return Bool(t.MustBytes())
}

type Value interface{}

func Ident(in []byte) Ident_ {
	return Ident_([]string{string(in)})
}

type Ident_ []string

func (i Ident_) Append(in Ident_) Ident_ {
	return append(i, in...)
}

func (m Ident_) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.Join(m, "."))
}

func (m Ident_) String() string {
	return strings.Join(m, ".")
}

func String(in []byte) StringLiteral {
	if string(in) == `""` || string(in) == `''` {
		return StringLiteral(string(""))
	}
	return StringLiteral(string(in))
}

func Null(in []byte) Null_ {
	return Null_{}
}

type Null_ struct{}

func (a Null_) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

type StringLiteral string

func Int(in []byte) IntLiteral {
	i, err := strconv.ParseInt(string(in), 10, 64)
	if err != nil {
		log.Printf("annotation.IntLiteral: invalid input data %q, err = %v\n", string(in), err)
		return IntLiteral(0)
	}
	return IntLiteral(i)
}

type IntLiteral int64

func Float(in []byte) FloatLiteral {
	i, err := strconv.ParseFloat(string(in), 64)
	if err != nil {
		log.Printf("annotation.FloatLiteral: invalid input data %q, err = %v\n", string(in), err)
		return FloatLiteral(0)
	}
	return FloatLiteral(i)
}

type FloatLiteral float64

func Bool(in []byte) BoolLiteral {
	i, err := strconv.ParseBool(string(in))
	if err != nil {
		log.Printf("annotation.FloatLiteral: invalid input data %q, err = %v\n", string(in), err)
		return BoolLiteral(false)
	}
	return BoolLiteral(i)
}

type BoolLiteral bool

type Array []Value

func (a *Array) Add(in ...Value) {
	*a = append(*a, in...)
}

// func (a Array) MarshalJSON() ([]byte, error) {
// 	res := []interface{}{}
// 	for _, item := range a {
// 		res = append(res, item)
// 	}
// 	return json.Marshal(res)
// }

type Map struct {
	Keys   []string
	Values Array
}

func (m Map) MarshalJSON() ([]byte, error) {
	res := map[string]interface{}{}
	for i, key := range m.Keys {
		res[key] = m.Values[i]
	}
	return json.Marshal(res)
}

func (m *Map) Add(key string, val Value) {
	m.Keys = append(m.Keys, key)
	m.Values = append(m.Values, val)
}
