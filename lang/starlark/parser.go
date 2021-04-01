// Code generated by goyacc -p starlark -o parser.go parser.y. DO NOT EDIT.

//line parser.y:2
package starlark

import __yyfmt__ "fmt"

//line parser.y:2

import (
	"github.com/fader4/okdoc/syntax"
	"log"
	"strings"
)

//line parser.y:13
type starlarkSymType struct {
	yys   int
	token *syntax.TokenWithData

	Comment   *Comment
	Return    *Return
	Module    *Module
	Load      *Load
	Def       *Def
	DefField  *DefField
	DefFields []*DefField
	Assign    *Assign
	Assigns   []*Assign
	CallFunc  *CallFunc
	Operand   *Operand
	arrint    []int

	val  syntax.Value
	arr  syntax.Array
	map_ syntax.Map
}

const ident = 57346
const nullLiteral = 57347
const stringLiteral = 57348
const boolLiteral = 57349
const integerLiteral = 57350
const floatLiteral = 57351
const not = 57352
const or = 57353
const and = 57354
const returnKeyword = 57355
const commentInline = 57356
const def = 57357
const endDef = 57358
const load = 57359
const endLoad = 57360
const module = 57361
const endModule = 57362
const commentMultiline = 57363
const endCommentMultiline = 57364

var starlarkToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ident",
	"'('",
	"')'",
	"'='",
	"','",
	"'{'",
	"'}'",
	"'['",
	"']'",
	"'.'",
	"'\"'",
	"'\\''",
	"'*'",
	"':'",
	"'+'",
	"'-'",
	"'/'",
	"'>'",
	"'<'",
	"nullLiteral",
	"stringLiteral",
	"boolLiteral",
	"integerLiteral",
	"floatLiteral",
	"not",
	"or",
	"and",
	"returnKeyword",
	"commentInline",
	"def",
	"endDef",
	"load",
	"endLoad",
	"module",
	"endModule",
	"commentMultiline",
	"endCommentMultiline",
}

var starlarkStatenames = [...]string{}

const starlarkEofCode = 1
const starlarkErrCode = 2
const starlarkInitialStackSize = 16

//line yacctab:1
var starlarkExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const starlarkPrivate = 57344

const starlarkLast = 301

var starlarkAct = [...]int{
	83, 99, 59, 78, 64, 56, 88, 21, 36, 12,
	62, 58, 52, 29, 18, 91, 65, 86, 87, 92,
	89, 90, 20, 97, 65, 30, 47, 46, 84, 85,
	104, 68, 48, 52, 53, 103, 10, 8, 13, 45,
	11, 118, 19, 111, 9, 51, 110, 52, 82, 60,
	109, 52, 129, 30, 70, 82, 25, 74, 79, 94,
	93, 72, 15, 71, 69, 130, 96, 80, 95, 100,
	116, 93, 117, 101, 81, 102, 82, 67, 128, 66,
	93, 49, 23, 50, 24, 106, 108, 55, 26, 22,
	14, 121, 73, 16, 107, 2, 112, 1, 35, 33,
	114, 113, 98, 76, 63, 79, 119, 93, 123, 115,
	122, 34, 120, 17, 80, 93, 7, 124, 28, 125,
	100, 126, 105, 6, 93, 5, 4, 93, 52, 3,
	127, 91, 0, 86, 87, 92, 89, 90, 88, 0,
	0, 0, 0, 0, 84, 85, 0, 91, 0, 86,
	87, 92, 89, 90, 0, 0, 0, 32, 42, 75,
	84, 85, 44, 0, 43, 0, 0, 0, 0, 77,
	0, 0, 0, 0, 0, 0, 41, 37, 38, 39,
	40, 32, 42, 27, 0, 0, 44, 0, 43, 32,
	42, 0, 109, 31, 44, 0, 43, 0, 0, 0,
	41, 37, 38, 39, 40, 0, 0, 0, 41, 37,
	38, 39, 40, 32, 42, 0, 0, 0, 44, 0,
	43, 32, 42, 0, 0, 31, 44, 0, 43, 61,
	0, 0, 41, 37, 38, 39, 40, 0, 0, 0,
	41, 37, 38, 39, 40, 32, 42, 57, 0, 0,
	44, 0, 43, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 41, 37, 38, 39, 40, 32,
	42, 0, 0, 0, 44, 0, 43, 32, 42, 0,
	0, 54, 44, 0, 43, 0, 0, 0, 41, 37,
	38, 39, 40, 0, 0, 0, 41, 37, 38, 39,
	40,
}

