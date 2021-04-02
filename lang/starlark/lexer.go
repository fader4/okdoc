//line lexer.rl:1
package starlark

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
)

//line lexer.rl:171

//line lexer.go:16
const starlark_start int = 12
const starlark_first_final int = 12
const starlark_error int = 0

const starlark_en_singleQuoteString int = 55
const starlark_en_doubleQuoteString int = 57
const starlark_en_main int = 12

//line lexer.rl:174

func newTokenizer(data []byte) (*syntax.Lexer, error) {
	lex := syntax.NewLexer(data)

//line lexer.go:32
	{
		(lex.Cs) = starlark_start
		(lex.Top) = 0
		(lex.Ts) = 0
		(lex.Te) = 0
		(lex.Act) = 0
	}

//line lexer.rl:179

//line lexer.go:43
	{
		if (lex.P) == (lex.Pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.Cs {
		case 12:
			goto st12
		case 13:
			goto st13
		case 1:
			goto st1
		case 14:
			goto st14
		case 2:
			goto st2
		case 3:
			goto st3
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 4:
			goto st4
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 8:
			goto st8
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 9:
			goto st9
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto st40
		case 41:
			goto st41
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 45:
			goto st45
		case 46:
			goto st46
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 10:
			goto st10
		case 57:
			goto st57
		case 58:
			goto st58
		case 11:
			goto st11
		case 0:
			goto st0
		}

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof
		}
	_resume:
		switch lex.Cs {
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 1:
			goto st_case_1
		case 14:
			goto st_case_14
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 9:
			goto st_case_9
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 10:
			goto st_case_10
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 11:
			goto st_case_11
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line NONE:1
		switch lex.Act {
		case 9:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(def, "def")
			}
		case 10:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(not, "not")
			}
		case 11:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(module, "module")
			}
		case 12:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(load, "load")
			}
		case 13:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(returnKeyword, "return")
			}
		case 15:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseSymbol("op_and_punct")
			}
		case 17:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(or, "op_and_punct")
			}
		case 18:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(and, "op_and_punct")
			}
		case 19:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(nullLiteral, "literal")
			}
		case 20:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(ident, "ident")
			}
		case 21:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(boolLiteral, "literal", "bool")
			}
		default:
			{
				(lex.P) = (lex.Te) - 1
			}
		}

		goto st12
	tr2:
//line lexer.rl:129
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseToken(floatLiteral, "literal", "float")
		}
		goto st12
	tr5:
//line lexer.rl:79
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st12
	tr11:
//line lexer.rl:101
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(commentMultiline, "comment")
		}
		goto st12
	tr13:
//line lexer.rl:104
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st12
	tr20:
//line lexer.rl:165
		(lex.Te) = (lex.P) + 1

		goto st12
	tr21:
//line lexer.rl:25

		lex.ReleaseSymbol("SP")

//line lexer.rl:76
		(lex.Te) = (lex.P) + 1

		goto st12
	tr22:
//line lexer.rl:22

		lex.ReleaseSymbol("CR")

//line lexer.rl:77
		(lex.Te) = (lex.P) + 1

		goto st12
	tr25:
//line lexer.rl:104
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st12
	tr27:
//line lexer.rl:143
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('(', "bracket", "open_bracket")
			lex.BeginPairedChar(')')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 12
					(lex.Top)++
					goto st12
				}
			}
		}
		goto st12
	tr28:
//line lexer.rl:158
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
				lex.ReleaseSymbol("bracket", "close_bracket")
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
		}
		goto st12
	tr37:
//line lexer.rl:153
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('[', "bracket", "open_bracket")
			lex.BeginPairedChar(']')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 12
					(lex.Top)++
					goto st12
				}
			}
		}
		goto st12
	tr45:
//line lexer.rl:148
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('{', "bracket", "open_bracket")
			lex.BeginPairedChar('}')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 12
					(lex.Top)++
					goto st12
				}
			}
		}
		goto st12
	tr46:
//line lexer.rl:165
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st12
	tr47:
//line lexer.rl:129
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(floatLiteral, "literal", "float")
		}
		goto st12
	tr49:
//line lexer.rl:138
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 12
					(lex.Top)++
					goto st57
				}
			}
		}
		goto st12
	tr51:
//line lexer.rl:79
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st12
	tr52:
//line lexer.rl:83
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(commentInline, "comment")
		}
		goto st12
	tr53:
//line lexer.rl:133
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 12
					(lex.Top)++
					goto st55
				}
			}
		}
		goto st12
	tr55:
//line lexer.rl:104
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st12
	tr57:
//line lexer.rl:126
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(integerLiteral, "literal", "int")
		}
		goto st12
	tr58:
//line lexer.rl:120
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(ident, "ident")
		}
		goto st12
	st12:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:539
		switch data[(lex.P)] {
		case 0:
			goto tr19
		case 9:
			goto tr21
		case 10:
			goto tr22
		case 13:
			goto tr22
		case 32:
			goto tr21
		case 34:
			goto st16
		case 35:
			goto st18
		case 37:
			goto tr25
		case 39:
			goto st19
		case 40:
			goto tr27
		case 41:
			goto tr28
		case 45:
			goto tr29
		case 46:
			goto st24
		case 48:
			goto st22
		case 58:
			goto tr25
		case 70:
			goto st26
		case 78:
			goto st30
		case 84:
			goto st33
		case 91:
			goto tr37
		case 93:
			goto tr28
		case 95:
			goto tr33
		case 97:
			goto st35
		case 100:
			goto st37
		case 108:
			goto st39
		case 109:
			goto st42
		case 110:
			goto st47
		case 111:
			goto st49
		case 114:
			goto st50
		case 123:
			goto tr45
		case 125:
			goto tr28
		}
		switch {
		case data[(lex.P)] < 60:
			switch {
			case data[(lex.P)] > 47:
				if 49 <= data[(lex.P)] && data[(lex.P)] <= 57 {
					goto st23
				}
			case data[(lex.P)] >= 42:
				goto tr25
			}
		case data[(lex.P)] > 62:
			switch {
			case data[(lex.P)] > 90:
				if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
					goto tr33
				}
			case data[(lex.P)] >= 65:
				goto tr33
			}
		default:
			goto tr25
		}
		goto tr20
	tr19:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:165
		(lex.Act) = 30
		goto st13
	st13:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
//line lexer.go:637
		if data[(lex.P)] == 46 {
			goto st1
		}
		goto tr46
	st1:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof1
		}
	st_case_1:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto tr1
		}
		goto tr0
	tr1:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st14
	st14:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof14
		}
	st_case_14:
//line lexer.go:661
		switch data[(lex.P)] {
		case 69:
			goto st2
		case 101:
			goto st2
		}
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto tr1
		}
		goto tr47
	st2:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof2
		}
	st_case_2:
		switch data[(lex.P)] {
		case 43:
			goto st3
		case 45:
			goto st3
		}
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st15
		}
		goto tr2
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st15
		}
		goto tr2
	st15:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st15
		}
		goto tr47
	st16:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof16
		}
	st_case_16:
		if data[(lex.P)] == 34 {
			goto tr50
		}
		goto tr49
	tr50:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st17
	st17:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof17
		}
	st_case_17:
//line lexer.go:724
		if data[(lex.P)] == 34 {
			goto st4
		}
		goto tr51
	tr7:
//line lexer.rl:22

		lex.ReleaseSymbol("CR")

		goto st4
	st4:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
//line lexer.go:740
		switch data[(lex.P)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto st5
		case 39:
			goto st7
		}
		goto st4
	st5:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof5
		}
	st_case_5:
		switch data[(lex.P)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto st6
		case 39:
			goto st7
		}
		goto st4
	st6:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof6
		}
	st_case_6:
		switch data[(lex.P)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto tr11
		case 39:
			goto st7
		}
		goto st4
	st7:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
		switch data[(lex.P)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto st5
		case 39:
			goto st8
		}
		goto st4
	st8:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof8
		}
	st_case_8:
		switch data[(lex.P)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto st5
		case 39:
			goto tr11
		}
		goto st4
	st18:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof18
		}
	st_case_18:
		switch data[(lex.P)] {
		case 10:
			goto tr52
		case 13:
			goto tr52
		}
		goto st18
	st19:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof19
		}
	st_case_19:
		if data[(lex.P)] == 39 {
			goto tr54
		}
		goto tr53
	tr54:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st20
	st20:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof20
		}
	st_case_20:
//line lexer.go:847
		if data[(lex.P)] == 39 {
			goto st4
		}
		goto tr51
	tr29:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:104
		(lex.Act) = 15
		goto st21
	st21:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof21
		}
	st_case_21:
//line lexer.go:864
		switch data[(lex.P)] {
		case 0:
			goto st9
		case 46:
			goto st1
		case 48:
			goto st22
		}
		if 49 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st23
		}
		goto tr55
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
		if data[(lex.P)] == 46 {
			goto st1
		}
		goto tr13
	st22:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof22
		}
	st_case_22:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st22
		}
		goto tr57
	st23:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof23
		}
	st_case_23:
		if data[(lex.P)] == 46 {
			goto tr1
		}
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st23
		}
		goto tr57
	st24:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto tr1
		}
		goto tr55
	tr33:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:120
		(lex.Act) = 20
		goto st25
	tr62:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:123
		(lex.Act) = 21
		goto st25
	tr65:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:117
		(lex.Act) = 19
		goto st25
	tr68:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:113
		(lex.Act) = 18
		goto st25
	tr70:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:86
		(lex.Act) = 9
		goto st25
	tr73:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:95
		(lex.Act) = 12
		goto st25
	tr78:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:92
		(lex.Act) = 11
		goto st25
	tr80:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:89
		(lex.Act) = 10
		goto st25
	tr81:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:110
		(lex.Act) = 17
		goto st25
	tr86:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:98
		(lex.Act) = 13
		goto st25
	st25:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof25
		}
	st_case_25:
//line lexer.go:991
		if data[(lex.P)] == 95 {
			goto tr33
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr0
	st26:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof26
		}
	st_case_26:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 97:
			goto st27
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st27:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof27
		}
	st_case_27:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 108:
			goto st28
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st28:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof28
		}
	st_case_28:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 115:
			goto st29
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st29:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof29
		}
	st_case_29:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 101:
			goto tr62
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st30:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof30
		}
	st_case_30:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 111:
			goto st31
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st31:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof31
		}
	st_case_31:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 110:
			goto st32
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st32:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof32
		}
	st_case_32:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 101:
			goto tr65
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st33:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof33
		}
	st_case_33:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 114:
			goto st34
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st34:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof34
		}
	st_case_34:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 117:
			goto st29
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st35:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof35
		}
	st_case_35:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 110:
			goto st36
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st36:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof36
		}
	st_case_36:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 100:
			goto tr68
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st37:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof37
		}
	st_case_37:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 101:
			goto st38
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st38:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof38
		}
	st_case_38:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 102:
			goto tr70
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st39:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof39
		}
	st_case_39:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 111:
			goto st40
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st40:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof40
		}
	st_case_40:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 97:
			goto st41
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st41:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof41
		}
	st_case_41:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 100:
			goto tr73
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st42:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof42
		}
	st_case_42:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 111:
			goto st43
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st43:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof43
		}
	st_case_43:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 100:
			goto st44
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st44:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof44
		}
	st_case_44:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 117:
			goto st45
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st45:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof45
		}
	st_case_45:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 108:
			goto st46
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st46:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof46
		}
	st_case_46:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 101:
			goto tr78
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st47:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof47
		}
	st_case_47:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 111:
			goto st48
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st48:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof48
		}
	st_case_48:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 116:
			goto tr80
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st49:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof49
		}
	st_case_49:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 114:
			goto tr81
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st50:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof50
		}
	st_case_50:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 101:
			goto st51
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st51:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof51
		}
	st_case_51:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 116:
			goto st52
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st52:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof52
		}
	st_case_52:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 117:
			goto st53
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st53:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof53
		}
	st_case_53:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 114:
			goto st54
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	st54:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof54
		}
	st_case_54:
		switch data[(lex.P)] {
		case 95:
			goto tr33
		case 110:
			goto tr86
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr33
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto tr58
	tr15:
//line NONE:1
		switch lex.Act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 2:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(stringLiteral, "literal", "string")
			}
		}

		goto st55
	tr87:
//line lexer.rl:50
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar('\'') {
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
		}
		goto st55
	tr89:
//line lexer.rl:56
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st55
	st55:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof55
		}
	st_case_55:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:1748
		switch data[(lex.P)] {
		case 39:
			goto tr87
		case 92:
			goto st10
		}
		goto tr16
	tr16:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:56
		(lex.Act) = 2
		goto st56
	st56:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof56
		}
	st_case_56:
//line lexer.go:1768
		switch data[(lex.P)] {
		case 39:
			goto tr89
		case 92:
			goto st10
		}
		goto tr16
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
		goto tr16
	tr17:
//line NONE:1
		switch lex.Act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 4:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(stringLiteral, "literal", "string")
			}
		}

		goto st57
	tr90:
//line lexer.rl:62
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar('"') {
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
		}
		goto st57
	tr92:
//line lexer.rl:68
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st57
	st57:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof57
		}
	st_case_57:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:1826
		switch data[(lex.P)] {
		case 34:
			goto tr90
		case 92:
			goto st11
		}
		goto tr18
	tr18:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:68
		(lex.Act) = 4
		goto st58
	st58:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof58
		}
	st_case_58:
//line lexer.go:1846
		switch data[(lex.P)] {
		case 34:
			goto tr92
		case 92:
			goto st11
		}
		goto tr18
	st11:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
		goto tr18
	st_case_0:
	st0:
		(lex.Cs) = 0
		goto _out
	st_out:
	_test_eof12:
		(lex.Cs) = 12
		goto _test_eof
	_test_eof13:
		(lex.Cs) = 13
		goto _test_eof
	_test_eof1:
		(lex.Cs) = 1
		goto _test_eof
	_test_eof14:
		(lex.Cs) = 14
		goto _test_eof
	_test_eof2:
		(lex.Cs) = 2
		goto _test_eof
	_test_eof3:
		(lex.Cs) = 3
		goto _test_eof
	_test_eof15:
		(lex.Cs) = 15
		goto _test_eof
	_test_eof16:
		(lex.Cs) = 16
		goto _test_eof
	_test_eof17:
		(lex.Cs) = 17
		goto _test_eof
	_test_eof4:
		(lex.Cs) = 4
		goto _test_eof
	_test_eof5:
		(lex.Cs) = 5
		goto _test_eof
	_test_eof6:
		(lex.Cs) = 6
		goto _test_eof
	_test_eof7:
		(lex.Cs) = 7
		goto _test_eof
	_test_eof8:
		(lex.Cs) = 8
		goto _test_eof
	_test_eof18:
		(lex.Cs) = 18
		goto _test_eof
	_test_eof19:
		(lex.Cs) = 19
		goto _test_eof
	_test_eof20:
		(lex.Cs) = 20
		goto _test_eof
	_test_eof21:
		(lex.Cs) = 21
		goto _test_eof
	_test_eof9:
		(lex.Cs) = 9
		goto _test_eof
	_test_eof22:
		(lex.Cs) = 22
		goto _test_eof
	_test_eof23:
		(lex.Cs) = 23
		goto _test_eof
	_test_eof24:
		(lex.Cs) = 24
		goto _test_eof
	_test_eof25:
		(lex.Cs) = 25
		goto _test_eof
	_test_eof26:
		(lex.Cs) = 26
		goto _test_eof
	_test_eof27:
		(lex.Cs) = 27
		goto _test_eof
	_test_eof28:
		(lex.Cs) = 28
		goto _test_eof
	_test_eof29:
		(lex.Cs) = 29
		goto _test_eof
	_test_eof30:
		(lex.Cs) = 30
		goto _test_eof
	_test_eof31:
		(lex.Cs) = 31
		goto _test_eof
	_test_eof32:
		(lex.Cs) = 32
		goto _test_eof
	_test_eof33:
		(lex.Cs) = 33
		goto _test_eof
	_test_eof34:
		(lex.Cs) = 34
		goto _test_eof
	_test_eof35:
		(lex.Cs) = 35
		goto _test_eof
	_test_eof36:
		(lex.Cs) = 36
		goto _test_eof
	_test_eof37:
		(lex.Cs) = 37
		goto _test_eof
	_test_eof38:
		(lex.Cs) = 38
		goto _test_eof
	_test_eof39:
		(lex.Cs) = 39
		goto _test_eof
	_test_eof40:
		(lex.Cs) = 40
		goto _test_eof
	_test_eof41:
		(lex.Cs) = 41
		goto _test_eof
	_test_eof42:
		(lex.Cs) = 42
		goto _test_eof
	_test_eof43:
		(lex.Cs) = 43
		goto _test_eof
	_test_eof44:
		(lex.Cs) = 44
		goto _test_eof
	_test_eof45:
		(lex.Cs) = 45
		goto _test_eof
	_test_eof46:
		(lex.Cs) = 46
		goto _test_eof
	_test_eof47:
		(lex.Cs) = 47
		goto _test_eof
	_test_eof48:
		(lex.Cs) = 48
		goto _test_eof
	_test_eof49:
		(lex.Cs) = 49
		goto _test_eof
	_test_eof50:
		(lex.Cs) = 50
		goto _test_eof
	_test_eof51:
		(lex.Cs) = 51
		goto _test_eof
	_test_eof52:
		(lex.Cs) = 52
		goto _test_eof
	_test_eof53:
		(lex.Cs) = 53
		goto _test_eof
	_test_eof54:
		(lex.Cs) = 54
		goto _test_eof
	_test_eof55:
		(lex.Cs) = 55
		goto _test_eof
	_test_eof56:
		(lex.Cs) = 56
		goto _test_eof
	_test_eof10:
		(lex.Cs) = 10
		goto _test_eof
	_test_eof57:
		(lex.Cs) = 57
		goto _test_eof
	_test_eof58:
		(lex.Cs) = 58
		goto _test_eof
	_test_eof11:
		(lex.Cs) = 11
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 13:
				goto tr46
			case 1:
				goto tr0
			case 14:
				goto tr47
			case 2:
				goto tr2
			case 3:
				goto tr2
			case 15:
				goto tr47
			case 16:
				goto tr49
			case 17:
				goto tr51
			case 4:
				goto tr5
			case 5:
				goto tr5
			case 6:
				goto tr5
			case 7:
				goto tr5
			case 8:
				goto tr5
			case 18:
				goto tr52
			case 19:
				goto tr53
			case 20:
				goto tr51
			case 21:
				goto tr55
			case 9:
				goto tr13
			case 22:
				goto tr57
			case 23:
				goto tr57
			case 24:
				goto tr55
			case 25:
				goto tr0
			case 26:
				goto tr58
			case 27:
				goto tr58
			case 28:
				goto tr58
			case 29:
				goto tr58
			case 30:
				goto tr58
			case 31:
				goto tr58
			case 32:
				goto tr58
			case 33:
				goto tr58
			case 34:
				goto tr58
			case 35:
				goto tr58
			case 36:
				goto tr58
			case 37:
				goto tr58
			case 38:
				goto tr58
			case 39:
				goto tr58
			case 40:
				goto tr58
			case 41:
				goto tr58
			case 42:
				goto tr58
			case 43:
				goto tr58
			case 44:
				goto tr58
			case 45:
				goto tr58
			case 46:
				goto tr58
			case 47:
				goto tr58
			case 48:
				goto tr58
			case 49:
				goto tr58
			case 50:
				goto tr58
			case 51:
				goto tr58
			case 52:
				goto tr58
			case 53:
				goto tr58
			case 54:
				goto tr58
			case 56:
				goto tr89
			case 10:
				goto tr15
			case 58:
				goto tr92
			case 11:
				goto tr17
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:180

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
