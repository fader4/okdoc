package starlark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fader4/okdoc/syntax"
)

type Comment struct {
	*syntax.TokenWithData
	Multiline bool
}

func (r *Comment) MarshalJSON() ([]byte, error) {
	res := map[string]interface{}{
		"kind": "comment",
	}
	return json.MarshalIndent(res, " ", " ")
}

type Return struct {
	*syntax.TokenWithData
}

func (r *Return) MarshalJSON() ([]byte, error) {
	res := map[string]interface{}{
		"kind": "return",
	}
	return json.MarshalIndent(res, " ", " ")
}

type Def struct {
	*syntax.TokenWithData

	Name   syntax.Ident_
	Fields []*DefField
}

func (d *Def) MarshalJSON() ([]byte, error) {
	fields := []interface{}{}
	for _, field := range d.Fields {
		if field.Value != nil {
			fields = append(fields, syntax.Array{field.Key.String(), field.Value})
		}
		if field.ValueExpr != nil {
			fields = append(fields, syntax.Array{field.Key.String(), field.ValueExpr})
		}

	}
	res := map[string]interface{}{
		"kind":   "def",
		"name":   d.Name,
		"fields": fields,
	}
	return json.MarshalIndent(res, " ", " ")
}

type DefField struct {
	Key       *Operand
	Value     *Operand
	ValueExpr *Assign
	Varargs   bool
	Kwargs    bool
}

type Load struct {
	*syntax.TokenWithData

	Fields syntax.Array
}

func (l *Load) MarshalJSON() ([]byte, error) {
	fields := []interface{}{}
	for _, field := range l.Fields {
		fields = append(fields, field)
	}

	res := map[string]interface{}{
		"kind":   "load",
		"fields": fields,
	}
	return json.MarshalIndent(res, " ", " ")
}

type Module struct {
	*syntax.TokenWithData

	Export syntax.Ident_
	Name   syntax.StringLiteral
	Fields syntax.Array
}

func (m *Module) MarshalJSON() ([]byte, error) {
	fields := []interface{}{}
	for _, field := range m.Fields {
		fields = append(fields, field)
	}
	res := map[string]interface{}{
		"kind": "module",

		"export": m.Export.String(),
		"name":   m.Name,

		"fields": fields,
	}
	return json.MarshalIndent(res, " ", " ")
}

type CallFunc struct {
	Name syntax.Ident_
	// free fields => ((key1, val1), (key2, null), (key2, Assign) ...)
	// *vararrgs => (val1, val2, Assign, ...)
	// *kwargs => ((key1, val1), (key2, val2), (key3, Assign)...)
	//
	// Operand and mixed Assign fields
	Fields syntax.Array

	Varrarrgs bool
	Kwarrgs   bool
}

func (f CallFunc) FuncName() string {
	// без @
	return f.Name.String()[1:]
}

type Operand struct {
	Parent *Operand
	Value  syntax.Value
}

func reverseValues(in []syntax.Value) []syntax.Value {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}

func (o *Operand) Values() (res []syntax.Value) {
	next := o
	for {
		if next == nil {
			break
		}
		res = append(res, next.Value)
		next = next.Parent
	}
	return reverseValues(res)
}

func (o *Operand) String() string {
	buf := &bytes.Buffer{}
	for _, val := range o.Values() {
		switch v := val.(type) {
		case syntax.Ident_:
			buf.WriteString("." + v.String())
		case syntax.StringLiteral, syntax.IntLiteral, syntax.BoolLiteral, syntax.FloatLiteral:
			fmt.Fprintf(buf, ".%v", v)
		case *CallFunc:
			buf.WriteString("." + v.FuncName() + "(...)")
		case syntax.Array:
			buf.WriteString(".ListOrTuple")
		case syntax.Map:
			buf.WriteString(".Dict")
		default:
			fmt.Fprintf(buf, ".%T", v)
		}
	}
	res := buf.String()
	if res == "" {
		return res
	}
	return res[1:]
}

type Assign struct {
	Parent *Assign

	Left  *Operand
	Op    []int
	Right *Operand
}

func (a *Assign) String() string {
	return "Expr(...)"
}
