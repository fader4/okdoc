//line extractor.rl:1
package starlark

import (
	"github.com/fader4/okdoc/syntax"
)

//line extractor.rl:200

//line extractor.go:15
const starlark_extractor_start int = 29
const starlark_extractor_first_final int = 29
const starlark_extractor_error int = 0

const starlark_extractor_en_singleQuoteString int = 37
const starlark_extractor_en_doubleQuoteString int = 39
const starlark_extractor_en_defMethodBody int = 41
const starlark_extractor_en_loadBody int = 43
const starlark_extractor_en_moduleBody int = 44
const starlark_extractor_en_commentMultilineBody int = 45
const starlark_extractor_en_main int = 29

//line extractor.rl:203

func newPreprocessing(data []byte) (*syntax.Lexer, error) {
	lex := syntax.NewLexer(data)

//line extractor.go:35
	{
		(lex.Cs) = starlark_extractor_start
		(lex.Top) = 0
		(lex.Ts) = 0
		(lex.Te) = 0
		(lex.Act) = 0
	}

//line extractor.rl:208

//line extractor.go:46
	{
		if (lex.P) == (lex.Pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.Cs {
		case 29:
			goto st29
		case 30:
			goto st30
		case 1:
			goto st1
		case 31:
			goto st31
		case 32:
			goto st32
		case 2:
			goto st2
		case 33:
			goto st33
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
		case 10:
			goto st10
		case 11:
			goto st11
		case 34:
			goto st34
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
		case 35:
			goto st35
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 36:
			goto st36
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 37:
			goto st37
		case 38:
			goto st38
		case 25:
			goto st25
		case 39:
			goto st39
		case 40:
			goto st40
		case 26:
			goto st26
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
		case 27:
			goto st27
		case 47:
			goto st47
		case 28:
			goto st28
		case 0:
			goto st0
		}

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof
		}
	_resume:
		switch lex.Cs {
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 1:
			goto st_case_1
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 2:
			goto st_case_2
		case 33:
			goto st_case_33
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
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 34:
			goto st_case_34
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
		case 35:
			goto st_case_35
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 36:
			goto st_case_36
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 25:
			goto st_case_25
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 26:
			goto st_case_26
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
		case 27:
			goto st_case_27
		case 47:
			goto st_case_47
		case 28:
			goto st_case_28
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line extractor.rl:194
		(lex.P) = (lex.Te) - 1

		goto st29
	tr1:
//line extractor.rl:185
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(commentMultiline, "comment")
			lex.BeginPairedChar(endCommentMultiline)
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 29
					(lex.Top)++
					goto st45
				}
			}
		}
		goto st29
	tr2:
//line NONE:1
		switch lex.Act {
		case 34:
			{
				(lex.P) = (lex.Te) - 1

				lex.ReleaseToken(returnKeyword, "return")
			}
		default:
			{
				(lex.P) = (lex.Te) - 1
			}
		}

		goto st29
	tr13:
//line extractor.rl:180
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(module, "module")
			lex.BeginPairedChar(endModule)
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 29
					(lex.Top)++
					goto st44
				}
			}
		}
		goto st29
	tr19:
//line extractor.rl:170
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(def, "def")
			lex.BeginPairedChar(endDef)
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 29
					(lex.Top)++
					goto st41
				}
			}
		}
		goto st29
	tr23:
//line extractor.rl:175
		(lex.Te) = (lex.P) + 1
		{
			lex.ReleaseToken(load, "load")
			lex.BeginPairedChar(endLoad)
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 29
					(lex.Top)++
					goto st43
				}
			}
		}
		goto st29
	tr34:
//line extractor.rl:194
		(lex.Te) = (lex.P) + 1

		goto st29
	tr35:
//line extractor.rl:24

		lex.ReleaseSymbol("SP")

//line extractor.rl:164
		(lex.Te) = (lex.P) + 1

		goto st29
	tr36:
//line extractor.rl:21

		lex.ReleaseSymbol("CR")

//line extractor.rl:165
		(lex.Te) = (lex.P) + 1

		goto st29
	tr44:
//line extractor.rl:194
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st29
	tr46:
//line extractor.rl:167
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			lex.ReleaseToken(commentInline, "comment")
		}
		goto st29
	st29:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof29
		}
	st_case_29:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:359
		switch data[(lex.P)] {
		case 9:
			goto tr35
		case 10:
			goto tr36
		case 13:
			goto tr36
		case 32:
			goto tr35
		case 34:
			goto tr37
		case 35:
			goto st31
		case 39:
			goto tr39
		case 95:
			goto tr40
		case 100:
			goto tr41
		case 108:
			goto tr42
		case 114:
			goto tr43
		}
		switch {
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto tr40
			}
		case data[(lex.P)] >= 65:
			goto tr40
		}
		goto tr34
	tr37:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st30
	st30:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof30
		}
	st_case_30:
//line extractor.go:403
		if data[(lex.P)] == 34 {
			goto st1
		}
		goto tr44
	st1:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof1
		}
	st_case_1:
		if data[(lex.P)] == 34 {
			goto tr1
		}
		goto tr0
	st31:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof31
		}
	st_case_31:
		switch data[(lex.P)] {
		case 10:
			goto tr46
		case 13:
			goto tr46
		}
		goto st31
	tr39:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st32
	st32:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof32
		}
	st_case_32:
//line extractor.go:439
		if data[(lex.P)] == 39 {
			goto st2
		}
		goto tr44
	st2:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof2
		}
	st_case_2:
		if data[(lex.P)] == 39 {
			goto tr1
		}
		goto tr0
	tr27:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:190
		(lex.Act) = 34
		goto st33
	tr40:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:194
		(lex.Act) = 35
		goto st33
	st33:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof33
		}
	st_case_33:
//line extractor.go:472
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr2
	tr3:
//line extractor.rl:24

		lex.ReleaseSymbol("SP")

		goto st3
	st3:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof3
		}
	st_case_3:
//line extractor.go:507
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		}
		goto tr2
	tr5:
//line extractor.rl:24

		lex.ReleaseSymbol("SP")

		goto st4
	st4:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof4
		}
	st_case_4:
//line extractor.go:528
		switch data[(lex.P)] {
		case 9:
			goto tr5
		case 32:
			goto tr5
		case 109:
			goto st5
		}
		goto tr2
	st5:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof5
		}
	st_case_5:
		if data[(lex.P)] == 111 {
			goto st6
		}
		goto tr2
	st6:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof6
		}
	st_case_6:
		if data[(lex.P)] == 100 {
			goto st7
		}
		goto tr2
	st7:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof7
		}
	st_case_7:
		if data[(lex.P)] == 117 {
			goto st8
		}
		goto tr2
	st8:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof8
		}
	st_case_8:
		if data[(lex.P)] == 108 {
			goto st9
		}
		goto tr2
	st9:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof9
		}
	st_case_9:
		if data[(lex.P)] == 101 {
			goto st10
		}
		goto tr2
	tr12:
//line extractor.rl:24

		lex.ReleaseSymbol("SP")

		goto st10
	st10:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof10
		}
	st_case_10:
//line extractor.go:594
		switch data[(lex.P)] {
		case 9:
			goto tr12
		case 32:
			goto tr12
		case 40:
			goto tr13
		}
		goto tr2
	st11:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof11
		}
	st_case_11:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr2
	tr41:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:194
		(lex.Act) = 35
		goto st34
	st34:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof34
		}
	st_case_34:
//line extractor.go:644
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 101:
			goto st12
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr44
	st12:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof12
		}
	st_case_12:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 102:
			goto st13
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st13:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof13
		}
	st_case_13:
		switch data[(lex.P)] {
		case 9:
			goto tr16
		case 32:
			goto tr16
		case 61:
			goto st4
		case 95:
			goto st16
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
	tr16:
//line extractor.rl:24

		lex.ReleaseSymbol("SP")

		goto st14
	st14:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof14
		}
	st_case_14:
//line extractor.go:739
		switch data[(lex.P)] {
		case 9:
			goto tr16
		case 32:
			goto tr16
		case 61:
			goto st4
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
		goto tr0
	st15:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof15
		}
	st_case_15:
		switch data[(lex.P)] {
		case 40:
			goto tr19
		case 95:
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
		goto tr0
	st16:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof16
		}
	st_case_16:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 40:
			goto tr19
		case 61:
			goto st4
		case 95:
			goto st16
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st16
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
	tr42:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:194
		(lex.Act) = 35
		goto st35
	st35:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof35
		}
	st_case_35:
//line extractor.go:825
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 111:
			goto st17
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr44
	st17:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof17
		}
	st_case_17:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 97:
			goto st18
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 98 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st18:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof18
		}
	st_case_18:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 100:
			goto st19
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st19:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof19
		}
	st_case_19:
		switch data[(lex.P)] {
		case 9:
			goto tr22
		case 32:
			goto tr22
		case 40:
			goto tr23
		case 61:
			goto st4
		case 95:
			goto st11
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	tr22:
//line extractor.rl:24

		lex.ReleaseSymbol("SP")

		goto st20
	st20:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof20
		}
	st_case_20:
//line extractor.go:952
		switch data[(lex.P)] {
		case 9:
			goto tr22
		case 32:
			goto tr22
		case 40:
			goto tr23
		case 61:
			goto st4
		}
		goto tr0
	tr43:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:194
		(lex.Act) = 35
		goto st36
	st36:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof36
		}
	st_case_36:
//line extractor.go:976
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 101:
			goto st21
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr44
	st21:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof21
		}
	st_case_21:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 116:
			goto st22
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st22:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof22
		}
	st_case_22:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 117:
			goto st23
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st23:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof23
		}
	st_case_23:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 114:
			goto st24
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st24:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof24
		}
	st_case_24:
		switch data[(lex.P)] {
		case 9:
			goto tr3
		case 32:
			goto tr3
		case 61:
			goto st4
		case 95:
			goto st11
		case 110:
			goto tr27
		}
		switch {
		case data[(lex.P)] < 65:
			if 48 <= data[(lex.P)] && data[(lex.P)] <= 57 {
				goto st11
			}
		case data[(lex.P)] > 90:
			if 97 <= data[(lex.P)] && data[(lex.P)] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	tr28:
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

		goto st37
	tr52:
//line extractor.rl:42
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
		goto st37
	tr54:
//line extractor.rl:48
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st37
	st37:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof37
		}
	st_case_37:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:1162
		switch data[(lex.P)] {
		case 10:
			goto tr51
		case 13:
			goto tr51
		case 39:
			goto tr52
		case 92:
			goto st25
		}
		goto tr29
	tr29:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:48
		(lex.Act) = 2
		goto st38
	tr51:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:21

		lex.ReleaseSymbol("CR")

//line extractor.rl:48
		(lex.Act) = 2
		goto st38
	st38:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof38
		}
	st_case_38:
//line extractor.go:1197
		switch data[(lex.P)] {
		case 10:
			goto tr51
		case 13:
			goto tr51
		case 39:
			goto tr54
		case 92:
			goto st25
		}
		goto tr29
	st25:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof25
		}
	st_case_25:
		goto tr29
	tr30:
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

		goto st39
	tr56:
//line extractor.rl:52
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
		goto st39
	tr58:
//line extractor.rl:58
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st39
	st39:
//line NONE:1
		(lex.Ts) = 0

//line NONE:1
		(lex.Act) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof39
		}
	st_case_39:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:1255
		switch data[(lex.P)] {
		case 10:
			goto tr55
		case 13:
			goto tr55
		case 34:
			goto tr56
		case 92:
			goto st26
		}
		goto tr31
	tr31:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:58
		(lex.Act) = 4
		goto st40
	tr55:
//line NONE:1
		(lex.Te) = (lex.P) + 1

//line extractor.rl:21

		lex.ReleaseSymbol("CR")

//line extractor.rl:58
		(lex.Act) = 4
		goto st40
	st40:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof40
		}
	st_case_40:
//line extractor.go:1290
		switch data[(lex.P)] {
		case 10:
			goto tr55
		case 13:
			goto tr55
		case 34:
			goto tr58
		case 92:
			goto st26
		}
		goto tr31
	st26:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof26
		}
	st_case_26:
		goto tr31
	tr59:
//line extractor.rl:89
		(lex.Te) = (lex.P) + 1

		goto st41
	tr60:
//line extractor.rl:21

		lex.ReleaseSymbol("CR")

//line extractor.rl:62
		(lex.Te) = (lex.P) + 1

		goto st41
	tr61:
//line extractor.rl:69
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 41
					(lex.Top)++
					goto st39
				}
			}
		}
		goto st41
	tr62:
//line extractor.rl:64
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 41
					(lex.Top)++
					goto st37
				}
			}
		}
		goto st41
	tr63:
//line extractor.rl:75
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar(')')
		}
		goto st41
	tr65:
//line extractor.rl:84
		(lex.Te) = (lex.P)
		(lex.P)--
		{
			if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
			}
		}
		goto st41
	tr66:
//line extractor.rl:78
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar(endDef) {
				lex.ReleaseToken(endDef, "def")
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
		}
		goto st41
	st41:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof41
		}
	st_case_41:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:1375
		switch data[(lex.P)] {
		case 10:
			goto tr60
		case 13:
			goto tr60
		case 34:
			goto tr61
		case 39:
			goto tr62
		case 40:
			goto tr63
		case 41:
			goto st42
		}
		goto tr59
	st42:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof42
		}
	st_case_42:
		if data[(lex.P)] == 58 {
			goto tr66
		}
		goto tr65
	tr67:
//line extractor.rl:118
		(lex.Te) = (lex.P) + 1

		goto st43
	tr68:
//line extractor.rl:21

		lex.ReleaseSymbol("CR")

//line extractor.rl:93
		(lex.Te) = (lex.P) + 1

		goto st43
	tr69:
//line extractor.rl:100
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 43
					(lex.Top)++
					goto st39
				}
			}
		}
		goto st43
	tr70:
//line extractor.rl:95
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 43
					(lex.Top)++
					goto st37
				}
			}
		}
		goto st43
	tr71:
//line extractor.rl:106
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar(')')
		}
		goto st43
	tr72:
//line extractor.rl:109
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar(endLoad) {
				lex.ReleaseToken(endLoad, "load")
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
			if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
			}
		}
		goto st43
	st43:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof43
		}
	st_case_43:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:1460
		switch data[(lex.P)] {
		case 10:
			goto tr68
		case 13:
			goto tr68
		case 34:
			goto tr69
		case 39:
			goto tr70
		case 40:
			goto tr71
		case 41:
			goto tr72
		}
		goto tr67
	tr73:
//line extractor.rl:147
		(lex.Te) = (lex.P) + 1

		goto st44
	tr74:
//line extractor.rl:21

		lex.ReleaseSymbol("CR")

//line extractor.rl:122
		(lex.Te) = (lex.P) + 1

		goto st44
	tr75:
//line extractor.rl:129
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('"')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 44
					(lex.Top)++
					goto st39
				}
			}
		}
		goto st44
	tr76:
//line extractor.rl:124
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar('\'')
			{
				lex.StackGrowth()
				{
					(lex.Stack)[(lex.Top)] = 44
					(lex.Top)++
					goto st37
				}
			}
		}
		goto st44
	tr77:
//line extractor.rl:135
		(lex.Te) = (lex.P) + 1
		{
			lex.BeginPairedChar(')')
		}
		goto st44
	tr78:
//line extractor.rl:138
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar(endModule) {
				lex.ReleaseToken(endModule, "module")
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
			if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
			}
		}
		goto st44
	st44:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof44
		}
	st_case_44:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:1536
		switch data[(lex.P)] {
		case 10:
			goto tr74
		case 13:
			goto tr74
		case 34:
			goto tr75
		case 39:
			goto tr76
		case 40:
			goto tr77
		case 41:
			goto tr78
		}
		goto tr73
	tr32:
//line extractor.rl:158
		(lex.P) = (lex.Te) - 1

		goto st45
	tr33:
//line extractor.rl:152
		(lex.Te) = (lex.P) + 1
		{
			if lex.IsEndPairedChar(endCommentMultiline) {
				lex.ReleaseToken(endCommentMultiline, "comment")
				{
					(lex.Top)--
					(lex.Cs) = (lex.Stack)[(lex.Top)]
					goto _again
				}
			}
		}
		goto st45
	tr79:
//line extractor.rl:158
		(lex.Te) = (lex.P) + 1

		goto st45
	tr80:
//line extractor.rl:21

		lex.ReleaseSymbol("CR")

//line extractor.rl:151
		(lex.Te) = (lex.P) + 1

		goto st45
	tr83:
//line extractor.rl:158
		(lex.Te) = (lex.P)
		(lex.P)--

		goto st45
	st45:
//line NONE:1
		(lex.Ts) = 0

		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof45
		}
	st_case_45:
//line NONE:1
		(lex.Ts) = (lex.P)

//line extractor.go:1598
		switch data[(lex.P)] {
		case 10:
			goto tr80
		case 13:
			goto tr80
		case 34:
			goto tr81
		case 39:
			goto tr82
		}
		goto tr79
	tr81:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st46
	st46:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof46
		}
	st_case_46:
//line extractor.go:1620
		if data[(lex.P)] == 34 {
			goto st27
		}
		goto tr83
	st27:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof27
		}
	st_case_27:
		if data[(lex.P)] == 34 {
			goto tr33
		}
		goto tr32
	tr82:
//line NONE:1
		(lex.Te) = (lex.P) + 1

		goto st47
	st47:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof47
		}
	st_case_47:
//line extractor.go:1644
		if data[(lex.P)] == 39 {
			goto st28
		}
		goto tr83
	st28:
		if (lex.P)++; (lex.P) == (lex.Pe) {
			goto _test_eof28
		}
	st_case_28:
		if data[(lex.P)] == 39 {
			goto tr33
		}
		goto tr32
	st_case_0:
	st0:
		(lex.Cs) = 0
		goto _out
	st_out:
	_test_eof29:
		(lex.Cs) = 29
		goto _test_eof
	_test_eof30:
		(lex.Cs) = 30
		goto _test_eof
	_test_eof1:
		(lex.Cs) = 1
		goto _test_eof
	_test_eof31:
		(lex.Cs) = 31
		goto _test_eof
	_test_eof32:
		(lex.Cs) = 32
		goto _test_eof
	_test_eof2:
		(lex.Cs) = 2
		goto _test_eof
	_test_eof33:
		(lex.Cs) = 33
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
	_test_eof10:
		(lex.Cs) = 10
		goto _test_eof
	_test_eof11:
		(lex.Cs) = 11
		goto _test_eof
	_test_eof34:
		(lex.Cs) = 34
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
	_test_eof35:
		(lex.Cs) = 35
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
	_test_eof36:
		(lex.Cs) = 36
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
	_test_eof37:
		(lex.Cs) = 37
		goto _test_eof
	_test_eof38:
		(lex.Cs) = 38
		goto _test_eof
	_test_eof25:
		(lex.Cs) = 25
		goto _test_eof
	_test_eof39:
		(lex.Cs) = 39
		goto _test_eof
	_test_eof40:
		(lex.Cs) = 40
		goto _test_eof
	_test_eof26:
		(lex.Cs) = 26
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
	_test_eof27:
		(lex.Cs) = 27
		goto _test_eof
	_test_eof47:
		(lex.Cs) = 47
		goto _test_eof
	_test_eof28:
		(lex.Cs) = 28
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.P) == (lex.EOF) {
			switch lex.Cs {
			case 30:
				goto tr44
			case 1:
				goto tr0
			case 31:
				goto tr46
			case 32:
				goto tr44
			case 2:
				goto tr0
			case 33:
				goto tr2
			case 3:
				goto tr2
			case 4:
				goto tr2
			case 5:
				goto tr2
			case 6:
				goto tr2
			case 7:
				goto tr2
			case 8:
				goto tr2
			case 9:
				goto tr2
			case 10:
				goto tr2
			case 11:
				goto tr2
			case 34:
				goto tr44
			case 12:
				goto tr0
			case 13:
				goto tr0
			case 14:
				goto tr0
			case 15:
				goto tr0
			case 16:
				goto tr0
			case 35:
				goto tr44
			case 17:
				goto tr0
			case 18:
				goto tr0
			case 19:
				goto tr0
			case 20:
				goto tr0
			case 36:
				goto tr44
			case 21:
				goto tr0
			case 22:
				goto tr0
			case 23:
				goto tr0
			case 24:
				goto tr0
			case 38:
				goto tr54
			case 25:
				goto tr28
			case 40:
				goto tr58
			case 26:
				goto tr30
			case 42:
				goto tr65
			case 46:
				goto tr83
			case 27:
				goto tr32
			case 47:
				goto tr83
			case 28:
				goto tr32
			}
		}

	_out:
		{
		}
	}

//line extractor.rl:209

	if err := lex.Valid(); err != nil {
		return lex, err
	}

	return lex, nil
}