var starlarkPact = [...]int{
	-1000, -1000, 5, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 85, 55, 89, 18, -30, 84, 76, -1000, -1000,
	49, 83, 177, -1000, 18, 3, 2, 15, 75, -1000,
	38, 265, 82, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 241, 217, 0, -1000, -1000, 71, -1000, 14,
	209, 273, 88, 20, 273, 153, 68, -1000, 131, -1,
	47, -1000, -1000, 58, -1000, 6, 273, -1000, -1000, -1000,
	-1, 131, -1000, 82, 20, -1000, 67, 19, -1000, 115,
	131, -1000, 273, 273, -1000, -1000, -1000, -1000, 43, 39,
	36, -1000, -1000, 273, -1000, -1000, -8, 273, 64, -1000,
	34, -1000, 273, 87, 273, 185, -1, 131, 20, -1000,
	-1000, -1000, 20, -1000, -1, 131, -1000, 273, 273, -1000,
	72, 82, 40, -1, 131, -1000, -1, 131, -1000, 59,
	-1000,
}

var starlarkPgo = [...]int{
	0, 129, 126, 125, 123, 13, 118, 8, 116, 2,
	11, 113, 111, 5, 104, 4, 103, 1, 102, 99,
	14, 3, 98, 0, 97, 95,
}

var starlarkR1 = [...]int{
	0, 24, 25, 25, 25, 25, 25, 25, 8, 8,
	6, 6, 6, 5, 5, 5, 5, 5, 1, 1,
	2, 3, 11, 11, 11, 20, 20, 4, 4, 18,
	18, 18, 17, 17, 7, 7, 7, 7, 16, 16,
	16, 21, 21, 21, 21, 19, 19, 19, 19, 19,
	12, 12, 12, 12, 13, 13, 13, 13, 13, 22,
	22, 14, 14, 14, 15, 15, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 10, 10, 9,
	9, 9, 9, 9, 9, 9,
}

var starlarkR2 = [...]int{
	0, 1, 0, 2, 2, 2, 2, 2, 5, 6,
	1, 3, 2, 1, 2, 3, 3, 3, 1, 1,
	1, 4, 1, 3, 2, 1, 3, 8, 6, 1,
	3, 2, 3, 3, 3, 4, 6, 7, 3, 1,
	2, 1, 1, 3, 3, 1, 1, 1, 1, 1,
	3, 3, 2, 2, 1, 1, 3, 3, 2, 2,
	3, 1, 3, 2, 3, 3, 1, 1, 1, 1,
	2, 1, 1, 2, 2, 1, 1, 3, 3, 1,
	1, 1, 1, 1, 3, 3,
}

var starlarkChk = [...]int{
	-1000, -24, -25, -1, -2, -3, -4, -8, 32, 39,
	31, 35, 4, 33, 5, 7, 4, -11, -20, 24,
	4, 37, 5, 6, 8, 7, 5, 6, -6, -5,
	-9, 16, 4, -19, -12, -22, -7, 24, 25, 26,
	27, 23, 5, 11, 9, -20, 24, 24, 17, 6,
	8, 7, 13, -9, 16, 5, -13, 6, -10, -9,
	-13, 12, 10, -14, -15, 24, 8, 6, 17, -5,
	-9, -10, -7, 4, -9, 6, -16, 16, -21, -9,
	-10, 6, 8, -23, 29, 30, 18, 19, 7, 21,
	22, 16, 20, -23, 12, 10, 8, 17, -18, -17,
	-9, 6, 8, 16, 11, 7, -9, -10, -9, 7,
	7, 7, -9, -15, -9, -10, 6, 8, 7, -21,
	-7, 4, -13, -9, -10, -17, -9, -10, 6, 12,
	6,
}

