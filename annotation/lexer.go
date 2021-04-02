//line lexer.rl:1
package annotation

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
)

//line lexer.rl:138

//line lexer.go:16
const annotation_start int = 7
const annotation_first_final int = 7
const annotation_error int = 0

const annotation_en_singleQuoteString int = 33
const annotation_en_doubleQuoteString int = 35
const annotation_en_main int = 7

//line lexer.rl:141

func newTokenizer(data []byte) (*syntax.Lexer, error) {
	lex := syntax.NewLexer(data)

//line lexer.go:32
	{
		(lex.Cs) = annotation_start
		(lex.Top) = 0
		(lex.Ts) = 0
		(lex.Te) = 0
		(lex.Act) = 0
	}

//line lexer.rl:146

//line lexer.go:43
	{
		if (lex.P) == (lex.Pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.Cs {
		case 7:
			goto st7
		case 8:
			goto st8
		case 1:
			goto st1
		case 9:
			goto st9
		case 2:
			goto st2
		case 3:
			goto st3
		case 10:
			goto st10
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 4:
			goto st4
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
		case 5:
			goto st5
		case 35:
			goto st35
		case 36:
			goto st36
		case 6:
			goto st6
		case 0:
			goto st0
		}

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof
		}
	_resume:
		switch lex.Cs {
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 1:
			goto st_case_1
		case 9:
			goto st_case_9
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 4:
			goto st_case_4
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
		case 5:
			goto st_case_5
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 6:
			goto st_case_6
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line lexer.rl:132
		(lex.P) = (lex.Te) - 1

		goto st7
	tr2:
//line lexer.rl:96
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseToken(floatLiteral, "literal", "float")
		}
		goto st7
	tr11:
//line lexer.rl:132
		(lex.Te) = (lex.P) + 1

		goto st7
	tr12:
//line lexer.rl:25

		lex.ReleaseSymbol("SP")

//line lexer.rl:70
		(lex.Te) = (lex.P) + 1

		goto st7
	tr13:
//line lexer.rl:22

		lex.ReleaseSymbol("CR")

//line lexer.rl:71
		(lex.Te) = (lex.P) + 1

		goto st7
	tr16:
//line lexer.rl:110
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('(', "bracket", "open_bracket")
			lex.BeginPairedChar(')')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 7
					(lex.Top)++
					goto st7
				}
			}
		}
		goto st7
	tr17:
//line lexer.rl:125
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
		goto st7
	tr18:
//line lexer.rl:80
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st7
	tr23:
//line lexer.rl:77
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseSymbol("at")
		}
		goto st7
	tr28:
//line lexer.rl:120
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('[', "bracket", "open_bracket")
			lex.BeginPairedChar(']')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 7
					(lex.Top)++
					goto st7
				}
			}
		}
		goto st7
	tr29:
//line lexer.rl:115
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('{', "bracket", "open_bracket")
			lex.BeginPairedChar('}')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 7
					(lex.Top)++
					goto st7
				}
			}
		}
		goto st7
	tr30:
//line lexer.rl:132
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st7
	tr31:
//line lexer.rl:96
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(floatLiteral, "literal", "float")
		}
		goto st7
	tr33:
//line lexer.rl:105
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 7
					(lex.Top)++
					goto st35
				}
			}
		}
		goto st7
	tr34:
//line lexer.rl:73
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st7
	tr35:
//line lexer.rl:100
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 7
					(lex.Top)++
					goto st33
				}
			}
		}
		goto st7
	tr37:
//line lexer.rl:93
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(integerLiteral, "literal", "int")
		}
		goto st7
	tr38:
//line lexer.rl:80
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st7
	tr39:
//line NONE:1
		switch lex.Act {
		case 10:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(nullLiteral, "literal")
			}
		case 11:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(ident, "ident")
			}
		case 12:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(boolLiteral, "literal", "bool")
			}
		}

		goto st7
	tr40:
//line lexer.rl:87
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(ident, "ident")
		}
		goto st7
	st7:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:390
		switch data[(lex.P)] {
		case 0:
			goto tr10
		case 9:
			goto tr12
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 32:
			goto tr12
		case 34:
			goto st11
		case 39:
			goto st12
		case 40:
			goto tr16
		case 41:
			goto tr17
		case 44:
			goto tr18
		case 45:
			goto tr19
		case 46:
			goto st16
		case 48:
			goto st14
		case 61:
			goto tr18
		case 64:
			goto tr23
		case 70:
			goto st18
		case 78:
			goto st25
		case 84:
			goto st30
		case 91:
			goto tr28
		case 93:
			goto tr17
		case 95:
			goto tr24
		case 123:
			goto tr29
		case 125:
			goto tr17
		}
		switch {
		case data[(lex.P)] < 65:
			if 49 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr11
	tr10:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st8
	st8:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof8
		}
	st_case_8:
//line lexer.go:462
		if data[(lex.P)] == 46 {
			goto st1
		}
		goto tr30
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

		goto st9
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer.go:486
		switch data[(lex.P)] {
		case 69:
			goto st2
		case 101:
			goto st2
		}
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto tr1
		}
		goto tr31
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
			goto st10
		}
		goto tr2
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st10
		}
		goto tr2
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st10
		}
		goto tr31
	st11:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
		if data[(lex.P)] == 34 {
			goto tr34
		}
		goto tr33
	st12:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
		if data[(lex.P)] == 39 {
			goto tr34
		}
		goto tr35
	tr19:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st13
	st13:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
//line lexer.go:558
		switch data[(lex.P)] {
		case 0:
			goto st4
		case 46:
			goto st1
		case 48:
			goto st14
		}
		if 49 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st15
		}
		goto tr30
	st4:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
		if data[(lex.P)] == 46 {
			goto st1
		}
		goto tr0
	st14:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof14
		}
	st_case_14:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st14
		}
		goto tr37
	st15:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
		if data[(lex.P)] == 46 {
			goto tr1
		}
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto st15
		}
		goto tr37
	st16:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
			goto tr1
		}
		goto tr38
	tr24:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:87
		(lex.Act) = 11
		goto st17
	tr45:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:90
		(lex.Act) = 12
		goto st17
	tr51:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:84
		(lex.Act) = 10
		goto st17
	st17:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof17
		}
	st_case_17:
//line lexer.go:636
		if data[(lex.P)] == 95 {
			goto tr24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr39
	st18:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof18
		}
	st_case_18:
		switch data[(lex.P)] {
		case 65:
			goto st19
		case 95:
			goto tr24
		case 97:
			goto st22
		}
		switch {
		case data[(lex.P)] < 66:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st19:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof19
		}
	st_case_19:
		switch data[(lex.P)] {
		case 76:
			goto st20
		case 95:
			goto tr24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st20:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof20
		}
	st_case_20:
		switch data[(lex.P)] {
		case 83:
			goto st21
		case 95:
			goto tr24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st21:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof21
		}
	st_case_21:
		switch data[(lex.P)] {
		case 69:
			goto tr45
		case 95:
			goto tr24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st22:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof22
		}
	st_case_22:
		switch data[(lex.P)] {
		case 95:
			goto tr24
		case 108:
			goto st23
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st23:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof23
		}
	st_case_23:
		switch data[(lex.P)] {
		case 95:
			goto tr24
		case 115:
			goto st24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st24:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof24
		}
	st_case_24:
		switch data[(lex.P)] {
		case 95:
			goto tr24
		case 101:
			goto tr45
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st25:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof25
		}
	st_case_25:
		switch data[(lex.P)] {
		case 85:
			goto st26
		case 95:
			goto tr24
		case 117:
			goto st28
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st26:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof26
		}
	st_case_26:
		switch data[(lex.P)] {
		case 76:
			goto st27
		case 95:
			goto tr24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st27:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof27
		}
	st_case_27:
		switch data[(lex.P)] {
		case 76:
			goto tr51
		case 95:
			goto tr24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st28:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof28
		}
	st_case_28:
		switch data[(lex.P)] {
		case 95:
			goto tr24
		case 108:
			goto st29
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st29:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof29
		}
	st_case_29:
		switch data[(lex.P)] {
		case 95:
			goto tr24
		case 108:
			goto tr51
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st30:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof30
		}
	st_case_30:
		switch data[(lex.P)] {
		case 82:
			goto st31
		case 95:
			goto tr24
		case 114:
			goto st32
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st31:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof31
		}
	st_case_31:
		switch data[(lex.P)] {
		case 85:
			goto st21
		case 95:
			goto tr24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	st32:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof32
		}
	st_case_32:
		switch data[(lex.P)] {
		case 95:
			goto tr24
		case 117:
			goto st24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto tr24
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto tr40
	tr6:
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

		goto st33
	tr55:
//line lexer.rl:44
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
		goto st33
	tr57:
//line lexer.rl:50
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st33
	st33:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof33
		}
	st_case_33:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:1063
		switch data[(lex.P)] {
		case 39:
			goto tr55
		case 92:
			goto st5
		}
		goto tr7
	tr7:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:50
		(lex.Act) = 2
		goto st34
	st34:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof34
		}
	st_case_34:
