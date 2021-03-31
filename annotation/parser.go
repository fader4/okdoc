// Code generated by goyacc -p annotation -o parser.go parser.y. DO NOT EDIT.

//line parser.y:2
package annotation

import __yyfmt__ "fmt"

//line parser.y:2

import (
	_ "fmt"
	"github.com/fader4/okdoc/syntax"
	"log"
	"strings"
)

//line parser.y:14
type annotationSymType struct {
	yys         int
	token       *syntax.TokenWithData
	annot       *Annotation
	annotField  *AnnotationField
	annotFields []*AnnotationField

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
const annotation = 57352
const beginAnnotation = 57353
const endAnnotation = 57354

var annotationToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ident",
	"'('",
	"')'",
	"'@'",
	"'='",
	"','",
	"'{'",
	"'}'",
	"'['",
	"']'",
	"nullLiteral",
	"stringLiteral",
	"boolLiteral",
	"integerLiteral",
	"floatLiteral",
	"annotation",
	"beginAnnotation",
	"endAnnotation",
	"'.'",
}

var annotationStatenames = [...]string{}

const annotationEofCode = 1
const annotationErrCode = 2
const annotationInitialStackSize = 16

//line yacctab:1
var annotationExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const annotationPrivate = 57344

const annotationLast = 78

var annotationAct = [...]int{
	13, 11, 32, 25, 21, 45, 12, 10, 3, 9,
	19, 28, 20, 31, 18, 14, 15, 16, 17, 25,
	29, 36, 35, 6, 24, 38, 41, 34, 33, 21,
	47, 40, 39, 37, 46, 19, 21, 20, 25, 18,
	14, 15, 16, 17, 21, 29, 51, 52, 36, 35,
	53, 26, 50, 49, 34, 33, 48, 21, 44, 7,
	43, 22, 42, 19, 23, 20, 5, 18, 14, 15,
	16, 17, 1, 30, 27, 8, 4, 2,
}

var annotationPact = [...]int{
	-1000, -1000, 1, 62, -1000, 18, 53, -1000, 55, -1000,
	-1000, -1000, -1000, 16, -1000, -1000, -1000, -1000, -1000, 40,
	0, -1000, -1000, 25, 25, 58, -1000, 49, -1000, -3,
	21, -1000, -1000, -1000, -1000, -1000, -19, -1000, -19, -1000,
	-1000, -1000, -1000, -1000, 32, 25, -1000, 25, -1000, -1000,
	-1000, -19, -1000, -1000,
}

var annotationPgo = [...]int{
	0, 7, 0, 2, 77, 76, 9, 75, 11, 74,
	73, 6, 1, 72,
}

var annotationR1 = [...]int{
	0, 13, 4, 4, 5, 5, 5, 7, 7, 6,
	6, 6, 6, 6, 6, 6, 6, 2, 2, 1,
	1, 1, 1, 1, 11, 11, 10, 10, 3, 3,
	3, 3, 12, 12, 9, 9, 8, 8, 8, 8,
}

var annotationR2 = [...]int{
	0, 1, 0, 3, 1, 3, 4, 3, 1, 1,
	1, 1, 1, 3, 3, 3, 3, 1, 3, 1,
	1, 1, 1, 1, 3, 2, 1, 3, 1, 1,
	1, 1, 2, 3, 1, 3, 3, 3, 3, 3,
}

var annotationChk = [...]int{
	-1000, -13, -4, 7, -5, 4, 5, 6, -7, -6,
	-1, -12, -11, -2, 15, 16, 17, 18, 14, 10,
	12, 4, 6, 9, 8, 22, 11, -9, -8, -2,
	-10, 13, -3, -1, -11, -12, -2, -6, -2, -1,
	-11, -12, 4, 11, 9, 8, 13, 9, -8, -1,
	-11, -2, -12, -3,
}

var annotationDef = [...]int{
	2, -2, 1, 0, 3, 4, 0, 5, 0, 8,
	9, 10, 11, 12, 19, 20, 21, 22, 23, 0,
	0, 17, 6, 0, 0, 0, 32, 0, 34, 0,
	0, 25, 26, 28, 29, 30, 31, 7, 13, 14,
	15, 16, 18, 33, 0, 0, 24, 0, 35, 36,
	37, 38, 39, 27,
}

var annotationTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 6, 3, 3, 9, 3, 22, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 8, 3, 3, 7, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 12, 3, 13, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 10, 3, 11,
}