var starlarkDef = [...]int{
	2, -2, 1, 3, 4, 5, 6, 7, 18, 19,
	20, 0, 0, 0, 0, 0, 0, 0, 22, 25,
	0, 0, 0, 21, 24, 0, 0, 0, 0, 10,
	13, 0, 79, 80, 81, 82, 83, 45, 46, 47,
	48, 49, 0, 0, 0, 23, 26, 0, 8, 0,
	12, 0, 0, 14, 0, 0, 0, 52, 54, 55,
	0, 53, 59, 0, 61, 0, 0, 28, 9, 11,
	16, 17, 84, 85, 15, 34, 0, 0, 39, 41,
	42, 50, 58, 0, 66, 67, 68, 69, 0, 71,
	72, 75, 76, 0, 51, 60, 63, 0, 0, 29,
	0, 35, 40, 0, 0, 0, 56, 57, 78, 70,
	74, 73, 77, 62, 64, 65, 27, 31, 0, 38,
	0, 0, 0, 43, 44, 30, 32, 33, 36, 0,
	37,
}

var starlarkTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 14, 3, 3, 3, 3, 15,
	5, 6, 16, 18, 8, 19, 13, 20, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 17, 3,
	22, 7, 21, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 11, 3, 12, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 9, 3, 10,
}

var starlarkTok2 = [...]int{
	2, 3, 4, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40,
}

var starlarkTok3 = [...]int{
	0,
}

var starlarkErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	starlarkDebug        = 0
	starlarkErrorVerbose = false
)

type starlarkLexer interface {
	Lex(lval *starlarkSymType) int
	Error(s string)
}

type starlarkParser interface {
	Parse(starlarkLexer) int
	Lookahead() int
}

type starlarkParserImpl struct {
	lval  starlarkSymType
	stack [starlarkInitialStackSize]starlarkSymType
	char  int
}

func (p *starlarkParserImpl) Lookahead() int {
	return p.char
}

func starlarkNewParser() starlarkParser {
	return &starlarkParserImpl{}
}

const starlarkFlag = -1000