//line lexer.go:1083
		switch data[(lex.P)] {
		case 39:
			goto tr57
		case 92:
			goto st5
		}
		goto tr7
	st5:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof5
		}
	st_case_5:
		goto tr7
	tr8:
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

		goto st35
	tr58:
//line lexer.rl:56
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
		goto st35
	tr60:
//line lexer.rl:62
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st35
	st35:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof35
		}
	st_case_35:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:1141
		switch data[(lex.P)] {
		case 34:
			goto tr58
		case 92:
			goto st6
		}
		goto tr9
	tr9:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:62
		(lex.Act) = 4
		goto st36
	st36:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof36
		}
	st_case_36:
//line lexer.go:1161
		switch data[(lex.P)] {
		case 34:
			goto tr60
		case 92:
			goto st6
		}
		goto tr9
	st6:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof6
		}
	st_case_6:
		goto tr9
	st_case_0:
	st0:
		(lex.Cs) = 0
		goto _out
	st_out:
	_test_eof7:
		(lex.Cs) = 7
		goto _test_eof
	_test_eof8:
		(lex.Cs) = 8
		goto _test_eof
	_test_eof1:
		(lex.Cs) = 1
		goto _test_eof
	_test_eof9:
		(lex.Cs) = 9
		goto _test_eof
	_test_eof2:
		(lex.Cs) = 2
		goto _test_eof
	_test_eof3:
		(lex.Cs) = 3
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
	_test_eof4:
		(lex.Cs) = 4
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
	_test_eof5:
		(lex.Cs) = 5
		goto _test_eof
	_test_eof35:
		(lex.Cs) = 35
		goto _test_eof
	_test_eof36:
		(lex.Cs) = 36
		goto _test_eof
	_test_eof6:
		(lex.Cs) = 6
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 8:
				goto tr30
			case 1:
				goto tr0
			case 9:
				goto tr31
			case 2:
				goto tr2
			case 3:
				goto tr2
			case 10:
				goto tr31
			case 11:
				goto tr33
			case 12:
				goto tr35
			case 13:
				goto tr30
			case 4:
				goto tr0
			case 14:
				goto tr37
			case 15:
				goto tr37
			case 16:
				goto tr38
			case 17:
				goto tr39
			case 18:
				goto tr40
			case 19:
				goto tr40
			case 20:
				goto tr40
			case 21:
				goto tr40
			case 22:
				goto tr40
			case 23:
				goto tr40
			case 24:
				goto tr40
			case 25:
				goto tr40
			case 26:
				goto tr40
			case 27:
				goto tr40
			case 28:
				goto tr40
			case 29:
				goto tr40
			case 30:
				goto tr40
			case 31:
				goto tr40
			case 32:
				goto tr40
			case 34:
				goto tr57
			case 5:
				goto tr6
			case 36:
				goto tr60
			case 6:
				goto tr8
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:147

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
