//line py_lexer.rl:1
package okdoc

import _ "fmt"

//line py_lexer.rl:141

//line py_lexer.go:13
const py_lang_start int = 12
const py_lang_first_final int = 12
const py_lang_error int = 0

const py_lang_en_singleQuoteString int = 39
const py_lang_en_doubleQuoteString int = 41
const py_lang_en_main int = 12

//line py_lexer.rl:144

func pyLex(data []byte) (*lexer, error) {
	lex, eof := newLexer(data)

//line py_lexer.go:29
	{
		lex.cs = py_lang_start
		lex.top = 0
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line py_lexer.rl:149

//line py_lexer.go:40
	{
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 12:
			goto st12
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
		case 15:
			goto st15
		case 16:
			goto st16
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 8:
			goto st8
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
		case 10:
			goto st10
		case 41:
			goto st41
		case 42:
			goto st42
		case 11:
			goto st11
		case 0:
			goto st0
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 12:
			goto st_case_12
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
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
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
		case 10:
			goto st_case_10
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 11:
			goto st_case_11
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line py_lexer.rl:107
		(lex.p) = (lex.te) - 1
		{
			lex.beginPairedChar('"')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 12
					lex.top++
					goto st41
				}
			}
		}
		goto st12
	tr5:
//line py_lexer.rl:67
		lex.te = (lex.p) + 1
		{
			lex.releaseToken(commentMultiline)
		}
		goto st12
	tr6:
//line py_lexer.rl:102
		(lex.p) = (lex.te) - 1
		{
			lex.beginPairedChar('\'')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 12
					lex.top++
					goto st39
				}
			}
		}
		goto st12
	tr11:
//line py_lexer.rl:89
		(lex.p) = (lex.te) - 1
		{
			lex.releaseToken(ident)
		}
		goto st12
	tr17:
//line py_lexer.rl:135
		lex.te = (lex.p) + 1

		goto st12
	tr19:
//line py_lexer.rl:13

		lex.releaseNEL()

//line py_lexer.rl:62
		lex.te = (lex.p) + 1

		goto st12
	tr23:
//line py_lexer.rl:112
		lex.te = (lex.p) + 1
		{
			lex.releaseToken('(')
			lex.beginPairedChar(')')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 12
					lex.top++
					goto st12
				}
			}
		}
		goto st12
	tr24:
//line py_lexer.rl:128
		lex.te = (lex.p) + 1
		{
			if lex.isEndPairedChar(int(lex.data[lex.ts])) {
				lex.releaseToken(int(lex.data[lex.ts]))
				{
					lex.top--
					lex.cs = lex.stack[lex.top]
					goto _again
				}
			}
		}
		goto st12
	tr25:
//line py_lexer.rl:84
		lex.te = (lex.p) + 1
		{
			lex.releaseToken(int(lex.data[lex.ts]))
		}
		goto st12
	tr32:
//line py_lexer.rl:61
		lex.te = (lex.p)
		(lex.p)--

		goto st12
	tr33:
//line py_lexer.rl:107
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.beginPairedChar('"')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 12
					lex.top++
					goto st41
				}
			}
		}
		goto st12
	tr35:
//line py_lexer.rl:64
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(commentInline)
		}
		goto st12
	tr36:
//line py_lexer.rl:102
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.beginPairedChar('\'')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 12
					lex.top++
					goto st39
				}
			}
		}
		goto st12
	tr38:
//line NONE:1
		switch lex.act {
		case 9:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(loadKeyword)
			}
		case 10:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(moduleKeyword)
			}
		case 11:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(defKeyword)
			}
		case 12:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(returnKeyword)
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(ident)
			}
		}

		goto st12
	tr39:
//line py_lexer.rl:89
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(ident)
		}
		goto st12
	st12:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
//line NONE:1
		lex.ts = (lex.p)