var annotationTok2 = [...]int{
	2, 3, 4, 14, 15, 16, 17, 18, 19, 20,
	21,
}

var annotationTok3 = [...]int{
	0,
}

var annotationErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	annotationDebug        = 0
	annotationErrorVerbose = false
)

type annotationLexer interface {
	Lex(lval *annotationSymType) int
	Error(s string)
}

type annotationParser interface {
	Parse(annotationLexer) int
	Lookahead() int
}

type annotationParserImpl struct {
	lval  annotationSymType
	stack [annotationInitialStackSize]annotationSymType
	char  int
}

func (p *annotationParserImpl) Lookahead() int {
	return p.char
}

func annotationNewParser() annotationParser {
	return &annotationParserImpl{}
}

const annotationFlag = -1000

func annotationTokname(c int) string {
	if c >= 1 && c-1 < len(annotationToknames) {
		if annotationToknames[c-1] != "" {
			return annotationToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func annotationStatname(s int) string {
	if s >= 0 && s < len(annotationStatenames) {
		if annotationStatenames[s] != "" {
			return annotationStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func annotationErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !annotationErrorVerbose {
		return "syntax error"
	}

	for _, e := range annotationErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + annotationTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := annotationPact[state]
	for tok := TOKSTART; tok-1 < len(annotationToknames); tok++ {
		if n := base + tok; n >= 0 && n < annotationLast && annotationChk[annotationAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if annotationDef[state] == -2 {
		i := 0
		for annotationExca[i] != -1 || annotationExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; annotationExca[i] >= 0; i += 2 {
			tok := annotationExca[i]
			if tok < TOKSTART || annotationExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if annotationExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += annotationTokname(tok)
	}
	return res
}

func annotationlex1(lex annotationLexer, lval *annotationSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = annotationTok1[0]
		goto out
	}
	if char < len(annotationTok1) {
		token = annotationTok1[char]
		goto out
	}
	if char >= annotationPrivate {
		if char < annotationPrivate+len(annotationTok2) {
			token = annotationTok2[char-annotationPrivate]
			goto out
		}
	}
	for i := 0; i < len(annotationTok3); i += 2 {
		token = annotationTok3[i+0]
		if token == char {
			token = annotationTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = annotationTok2[1] /* unknown char */
	}
	if annotationDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", annotationTokname(token), uint(char))
	}
	return char, token
}

func annotationParse(annotationlex annotationLexer) int {
	return annotationNewParser().Parse(annotationlex)
}

func (annotationrcvr *annotationParserImpl) Parse(annotationlex annotationLexer) int {
	var annotationn int
	var annotationVAL annotationSymType
	var annotationDollar []annotationSymType
	_ = annotationDollar // silence set and not used
	annotationS := annotationrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	annotationstate := 0
	annotationrcvr.char = -1
	annotationtoken := -1 // annotationrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		annotationstate = -1
		annotationrcvr.char = -1
		annotationtoken = -1
	}()
	annotationp := -1
	goto annotationstack

ret0:
	return 0

ret1:
	return 1

annotationstack:
	/* put a state and value onto the stack */
	if annotationDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", annotationTokname(annotationtoken), annotationStatname(annotationstate))
	}

	annotationp++
	if annotationp >= len(annotationS) {
		nyys := make([]annotationSymType, len(annotationS)*2)
		copy(nyys, annotationS)
		annotationS = nyys
	}
	annotationS[annotationp] = annotationVAL
	annotationS[annotationp].yys = annotationstate

annotationnewstate:
	annotationn = annotationPact[annotationstate]
	if annotationn <= annotationFlag {
		goto annotationdefault /* simple state */
	}
	if annotationrcvr.char < 0 {
		annotationrcvr.char, annotationtoken = annotationlex1(annotationlex, &annotationrcvr.lval)
	}
	annotationn += annotationtoken
	if annotationn < 0 || annotationn >= annotationLast {
		goto annotationdefault
	}
	annotationn = annotationAct[annotationn]
	if annotationChk[annotationn] == annotationtoken { /* valid shift */
		annotationrcvr.char = -1
		annotationtoken = -1
		annotationVAL = annotationrcvr.lval
		annotationstate = annotationn
		if Errflag > 0 {
			Errflag--
		}
		goto annotationstack
	}

annotationdefault:
	/* default state action */
	annotationn = annotationDef[annotationstate]
	if annotationn == -2 {
		if annotationrcvr.char < 0 {
			annotationrcvr.char, annotationtoken = annotationlex1(annotationlex, &annotationrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if annotationExca[xi+0] == -1 && annotationExca[xi+1] == annotationstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			annotationn = annotationExca[xi+0]
			if annotationn < 0 || annotationn == annotationtoken {
				break
			}
		}
		annotationn = annotationExca[xi+1]
		if annotationn < 0 {
			goto ret0
		}
	}
	if annotationn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			annotationlex.Error(annotationErrorMessage(annotationstate, annotationtoken))
			Nerrs++
			if annotationDebug >= 1 {
				__yyfmt__.Printf("%s", annotationStatname(annotationstate))
				__yyfmt__.Printf(" saw %s\n", annotationTokname(annotationtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for annotationp >= 0 {
				annotationn = annotationPact[annotationS[annotationp].yys] + annotationErrCode
				if annotationn >= 0 && annotationn < annotationLast {
					annotationstate = annotationAct[annotationn] /* simulate a shift of "error" */
					if annotationChk[annotationstate] == annotationErrCode {
						goto annotationstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if annotationDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", annotationS[annotationp].yys)
				}
				annotationp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if annotationDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", annotationTokname(annotationtoken))
			}
			if annotationtoken == annotationEofCode {
				goto ret1
			}
			annotationrcvr.char = -1
			annotationtoken = -1
			goto annotationnewstate /* try again in the same state */
		}
	}

	/* reduction by production annotationn */
	if annotationDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", annotationn, annotationStatname(annotationstate))
	}

	annotationnt := annotationn
	annotationpt := annotationp
	_ = annotationpt // guard against "declared and not used"

	annotationp -= annotationR2[annotationn]
	// annotationp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if annotationp+1 >= len(annotationS) {
		nyys := make([]annotationSymType, len(annotationS)*2)
		copy(nyys, annotationS)
		annotationS = nyys
	}
	annotationVAL = annotationS[annotationp+1]

	/* consult goto table to find next state */
	annotationn = annotationR1[annotationn]
	annotationg := annotationPgo[annotationn]
	annotationj := annotationg + annotationS[annotationp].yys + 1

	if annotationj >= annotationLast {
		annotationstate = annotationAct[annotationg]
	} else {
		annotationstate = annotationAct[annotationj]
		if annotationChk[annotationstate] != -annotationn {
			annotationstate = annotationAct[annotationg]
		}
	}
	// dummy call; replaced with literal code
	switch annotationnt {

	case 2:
		annotationDollar = annotationS[annotationpt-0 : annotationpt+1]
//line parser.y:45
		{
			annotationVAL.annot = annotationlex.(*annotationLex).annot
		}
	case 3:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:47
		{
			annotationlex.(*annotationLex).annot.name = annotationDollar[3].annot.name
			annotationlex.(*annotationLex).annot.fields = annotationDollar[3].annot.fields
			annotationVAL.annot = annotationlex.(*annotationLex).annot
		}
	case 4:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:53
		{
			annotationVAL.annot = &Annotation{name: annotationDollar[1].token.Ident()}
		}
	case 5:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:55
		{
			annotationVAL.annot = &Annotation{name: annotationDollar[1].token.Ident()}
		}
	case 6:
		annotationDollar = annotationS[annotationpt-4 : annotationpt+1]
//line parser.y:57
		{
			annotationVAL.annot = &Annotation{name: annotationDollar[1].token.Ident(), fields: annotationDollar[3].annotFields}
		}
	case 7:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:61
		{
			annotationVAL.annotFields = append(annotationDollar[1].annotFields, annotationDollar[3].annotField)
		}
	case 8:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:63
		{
			annotationVAL.annotFields = []*AnnotationField{annotationDollar[1].annotField}
		}
	case 9:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:68
		{
			annotationVAL.annotField = &AnnotationField{
				Value: annotationDollar[1].val,
			}
		}
	case 10:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:73
		{
			annotationVAL.annotField = &AnnotationField{
				Value: annotationDollar[1].map_,
			}
		}
	case 11:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:78
		{
			annotationVAL.annotField = &AnnotationField{
				Value: annotationDollar[1].arr,
			}
		}
	case 12:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:83
		{
			annotationVAL.annotField = &AnnotationField{
				Value: annotationDollar[1].val,
			}
		}
	case 13:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:88
		{
			annotationVAL.annotField = &AnnotationField{
				Value: annotationDollar[1].val,
			}
		}
	case 14:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:93
		{
			annotationVAL.annotField = &AnnotationField{
				Key:   annotationDollar[1].val,
				Value: annotationDollar[3].val,
			}
		}
	case 15:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:99
		{
			annotationVAL.annotField = &AnnotationField{
				Key:   annotationDollar[1].val,
				Value: annotationDollar[3].arr,
			}
		}
	case 16:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:105
		{
			annotationVAL.annotField = &AnnotationField{
				Key:   annotationDollar[1].val,
				Value: annotationDollar[3].map_,
			}
		}
	case 17:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:113
		{
			annotationVAL.val = annotationDollar[1].token.Ident()
		}
	case 18:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:115
		{
			annotationVAL.val = annotationDollar[1].val.(syntax.Ident_).Append(annotationDollar[3].token.Ident())
		}
	case 19:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:119
		{
			annotationVAL.val = annotationDollar[1].token.String()
		}
	case 20:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:121
		{
			annotationVAL.val = annotationDollar[1].token.Bool()
		}
	case 21:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:123
		{
			annotationVAL.val = annotationDollar[1].token.Int()
		}
	case 22:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:125
		{
			annotationVAL.val = annotationDollar[1].token.Float()
		}
	case 23:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:127
		{
			annotationVAL.val = annotationDollar[1].token.Null()
		}
	case 24:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:131
		{
			annotationVAL.arr = annotationDollar[2].arr
		}
	case 25:
		annotationDollar = annotationS[annotationpt-2 : annotationpt+1]
//line parser.y:133
		{
			annotationVAL.arr = syntax.Array{}
		}
	case 26:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:137
		{
			annotationVAL.arr = syntax.Array{annotationDollar[1].val}
		}
	case 27:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:139
		{
			annotationDollar[1].arr.Add(annotationDollar[3].val)
			annotationVAL.arr = annotationDollar[1].arr
		}
	case 28:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:144
		{
			annotationVAL.val = annotationDollar[1].val
		}
	case 29:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:146
		{
			annotationVAL.val = annotationDollar[1].arr
		}
	case 30:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:148
		{
			annotationVAL.val = annotationDollar[1].map_
		}
	case 31:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:150
		{
			annotationVAL.val = annotationDollar[1].val
		}
	case 32:
		annotationDollar = annotationS[annotationpt-2 : annotationpt+1]
//line parser.y:154
		{
			annotationVAL.map_ = syntax.Map{}
		}
	case 33:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:156
		{
			annotationVAL.map_ = syntax.Map{}
			for _, item := range annotationDollar[2].arr {
				key := item.(syntax.Array)[0]
				switch in := key.(type) {
				case syntax.StringLiteral:
					annotationVAL.map_.Keys = append(annotationVAL.map_.Keys, string(in))
				case syntax.Ident_:
					if len(in) > 0 {
						annotationVAL.map_.Keys = append(annotationVAL.map_.Keys, "@"+strings.Join(in, "."))
					} else {
						annotationVAL.map_.Keys = append(annotationVAL.map_.Keys, "")
					}
				case nil:
					annotationVAL.map_.Keys = append(annotationVAL.map_.Keys, "")
				default:
					log.Printf("Annotation#Fields: not supported key type %T\n", key)
				}
				annotationVAL.map_.Values.Add(item.(syntax.Array)[1])
			}
		}
	case 34:
		annotationDollar = annotationS[annotationpt-1 : annotationpt+1]
//line parser.y:178
		{
			annotationVAL.arr = syntax.Array{annotationDollar[1].arr}
		}
	case 35:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:180
		{
			annotationDollar[1].arr.Add(annotationDollar[3].arr)
			annotationVAL.arr = annotationDollar[1].arr
		}
	case 36:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:185
		{
			annotationVAL.arr = syntax.Array{annotationDollar[1].val, annotationDollar[3].val}
		}
	case 37:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:187
		{
			annotationVAL.arr = syntax.Array{annotationDollar[1].val, annotationDollar[3].arr}
		}
	case 38:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:189
		{
			annotationVAL.arr = syntax.Array{annotationDollar[1].val, annotationDollar[3].val}
		}
	case 39:
		annotationDollar = annotationS[annotationpt-3 : annotationpt+1]
//line parser.y:191
		{
			annotationVAL.arr = syntax.Array{annotationDollar[1].val, annotationDollar[3].map_}
		}
	}
	goto annotationstack /* stack new state and value */
}