func starlarkTokname(c int) string {
	if c >= 1 && c-1 < len(starlarkToknames) {
		if starlarkToknames[c-1] != "" {
			return starlarkToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func starlarkStatname(s int) string {
	if s >= 0 && s < len(starlarkStatenames) {
		if starlarkStatenames[s] != "" {
			return starlarkStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func starlarkErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !starlarkErrorVerbose {
		return "syntax error"
	}

	for _, e := range starlarkErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + starlarkTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := starlarkPact[state]
	for tok := TOKSTART; tok-1 < len(starlarkToknames); tok++ {
		if n := base + tok; n >= 0 && n < starlarkLast && starlarkChk[starlarkAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if starlarkDef[state] == -2 {
		i := 0
		for starlarkExca[i] != -1 || starlarkExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; starlarkExca[i] >= 0; i += 2 {
			tok := starlarkExca[i]
			if tok < TOKSTART || starlarkExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if starlarkExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += starlarkTokname(tok)
	}
	return res
}

func starlarklex1(lex starlarkLexer, lval *starlarkSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = starlarkTok1[0]
		goto out
	}
	if char < len(starlarkTok1) {
		token = starlarkTok1[char]
		goto out
	}
	if char >= starlarkPrivate {
		if char < starlarkPrivate+len(starlarkTok2) {
			token = starlarkTok2[char-starlarkPrivate]
			goto out
		}
	}
	for i := 0; i < len(starlarkTok3); i += 2 {
		token = starlarkTok3[i+0]
		if token == char {
			token = starlarkTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = starlarkTok2[1] /* unknown char */
	}
	if starlarkDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", starlarkTokname(token), uint(char))
	}
	return char, token
}

func starlarkParse(starlarklex starlarkLexer) int {
	return starlarkNewParser().Parse(starlarklex)
}

func (starlarkrcvr *starlarkParserImpl) Parse(starlarklex starlarkLexer) int {
	var starlarkn int
	var starlarkVAL starlarkSymType
	var starlarkDollar []starlarkSymType
	_ = starlarkDollar // silence set and not used
	starlarkS := starlarkrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	starlarkstate := 0
	starlarkrcvr.char = -1
	starlarktoken := -1 // starlarkrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		starlarkstate = -1
		starlarkrcvr.char = -1
		starlarktoken = -1
	}()
	starlarkp := -1
	goto starlarkstack

ret0:
	return 0

ret1:
	return 1

starlarkstack:
	/* put a state and value onto the stack */
	if starlarkDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", starlarkTokname(starlarktoken), starlarkStatname(starlarkstate))
	}

	starlarkp++
	if starlarkp >= len(starlarkS) {
		nyys := make([]starlarkSymType, len(starlarkS)*2)
		copy(nyys, starlarkS)
		starlarkS = nyys
	}
	starlarkS[starlarkp] = starlarkVAL
	starlarkS[starlarkp].yys = starlarkstate

starlarknewstate:
	starlarkn = starlarkPact[starlarkstate]
	if starlarkn <= starlarkFlag {
		goto starlarkdefault /* simple state */
	}
	if starlarkrcvr.char < 0 {
		starlarkrcvr.char, starlarktoken = starlarklex1(starlarklex, &starlarkrcvr.lval)
	}
	starlarkn += starlarktoken
	if starlarkn < 0 || starlarkn >= starlarkLast {
		goto starlarkdefault
	}
	starlarkn = starlarkAct[starlarkn]
	if starlarkChk[starlarkn] == starlarktoken { /* valid shift */
		starlarkrcvr.char = -1
		starlarktoken = -1
		starlarkVAL = starlarkrcvr.lval
		starlarkstate = starlarkn
		if Errflag > 0 {
			Errflag--
		}
		goto starlarkstack
	}

starlarkdefault:
	/* default state action */
	starlarkn = starlarkDef[starlarkstate]
	if starlarkn == -2 {
		if starlarkrcvr.char < 0 {
			starlarkrcvr.char, starlarktoken = starlarklex1(starlarklex, &starlarkrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if starlarkExca[xi+0] == -1 && starlarkExca[xi+1] == starlarkstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			starlarkn = starlarkExca[xi+0]
			if starlarkn < 0 || starlarkn == starlarktoken {
				break
			}
		}
		starlarkn = starlarkExca[xi+1]
		if starlarkn < 0 {
			goto ret0
		}
	}
	if starlarkn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			starlarklex.Error(starlarkErrorMessage(starlarkstate, starlarktoken))
			Nerrs++
			if starlarkDebug >= 1 {
				__yyfmt__.Printf("%s", starlarkStatname(starlarkstate))
				__yyfmt__.Printf(" saw %s\n", starlarkTokname(starlarktoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for starlarkp >= 0 {
				starlarkn = starlarkPact[starlarkS[starlarkp].yys] + starlarkErrCode
				if starlarkn >= 0 && starlarkn < starlarkLast {
					starlarkstate = starlarkAct[starlarkn] /* simulate a shift of "error" */
					if starlarkChk[starlarkstate] == starlarkErrCode {
						goto starlarkstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if starlarkDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", starlarkS[starlarkp].yys)
				}
				starlarkp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if starlarkDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", starlarkTokname(starlarktoken))
			}
			if starlarktoken == starlarkEofCode {
				goto ret1
			}
			starlarkrcvr.char = -1
			starlarktoken = -1
			goto starlarknewstate /* try again in the same state */
		}
	}

	/* reduction by production starlarkn */
	if starlarkDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", starlarkn, starlarkStatname(starlarkstate))
	}

	starlarknt := starlarkn
	starlarkpt := starlarkp
	_ = starlarkpt // guard against "declared and not used"

	starlarkp -= starlarkR2[starlarkn]
	// starlarkp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if starlarkp+1 >= len(starlarkS) {
		nyys := make([]starlarkSymType, len(starlarkS)*2)
		copy(nyys, starlarkS)
		starlarkS = nyys
	}
	starlarkVAL = starlarkS[starlarkp+1]

	/* consult goto table to find next state */
	starlarkn = starlarkR1[starlarkn]
	starlarkg := starlarkPgo[starlarkn]
	starlarkj := starlarkg + starlarkS[starlarkp].yys + 1

	if starlarkj >= starlarkLast {
		starlarkstate = starlarkAct[starlarkg]
	} else {
		starlarkstate = starlarkAct[starlarkj]
		if starlarkChk[starlarkstate] != -starlarkn {
			starlarkstate = starlarkAct[starlarkg]
		}
	}
	// dummy call; replaced with literal code
	switch starlarknt {

	case 2:
		starlarkDollar = starlarkS[starlarkpt-0 : starlarkpt+1]
//line parser.y:74
		{
		}
	case 3:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:75
		{
			starlarklex.(*starlarkLex).Comment = starlarkDollar[2].Comment
		}
	case 4:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:77
		{
			starlarklex.(*starlarkLex).Return = starlarkDollar[2].Return
		}
	case 5:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:79
		{
			starlarklex.(*starlarkLex).Load = starlarkDollar[2].Load
		}
	case 6:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:81
		{
			starlarklex.(*starlarkLex).Module = starlarkDollar[2].Module
		}
	case 7:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:83
		{
			starlarklex.(*starlarkLex).Def = starlarkDollar[2].Def
		}
	case 8:
		starlarkDollar = starlarkS[starlarkpt-5 : starlarkpt+1]
//line parser.y:92
		{
			starlarkVAL.Def = &Def{TokenWithData: starlarkDollar[1].token, Name: starlarkDollar[2].token.Ident()}
		}
	case 9:
		starlarkDollar = starlarkS[starlarkpt-6 : starlarkpt+1]
//line parser.y:94
		{
			starlarkVAL.Def = &Def{TokenWithData: starlarkDollar[1].token, Name: starlarkDollar[2].token.Ident(), Fields: starlarkDollar[4].DefFields}
		}
	case 10:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:99
		{
			starlarkVAL.DefFields = []*DefField{starlarkDollar[1].DefField}
		}
	case 11:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:101
		{
			starlarkVAL.DefFields = append(starlarkDollar[1].DefFields, starlarkDollar[3].DefField)
		}
	case 12:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:103
		{
			starlarkVAL.DefFields = starlarkDollar[1].DefFields
		}
	case 13:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:108
		{
			starlarkVAL.DefField = &DefField{
				Value: starlarkDollar[1].Operand,
			}
		}
	case 14:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:113
		{
			starlarkVAL.DefField = &DefField{
				Value:   starlarkDollar[2].Operand,
				Varargs: true,
			}
		}
	case 15:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:119
		{
			starlarkVAL.DefField = &DefField{
				Value:  starlarkDollar[3].Operand,
				Kwargs: true,
			}
		}
	case 16:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:125
		{
			starlarkVAL.DefField = &DefField{
				Key:   starlarkDollar[1].Operand,
				Value: starlarkDollar[3].Operand,
			}
		}
	case 17:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:131
		{
			starlarkVAL.DefField = &DefField{
				Key:       starlarkDollar[1].Operand,
				ValueExpr: starlarkDollar[3].Assign,
			}
		}
	case 18:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:144
		{
			starlarkVAL.Comment = &Comment{starlarkDollar[1].token, false}
		}
	case 19:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:146
		{
			starlarkVAL.Comment = &Comment{starlarkDollar[1].token, true}
		}
	case 20:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:155
		{
			starlarkVAL.Return = &Return{starlarkDollar[1].token}
		}
	case 21:
		starlarkDollar = starlarkS[starlarkpt-4 : starlarkpt+1]
//line parser.y:164
		{
			starlarkVAL.Load = &Load{Fields: starlarkDollar[3].arr}
		}
	case 22:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:169
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].val}
		}
	case 23:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:171
		{
			starlarkVAL.arr = append(starlarkDollar[1].arr, starlarkDollar[3].val)
		}
	case 24:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:173
		{
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 25:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:178
		{
			starlarkVAL.val = syntax.Array{syntax.Null_{}, starlarkDollar[1].token.String()}
		}
	case 26:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:181
		{
			starlarkVAL.val = syntax.Array{starlarkDollar[1].token.Ident(), starlarkDollar[3].token.String()}
		}
	case 27:
		starlarkDollar = starlarkS[starlarkpt-8 : starlarkpt+1]
//line parser.y:190
		{
			starlarkVAL.Module = &Module{
				Export: starlarkDollar[1].token.Ident(),
				Name:   starlarkDollar[5].token.String(),
				Fields: starlarkDollar[7].arr,
			}
		}
	case 28:
		starlarkDollar = starlarkS[starlarkpt-6 : starlarkpt+1]
//line parser.y:196
		{
			starlarkVAL.Module = &Module{
				Export: starlarkDollar[1].token.Ident(),
				Name:   starlarkDollar[5].token.String(),
			}
		}
	case 29:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:204
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].arr}
		}
	case 30:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:207
		{
			starlarkVAL.arr = append(starlarkDollar[1].arr, starlarkDollar[3].arr)
		}
	case 31:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:210
		{
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 32:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:215
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].Operand, starlarkDollar[3].Operand}
		}
	case 33:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:218
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].Operand, starlarkDollar[3].Assign}
		}
	case 34:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:223
		{
			starlarkVAL.CallFunc = &CallFunc{
				Name: starlarkDollar[1].token.Ident(),
			}
		}
	case 35:
		starlarkDollar = starlarkS[starlarkpt-4 : starlarkpt+1]