//line py_lexer.go:382
		switch lex.data[(lex.p)] {
		case 9:
			goto tr18
		case 10:
			goto tr19
		case 13:
			goto tr19
		case 32:
			goto tr18
		case 34:
			goto tr20
		case 35:
			goto st15
		case 39:
			goto tr22
		case 40:
			goto tr23
		case 41:
			goto tr24
		case 61:
			goto tr25
		case 70:
			goto st18
		case 84:
			goto st22
		case 100:
			goto st24
		case 108:
			goto st26
		case 109:
			goto st29
		case 114:
			goto st34
		}
		switch {
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		case lex.data[(lex.p)] >= 65:
			goto tr12
		}
		goto tr17
	tr18:
//line py_lexer.rl:16

		lex.releaseWhiteSpace()

		goto st13
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
//line py_lexer.go:437
		switch lex.data[(lex.p)] {
		case 9:
			goto tr18
		case 32:
			goto tr18
		}
		goto tr32
	tr20:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st14
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
//line py_lexer.go:455
		if lex.data[(lex.p)] == 34 {
			goto st1
		}
		goto tr33
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if lex.data[(lex.p)] == 34 {
			goto st2
		}
		goto tr0
	tr2:
//line py_lexer.rl:25
		lex.countNewLinesInComments++
		goto st2
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
//line py_lexer.go:478
		switch lex.data[(lex.p)] {
		case 10:
			goto tr2
		case 13:
			goto tr2
		case 34:
			goto st3
		}
		goto st2
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr2
		case 13:
			goto tr2
		case 34:
			goto st4
		}
		goto st2
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr2
		case 13:
			goto tr2
		case 34:
			goto tr5
		}
		goto st2
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr35
		case 13:
			goto tr35
		}
		goto st15
	tr22:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st16
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
//line py_lexer.go:538
		if lex.data[(lex.p)] == 39 {
			goto st5
		}
		goto tr36
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if lex.data[(lex.p)] == 39 {
			goto st6
		}
		goto tr6
	tr8:
//line py_lexer.rl:24
		lex.countNewLinesInComments++
		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line py_lexer.go:561
		switch lex.data[(lex.p)] {
		case 10:
			goto tr8
		case 13:
			goto tr8
		case 39:
			goto st7
		}
		goto st6
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr8
		case 13:
			goto tr8
		case 39:
			goto st8
		}
		goto st6
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr8
		case 13:
			goto tr8
		case 39:
			goto tr5
		}
		goto st6
	tr12:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:89
		lex.act = 14
		goto st17
	tr46:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:78
		lex.act = 11
		goto st17
	tr49:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:72
		lex.act = 9
		goto st17
	tr54:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:75
		lex.act = 10
		goto st17
	tr59:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:81
		lex.act = 12
		goto st17
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
//line py_lexer.go:639
		if lex.data[(lex.p)] == 95 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr38
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 97:
			goto st19
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 108:
			goto st20
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 115:
			goto tr42
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	tr42:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st21
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
//line py_lexer.go:738
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 101:
			goto st9
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		if lex.data[(lex.p)] == 95 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr11
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 114:
			goto st23
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 117:
			goto tr42
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 101:
			goto st25
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 102:
			goto tr46
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 111:
			goto st27
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 97:
			goto st28
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 100:
			goto tr49
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 111:
			goto st30
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 100:
			goto st31
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 117:
			goto st32
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 108:
			goto st33
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 101:
			goto tr54
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 101:
			goto st35
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 116:
			goto st36
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 117:
			goto st37
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 114:
			goto st38
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr12
		case 110:
			goto tr59
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr12
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr39
	tr13:
//line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 2:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(stringLiteral)
			}
		}

		goto st39
	tr60:
//line py_lexer.rl:35
		lex.te = (lex.p) + 1
		{
			if lex.isEndPairedChar('\'') {
				{
					lex.top--
					lex.cs = lex.stack[lex.top]
					goto _again
				}
			}
		}
		goto st39
	tr62:
//line py_lexer.rl:41
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(stringLiteral)
		}
		goto st39
	st39:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
//line NONE:1
		lex.ts = (lex.p)

