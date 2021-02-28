//line py_lexer.rl:1
package okdoc

import _ "fmt"

//line py_lexer.rl:130

//line py_lexer.go:13
const py_lang_start int = 14
const py_lang_first_final int = 14
const py_lang_error int = 0

const py_lang_en_singleQuoteString int = 42
const py_lang_en_doubleQuoteString int = 44
const py_lang_en_main int = 14

//line py_lexer.rl:133

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

//line py_lexer.rl:138

//line py_lexer.go:40
	{
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 14:
			goto st14
		case 15:
			goto st15
		case 1:
			goto st1
		case 16:
			goto st16
		case 2:
			goto st2
		case 3:
			goto st3
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 4:
			goto st4
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 20:
			goto st20
		case 21:
			goto st21
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 11:
			goto st11
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
		case 12:
			goto st12
		case 44:
			goto st44
		case 45:
			goto st45
		case 13:
			goto st13
		case 0:
			goto st0
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 1:
			goto st_case_1
		case 16:
			goto st_case_16
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
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
		case 12:
			goto st_case_12
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 13:
			goto st_case_13
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line py_lexer.rl:124
		(lex.p) = (lex.te) - 1

		goto st14
	tr2:
//line py_lexer.rl:88
		(lex.p) = (lex.te) - 1
		{
			lex.releaseToken(floatLiteral)
		}
		goto st14
	tr5:
//line py_lexer.rl:97
		(lex.p) = (lex.te) - 1
		{
			lex.beginPairedChar('"')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 14
					lex.top++
					goto st44
				}
			}
		}
		goto st14
	tr9:
//line py_lexer.rl:60
		lex.te = (lex.p) + 1
		{
			lex.releaseToken(commentMultiline)
		}
		goto st14
	tr10:
//line py_lexer.rl:92
		(lex.p) = (lex.te) - 1
		{
			lex.beginPairedChar('\'')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 14
					lex.top++
					goto st42
				}
			}
		}
		goto st14
	tr19:
//line py_lexer.rl:124
		lex.te = (lex.p) + 1

		goto st14
	tr21:
//line py_lexer.rl:55
		lex.te = (lex.p) + 1

		goto st14
	tr25:
//line py_lexer.rl:102
		lex.te = (lex.p) + 1
		{
			lex.releaseToken('(')
			lex.beginPairedChar(')')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 14
					lex.top++
					goto st14
				}
			}
		}
		goto st14
	tr26:
//line py_lexer.rl:117
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
		goto st14
	tr27:
//line py_lexer.rl:74
		lex.te = (lex.p) + 1
		{
			lex.releaseToken(int(lex.data[lex.ts]))
		}
		goto st14
	tr34:
//line py_lexer.rl:107
		lex.te = (lex.p) + 1
		{
			lex.releaseToken('[')
			lex.beginPairedChar(']')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 14
					lex.top++
					goto st14
				}
			}
		}
		goto st14
	tr38:
//line py_lexer.rl:112
		lex.te = (lex.p) + 1
		{
			lex.releaseToken('{')
			lex.beginPairedChar('}')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 14
					lex.top++
					goto st14
				}
			}
		}
		goto st14
	tr39:
//line py_lexer.rl:124
		lex.te = (lex.p)
		(lex.p)--

		goto st14
	tr41:
//line py_lexer.rl:88
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(floatLiteral)
		}
		goto st14
	tr43:
//line py_lexer.rl:54
		lex.te = (lex.p)
		(lex.p)--

		goto st14
	tr44:
//line py_lexer.rl:97
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.beginPairedChar('"')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 14
					lex.top++
					goto st44
				}
			}
		}
		goto st14
	tr46:
//line py_lexer.rl:57
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(commentInline)
		}
		goto st14
	tr47:
//line py_lexer.rl:92
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.beginPairedChar('\'')
			{
				lex.stackGrowth()
				{
					lex.stack[lex.top] = 14
					lex.top++
					goto st42
				}
			}
		}
		goto st14
	tr49:
//line py_lexer.rl:85
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(integerLiteral)
		}
		goto st14
	tr50:
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

				lex.releaseToken(defKeyword)
			}
		case 11:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(returnKeyword)
			}
		case 13:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(identToken)
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1

				lex.releaseToken(boolLiteral)
			}
		}

		goto st14
	tr51:
//line py_lexer.rl:79
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(identToken)
		}
		goto st14
	st14:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
//line NONE:1
		lex.ts = (lex.p)

