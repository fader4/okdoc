//line lexer.rl:1
package starlark

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
)

//line lexer.rl:137

//line lexer.go:16
const starlark_start int = 6
const starlark_first_final int = 6
const starlark_error int = 0

const starlark_en_singleQuoteString int = 24
const starlark_en_doubleQuoteString int = 26
const starlark_en_main int = 6

//line lexer.rl:140

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

//line lexer.rl:145

//line lexer.go:43
	{
		if (lex.P) == (lex.Pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.Cs {
		case 6:
			goto st6
		case 7:
			goto st7
		case 1:
			goto st1
		case 8:
			goto st8
		case 2:
			goto st2
		case 3:
			goto st3
		case 9:
			goto st9
		case 10:
			goto st10
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 4:
			goto st4
		case 26:
			goto st26
		case 27:
			goto st27
		case 5:
			goto st5
		case 0:
			goto st0
		}

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof
		}
	_resume:
		switch lex.Cs {
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 1:
			goto st_case_1
		case 8:
			goto st_case_8
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 4:
			goto st_case_4
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 5:
			goto st_case_5
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line lexer.rl:131
		(lex.P) = (lex.Te) - 1

		goto st6
	tr2:
//line lexer.rl:95
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseToken(floatLiteral, "literal", "float")
		}
		goto st6
	tr10:
//line lexer.rl:131
		(lex.Te) = (lex.P) + 1

		goto st6
	tr11:
//line lexer.rl:25

		lex.ReleaseSymbol("SP")

//line lexer.rl:69
		(lex.Te) = (lex.P) + 1

		goto st6
	tr12:
//line lexer.rl:22

		lex.ReleaseSymbol("CR")

//line lexer.rl:70
		(lex.Te) = (lex.P) + 1

		goto st6
	tr13:
//line lexer.rl:104
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 6
					(lex.Top)++
					goto st26
				}
			}
		}
		goto st6
	tr15:
//line lexer.rl:99
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 6
					(lex.Top)++
					goto st24
				}
			}
		}
		goto st6
	tr16:
//line lexer.rl:109
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('(', "bracket", "open_bracket")
			lex.BeginPairedChar(')')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 6
					(lex.Top)++
					goto st6
				}
			}
		}
		goto st6
	tr17:
//line lexer.rl:124
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
		goto st6
	tr18:
//line lexer.rl:79
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st6
	tr26:
//line lexer.rl:119
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('[', "bracket", "open_bracket")
			lex.BeginPairedChar(']')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 6
					(lex.Top)++
					goto st6
				}
			}
		}
		goto st6
	tr27:
//line lexer.rl:114
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('{', "bracket", "open_bracket")
			lex.BeginPairedChar('}')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 6
					(lex.Top)++
					goto st6
				}
			}
		}
		goto st6
	tr28:
//line lexer.rl:131
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st6
	tr30:
//line lexer.rl:95
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(floatLiteral, "literal", "float")
		}
		goto st6
	tr32:
//line lexer.rl:72
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(commentInline, "comment")
		}
		goto st6
	tr33:
//line lexer.rl:79
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st6
	tr34:
//line lexer.rl:92
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(integerLiteral, "literal", "int")
		}
		goto st6
	tr35:
//line NONE:1
		switch lex.Act {
		case 9:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(nullLiteral, "literal")
			}
		case 10:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(ident, "ident")
			}
		case 11:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(boolLiteral, "literal", "bool")
			}
		}

		goto st6
	tr36:
//line lexer.rl:86
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(ident, "ident")
		}
		goto st6
	st6:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof6
		}
	st_case_6:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:346
		switch data[(lex.P)] {
		case 0:
			goto tr9
		case 9:
			goto tr11
		case 10:
			goto tr12
		case 13:
			goto tr12
		case 32:
			goto tr11
		case 34:
			goto tr13
		case 35:
			goto st10
		case 39:
			goto tr15
		case 40:
			goto tr16
		case 41:
			goto tr17
		case 44:
			goto tr18
		case 46:
			goto st11
		case 48:
			goto st12
		case 58:
			goto tr18
		case 61:
			goto tr18
		case 70:
			goto st15
		case 78:
			goto st19
		case 84:
			goto st22
		case 91:
			goto tr26
		case 93:
			goto tr17
		case 95:
			goto tr22
		case 123:
			goto tr27
		case 125:
			goto tr17
		}
		switch {
		case data[(lex.P)] < 65:
			if 49 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st13
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr10
	tr9:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st7
	st7:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
//line lexer.go:418
		if data[(lex.P)] == 46 {
			goto st1
		}
		goto tr28
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

		goto st8
	st8:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof8
		}
	st_case_8:
//line lexer.go:442
		switch data[(lex.P)] {
		case 69:
			goto st2
		case 101:
			goto st2
		}
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto tr1
		}
		goto tr30
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
			goto st9
		}
		goto tr2
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st9
		}
		goto tr2
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st9
		}
		goto tr30
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
		switch data[(lex.P)] {
		case 10:
			goto tr32
		case 13:
			goto tr32
		}
		goto st10
	st11:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto tr1
		}
		goto tr33
	st12:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st12
		}
		goto tr34
	st13:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
		if data[(lex.P)] == 46 {
			goto tr1
		}
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st13
		}
		goto tr34
	tr22:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:86
		(lex.Act) = 10
		goto st14
	tr40:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:89
		(lex.Act) = 11
		goto st14
	tr43:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:83
		(lex.Act) = 9
		goto st14
	st14:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof14
		}
	st_case_14:
//line lexer.go:554
		if data[(lex.P)] == 95 {
			goto tr22
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr35
	st15:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 97:
			goto st16
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st16:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof16
		}
	st_case_16:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 108:
			goto st17
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st17:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof17
		}
	st_case_17:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 115:
			goto st18
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st18:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof18
		}
	st_case_18:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 101:
			goto tr40
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st19:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof19
		}
	st_case_19:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 117:
			goto st20
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st20:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof20
		}
	st_case_20:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 108:
			goto st21
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st21:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof21
		}
	st_case_21:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 108:
			goto tr43
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st22:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof22
		}
	st_case_22:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 114:
			goto st23
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	st23:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof23
		}
	st_case_23:
		switch data[(lex.P)] {
		case 95:
			goto tr22
		case 117:
			goto st18
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr22
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto tr36
	tr5:
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

		goto st24
	tr45:
//line lexer.rl:43
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
		goto st24
	tr47:
//line lexer.rl:49
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st24
	st24:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof24
		}
	st_case_24:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:831
		switch data[(lex.P)] {
		case 39:
			goto tr45
		case 92:
			goto st4
		}
		goto tr6
	tr6:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:49
		(lex.Act) = 2
		goto st25
	st25:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof25
		}
	st_case_25:
//line lexer.go:851
		switch data[(lex.P)] {
		case 39:
			goto tr47
		case 92:
			goto st4
		}
		goto tr6
	st4:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
		goto tr6
	tr7:
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

		goto st26
	tr48:
//line lexer.rl:55
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
		goto st26
	tr50:
//line lexer.rl:61
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st26
	st26:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof26
		}
	st_case_26:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:909
		switch data[(lex.P)] {
		case 34:
			goto tr48
		case 92:
			goto st5
		}
		goto tr8
	tr8:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:61
		(lex.Act) = 4
		goto st27
	st27:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof27
		}
	st_case_27:
//line lexer.go:929
		switch data[(lex.P)] {
		case 34:
			goto tr50
		case 92:
			goto st5
		}
		goto tr8
	st5:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof5
		}
	st_case_5:
		goto tr8
	st_case_0:
	st0:
		(lex.Cs) = 0
		goto _out
	st_out:
	_test_eof6:
		(lex.Cs) = 6
		goto _test_eof
	_test_eof7:
		(lex.Cs) = 7
		goto _test_eof
	_test_eof1:
		(lex.Cs) = 1
		goto _test_eof
	_test_eof8:
		(lex.Cs) = 8
		goto _test_eof
	_test_eof2:
		(lex.Cs) = 2
		goto _test_eof
	_test_eof3:
		(lex.Cs) = 3
		goto _test_eof
	_test_eof9:
		(lex.Cs) = 9
		goto _test_eof
	_test_eof10:
		(lex.Cs) = 10
		goto _test_eof
	_test_eof11:
		(lex.Cs) = 11
		goto _test_eof
	_test_eof12:
		(lex.Cs) = 12
		goto _test_eof
	_test_eof13:
		(lex.Cs) = 13
		goto _test_eof
	_test_eof14:
		(lex.Cs) = 14
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
	_test_eof4:
		(lex.Cs) = 4
		goto _test_eof
	_test_eof26:
		(lex.Cs) = 26
		goto _test_eof
	_test_eof27:
		(lex.Cs) = 27
		goto _test_eof
	_test_eof5:
		(lex.Cs) = 5
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 7:
				goto tr28
			case 1:
				goto tr0
			case 8:
				goto tr30
			case 2:
				goto tr2
			case 3:
				goto tr2
			case 9:
				goto tr30
			case 10:
				goto tr32
			case 11:
				goto tr33
			case 12:
				goto tr34
			case 13:
				goto tr34
			case 14:
				goto tr35
			case 15:
				goto tr36
			case 16:
				goto tr36
			case 17:
				goto tr36
			case 18:
				goto tr36
			case 19:
				goto tr36
			case 20:
				goto tr36
			case 21:
				goto tr36
			case 22:
				goto tr36
			case 23:
				goto tr36
			case 25:
				goto tr47
			case 4:
				goto tr5
			case 27:
				goto tr50
			case 5:
				goto tr7
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:146

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