//line py_lexer.go:1231
		switch lex.data[(lex.p)] {
		case 39:
			goto tr60
		case 92:
			goto st10
		}
		goto tr14
	tr14:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:41
		lex.act = 2
		goto st40
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
//line py_lexer.go:1251
		switch lex.data[(lex.p)] {
		case 39:
			goto tr62
		case 92:
			goto st10
		}
		goto tr14
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		goto tr14
	tr15:
//line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 4:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(stringLiteral)
			}
		}

		goto st41
	tr63:
//line py_lexer.rl:47
		lex.te = (lex.p) + 1
		{
			if lex.isEndPairedChar('"') {
				{
					lex.top--
					lex.cs = lex.stack[lex.top]
					goto _again
				}
			}
		}
		goto st41
	tr65:
//line py_lexer.rl:53
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(stringLiteral)
		}
		goto st41
	st41:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
//line NONE:1
		lex.ts = (lex.p)

//line py_lexer.go:1309
		switch lex.data[(lex.p)] {
		case 34:
			goto tr63
		case 92:
			goto st11
		}
		goto tr16
	tr16:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:53
		lex.act = 4
		goto st42
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
//line py_lexer.go:1329
		switch lex.data[(lex.p)] {
		case 34:
			goto tr65
		case 92:
			goto st11
		}
		goto tr16
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		goto tr16
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	st_out:
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof18:
		lex.cs = 18
		goto _test_eof
	_test_eof19:
		lex.cs = 19
		goto _test_eof
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof22:
		lex.cs = 22
		goto _test_eof
	_test_eof23:
		lex.cs = 23
		goto _test_eof
	_test_eof24:
		lex.cs = 24
		goto _test_eof
	_test_eof25:
		lex.cs = 25
		goto _test_eof
	_test_eof26:
		lex.cs = 26
		goto _test_eof
	_test_eof27:
		lex.cs = 27
		goto _test_eof
	_test_eof28:
		lex.cs = 28
		goto _test_eof
	_test_eof29:
		lex.cs = 29
		goto _test_eof
	_test_eof30:
		lex.cs = 30
		goto _test_eof
	_test_eof31:
		lex.cs = 31
		goto _test_eof
	_test_eof32:
		lex.cs = 32
		goto _test_eof
	_test_eof33:
		lex.cs = 33
		goto _test_eof
	_test_eof34:
		lex.cs = 34
		goto _test_eof
	_test_eof35:
		lex.cs = 35
		goto _test_eof
	_test_eof36:
		lex.cs = 36
		goto _test_eof
	_test_eof37:
		lex.cs = 37
		goto _test_eof
	_test_eof38:
		lex.cs = 38
		goto _test_eof
	_test_eof39:
		lex.cs = 39
		goto _test_eof
	_test_eof40:
		lex.cs = 40
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof41:
		lex.cs = 41
		goto _test_eof
	_test_eof42:
		lex.cs = 42
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 13:
				goto tr32
			case 14:
				goto tr33
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr0
			case 4:
				goto tr0
			case 15:
				goto tr35
			case 16:
				goto tr36
			case 5:
				goto tr6
			case 6:
				goto tr6
			case 7:
				goto tr6
			case 8:
				goto tr6
			case 17:
				goto tr38
			case 18:
				goto tr39
			case 19:
				goto tr39
			case 20:
				goto tr39
			case 21:
				goto tr39
			case 9:
				goto tr11
			case 22:
				goto tr39
			case 23:
				goto tr39
			case 24:
				goto tr39
			case 25:
				goto tr39
			case 26:
				goto tr39
			case 27:
				goto tr39
			case 28:
				goto tr39
			case 29:
				goto tr39
			case 30:
				goto tr39
			case 31:
				goto tr39
			case 32:
				goto tr39
			case 33:
				goto tr39
			case 34:
				goto tr39
			case 35:
				goto tr39
			case 36:
				goto tr39
			case 37:
				goto tr39
			case 38:
				goto tr39
			case 40:
				goto tr62
			case 10:
				goto tr13
			case 42:
				goto tr65
			case 11:
				goto tr15
			}
		}

	_out:
		{
		}
	}

//line py_lexer.rl:150

	return lex.validAndReturn()
}