//line parser.y:228
		{
			starlarkVAL.CallFunc = &CallFunc{
				Name:   starlarkDollar[1].token.Ident(),
				Fields: starlarkDollar[3].arr,
			}
		}
	case 36:
		starlarkDollar = starlarkS[starlarkpt-6 : starlarkpt+1]
//line parser.y:233
		{
			if starlarkDollar[5].CallFunc.FuncName() != "dict" {
				panic("should be only dictionary for kwargs")
			}
			starlarkVAL.CallFunc = &CallFunc{
				Name:    starlarkDollar[1].token.Ident(),
				Fields:  starlarkDollar[5].CallFunc.Fields,
				Kwarrgs: true,
			}
		}
	case 37:
		starlarkDollar = starlarkS[starlarkpt-7 : starlarkpt+1]
//line parser.y:242
		{
			starlarkVAL.CallFunc = &CallFunc{
				Name:      starlarkDollar[1].token.Ident(),
				Fields:    starlarkDollar[5].arr,
				Varrarrgs: true,
			}
		}
	case 38:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:251
		{
			starlarkVAL.arr = append(starlarkDollar[1].arr, starlarkDollar[3].val)
		}
	case 39:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:253
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].val}
		}
	case 40:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:255
		{
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 41:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:260
		{
			starlarkVAL.val = syntax.Array{starlarkDollar[1].Operand, syntax.Null_{}}
		}
	case 42:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:263
		{
			starlarkVAL.val = syntax.Array{syntax.Null_{}, starlarkDollar[1].Assign}
		}
	case 43:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:266
		{
			starlarkVAL.val = syntax.Array{starlarkDollar[1].Operand, starlarkDollar[3].Operand}
		}
	case 44:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:269
		{
			starlarkVAL.val = syntax.Array{starlarkDollar[1].Operand, starlarkDollar[3].Assign}
		}
	case 45:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:279
		{
			starlarkVAL.val = starlarkDollar[1].token.String()
		}
	case 46:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:281
		{
			starlarkVAL.val = starlarkDollar[1].token.Bool()
		}
	case 47:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:283
		{
			starlarkVAL.val = starlarkDollar[1].token.Int()
		}
	case 48:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:285
		{
			starlarkVAL.val = starlarkDollar[1].token.Float()
		}
	case 49:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:287
		{
			starlarkVAL.val = starlarkDollar[1].token.Null()
		}
	case 50:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:292
		{
			// TODO: it is tuple
			starlarkVAL.arr = starlarkDollar[2].arr
		}
	case 51:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:295
		{
			// TODO: it is list
			starlarkVAL.arr = starlarkDollar[2].arr
		}
	case 52:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:298
		{
			starlarkVAL.arr = syntax.Array{}
		}
	case 53:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:300
		{
			starlarkVAL.arr = syntax.Array{}
		}
	case 54:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:305
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].Assign}
		}
	case 55:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:307
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].Operand}
		}
	case 56:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:309
		{
			starlarkDollar[1].arr.Add(starlarkDollar[3].Operand)
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 57:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:312
		{
			starlarkDollar[1].arr.Add(starlarkDollar[3].Assign)
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 58:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:315
		{
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 59:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:320
		{
			starlarkVAL.map_ = syntax.Map{}
		}
	case 60:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:322
		{
			starlarkVAL.map_ = syntax.Map{}
			for _, item := range starlarkDollar[2].arr {
				key := item.(syntax.Array)[0]
				switch in := key.(type) {
				case syntax.StringLiteral:
					starlarkVAL.map_.Keys = append(starlarkVAL.map_.Keys, string(in))
				case syntax.Ident_:
					if len(in) > 0 {
						starlarkVAL.map_.Keys = append(starlarkVAL.map_.Keys, "@"+strings.Join(in, "."))
					} else {
						starlarkVAL.map_.Keys = append(starlarkVAL.map_.Keys, "")
					}
				case nil:
					starlarkVAL.map_.Keys = append(starlarkVAL.map_.Keys, "")
				default:
					log.Printf("starlark#Fields: not supported key type %T\n", key)
				}
				starlarkVAL.map_.Values.Add(item.(syntax.Array)[1])
			}
		}
	case 61:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:345
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].arr}
		}
	case 62:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:347
		{
			starlarkDollar[1].arr.Add(starlarkDollar[3].arr)
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 63:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:350
		{
			starlarkVAL.arr = starlarkDollar[1].arr
		}
	case 64:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:355
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].token.String(), starlarkDollar[3].Operand}
		}
	case 65:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:357
		{
			starlarkVAL.arr = syntax.Array{starlarkDollar[1].token.String(), starlarkDollar[3].Assign}
		}
	case 66:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:362
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 67:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:364
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 68:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:366
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 69:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:368
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 70:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:370
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol, starlarkDollar[2].token.Symbol}
		}
	case 71:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:372
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 72:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:374
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 73:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:376
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol, starlarkDollar[2].token.Symbol}
		}
	case 74:
		starlarkDollar = starlarkS[starlarkpt-2 : starlarkpt+1]
