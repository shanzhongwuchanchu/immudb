// Code generated by goyacc -l -o sql_parser.go sql_grammar.y. DO NOT EDIT.
package sql

import __yyfmt__ "fmt"

import "fmt"

func setResult(l yyLexer, stmts []SQLStmt) {
	l.(*lexer).result = stmts
}

type yySymType struct {
	yys      int
	stmts    []SQLStmt
	stmt     SQLStmt
	colsSpec []*ColSpec
	colSpec  *ColSpec
	cols     []*ColSelector
	rows     []*RowSpec
	row      *RowSpec
	values   []ValueExp
	value    ValueExp
	id       string
	number   uint64
	str      string
	boolean  bool
	blob     []byte
	sqlType  SQLValueType
	aggFn    AggregateFn
	ids      []string
	col      *ColSelector
	sel      Selector
	sels     []Selector
	distinct bool
	ds       DataSource
	tableRef *tableRef
	joins    []*JoinSpec
	join     *JoinSpec
	joinType JoinType
	exp      ValueExp
	binExp   ValueExp
	err      error
	ordcols  []*OrdCol
	opt_ord  bool
	logicOp  LogicOperator
	cmpOp    CmpOperator
	pparam   int
	update   *colUpdate
	updates  []*colUpdate
}

const CREATE = 57346
const USE = 57347
const DATABASE = 57348
const SNAPSHOT = 57349
const SINCE = 57350
const UP = 57351
const TO = 57352
const TABLE = 57353
const UNIQUE = 57354
const INDEX = 57355
const ON = 57356
const ALTER = 57357
const ADD = 57358
const COLUMN = 57359
const PRIMARY = 57360
const KEY = 57361
const BEGIN = 57362
const TRANSACTION = 57363
const COMMIT = 57364
const ROLLBACK = 57365
const INSERT = 57366
const UPSERT = 57367
const INTO = 57368
const VALUES = 57369
const DELETE = 57370
const UPDATE = 57371
const SET = 57372
const SELECT = 57373
const DISTINCT = 57374
const FROM = 57375
const BEFORE = 57376
const TX = 57377
const JOIN = 57378
const HAVING = 57379
const WHERE = 57380
const GROUP = 57381
const BY = 57382
const LIMIT = 57383
const ORDER = 57384
const ASC = 57385
const DESC = 57386
const AS = 57387
const NOT = 57388
const LIKE = 57389
const IF = 57390
const EXISTS = 57391
const IN = 57392
const IS = 57393
const AUTO_INCREMENT = 57394
const NULL = 57395
const NPARAM = 57396
const PPARAM = 57397
const JOINTYPE = 57398
const LOP = 57399
const CMPOP = 57400
const IDENTIFIER = 57401
const TYPE = 57402
const NUMBER = 57403
const VARCHAR = 57404
const BOOLEAN = 57405
const BLOB = 57406
const AGGREGATE_FUNC = 57407
const ERROR = 57408
const STMT_SEPARATOR = 57409

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CREATE",
	"USE",
	"DATABASE",
	"SNAPSHOT",
	"SINCE",
	"UP",
	"TO",
	"TABLE",
	"UNIQUE",
	"INDEX",
	"ON",
	"ALTER",
	"ADD",
	"COLUMN",
	"PRIMARY",
	"KEY",
	"BEGIN",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"INSERT",
	"UPSERT",
	"INTO",
	"VALUES",
	"DELETE",
	"UPDATE",
	"SET",
	"SELECT",
	"DISTINCT",
	"FROM",
	"BEFORE",
	"TX",
	"JOIN",
	"HAVING",
	"WHERE",
	"GROUP",
	"BY",
	"LIMIT",
	"ORDER",
	"ASC",
	"DESC",
	"AS",
	"NOT",
	"LIKE",
	"IF",
	"EXISTS",
	"IN",
	"IS",
	"AUTO_INCREMENT",
	"NULL",
	"NPARAM",
	"PPARAM",
	"JOINTYPE",
	"LOP",
	"CMPOP",
	"IDENTIFIER",
	"TYPE",
	"NUMBER",
	"VARCHAR",
	"BOOLEAN",
	"BLOB",
	"AGGREGATE_FUNC",
	"ERROR",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"STMT_SEPARATOR",
	"'('",
	"')'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 94,
	47, 120,
	50, 120,
	-2, 109,
	-1, 152,
	36, 87,
	-2, 82,
	-1, 184,
	36, 87,
	-2, 84,
}

const yyPrivate = 57344

const yyLast = 313

