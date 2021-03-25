//line extractor.rl:1
package annotation

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
)

//line extractor.rl:109

//line extractor.go:16
const annotation_extractor_start int = 6
const annotation_extractor_first_final int = 6
const annotation_extractor_error int = 0

const annotation_extractor_en_singleQuoteString int = 14
const annotation_extractor_en_doubleQuoteString int = 16
const annotation_extractor_en_annotationBody int = 18
const annotation_extractor_en_main int = 6

//line extractor.rl:112

func newPreprocessing(data []byte) (*syntax.Lexer, error) {
	lex := syntax.NewLexer(data)

//line extractor.go:33
	{
		(lex.Cs) = annotation_extractor_start
		(lex.Top) = 0
		(lex.Ts) = 0
		(lex.Te) = 0
		(lex.Act) = 0
	}

//line extractor.rl:117

//line extractor.go:44
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
		case 9:
			goto st9
		case 10:
			goto st10
		case 11:
			goto st11
		case 12:
			goto st12
		case 3:
			goto st3
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 4:
			goto st4
		case 16:
			goto st16
		case 17:
			goto st17
		case 5:
			goto st5
		case 18:
			goto st18
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
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 3:
			goto st_case_3
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 4:
			goto st_case_4
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 5:
			goto st_case_5
		case 18:
			goto st_case_18
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line NONE:1
		switch lex.Act {
		case 13:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(beginAnnotation, "annotation")
				lex.ReleaseToken(endAnnotation, "annotation")
			}
		default:
			{
				(lex.P) = (lex.Te) - 1
			}
		}

		goto st6
	tr1:
//line extractor.rl:93
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(beginAnnotation, "annotation")
			lex.ReleaseToken(endAnnotation, "annotation")
		}
		goto st6
	tr2:
//line extractor.rl:97
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(beginAnnotation, "annotation")
			lex.BeginPairedChar(endAnnotation)
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 6
					(lex.Top)++
					goto st18
				}
			}
		}
		goto st6
	tr4:
//line extractor.rl:103
		(lex.P) = (lex.Te) - 1

		goto st6
	tr11:
//line extractor.rl:103
		(lex.Te) = (lex.P) + 1

		goto st6
	tr12:
//line extractor.rl:25

		lex.ReleaseSymbol("SP")

//line extractor.rl:89
		(lex.Te) = (lex.P) + 1

		goto st6
	tr13:
//line extractor.rl:22

		lex.ReleaseSymbol("CR")

//line extractor.rl:90
		(lex.Te) = (lex.P) + 1

		goto st6
	tr15:
//line extractor.rl:103
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st6
	tr19:
//line extractor.rl:93
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(beginAnnotation, "annotation")
			lex.ReleaseToken(endAnnotation, "annotation")
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

//line extractor.go:224
		switch data[(lex.P)] {
		case 9:
			goto tr12
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 32:
			goto tr12
		case 64:
			goto tr14
		}
		goto tr11
	tr14:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:103
		(lex.Act) = 15
		goto st7
	st7:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
//line extractor.go:250
		switch data[(lex.P)] {
		case 70:
			goto st2
		case 84:
			goto st3
		}
		if 65 <= data[(lex.P)] && data[(lex.P)] <= 90 {
			goto st1
		}
		goto tr15
	st1:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof1
		}
	st_case_1:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	st8:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof8
		}
	st_case_8:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	st2:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof2
		}
	st_case_2:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		case 97:
			goto st9
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		case 108:
			goto st10
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		case 115:
			goto st11
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	st11:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		case 101:
			goto tr22
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	tr22:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:93
		(lex.Act) = 13
		goto st12
	st12:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
//line extractor.go:425
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
		goto tr19
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		case 114:
			goto st13
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	st13:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
		switch data[(lex.P)] {
		case 40:
			goto tr2
		case 95:
			goto st8
		case 117:
			goto st11
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st8
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto tr1
	tr7:
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

		goto st14
	tr23:
//line extractor.rl:38
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
		goto st14
	tr25:
//line extractor.rl:44
		(lex.Te) = (lex.P)
		(lex.P)--

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

//line extractor.go:534
		switch data[(lex.P)] {
		case 39:
			goto tr23
		case 92:
			goto st4
		}
		goto tr8
	tr8:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:44
		(lex.Act) = 2
		goto st15
	st15:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
//line extractor.go:554
		switch data[(lex.P)] {
		case 39:
			goto tr25
		case 92:
			goto st4
		}
		goto tr8
	st4:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
		goto tr8
	tr9:
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

		goto st16
	tr26:
//line extractor.rl:48
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
		goto st16
	tr28:
//line extractor.rl:54
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st16
	st16:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof16
		}
	st_case_16:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:608
		switch data[(lex.P)] {
		case 34:
			goto tr26
		case 92:
			goto st5
		}
		goto tr10
	tr10:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:54
		(lex.Act) = 4
		goto st17
	st17:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof17
		}
	st_case_17:
//line extractor.go:628
		switch data[(lex.P)] {
		case 34:
			goto tr28
		case 92:
			goto st5
		}
		goto tr10
	st5:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof5
		}
	st_case_5:
		goto tr10
	tr29:
//line extractor.rl:83
		(lex.Te) = (lex.P) + 1

		goto st18
	tr30:
//line extractor.rl:22

		lex.ReleaseSymbol("CR")

//line extractor.rl:58
		(lex.Te) = (lex.P) + 1

		goto st18
	tr31:
//line extractor.rl:65
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 18
					(lex.Top)++
					goto st16
				}
			}
		}
		goto st18
	tr32:
//line extractor.rl:60
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 18
					(lex.Top)++
					goto st14
				}
			}
		}
		goto st18
	tr33:
//line extractor.rl:71
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar(')')
		}
		goto st18
	tr34:
//line extractor.rl:74
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
		goto st18
	st18:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof18
		}
	st_case_18:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:702
		switch data[(lex.P)] {
		case 10:
			goto tr30
		case 13:
			goto tr30
		case 34:
			goto tr31
		case 39:
			goto tr32
		case 40:
			goto tr33
		case 41:
			goto tr34
		}
		goto tr29
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
	_test_eof3:
		(lex.Cs) = 3
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
	_test_eof4:
		(lex.Cs) = 4
		goto _test_eof
	_test_eof16:
		(lex.Cs) = 16
		goto _test_eof
	_test_eof17:
		(lex.Cs) = 17
		goto _test_eof
	_test_eof5:
		(lex.Cs) = 5
		goto _test_eof
	_test_eof18:
		(lex.Cs) = 18
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 7:
				goto tr15
			case 1:
				goto tr0
			case 8:
				goto tr19
			case 2:
				goto tr4
			case 9:
				goto tr19
			case 10:
				goto tr19
			case 11:
				goto tr19
			case 12:
				goto tr19
			case 3:
				goto tr4
			case 13:
				goto tr19
			case 15:
				goto tr25
			case 4:
				goto tr7
			case 17:
				goto tr28
			case 5:
				goto tr9
			}
		}

	_out:
		{
		}
	}

//line extractor.rl:118

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
