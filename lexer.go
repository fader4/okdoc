//line lexer.rl:1
package okdoc

import "fmt"

//line lexer.go:9
const okdoc_start int = 15
const okdoc_first_final int = 15
const okdoc_error int = 0

const okdoc_en_string int = 47
const okdoc_en_rstring int = 49
const okdoc_en_expr int = 15

//line lexer.rl:66

type lexer struct {
	// It must be an array containting the data to process.
	data []byte

	// Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
	pe int

	// Data pointer. In C/D code this variable is expected to be a pointer to the character data to process. It should be initialized to the beginning of the data block on every run of the machine. In Java and Ruby it is used as an offset to data and must be an integer. In this case it should be initialized to zero on every run of the machine.
	p int

	// This must be a pointer to character data. In Java and Ruby code this must be an integer. See Section 6.3 for more information.
	ts int

	// Also a pointer to character data.
	te int

	// This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
	act int

	// Current state. This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently. This variable may be modified from outside the execution loop, but not from within.
	cs int

	// This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
	top int

	// stack of waiting boundaries to close
	stackBoundary []int

	// This must be an array of integers. It is used to store integer values representing states. If the stack must resize dynamically the Pre-push and Post-Pop statements can be used to do this (Sections 5.6 and 5.7).
	stack [64]int
}