var yyAct = [...]int{
	219, 258, 54, 131, 91, 197, 200, 113, 6, 218,
	75, 183, 67, 122, 196, 61, 88, 70, 17, 230,
	233, 129, 129, 129, 193, 241, 129, 235, 236, 234,
	215, 194, 232, 96, 130, 206, 98, 188, 201, 32,
	109, 107, 108, 180, 157, 99, 106, 156, 102, 103,
	104, 105, 55, 202, 115, 198, 97, 96, 128, 205,
	98, 101, 162, 145, 109, 107, 108, 124, 80, 93,
	106, 140, 102, 103, 104, 105, 55, 138, 139, 53,
	97, 119, 110, 90, 78, 101, 140, 66, 134, 135,
	137, 136, 138, 139, 79, 179, 147, 143, 144, 65,
	127, 19, 146, 134, 135, 137, 136, 257, 158, 79,
	49, 56, 56, 151, 140, 149, 252, 55, 152, 214,
	116, 139, 51, 118, 154, 140, 68, 155, 140, 150,
	153, 134, 135, 137, 136, 168, 169, 170, 171, 172,
	173, 161, 134, 135, 137, 136, 233, 137, 136, 216,
	159, 129, 74, 181, 178, 111, 56, 166, 77, 126,
	85, 190, 55, 160, 187, 56, 89, 189, 164, 71,
	148, 123, 76, 191, 125, 120, 117, 204, 82, 195,
	199, 72, 57, 32, 44, 41, 36, 112, 186, 242,
	229, 175, 203, 213, 38, 207, 208, 228, 174, 210,
	140, 176, 123, 142, 177, 81, 58, 259, 260, 245,
	132, 37, 251, 239, 221, 222, 224, 225, 226, 220,
	68, 238, 231, 209, 84, 63, 62, 73, 30, 34,
	240, 17, 48, 165, 243, 39, 163, 29, 28, 246,
	20, 2, 248, 211, 86, 64, 10, 11, 250, 249,
	253, 114, 60, 167, 255, 256, 133, 12, 83, 59,
	261, 35, 7, 262, 8, 9, 13, 14, 31, 21,
	15, 16, 40, 17, 22, 24, 23, 27, 43, 92,
	45, 46, 47, 25, 26, 18, 69, 141, 227, 212,
	244, 254, 192, 223, 95, 94, 237, 185, 184, 182,
	42, 33, 52, 50, 100, 217, 247, 87, 121, 5,
	4, 3, 1,
}

var yyPact = [...]int{
	242, -1000, -1000, 28, -1000, -1000, -1000, 219, -1000, -1000,
	263, 277, 266, 212, 211, 195, 124, 197, -1000, 242,
	-1000, 127, 146, 146, 259, 126, 270, 125, 124, 124,
	124, 202, 38, 52, -1000, -1000, -1000, 123, 160, 245,
	146, -1000, 192, 190, 229, 25, 13, 182, 110, 122,
	194, -1000, 85, 113, -1000, 10, 37, -6, 156, 119,
	244, -1000, 189, 99, 227, 107, 107, 274, 11, 88,
	-1000, 129, -1000, -20, 97, -1000, -1000, 117, 53, 116,
	112, -1000, -7, 115, 98, -1000, 112, -17, 84, -1000,
	-41, 169, 243, 35, 157, -1000, 11, 11, -11, -1000,
	-1000, 11, -1000, -1000, -1000, -1000, 22, 111, -1000, -1000,
	274, 110, 11, 274, 192, 200, 113, -1000, -28, -31,
	36, 83, -1000, 103, 107, -12, -1000, -1000, 209, 109,
	206, -1000, 96, 239, 11, 11, 11, 11, 11, 11,
	145, 154, -1000, 63, 77, 200, 20, -32, -1000, 169,
	-1000, 35, 132, 113, -38, -1000, -1000, -1000, 108, 143,
	-52, -44, 107, -19, -1000, -19, -1000, -21, 77, 77,
	149, 149, 63, 74, -1000, 139, 11, -15, -40, -1000,
	-1000, -1000, 182, -1000, 132, 187, -1000, -1000, 113, -1000,
	224, -1000, 141, 58, -1000, -45, 82, -1000, 11, 82,
	-1000, -1000, 107, -1000, 63, -13, -1000, 177, -1000, -20,
	-1000, -21, 144, -1000, -58, -1000, -19, -43, 79, 35,
	-46, -48, -47, 184, 173, 274, -50, -1000, -1000, 136,
	-1000, -1000, -1000, 11, -1000, -1000, -1000, 167, 11, 106,
	235, -1000, -1000, 35, 169, 172, 35, 49, -1000, 11,
	-1000, 106, 106, 35, 40, 164, -1000, 106, -1000, -1000,
	-1000, 164, -1000,
}