//line py_lexer.go:435
		switch lex.data[(lex.p)] {
		case 0:
			goto tr18
		case 9:
			goto st18
		case 10:
			goto tr21
		case 13:
			goto tr21
		case 32:
			goto st18
		case 34:
			goto tr22
		case 35:
			goto st20
		case 39:
			goto tr24
		case 40:
			goto tr25
		case 41:
			goto tr26
		case 42:
			goto tr27
		case 46:
			goto st22
		case 48:
			goto st23
		case 58:
			goto tr27
		case 61:
			goto tr27
		case 70:
			goto st26
		case 84:
			goto st30
		case 91:
			goto tr34
		case 93:
			goto tr26
		case 100:
			goto st32
		case 108:
			goto st34
		case 114:
			goto st37
		case 123:
			goto tr38
		case 125:
			goto tr26
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st24
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr19
	tr18:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st15
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
//line py_lexer.go:509
		if lex.data[(lex.p)] == 46 {
			goto st1
		}
		goto tr39
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr1
		}
		goto tr0
	tr1:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st16
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
//line py_lexer.go:533
		switch lex.data[(lex.p)] {
		case 69:
			goto st2
		case 101:
			goto st2
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr1
		}
		goto tr41
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 43:
			goto st3
		case 45:
			goto st3
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st17
		}
		goto tr2
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st17
		}
		goto tr2
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st17
		}
		goto tr41
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 9:
			goto st18
		case 32:
			goto st18
		}
		goto tr43
	tr22:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st19
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
//line py_lexer.go:599
		if lex.data[(lex.p)] == 34 {
			goto st4
		}
		goto tr44
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		if lex.data[(lex.p)] == 34 {
			goto st5
		}
		goto tr5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if lex.data[(lex.p)] == 34 {
			goto st6
		}
		goto st5
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if lex.data[(lex.p)] == 34 {
			goto st7
		}
		goto st5
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if lex.data[(lex.p)] == 34 {
			goto tr9
		}
		goto st5
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr46
		case 13:
			goto tr46
		}
		goto st20
	tr24:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st21
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
//line py_lexer.go:662
		if lex.data[(lex.p)] == 39 {
			goto st8
		}
		goto tr47
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		if lex.data[(lex.p)] == 39 {
			goto st9
		}
		goto tr10
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		if lex.data[(lex.p)] == 39 {
			goto st10
		}
		goto st9
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		if lex.data[(lex.p)] == 39 {
			goto st11
		}
		goto st9
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		if lex.data[(lex.p)] == 39 {
			goto tr9
		}
		goto st9
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr1
		}
		goto tr39
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st23
		}
		goto tr49
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		if lex.data[(lex.p)] == 46 {
			goto tr1
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st24
		}
		goto tr49
	tr31:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:79
		lex.act = 13
		goto st25
	tr55:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:82
		lex.act = 14
		goto st25
	tr58:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:68
		lex.act = 10
		goto st25
	tr61:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:65
		lex.act = 9
		goto st25
	tr66:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:71
		lex.act = 11
		goto st25
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
//line py_lexer.go:773
		if lex.data[(lex.p)] == 95 {
			goto tr31
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr50
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 97:
			goto st27
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 108:
			goto st28
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 115:
			goto st29
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 101:
			goto tr55
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 114:
			goto st31
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 117:
			goto st29
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 101:
			goto st33
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 102:
			goto tr58
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 111:
			goto st35
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 97:
			goto st36
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 100:
			goto tr61
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 101:
			goto st38
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 116:
			goto st39
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 117:
			goto st40
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 114:
			goto st41
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr31
		case 110:
			goto tr66
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr51
	tr14:
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

		goto st42
	tr67:
//line py_lexer.rl:28
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
		goto st42
	tr69:
//line py_lexer.rl:34
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(stringLiteral)
		}
		goto st42
	st42:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
//line NONE:1
		lex.ts = (lex.p)

//line py_lexer.go:1218
		switch lex.data[(lex.p)] {
		case 39:
			goto tr67
		case 92:
			goto st12
		}
		goto tr15
	tr15:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:34
		lex.act = 2
		goto st43
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
//line py_lexer.go:1238
		switch lex.data[(lex.p)] {
		case 39:
			goto tr69
		case 92:
			goto st12
		}
		goto tr15
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		goto tr15
	tr16:
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

		goto st44
	tr70:
//line py_lexer.rl:40
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
		goto st44
	tr72:
//line py_lexer.rl:46
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.releaseToken(stringLiteral)
		}
		goto st44
	st44:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
//line NONE:1
		lex.ts = (lex.p)

//line py_lexer.go:1296
		switch lex.data[(lex.p)] {
		case 34:
			goto tr70
		case 92:
			goto st13
		}
		goto tr17
	tr17:
//line NONE:1
		lex.te = (lex.p) + 1

//line py_lexer.rl:46
		lex.act = 4
		goto st45
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
//line py_lexer.go:1316
		switch lex.data[(lex.p)] {
		case 34:
			goto tr72
		case 92:
			goto st13
		}
		goto tr17
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		goto tr17
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	st_out:
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
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
	_test_eof4:
		lex.cs = 4
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
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof11:
		lex.cs = 11
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
	_test_eof41:
		lex.cs = 41
		goto _test_eof
	_test_eof42:
		lex.cs = 42
		goto _test_eof
	_test_eof43:
		lex.cs = 43
		goto _test_eof
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof44:
		lex.cs = 44
		goto _test_eof
	_test_eof45:
		lex.cs = 45
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 15:
				goto tr39
			case 1:
				goto tr0
			case 16:
				goto tr41
			case 2:
				goto tr2
			case 3:
				goto tr2
			case 17:
				goto tr41
			case 18:
				goto tr43
			case 19:
				goto tr44
			case 4:
				goto tr5
			case 5:
				goto tr5
			case 6:
				goto tr5
			case 7:
				goto tr5
			case 20:
				goto tr46
			case 21:
				goto tr47
			case 8:
				goto tr10
			case 9:
				goto tr10
			case 10:
				goto tr10
			case 11:
				goto tr10
			case 22:
				goto tr39
			case 23:
				goto tr49
			case 24:
				goto tr49
			case 25:
				goto tr50
			case 26:
				goto tr51
			case 27:
				goto tr51
			case 28:
				goto tr51
			case 29:
				goto tr51
			case 30:
				goto tr51
			case 31:
				goto tr51
			case 32:
				goto tr51
			case 33:
				goto tr51
			case 34:
				goto tr51
			case 35:
				goto tr51
			case 36:
				goto tr51
			case 37:
				goto tr51
			case 38:
				goto tr51
			case 39:
				goto tr51
			case 40:
				goto tr51
			case 41:
				goto tr51
			case 43:
				goto tr69
			case 12:
				goto tr14
			case 45:
				goto tr72
			case 13:
				goto tr16
			}
		}

	_out:
		{
		}
	}

//line py_lexer.rl:139

	return lex.validAndReturn()
}
