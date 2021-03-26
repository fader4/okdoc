package annotation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TypesMarshal(t *testing.T) {
	m1 := Map{
		Keys: []string{
			"str",
			"bool",
			"int",
			"float",
			"null",
			"arr",
		},
		Values: []Value{
			StringLiteral("a"),
			BoolLiteral(true),
			IntLiteral(123),
			FloatLiteral(123.123),
			Null_{},
			Array{
				StringLiteral("a"),
				BoolLiteral(true),
				IntLiteral(123),
				FloatLiteral(123.123),
				Null_{},
				Array{
					StringLiteral("a"),
					BoolLiteral(true),
					IntLiteral(123),
					FloatLiteral(123.123),
					Null_{},
				},
			},
		},
	}
	cases := []struct {
		in       Value
		expected string
	}{
		{Ident_{"a"}, `{"$ident":"a"}`},
		{Ident_{"a", "b"}, `{"$ident":"a.b"}`},
		{Null_{}, `null`},
		{StringLiteral("a"), `"a"`},
		{BoolLiteral(true), `true`},
		{IntLiteral(123), `123`},
		{FloatLiteral(123.123), `123.123`},
		{Array{
			StringLiteral("a"),
			Null_{},
			BoolLiteral(true),
			IntLiteral(123),
			FloatLiteral(123.123),
		}, `["a",null,true,123,123.123]`},
		{Array{
			StringLiteral("a"),
			BoolLiteral(true),
			IntLiteral(123),
			FloatLiteral(123.123),
			Array{
				StringLiteral("a"),
				BoolLiteral(true),
				IntLiteral(123),
				FloatLiteral(123.123),
			},
		}, `["a",true,123,123.123,["a",true,123,123.123]]`},
		{m1, `{"arr":["a",true,123,123.123,null,["a",true,123,123.123,null]],"bool":true,"float":123.123,"int":123,"null":null,"str":"a"}`},
		{Map{
			Keys: []string{
				"m2",
			},
			Values: []Value{
				m1,
			},
		}, `{"m2":{"arr":["a",true,123,123.123,null,["a",true,123,123.123,null]],"bool":true,"float":123.123,"int":123,"null":null,"str":"a"}}`},
	}

	for _, case_ := range cases {
		t.Run("", func(t *testing.T) {
			res, err := json.Marshal(case_.in)
			assert.NoError(t, err)
			assert.EqualValues(t, case_.expected, string(res))
		})
	}

}