var yyPgo = [...]int{
	0, 312, 241, 311, 310, 8, 309, 308, 13, 16,
	6, 307, 306, 14, 5, 9, 305, 304, 45, 303,
	302, 2, 301, 7, 251, 300, 15, 299, 11, 298,
	297, 0, 12, 296, 295, 294, 293, 3, 292, 10,
	291, 290, 1, 4, 211, 289, 288, 287, 17, 286,
	285,
}

var yyR1 = [...]int{
	0, 1, 2, 2, 50, 50, 3, 3, 3, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 25,
	25, 44, 44, 10, 10, 6, 6, 6, 6, 49,
	49, 48, 11, 11, 13, 13, 14, 9, 9, 12,
	12, 16, 16, 15, 15, 17, 17, 17, 17, 17,
	17, 17, 17, 7, 7, 8, 38, 38, 45, 45,
	46, 46, 46, 5, 22, 22, 19, 19, 20, 20,
	18, 18, 18, 21, 21, 21, 23, 23, 24, 24,
	26, 26, 27, 27, 28, 28, 29, 30, 30, 32,
	32, 36, 36, 33, 33, 37, 37, 41, 41, 43,
	43, 40, 40, 42, 42, 42, 39, 39, 39, 31,
	31, 31, 31, 31, 31, 31, 31, 34, 34, 34,
	47, 47, 35, 35, 35, 35, 35, 35, 35, 35,
}

var yyR2 = [...]int{
	0, 1, 2, 3, 0, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 4, 11, 8, 9, 6, 0,
	3, 0, 3, 1, 3, 8, 8, 6, 7, 1,
	3, 3, 0, 1, 1, 3, 3, 1, 3, 1,
	3, 0, 1, 1, 3, 1, 1, 1, 1, 3,
	2, 1, 1, 1, 3, 5, 0, 3, 0, 1,
	0, 1, 2, 12, 0, 1, 1, 1, 2, 4,
	1, 4, 4, 1, 3, 5, 3, 4, 1, 3,
	0, 3, 0, 1, 1, 2, 6, 0, 1, 0,
	2, 0, 3, 0, 2, 0, 2, 0, 3, 0,
	4, 2, 4, 0, 1, 1, 0, 1, 2, 1,
	1, 2, 2, 4, 4, 6, 6, 1, 1, 3,
	0, 1, 3, 3, 3, 3, 3, 3, 3, 4,
}

var yyChk = [...]int{
	-1000, -1, -2, -3, -4, -6, -5, 20, 22, 23,
	4, 5, 15, 24, 25, 28, 29, 31, -50, 73,
	21, 6, 11, 13, 12, 6, 7, 11, 26, 26,
	33, -24, 59, -22, 32, -2, 59, -44, 48, -44,
	13, 59, -25, 8, 59, -24, -24, -24, 30, 72,
	-19, 70, -20, -18, -21, 65, 59, 59, 46, 14,
	-44, -26, 34, 35, 16, 74, 74, -32, 38, -49,
	-48, 59, 59, 33, 67, -39, 59, 45, 74, 72,
	74, 49, 59, 14, 35, 61, 17, -11, -9, 59,
	-9, -43, 5, -31, -34, -35, 46, 69, 49, -18,
	-17, 74, 61, 62, 63, 64, 59, 54, 55, 53,
	-32, 67, 58, -23, -24, 74, -18, 59, 70, -21,
	59, -7, -8, 59, 74, 59, 61, -8, 75, 67,
	75, -37, 41, 13, 68, 69, 71, 70, 57, 58,
	51, -47, 46, -31, -31, 74, -31, 74, 59, -43,
	-48, -31, -43, -26, -5, -39, 75, 75, 72, 67,
	60, -9, 74, 27, 59, 27, 61, 14, -31, -31,
	-31, -31, -31, -31, 53, 46, 47, 50, -5, 75,
	75, -37, -27, -28, -29, -30, 56, -39, 75, 59,
	18, -8, -38, 76, 75, -9, -13, -14, 74, -13,
	-10, 59, 74, 53, -31, 74, 75, -32, -28, 36,
	-39, 19, -45, 52, 61, 75, 67, -16, -15, -31,
	-9, -5, -15, -36, 39, -23, -10, -46, 53, 46,
	77, -14, 75, 67, 75, 75, 75, -33, 37, 40,
	-43, 75, 53, -31, -41, 42, -31, -12, -21, 14,
	-37, 40, 67, -31, -40, -21, -21, 67, -42, 43,
	44, -21, -42,
}

