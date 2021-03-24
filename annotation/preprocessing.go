//line preprocessing.rl:1
package annotation

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
)

//line preprocessing.rl:105

//line preprocessing.go:16
const annotation_preprocessing_start int = 13
const annotation_preprocessing_first_final int = 13
const annotation_preprocessing_error int = 0

const annotation_preprocessing_en_singleQuoteString int = 22
const annotation_preprocessing_en_doubleQuoteString int = 24
const annotation_preprocessing_en_main int = 13

//line preprocessing.rl:108

func newPreprocessing(data []byte) (*syntax.Lexer, error) {
	lex := syntax.NewLexer(data)

//line preprocessing.go:32
	{
		(lex.Cs) = annotation_preprocessing_start
		(lex.Top) = 0
		(lex.Ts) = 0
		(lex.Te) = 0
		(lex.Act) = 0
	}

//line preprocessing.rl:113

//line preprocessing.go:43
	{
		if (lex.P) == (lex.Pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.Cs {
		case 13:
			goto st13
		case 14:
			goto st14
		case 1:
			goto st1
		case 2:
			goto st2
		case 3:
			goto st3
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
		case 10:
			goto st10
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 11:
			goto st11
		case 24:
			goto st24
		case 25:
			goto st25
		case 12:
			goto st12
		case 0:
			goto st0
		}

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof
		}
	_resume:
		switch lex.Cs {
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
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
		case 10:
			goto st_case_10
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 11:
			goto st_case_11
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 12:
			goto st_case_12
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line preprocessing.rl:69
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseSymbol("at")
		}
		goto st13
	tr1:
//line preprocessing.rl:22

		lex.ReleaseSymbol("CR")

//line preprocessing.rl:78
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(1, "annotation")
		}
		goto st13
	tr5:
//line preprocessing.rl:82
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(2, "annotation")
		}
		goto st13
	tr11:
//line preprocessing.rl:74
		(lex.P) = (lex.Te) - 1
		{
			lex.ReleaseToken(ident, "ident")
		}
		goto st13
	tr17:
//line preprocessing.rl:99
		(lex.Te) = (lex.P) + 1

		goto st13
	tr18:
//line preprocessing.rl:25

		lex.ReleaseSymbol("SP")

//line preprocessing.rl:66
		(lex.Te) = (lex.P) + 1

		goto st13
	tr19:
//line preprocessing.rl:22

		lex.ReleaseSymbol("CR")

//line preprocessing.rl:67
		(lex.Te) = (lex.P) + 1

		goto st13
	tr20:
//line preprocessing.rl:87
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken('(', "bracket", "open_bracket")
			lex.BeginPairedChar(')')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 13
					(lex.Top)++
					goto st13
				}
			}
		}
		goto st13
	tr21:
//line preprocessing.rl:92
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
		goto st13
	tr25:
//line preprocessing.rl:69
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseSymbol("at")
		}
		goto st13
	tr28:
//line preprocessing.rl:74
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(ident, "ident")
		}
		goto st13
	st13:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
//line NONE:1
		(lex.Ts) = (lex.P)

//line preprocessing.go:266
		switch data[(lex.P)] {
		case 9:
			goto tr18
		case 10:
			goto tr19
		case 13:
			goto tr19
		case 32:
			goto tr18
		case 40:
			goto tr20
		case 41:
			goto tr21
		case 64:
			goto tr22
		case 70:
			goto st16
		case 84:
			goto st20
		case 95:
			goto st15
		}
		switch {
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		case data[(lex.P)] >= 65:
			goto st15
		}
		goto tr17
	tr22:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st14
	st14:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof14
		}
	st_case_14:
//line preprocessing.go:308
		switch data[(lex.P)] {
		case 70:
			goto st3
		case 84:
			goto st8
		case 95:
			goto st1
		}
		switch {
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		case data[(lex.P)] >= 65:
			goto st1
		}
		goto tr25
	st1:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof1
		}
	st_case_1:
		switch data[(lex.P)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 40:
			goto st2
		case 95:
			goto st1
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	tr4:
//line preprocessing.rl:37
		lex.ReleaseNewLine("CR", "CR_nested")
		goto st2
	st2:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof2
		}
	st_case_2:
//line preprocessing.go:363
		switch data[(lex.P)] {
		case 10:
			goto tr4
		case 13:
			goto tr4
		case 41:
			goto tr5
		}
		goto st2
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
		switch data[(lex.P)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 40:
			goto st2
		case 95:
			goto st1
		case 97:
			goto st4
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st4:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
		switch data[(lex.P)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 40:
			goto st2
		case 95:
			goto st1
		case 108:
			goto st5
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st5:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof5
		}
	st_case_5:
		switch data[(lex.P)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 40:
			goto st2
		case 95:
			goto st1
		case 115:
			goto st6
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st6:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof6
		}
	st_case_6:
		switch data[(lex.P)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 40:
			goto st2
		case 95:
			goto st1
		case 101:
			goto st7
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st7:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
		if data[(lex.P)] == 95 {
			goto st1
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st8:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof8
		}
	st_case_8:
		switch data[(lex.P)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 40:
			goto st2
		case 95:
			goto st1
		case 114:
			goto st9
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
		switch data[(lex.P)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 40:
			goto st2
		case 95:
			goto st1
		case 117:
			goto st6
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st1
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st15:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
		if data[(lex.P)] == 95 {
			goto st15
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr28
	st16:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof16
		}
	st_case_16:
		switch data[(lex.P)] {
		case 95:
			goto st15
		case 97:
			goto st17
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr28
	st17:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof17
		}
	st_case_17:
		switch data[(lex.P)] {
		case 95:
			goto st15
		case 108:
			goto st18
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr28
	st18:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof18
		}
	st_case_18:
		switch data[(lex.P)] {
		case 95:
			goto st15
		case 115:
			goto tr31
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr28
	tr31:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st19
	st19:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof19
		}
	st_case_19:
//line preprocessing.go:677
		switch data[(lex.P)] {
		case 95:
			goto st15
		case 101:
			goto st10
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr28
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
		if data[(lex.P)] == 95 {
			goto st15
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr11
	st20:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof20
		}
	st_case_20:
		switch data[(lex.P)] {
		case 95:
			goto st15
		case 114:
			goto st21
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr28
	st21:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof21
		}
	st_case_21:
		switch data[(lex.P)] {
		case 95:
			goto st15
		case 117:
			goto tr31
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st15
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr28
	tr13:
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

		goto st22
	tr34:
//line preprocessing.rl:40
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
		goto st22
	tr36:
//line preprocessing.rl:46
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(stringLiteral, "literal", "string")
		}
		goto st22
	st22:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof22
		}
	st_case_22:
//line NONE:1
		(lex.Ts) = (lex.P)

//line preprocessing.go:810
		switch data[(lex.P)] {
		case 39:
			goto tr34
		case 92:
			goto st11
		}
		goto tr14
	tr14:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line preprocessing.rl:46
		(lex.Act) = 2
		goto st23
	st23:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof23
		}
	st_case_23:
//line preprocessing.go:830
		switch data[(lex.P)] {
		case 39:
			goto tr36
		case 92:
			goto st11
		}
		goto tr14
	st11:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
		goto tr14
	tr15:
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

		goto st24
	tr37:
//line preprocessing.rl:52
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
		goto st24
	tr39:
//line preprocessing.rl:58
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

//line preprocessing.go:888
		switch data[(lex.P)] {
		case 34:
			goto tr37
		case 92:
			goto st12
		}
		goto tr16
	tr16:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line preprocessing.rl:58
		(lex.Act) = 4
		goto st25
	st25:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof25
		}
	st_case_25:
//line preprocessing.go:908
		switch data[(lex.P)] {
		case 34:
			goto tr39
		case 92:
			goto st12
		}
		goto tr16
	st12:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
		goto tr16
	st_case_0:
	st0:
		(lex.Cs) = 0
		goto _out
	st_out:
	_test_eof13:
		(lex.Cs) = 13
		goto _test_eof
	_test_eof14:
		(lex.Cs) = 14
		goto _test_eof
	_test_eof1:
		(lex.Cs) = 1
		goto _test_eof
	_test_eof2:
		(lex.Cs) = 2
		goto _test_eof
	_test_eof3:
		(lex.Cs) = 3
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
	_test_eof9:
		(lex.Cs) = 9
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
	_test_eof10:
		(lex.Cs) = 10
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
	_test_eof11:
		(lex.Cs) = 11
		goto _test_eof
	_test_eof24:
		(lex.Cs) = 24
		goto _test_eof
	_test_eof25:
		(lex.Cs) = 25
		goto _test_eof
	_test_eof12:
		(lex.Cs) = 12
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 14:
				goto tr25
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr0
			case 4:
				goto tr0
			case 5:
				goto tr0
			case 6:
				goto tr0
			case 7:
				goto tr0
			case 8:
				goto tr0
			case 9:
				goto tr0
			case 15:
				goto tr28
			case 16:
				goto tr28
			case 17:
				goto tr28
			case 18:
				goto tr28
			case 19:
				goto tr28
			case 10:
				goto tr11
			case 20:
				goto tr28
			case 21:
				goto tr28
			case 23:
				goto tr36
			case 11:
				goto tr13
			case 25:
				goto tr39
			case 12:
				goto tr15
			}
		}

	_out:
		{
		}
	}

//line preprocessing.rl:114

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