//line parser.y:378
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol, starlarkDollar[2].token.Symbol}
		}
	case 75:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:380
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 76:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:382
		{
			starlarkVAL.arrint = []int{starlarkDollar[1].token.Symbol}
		}
	case 77:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:387
		{
			starlarkVAL.Assign = &Assign{Left: starlarkDollar[1].Operand, Op: starlarkDollar[2].arrint, Right: starlarkDollar[3].Operand}
		}
	case 78:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:389
		{
			starlarkVAL.Assign = &Assign{Parent: starlarkDollar[1].Assign, Op: starlarkDollar[2].arrint, Right: starlarkDollar[3].Operand}
		}
	case 79:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:394
		{
			starlarkVAL.Operand = &Operand{
				Value: starlarkDollar[1].token.Ident(),
			}
		}
	case 80:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:399
		{
			starlarkVAL.Operand = &Operand{
				Value: starlarkDollar[1].val,
			}
		}
	case 81:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:404
		{
			starlarkVAL.Operand = &Operand{
				Value: starlarkDollar[1].arr,
			}
		}
	case 82:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:409
		{
			starlarkVAL.Operand = &Operand{
				Value: starlarkDollar[1].map_,
			}
		}
	case 83:
		starlarkDollar = starlarkS[starlarkpt-1 : starlarkpt+1]
//line parser.y:414
		{
			starlarkVAL.Operand = &Operand{
				Value: starlarkDollar[1].CallFunc,
			}
		}
	case 84:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:420
		{
			starlarkVAL.Operand = &Operand{
				Parent: starlarkDollar[1].Operand,
				Value:  starlarkDollar[3].CallFunc,
			}
		}
	case 85:
		starlarkDollar = starlarkS[starlarkpt-3 : starlarkpt+1]
//line parser.y:427
		{
			starlarkVAL.Operand = &Operand{
				Parent: starlarkDollar[1].Operand,
				Value:  starlarkDollar[3].token.Ident(),
			}
		}
	}
	goto starlarkstack /* stack new state and value */
}
