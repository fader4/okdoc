//line lexer.rl:1
package meta

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
)

//line lexer.rl:135

//line lexer.go:16
const meta_start int = 4
const meta_first_final int = 4
const meta_error int = 0

const meta_en_singleQuoteString int = 12
const meta_en_doubleQuoteString int = 14
const meta_en_main int = 4

//line lexer.rl:138

func newTokenizer(data []byte) (*syntax.Lexer, error) {
	lex := syntax.NewLexer(data)

//line lexer.go:32
	{
		(lex.Cs) = meta_start
		(lex.Top) = 0
		(lex.Ts) = 0
		(lex.Te) = 0
		(lex.Act) = 0
	}

//line lexer.rl:143

//line lexer.go:43
	{
		if (lex.P) == (lex.Pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.Cs {
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
		case 9:
			goto st9
		case 1:
			goto st1
		case 10:
			goto st10
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 2:
			goto st2
		case 14:
			goto st14
		case 15:
			goto st15
		case 3:
			goto st3
		case 0:
			goto st0
		}

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof
		}
	_resume:
		switch lex.Cs {
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
		case 9:
			goto st_case_9
		case 1:
			goto st_case_1
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 2:
			goto st_case_2
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 3:
			goto st_case_3
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line lexer.rl:83
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseToken(ident, "ident")
		}
		goto st4
	tr6:
//line lexer.rl:129
		(lex.Te) = (lex.P) + 1

		goto st4
	tr7:
//line lexer.rl:25

		lex.ReleaseSymbol("SP")

//line lexer.rl:68
		(lex.Te) = (lex.P) + 1

		goto st4
	tr8:
//line lexer.rl:22

		lex.ReleaseSymbol("CR")

//line lexer.rl:69
		(lex.Te) = (lex.P) + 1

		goto st4
	tr9:
//line lexer.rl:101
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 4
					(lex.Top)++
					goto st14
				}
			}
		}
		goto st4
	tr10:
//line lexer.rl:96
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 4
					(lex.Top)++
					goto st12
				}
			}
		}
		goto st4
	tr11:
//line lexer.rl:106
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('(', "bracket", "open_bracket")
			lex.BeginPairedChar(')')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 4
					(lex.Top)++
					goto st4
				}
			}
		}
		goto st4
	tr12:
//line lexer.rl:122
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
		goto st4
	tr13:
//line lexer.rl:78
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseSymbol("op_and_punct")
		}
		goto st4
	tr16:
//line lexer.rl:83
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(ident, "ident")
		}
		goto st4
	st4:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:216
		switch data[(lex.P)] {
		case 9:
			goto tr7
		case 10:
			goto tr8
		case 13:
			goto tr8
		case 32:
			goto tr7
		case 34:
			goto tr9
		case 39:
			goto tr10
		case 40:
			goto tr11
		case 41:
			goto tr12
		case 61:
			goto tr13
		case 70:
			goto st6
		case 84:
			goto st10
		case 95:
			goto st5
		}
		switch {
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		case data[(lex.P)] >= 65:
			goto st5
		}
		goto tr6
	st5:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof5
		}
	st_case_5:
		if data[(lex.P)] == 95 {
			goto st5
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr16
	st6:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof6
		}
	st_case_6:
		switch data[(lex.P)] {
		case 95:
			goto st5
		case 97:
			goto st7
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr16
	st7:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
		switch data[(lex.P)] {
		case 95:
			goto st5
		case 108:
			goto st8
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr16
	st8:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof8
		}
	st_case_8:
		switch data[(lex.P)] {
		case 95:
			goto st5
		case 115:
			goto tr19
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr16
	tr19:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st9
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer.go:355
		switch data[(lex.P)] {
		case 95:
			goto st5
		case 101:
			goto st1
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr16
	st1:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof1
		}
	st_case_1:
		if data[(lex.P)] == 95 {
			goto st5
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr0
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
		switch data[(lex.P)] {
		case 95:
			goto st5
		case 114:
			goto st11
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr16
	st11:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
		switch data[(lex.P)] {
		case 95:
			goto st5
		case 117:
			goto tr19
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st5
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr16
	tr2:
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

				lex.ReleaseToken(stringLiteral, "string")
			}
		}

		goto st12
	tr22:
//line lexer.rl:42
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
		goto st12
	tr24:
//line lexer.rl:48
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "string")
		}
		goto st12
	st12:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:488
		switch data[(lex.P)] {
		case 39:
			goto tr22
		case 92:
			goto st2
		}
		goto tr3
	tr3:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:48
		(lex.Act) = 2
		goto st13
	st13:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
//line lexer.go:508
		switch data[(lex.P)] {
		case 39:
			goto tr24
		case 92:
			goto st2
		}
		goto tr3
	st2:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof2
		}
	st_case_2:
		goto tr3
	tr4:
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

				lex.ReleaseToken(stringLiteral, "string")
			}
		}

		goto st14
	tr25:
//line lexer.rl:54
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
		goto st14
	tr27:
//line lexer.rl:60
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "string")
		}
		goto st14
	st14:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof14
		}
	st_case_14:
//line NONE:1
		(lex.Ts) = (lex.P)

//line lexer.go:566
		switch data[(lex.P)] {
		case 34:
			goto tr25
		case 92:
			goto st3
		}
		goto tr5
	tr5:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line lexer.rl:60
		(lex.Act) = 4
		goto st15
	st15:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
//line lexer.go:586
		switch data[(lex.P)] {
		case 34:
			goto tr27
		case 92:
			goto st3
		}
		goto tr5
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
		goto tr5
	st_case_0:
	st0:
		(lex.Cs) = 0
		goto _out
	st_out:
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
	_test_eof9:
		(lex.Cs) = 9
		goto _test_eof
	_test_eof1:
		(lex.Cs) = 1
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
	_test_eof2:
		(lex.Cs) = 2
		goto _test_eof
	_test_eof14:
		(lex.Cs) = 14
		goto _test_eof
	_test_eof15:
		(lex.Cs) = 15
		goto _test_eof
	_test_eof3:
		(lex.Cs) = 3
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 5:
				goto tr16
			case 6:
				goto tr16
			case 7:
				goto tr16
			case 8:
				goto tr16
			case 9:
				goto tr16
			case 1:
				goto tr0
			case 10:
				goto tr16
			case 11:
				goto tr16
			case 13:
				goto tr24
			case 2:
				goto tr2
			case 15:
				goto tr27
			case 3:
				goto tr4
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:144

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