var yyDef = [...]int{
	0, -2, 1, 4, 6, 7, 8, 0, 10, 11,
	0, 0, 0, 0, 0, 0, 0, 64, 2, 5,
	9, 0, 21, 21, 0, 0, 19, 0, 0, 0,
	0, 0, 78, 0, 65, 3, 12, 0, 0, 0,
	21, 13, 80, 0, 0, 0, 0, 89, 0, 0,
	0, 66, 67, 106, 70, 0, 73, 0, 0, 0,
	0, 14, 0, 0, 0, 32, 0, 99, 0, 89,
	29, 0, 79, 0, 0, 68, 107, 0, 0, 0,
	0, 22, 0, 0, 0, 20, 0, 0, 33, 37,
	0, 95, 0, 90, -2, 110, 0, 0, 0, 117,
	118, 0, 45, 46, 47, 48, 73, 0, 51, 52,
	99, 0, 0, 99, 80, 0, 106, 108, 0, 0,
	74, 0, 53, 0, 0, 0, 81, 18, 0, 0,
	0, 27, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 121, 111, 112, 0, 0, 0, 50, 95,
	30, 31, -2, 106, 0, 69, 71, 72, 0, 0,
	56, 0, 0, 0, 38, 0, 96, 0, 122, 123,
	124, 125, 126, 127, 128, 0, 0, 0, 0, 119,
	49, 28, 89, 83, -2, 0, 88, 76, 106, 75,
	0, 54, 58, 0, 16, 0, 25, 34, 41, 26,
	100, 23, 0, 129, 113, 0, 114, 91, 85, 0,
	77, 0, 60, 59, 0, 17, 0, 0, 42, 43,
	0, 0, 0, 93, 0, 99, 0, 55, 61, 0,
	57, 35, 36, 0, 24, 115, 116, 97, 0, 0,
	0, 15, 62, 44, 95, 0, 94, 92, 39, 0,
	63, 0, 0, 86, 98, 103, 40, 0, 101, 104,
	105, 103, 102,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	74, 75, 70, 68, 67, 69, 72, 71, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 76, 3, 77,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 73,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmts = yyDollar[1].stmts
			setResult(yylex, yyDollar[1].stmts)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &BeginTransactionStmt{}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &CommitStmt{}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &RollbackStmt{}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &CreateDatabaseStmt{DB: yyDollar[3].id}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[3].id}
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &UseSnapshotStmt{sinceTx: yyDollar[3].number, asBefore: yyDollar[4].number}
		}
	case 15:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			yyVAL.stmt = &CreateTableStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[4].id, colsSpec: yyDollar[6].colsSpec, pkColNames: yyDollar[10].ids}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[5].id, cols: yyDollar[7].ids}
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{unique: true, ifNotExists: yyDollar[4].boolean, table: yyDollar[6].id, cols: yyDollar[8].ids}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &AddColumnStmt{table: yyDollar[3].id, colSpec: yyDollar[6].colSpec}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = yyDollar[2].ids
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{isInsert: true, tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 26:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &DeleteFromStmt{tableRef: yyDollar[3].tableRef, where: yyDollar[4].exp, indexOn: yyDollar[5].ids, limit: int(yyDollar[6].number)}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &UpdateStmt{tableRef: yyDollar[2].tableRef, updates: yyDollar[4].updates, where: yyDollar[5].exp, indexOn: yyDollar[6].ids, limit: int(yyDollar[7].number)}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.updates = []*colUpdate{yyDollar[1].update}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.updates = append(yyDollar[1].updates, yyDollar[3].update)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.update = &colUpdate{col: yyDollar[1].id, op: yyDollar[2].cmpOp, val: yyDollar[3].exp}
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = yyDollar[1].ids
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = []*RowSpec{yyDollar[1].row}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].row)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.row = &RowSpec{Values: yyDollar[2].values}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = append(yyDollar[1].ids, yyDollar[3].id)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.cols = []*ColSelector{yyDollar[1].col}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = append(yyDollar[1].cols, yyDollar[3].col)
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.values = nil
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = yyDollar[1].values
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = []ValueExp{yyDollar[1].exp}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].exp)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Number{val: int64(yyDollar[1].number)}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Varchar{val: yyDollar[1].str}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Bool{val: yyDollar[1].boolean}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Blob{val: yyDollar[1].blob}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.value = &SysFn{fn: yyDollar[1].id}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.value = &Param{id: yyDollar[2].id}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Param{id: fmt.Sprintf("param%d", yyDollar[1].pparam), pos: yyDollar[1].pparam}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &NullValue{t: AnyType}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.colsSpec = []*ColSpec{yyDollar[1].colSpec}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colsSpec = append(yyDollar[1].colsSpec, yyDollar[3].colSpec)
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.colSpec = &ColSpec{colName: yyDollar[1].id, colType: yyDollar[2].sqlType, maxLen: int(yyDollar[3].number), autoIncrement: yyDollar[4].boolean, notNull: yyDollar[5].boolean}
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 63:
		yyDollar = yyS[yypt-12 : yypt+1]
		{
			yyVAL.stmt = &SelectStmt{
				distinct:  yyDollar[2].distinct,
				selectors: yyDollar[3].sels,
				ds:        yyDollar[5].ds,
				indexOn:   yyDollar[6].ids,
				joins:     yyDollar[7].joins,
				where:     yyDollar[8].exp,
				groupBy:   yyDollar[9].cols,
				having:    yyDollar[10].exp,
				orderBy:   yyDollar[11].ordcols,
				limit:     int(yyDollar[12].number),
			}
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = nil
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = yyDollar[1].sels
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].sel.setAlias(yyDollar[2].id)
			yyVAL.sels = []Selector{yyDollar[1].sel}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[3].sel.setAlias(yyDollar[4].id)
			yyVAL.sels = append(yyDollar[1].sels, yyDollar[3].sel)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sel = yyDollar[1].col
		}
	case 71:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, col: "*"}
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, db: yyDollar[3].col.db, table: yyDollar[3].col.table, col: yyDollar[3].col.col}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.col = &ColSelector{col: yyDollar[1].id}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.col = &ColSelector{table: yyDollar[1].id, col: yyDollar[3].id}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.col = &ColSelector{db: yyDollar[1].id, table: yyDollar[3].id, col: yyDollar[5].id}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].tableRef.asBefore = yyDollar[2].number
			yyDollar[1].tableRef.as = yyDollar[3].id
			yyVAL.ds = yyDollar[1].tableRef
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[2].stmt.(*SelectStmt).as = yyDollar[4].id
			yyVAL.ds = yyDollar[2].stmt.(DataSource)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{table: yyDollar[1].id}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{db: yyDollar[1].id, table: yyDollar[3].id}
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joins = nil
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = yyDollar[1].joins
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = []*JoinSpec{yyDollar[1].join}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.joins = append([]*JoinSpec{yyDollar[1].join}, yyDollar[2].joins...)
		}
	case 86:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.join = &JoinSpec{joinType: yyDollar[1].joinType, ds: yyDollar[3].ds, indexOn: yyDollar[4].ids, cond: yyDollar[6].exp}
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joinType = InnerJoin
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joinType = yyDollar[1].joinType
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.cols = nil
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = yyDollar[3].cols
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ordcols = nil
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ordcols = yyDollar[3].ordcols
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ids = yyDollar[4].ids
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ordcols = []*OrdCol{{sel: yyDollar[1].col, descOrder: yyDollar[2].opt_ord}}
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ordcols = append(yyDollar[1].ordcols, &OrdCol{sel: yyDollar[3].col, descOrder: yyDollar[4].opt_ord})
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = true
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.id = yyDollar[1].id
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.id = yyDollar[2].id
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].exp
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].binExp
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NotBoolExp{exp: yyDollar[2].exp}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NumExp{left: &Number{val: 0}, op: SUBSOP, right: yyDollar[2].exp}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &LikeBoolExp{val: yyDollar[1].exp, notLike: yyDollar[2].boolean, pattern: yyDollar[4].exp}
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &ExistsBoolExp{q: (yyDollar[3].stmt).(*SelectStmt)}
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InSubQueryExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, q: yyDollar[5].stmt.(*SelectStmt)}
		}
	case 116:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InListExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, values: yyDollar[5].values}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].sel
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].value
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 120:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: ADDOP, right: yyDollar[3].exp}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: SUBSOP, right: yyDollar[3].exp}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: DIVOP, right: yyDollar[3].exp}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: MULTOP, right: yyDollar[3].exp}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &BinBoolExp{left: yyDollar[1].exp, op: yyDollar[2].logicOp, right: yyDollar[3].exp}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: yyDollar[2].cmpOp, right: yyDollar[3].exp}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: EQ, right: &NullValue{t: AnyType}}
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: NE, right: &NullValue{t: AnyType}}
		}
	}
	goto yystack /* stack new state and value */
}