func newLexer(data []byte) *lexer {
	lex := &lexer{
		data: data,
		pe:   len(data),
	}

//line lexer.go:64
	{
		lex.cs = okdoc_start
		lex.top = 0
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line lexer.rl:110
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe

	// found token for parser
	foundToken := 0

	releaseToken := func(tok int, humanName string) {
		if tok == -1 {
			tok = int(lex.data[lex.ts])
		}
		foundToken = tok
		out.token = string(lex.data[lex.ts:lex.te])

		if yyDebug == 1 {
			fmt.Printf("lexer: Release %q: %q \n", humanName, string(lex.data[lex.ts:lex.te]))
		}
	}

//line lexer.go:97
	{
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 15:
			goto st15
		case 16:
			goto st16
		case 1:
			goto st1
		case 17:
			goto st17
		case 2:
			goto st2
		case 3:
			goto st3
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 4:
			goto st4
		case 5:
			goto st5
		case 6:
			goto st6
		case 7:
			goto st7
		case 21:
			goto st21
		case 22:
			goto st22
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 11:
			goto st11
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
		case 12:
			goto st12
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
		case 13:
			goto st13
		case 49:
			goto st49
		case 50:
			goto st50
		case 14:
			goto st14
		case 0:
			goto st0
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 1:
			goto st_case_1
		case 17:
			goto st_case_17
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
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
		case 12:
			goto st_case_12
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
		case 13:
			goto st_case_13
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 14:
			goto st_case_14
		case 0:
			goto st_case_0
		}
		goto st_out
	tr0:
//line lexer.rl:179
		(lex.p) = (lex.te) - 1

		goto st15
	tr2:
//line lexer.rl:167
		(lex.p) = (lex.te) - 1
		{
			releaseToken(floatLiteral, "floatLiteral")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	tr5:
//line lexer.rl:173
		(lex.p) = (lex.te) - 1
		{
			lex.beginBoundary('"')
			{
				if lex.nostack() {
					return foundToken
				}
				{
					lex.stack[lex.top] = 15
					lex.top++
					goto st49
				}
			}
		}
		goto st15
	tr10:
//line lexer.rl:150
		lex.te = (lex.p) + 1
		{
			releaseToken(commentMultiline, "commentMultiline")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	tr11:
//line lexer.rl:172
		(lex.p) = (lex.te) - 1
		{
			lex.beginBoundary('\'')
			{
				if lex.nostack() {
					return foundToken
				}
				{
					lex.stack[lex.top] = 15
					lex.top++
					goto st47
				}
			}
		}
		goto st15
	tr16:
//line lexer.rl:155
		(lex.p) = (lex.te) - 1
		{
			releaseToken(identToken, "identToken")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	tr23:
//line lexer.rl:179
		lex.te = (lex.p) + 1

		goto st15
	tr25:
//line lexer.rl:18
		out.totallines++
//line lexer.rl:139
		lex.te = (lex.p) + 1

		goto st15
	tr39:
//line lexer.rl:179
		lex.te = (lex.p)
		(lex.p)--

		goto st15
	tr41:
//line lexer.rl:167
		lex.te = (lex.p)
		(lex.p)--
		{
			releaseToken(floatLiteral, "floatLiteral")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	tr43:
//line lexer.rl:135
		lex.te = (lex.p)
		(lex.p)--
		{
			releaseToken(whiteSpace, "whiteSpace")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	tr44:
//line lexer.rl:173
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.beginBoundary('"')
			{
				if lex.nostack() {
					return foundToken
				}
				{
					lex.stack[lex.top] = 15
					lex.top++
					goto st49
				}
			}
		}
		goto st15
	tr46:
//line lexer.rl:146
		lex.te = (lex.p)
		(lex.p)--
		{
			releaseToken(commentInline, "commentInline")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	tr47:
//line lexer.rl:172
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.beginBoundary('\'')
			{
				if lex.nostack() {
					return foundToken
				}
				{
					lex.stack[lex.top] = 15
					lex.top++
					goto st47
				}
			}
		}
		goto st15
	tr49:
//line lexer.rl:163
		lex.te = (lex.p)
		(lex.p)--
		{
			releaseToken(integerLiteral, "integerLiteral")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	tr50:
//line NONE:1
		switch lex.act {
		case 7:
			{
				(lex.p) = (lex.te) - 1

				releaseToken(defKeyword, "defKeyword")
				{
					(lex.p)++
					lex.cs = 15
					goto _out
				}
			}
		case 10:
			{
				(lex.p) = (lex.te) - 1

				releaseToken(identToken, "identToken")
				{
					(lex.p)++
					lex.cs = 15
					goto _out
				}
			}
		case 11:
			{
				(lex.p) = (lex.te) - 1

				releaseToken(boolLiteral, "boolLiteral")
				{
					(lex.p)++
					lex.cs = 15
					goto _out
				}
			}
		}

		goto st15
	tr51:
//line lexer.rl:155
		lex.te = (lex.p)
		(lex.p)--
		{
			releaseToken(identToken, "identToken")
			{
				(lex.p)++
				lex.cs = 15
				goto _out
			}
		}
		goto st15
	st15:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
//line NONE:1
		lex.ts = (lex.p)

//line lexer.go:468
		switch lex.data[(lex.p)] {
		case 0:
			goto tr22
		case 9:
			goto st19
		case 10:
			goto tr25
		case 13:
			goto tr25
		case 32:
			goto st19
		case 34:
			goto tr26
		case 35:
			goto st21
		case 39:
			goto tr28
		case 46:
			goto st23
		case 48:
			goto st24
		case 70:
			goto st27
		case 84:
			goto st31
		case 100:
			goto st33
		case 101:
			goto st35
		case 105:
			goto tr36
		case 114:
			goto st39
		case 119:
			goto st44
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st25
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr23
	tr22:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st16
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
//line lexer.go:528
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

		goto st17
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
//line lexer.go:552
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
			goto st18
		}
		goto tr2
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st18
		}
		goto tr2
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st18
		}
		goto tr41
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 9:
			goto st19
		case 32:
			goto st19
		}
		goto tr43
	tr26:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st20
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
//line lexer.go:618
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
	tr7:
//line lexer.rl:18
		out.totallines++
		goto st5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line lexer.go:641
		switch lex.data[(lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto st6
		}
		goto st5
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto st7
		}
		goto st5
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 34:
			goto tr10
		}
		goto st5
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr46
		case 13:
			goto tr46
		}
		goto st21
	tr28:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st22
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
//line lexer.go:701
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
		goto tr11
	tr13:
//line lexer.rl:18
		out.totallines++
		goto st9
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer.go:724
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 39:
			goto st10
		}
		goto st9
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 39:
			goto st11
		}
		goto st9
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 39:
			goto tr10
		}
		goto st9
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr1
		}
		goto tr39
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st24
		}
		goto tr49
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		if lex.data[(lex.p)] == 46 {
			goto tr1
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto st25
		}
		goto tr49
	tr17:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:155
		lex.act = 10
		goto st26
	tr55:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:159
		lex.act = 11
		goto st26
	tr58:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:141
		lex.act = 7
		goto st26
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
//line lexer.go:818
		if lex.data[(lex.p)] == 95 {
			goto tr17
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr50
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 97:
			goto st28
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 98 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 108:
			goto st29
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 115:
			goto st30
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 101:
			goto tr55
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 114:
			goto st32
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 117:
			goto st30
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 101:
			goto st34
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 102:
			goto tr58
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 108:
			goto st36
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 115:
			goto tr60
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	tr60:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st37
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
//line lexer.go:1085
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 101:
			goto st12
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		if lex.data[(lex.p)] == 95 {
			goto tr17
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr16
	tr36:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st38
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
//line lexer.go:1136
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 102:
			goto st12
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 101:
			goto st40
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 116:
			goto st41
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 117:
			goto st42
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 114:
			goto tr65
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	tr65:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st43
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
//line lexer.go:1262
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 110:
			goto st12
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 104:
			goto st45
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 105:
			goto st46
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch lex.data[(lex.p)] {
		case 95:
			goto tr17
		case 108:
			goto tr60
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr17
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto tr17
			}
		default:
			goto tr17
		}
		goto tr51
	tr18:
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

				releaseToken(stringLiteral, "stringLiteral for single quotes")
				{
					(lex.p)++
					lex.cs = 47
					goto _out
				}
			}
		}

		goto st47
	tr68:
