//line preprocessing.rl:1
package annotation

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
)

//line preprocessing.rl:104

//line preprocessing.go:16
const annotation_preprocessing_start int = 11
const annotation_preprocessing_first_final int = 11
const annotation_preprocessing_error int = 0

const annotation_preprocessing_en_singleQuoteString int = 13
const annotation_preprocessing_en_doubleQuoteString int = 15
const annotation_preprocessing_en_annotationBody int = 17
const annotation_preprocessing_en_main int = 11

//line preprocessing.rl:107

func newPreprocessing(data []byte) (*syntax.Lexer, error) {
	lex := syntax.NewLexer(data)

//line preprocessing.go:33
	{
		(lex.Cs) = annotation_preprocessing_start
		(lex.Top) = 0
		(lex.Ts) = 0
		(lex.Te) = 0
		(lex.Act) = 0
	}

//line preprocessing.rl:112

//line preprocessing.go:44
	{
		if (lex.P) == (lex.Pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.Cs {
		case 11:
			goto st11
		case 12:
			goto st12
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
		case 13:
			goto st13
		case 14:
			goto st14
		case 9:
			goto st9
		case 15:
			goto st15
		case 16:
			goto st16
		case 10:
			goto st10
		case 17:
			goto st17
		case 0:
			goto st0
		}

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof
		}
	_resume:
		switch lex.Cs {
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
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
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 9:
			goto st_case_9
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 10:
			goto st_case_10
		case 17:
			goto st_case_17
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line preprocessing.rl:98
		(lex.P) = (lex.Te) - 1

		goto st11
	tr1:
//line preprocessing.rl:92
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(beginAnnotation, "annotation")
			lex.BeginPairedChar(endAnnotation)
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 11
					(lex.Top)++
					goto st17
				}
			}
		}
		goto st11
	tr12:
//line preprocessing.rl:98
		(lex.Te) = (lex.P) + 1

		goto st11
	tr13:
//line preprocessing.rl:25

		lex.ReleaseSymbol("SP")

//line preprocessing.rl:89
		(lex.Te) = (lex.P) + 1

		goto st11
	tr14:
//line preprocessing.rl:22

		lex.ReleaseSymbol("CR")

//line preprocessing.rl:90
		(lex.Te) = (lex.P) + 1

		goto st11
	tr16:
//line preprocessing.rl:98
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st11
	st11:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
//line NONE:1
		(lex.Ts) = (lex.P)

//line preprocessing.go:188
		switch data[(lex.P)] {
		case 9:
			goto tr13
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 32:
			goto tr13
		case 64:
			goto tr15
		}
		goto tr12
	tr15:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st12
	st12:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
//line preprocessing.go:212
		switch data[(lex.P)] {
		case 70:
			goto st2
		case 84:
			goto st7
		}
		if 65 <= data[(lex.P)] && data[(lex.P)] <= 90 {
			goto st1
		}
		goto tr16
	st1:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof1
		}
	st_case_1:
		switch data[(lex.P)] {
		case 40:
			goto tr1
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
	st2:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof2
		}
	st_case_2:
		switch data[(lex.P)] {
		case 40:
			goto tr1
		case 95:
			goto st1
		case 97:
			goto st3
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
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
		switch data[(lex.P)] {
		case 40:
			goto tr1
		case 95:
			goto st1
		case 108:
			goto st4
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
	st4:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
		switch data[(lex.P)] {
		case 40:
			goto tr1
		case 95:
			goto st1
		case 115:
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
		case 40:
			goto tr1
		case 95:
			goto st1
		case 101:
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
	st7:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
		switch data[(lex.P)] {
		case 40:
			goto tr1
		case 95:
			goto st1
		case 114:
			goto st8
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
		case 40:
			goto tr1
		case 95:
			goto st1
		case 117:
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
	tr8:
//line NONE:1
		switch lex.Act {
		case 0:
			{
				{
					goto st0
				}
			}
		default:
			{
				(lex.P) = (lex.Te) - 1
			}
		}

		goto st13
	tr19:
//line preprocessing.rl:38
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
		goto st13
	tr21:
//line preprocessing.rl:44
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st13
	st13:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
//line NONE:1
		(lex.Ts) = (lex.P)

//line preprocessing.go:464
		switch data[(lex.P)] {
		case 39:
			goto tr19
		case 92:
			goto st9
		}
		goto tr9
	tr9:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line preprocessing.rl:44
		(lex.Act) = 2
		goto st14
	st14:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof14
		}
	st_case_14:
//line preprocessing.go:484
		switch data[(lex.P)] {
		case 39:
			goto tr21
		case 92:
			goto st9
		}
		goto tr9
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
		goto tr9
	tr10:
//line NONE:1
		switch lex.Act {
		case 0:
			{
				{
					goto st0
				}
			}
		default:
			{
				(lex.P) = (lex.Te) - 1
			}
		}

		goto st15
	tr22:
//line preprocessing.rl:48
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
		goto st15
	tr24:
//line preprocessing.rl:54
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st15
	st15:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
//line NONE:1
		(lex.Ts) = (lex.P)

//line preprocessing.go:538
		switch data[(lex.P)] {
		case 34:
			goto tr22
		case 92:
			goto st10
		}
		goto tr11
	tr11:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line preprocessing.rl:54
		(lex.Act) = 4
		goto st16
	st16:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof16
		}
	st_case_16:
//line preprocessing.go:558
		switch data[(lex.P)] {
		case 34:
			goto tr24
		case 92:
			goto st10
		}
		goto tr11
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
		goto tr11
	tr25:
//line preprocessing.rl:83
		(lex.Te) = (lex.P) + 1

		goto st17
	tr26:
//line preprocessing.rl:22

		lex.ReleaseSymbol("CR")

//line preprocessing.rl:58
		(lex.Te) = (lex.P) + 1

		goto st17
	tr27:
//line preprocessing.rl:65
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 17
					(lex.Top)++
					goto st15
				}
			}
		}
		goto st17
	tr28:
//line preprocessing.rl:60
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 17
					(lex.Top)++
					goto st13
				}
			}
		}
		goto st17
	tr29:
//line preprocessing.rl:71
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar(')')
		}
		goto st17
	tr30:
//line preprocessing.rl:74
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar(endAnnotation) {
				lex.ReleaseToken(endAnnotation, "annotation")
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
			if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
			}
		}
		goto st17
	st17:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof17
		}
	st_case_17:
//line NONE:1
		(lex.Ts) = (lex.P)

//line preprocessing.go:632
		switch data[(lex.P)] {
		case 10:
			goto tr26
		case 13:
			goto tr26
		case 34:
			goto tr27
		case 39:
			goto tr28
		case 40:
			goto tr29
		case 41:
			goto tr30
		}
		goto tr25
	st_case_0:
	st0:
		(lex.Cs) = 0
		goto _out
	st_out:
	_test_eof11:
		(lex.Cs) = 11
		goto _test_eof
	_test_eof12:
		(lex.Cs) = 12
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
	_test_eof13:
		(lex.Cs) = 13
		goto _test_eof
	_test_eof14:
		(lex.Cs) = 14
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
	_test_eof10:
		(lex.Cs) = 10
		goto _test_eof
	_test_eof17:
		(lex.Cs) = 17
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 12:
				goto tr16
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
			case 14:
				goto tr21
			case 9:
				goto tr8
			case 16:
				goto tr24
			case 10:
				goto tr10
			}
		}

	_out:
		{
		}
	}

//line preprocessing.rl:113

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