//line lexer.rl:34
		lex.te = (lex.p) + 1
		{
			if !lex.isEndBoundary('\'') {
				(lex.p)--

			}
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st47
	tr70:
//line lexer.rl:41
		lex.te = (lex.p)
		(lex.p)--
		{
			releaseToken(stringLiteral, "stringLiteral for single quotes")
			{
				(lex.p)++
				lex.cs = 47
				goto _out
			}
		}
		goto st47
	st47:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
//line NONE:1
		lex.ts = (lex.p)

//line lexer.go:1402
		switch lex.data[(lex.p)] {
		case 39:
			goto tr68
		case 92:
			goto st13
		}
		goto tr19
	tr19:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:41
		lex.act = 2
		goto st48
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
//line lexer.go:1422
		switch lex.data[(lex.p)] {
		case 39:
			goto tr70
		case 92:
			goto st13
		}
		goto tr19
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		goto tr19
	tr20:
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

				releaseToken(stringLiteral, "stringLiteral for double quotes")
				{
					(lex.p)++
					lex.cs = 49
					goto _out
				}
			}
		}

		goto st49
	tr71:
//line lexer.rl:48
		lex.te = (lex.p) + 1
		{
			if !lex.isEndBoundary('"') {
				(lex.p)--

			}
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st49
	tr73:
//line lexer.rl:55
		lex.te = (lex.p)
		(lex.p)--
		{
			releaseToken(stringLiteral, "stringLiteral for double quotes")
			{
				(lex.p)++
				lex.cs = 49
				goto _out
			}
		}
		goto st49
	st49:
//line NONE:1
		lex.ts = 0

//line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof49
		}
	st_case_49:
//line NONE:1
		lex.ts = (lex.p)

//line lexer.go:1484
		switch lex.data[(lex.p)] {
		case 34:
			goto tr71
		case 92:
			goto st14
		}
		goto tr21
	tr21:
//line NONE:1
		lex.te = (lex.p) + 1

//line lexer.rl:55
		lex.act = 4
		goto st50
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
//line lexer.go:1504
		switch lex.data[(lex.p)] {
		case 34:
			goto tr73
		case 92:
			goto st14
		}
		goto tr21
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		goto tr21
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	st_out:
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
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
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof22:
		lex.cs = 22
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
	_test_eof12:
		lex.cs = 12
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
	_test_eof44:
		lex.cs = 44
		goto _test_eof
	_test_eof45:
		lex.cs = 45
		goto _test_eof
	_test_eof46:
		lex.cs = 46
		goto _test_eof
	_test_eof47:
		lex.cs = 47
		goto _test_eof
	_test_eof48:
		lex.cs = 48
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof49:
		lex.cs = 49
		goto _test_eof
	_test_eof50:
		lex.cs = 50
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 16:
				goto tr39
			case 1:
				goto tr0
			case 17:
				goto tr41
			case 2:
				goto tr2
			case 3:
				goto tr2
			case 18:
				goto tr41
			case 19:
				goto tr43
			case 20:
				goto tr44
			case 4:
				goto tr5
			case 5:
				goto tr5
			case 6:
				goto tr5
			case 7:
				goto tr5
			case 21:
				goto tr46
			case 22:
				goto tr47
			case 8:
				goto tr11
			case 9:
				goto tr11
			case 10:
				goto tr11
			case 11:
				goto tr11
			case 23:
				goto tr39
			case 24:
				goto tr49
			case 25:
				goto tr49
			case 26:
				goto tr50
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
			case 12:
				goto tr16
			case 38:
				goto tr51
			case 39:
				goto tr51
			case 40:
				goto tr51
			case 41:
				goto tr51
			case 42:
				goto tr51
			case 43:
				goto tr51
			case 44:
				goto tr51
			case 45:
				goto tr51
			case 46:
				goto tr51
			case 48:
				goto tr70
			case 13:
				goto tr18
			case 50:
				goto tr73
			case 14:
				goto tr20
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:184

	return foundToken
}

func (lex *lexer) beginBoundary(fin int) {
	// Push
	lex.stackBoundary = append(lex.stackBoundary, fin)
}

func (lex *lexer) endBoundary() int {
	// Pop
	top := len(lex.stackBoundary) - 1
	last := lex.stackBoundary[top]
	lex.stackBoundary = lex.stackBoundary[:top]

	return last
}

func (lex *lexer) isEndBoundary(sym int) bool {
	if lex.top == 0 || len(lex.stackBoundary) == 0 {
		fmt.Println("ERR: does not close anything")
		return false
	}
	prevsym := lex.endBoundary()
	if prevsym != sym {
		fmt.Println("ERR: does not close")
		return false
	}
	return true
}

func (lex *lexer) nostack() bool {
	if lex.top != len(lex.stack) {
		return false
	}
	fmt.Println("ERR: exceeds recursion limit")
	return true
}

func (lex *lexer) Error(e string) {
	fmt.Println("lexer: Error:", e)
}
