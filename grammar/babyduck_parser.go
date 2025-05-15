// Code generated from BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // BabyDuck

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BabyDuckParser struct {
	*antlr.BaseParser
}

var BabyDuckParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func babyduckParserInit() {
	staticData := &BabyDuckParserStaticData
	staticData.LiteralNames = []string{
		"", "'program'", "'main'", "'end'", "'var'", "'{'", "'}'", "'='", "'while'",
		"'do'", "'if'", "'else'", "'print'", "'>'", "'<'", "'!='", "'+'", "'-'",
		"'*'", "'/'", "'void'", "'int'", "'float'", "'('", "')'", "'['", "']'",
		"':'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "VOID", "INTTYPE", "FLOATTYPE", "LPAREN", "RPAREN", "LBRACKET",
		"RBRACKET", "COLON", "COMMA", "SEMICOLON", "ID", "INT", "FLOAT", "STRING",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "vars", "var_decl", "id_list", "type", "body", "statement",
		"assign", "cycle", "condition", "else_part", "print_stmt", "printexpr",
		"constant", "expression", "relational", "relop", "exp", "addop", "term",
		"mulop", "factor", "parexpr", "factorsign", "value", "funcs", "func",
		"param_list", "param", "funcbody", "f_call", "arg_list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 257, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 69, 8, 0, 1, 0, 5, 0, 72, 8, 0, 10, 0,
		12, 0, 75, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 4, 1, 83, 8, 1, 11,
		1, 12, 1, 84, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 95,
		8, 3, 10, 3, 12, 3, 98, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 104, 8, 5,
		10, 5, 12, 5, 107, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		116, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 3, 10, 141, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 148,
		8, 11, 10, 11, 12, 11, 151, 9, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3,
		12, 158, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 164, 8, 14, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 175, 8, 17,
		10, 17, 12, 17, 178, 9, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5,
		19, 186, 8, 19, 10, 19, 12, 19, 189, 9, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		3, 21, 195, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 3, 23, 202, 8, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 208, 8, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 3, 26, 216, 8, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 5, 27, 224, 8, 27, 10, 27, 12, 27, 227, 9, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 3, 29, 235, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 30, 3, 30, 244, 8, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 5, 31, 252, 8, 31, 10, 31, 12, 31, 255, 9, 31, 1, 31, 0, 0,
		32, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 0, 5, 1, 0, 21,
		22, 1, 0, 31, 32, 1, 0, 13, 15, 1, 0, 16, 17, 1, 0, 18, 19, 247, 0, 64,
		1, 0, 0, 0, 2, 80, 1, 0, 0, 0, 4, 86, 1, 0, 0, 0, 6, 91, 1, 0, 0, 0, 8,
		99, 1, 0, 0, 0, 10, 101, 1, 0, 0, 0, 12, 115, 1, 0, 0, 0, 14, 117, 1, 0,
		0, 0, 16, 122, 1, 0, 0, 0, 18, 130, 1, 0, 0, 0, 20, 140, 1, 0, 0, 0, 22,
		142, 1, 0, 0, 0, 24, 157, 1, 0, 0, 0, 26, 159, 1, 0, 0, 0, 28, 161, 1,
		0, 0, 0, 30, 165, 1, 0, 0, 0, 32, 168, 1, 0, 0, 0, 34, 170, 1, 0, 0, 0,
		36, 179, 1, 0, 0, 0, 38, 181, 1, 0, 0, 0, 40, 190, 1, 0, 0, 0, 42, 194,
		1, 0, 0, 0, 44, 196, 1, 0, 0, 0, 46, 201, 1, 0, 0, 0, 48, 207, 1, 0, 0,
		0, 50, 209, 1, 0, 0, 0, 52, 211, 1, 0, 0, 0, 54, 220, 1, 0, 0, 0, 56, 228,
		1, 0, 0, 0, 58, 232, 1, 0, 0, 0, 60, 240, 1, 0, 0, 0, 62, 248, 1, 0, 0,
		0, 64, 65, 5, 1, 0, 0, 65, 66, 5, 30, 0, 0, 66, 68, 5, 29, 0, 0, 67, 69,
		3, 2, 1, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 73, 1, 0, 0, 0,
		70, 72, 3, 50, 25, 0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1,
		0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76,
		77, 5, 2, 0, 0, 77, 78, 3, 10, 5, 0, 78, 79, 5, 3, 0, 0, 79, 1, 1, 0, 0,
		0, 80, 82, 5, 4, 0, 0, 81, 83, 3, 4, 2, 0, 82, 81, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 3, 1, 0, 0, 0,
		86, 87, 3, 6, 3, 0, 87, 88, 5, 27, 0, 0, 88, 89, 3, 8, 4, 0, 89, 90, 5,
		29, 0, 0, 90, 5, 1, 0, 0, 0, 91, 96, 5, 30, 0, 0, 92, 93, 5, 28, 0, 0,
		93, 95, 5, 30, 0, 0, 94, 92, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1,
		0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 7, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99,
		100, 7, 0, 0, 0, 100, 9, 1, 0, 0, 0, 101, 105, 5, 5, 0, 0, 102, 104, 3,
		12, 6, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0,
		0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108,
		109, 5, 6, 0, 0, 109, 11, 1, 0, 0, 0, 110, 116, 3, 14, 7, 0, 111, 116,
		3, 16, 8, 0, 112, 116, 3, 60, 30, 0, 113, 116, 3, 22, 11, 0, 114, 116,
		3, 18, 9, 0, 115, 110, 1, 0, 0, 0, 115, 111, 1, 0, 0, 0, 115, 112, 1, 0,
		0, 0, 115, 113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 13, 1, 0, 0, 0,
		117, 118, 5, 30, 0, 0, 118, 119, 5, 7, 0, 0, 119, 120, 3, 28, 14, 0, 120,
		121, 5, 29, 0, 0, 121, 15, 1, 0, 0, 0, 122, 123, 5, 8, 0, 0, 123, 124,
		5, 23, 0, 0, 124, 125, 3, 28, 14, 0, 125, 126, 5, 24, 0, 0, 126, 127, 5,
		9, 0, 0, 127, 128, 3, 10, 5, 0, 128, 129, 5, 29, 0, 0, 129, 17, 1, 0, 0,
		0, 130, 131, 5, 10, 0, 0, 131, 132, 5, 23, 0, 0, 132, 133, 3, 28, 14, 0,
		133, 134, 5, 24, 0, 0, 134, 135, 3, 10, 5, 0, 135, 136, 3, 20, 10, 0, 136,
		137, 5, 29, 0, 0, 137, 19, 1, 0, 0, 0, 138, 139, 5, 11, 0, 0, 139, 141,
		3, 10, 5, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 21, 1, 0,
		0, 0, 142, 143, 5, 12, 0, 0, 143, 144, 5, 23, 0, 0, 144, 149, 3, 24, 12,
		0, 145, 146, 5, 28, 0, 0, 146, 148, 3, 24, 12, 0, 147, 145, 1, 0, 0, 0,
		148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 24, 0, 0, 153, 154,
		5, 29, 0, 0, 154, 23, 1, 0, 0, 0, 155, 158, 3, 34, 17, 0, 156, 158, 5,
		33, 0, 0, 157, 155, 1, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158, 25, 1, 0, 0,
		0, 159, 160, 7, 1, 0, 0, 160, 27, 1, 0, 0, 0, 161, 163, 3, 34, 17, 0, 162,
		164, 3, 30, 15, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 29,
		1, 0, 0, 0, 165, 166, 3, 32, 16, 0, 166, 167, 3, 34, 17, 0, 167, 31, 1,
		0, 0, 0, 168, 169, 7, 2, 0, 0, 169, 33, 1, 0, 0, 0, 170, 176, 3, 38, 19,
		0, 171, 172, 3, 36, 18, 0, 172, 173, 3, 38, 19, 0, 173, 175, 1, 0, 0, 0,
		174, 171, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176,
		177, 1, 0, 0, 0, 177, 35, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 179, 180, 7,
		3, 0, 0, 180, 37, 1, 0, 0, 0, 181, 187, 3, 42, 21, 0, 182, 183, 3, 40,
		20, 0, 183, 184, 3, 42, 21, 0, 184, 186, 1, 0, 0, 0, 185, 182, 1, 0, 0,
		0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188,
		39, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 7, 4, 0, 0, 191, 41, 1,
		0, 0, 0, 192, 195, 3, 44, 22, 0, 193, 195, 3, 46, 23, 0, 194, 192, 1, 0,
		0, 0, 194, 193, 1, 0, 0, 0, 195, 43, 1, 0, 0, 0, 196, 197, 5, 23, 0, 0,
		197, 198, 3, 28, 14, 0, 198, 199, 5, 24, 0, 0, 199, 45, 1, 0, 0, 0, 200,
		202, 3, 36, 18, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 204, 3, 48, 24, 0, 204, 47, 1, 0, 0, 0, 205, 208, 5, 30,
		0, 0, 206, 208, 3, 26, 13, 0, 207, 205, 1, 0, 0, 0, 207, 206, 1, 0, 0,
		0, 208, 49, 1, 0, 0, 0, 209, 210, 3, 52, 26, 0, 210, 51, 1, 0, 0, 0, 211,
		212, 5, 20, 0, 0, 212, 213, 5, 30, 0, 0, 213, 215, 5, 23, 0, 0, 214, 216,
		3, 54, 27, 0, 215, 214, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1,
		0, 0, 0, 217, 218, 5, 24, 0, 0, 218, 219, 3, 58, 29, 0, 219, 53, 1, 0,
		0, 0, 220, 225, 3, 56, 28, 0, 221, 222, 5, 28, 0, 0, 222, 224, 3, 56, 28,
		0, 223, 221, 1, 0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225,
		226, 1, 0, 0, 0, 226, 55, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 229, 5,
		30, 0, 0, 229, 230, 5, 27, 0, 0, 230, 231, 3, 8, 4, 0, 231, 57, 1, 0, 0,
		0, 232, 234, 5, 25, 0, 0, 233, 235, 3, 2, 1, 0, 234, 233, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 3, 10, 5, 0, 237, 238,
		5, 26, 0, 0, 238, 239, 5, 29, 0, 0, 239, 59, 1, 0, 0, 0, 240, 241, 5, 30,
		0, 0, 241, 243, 5, 23, 0, 0, 242, 244, 3, 62, 31, 0, 243, 242, 1, 0, 0,
		0, 243, 244, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 5, 24, 0, 0, 246,
		247, 5, 29, 0, 0, 247, 61, 1, 0, 0, 0, 248, 253, 3, 28, 14, 0, 249, 250,
		5, 28, 0, 0, 250, 252, 3, 28, 14, 0, 251, 249, 1, 0, 0, 0, 252, 255, 1,
		0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 63, 1, 0, 0,
		0, 255, 253, 1, 0, 0, 0, 20, 68, 73, 84, 96, 105, 115, 140, 149, 157, 163,
		176, 187, 194, 201, 207, 215, 225, 234, 243, 253,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BabyDuckParserInit initializes any static state used to implement BabyDuckParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBabyDuckParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BabyDuckParserInit() {
	staticData := &BabyDuckParserStaticData
	staticData.once.Do(babyduckParserInit)
}

// NewBabyDuckParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBabyDuckParser(input antlr.TokenStream) *BabyDuckParser {
	BabyDuckParserInit()
	this := new(BabyDuckParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BabyDuckParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "BabyDuck.g4"

	return this
}

// BabyDuckParser tokens.
const (
	BabyDuckParserEOF       = antlr.TokenEOF
	BabyDuckParserT__0      = 1
	BabyDuckParserT__1      = 2
	BabyDuckParserT__2      = 3
	BabyDuckParserT__3      = 4
	BabyDuckParserT__4      = 5
	BabyDuckParserT__5      = 6
	BabyDuckParserT__6      = 7
	BabyDuckParserT__7      = 8
	BabyDuckParserT__8      = 9
	BabyDuckParserT__9      = 10
	BabyDuckParserT__10     = 11
	BabyDuckParserT__11     = 12
	BabyDuckParserT__12     = 13
	BabyDuckParserT__13     = 14
	BabyDuckParserT__14     = 15
	BabyDuckParserT__15     = 16
	BabyDuckParserT__16     = 17
	BabyDuckParserT__17     = 18
	BabyDuckParserT__18     = 19
	BabyDuckParserVOID      = 20
	BabyDuckParserINTTYPE   = 21
	BabyDuckParserFLOATTYPE = 22
	BabyDuckParserLPAREN    = 23
	BabyDuckParserRPAREN    = 24
	BabyDuckParserLBRACKET  = 25
	BabyDuckParserRBRACKET  = 26
	BabyDuckParserCOLON     = 27
	BabyDuckParserCOMMA     = 28
	BabyDuckParserSEMICOLON = 29
	BabyDuckParserID        = 30
	BabyDuckParserINT       = 31
	BabyDuckParserFLOAT     = 32
	BabyDuckParserSTRING    = 33
	BabyDuckParserWS        = 34
)

// BabyDuckParser rules.
const (
	BabyDuckParserRULE_program    = 0
	BabyDuckParserRULE_vars       = 1
	BabyDuckParserRULE_var_decl   = 2
	BabyDuckParserRULE_id_list    = 3
	BabyDuckParserRULE_type       = 4
	BabyDuckParserRULE_body       = 5
	BabyDuckParserRULE_statement  = 6
	BabyDuckParserRULE_assign     = 7
	BabyDuckParserRULE_cycle      = 8
	BabyDuckParserRULE_condition  = 9
	BabyDuckParserRULE_else_part  = 10
	BabyDuckParserRULE_print_stmt = 11
	BabyDuckParserRULE_printexpr  = 12
	BabyDuckParserRULE_constant   = 13
	BabyDuckParserRULE_expression = 14
	BabyDuckParserRULE_relational = 15
	BabyDuckParserRULE_relop      = 16
	BabyDuckParserRULE_exp        = 17
	BabyDuckParserRULE_addop      = 18
	BabyDuckParserRULE_term       = 19
	BabyDuckParserRULE_mulop      = 20
	BabyDuckParserRULE_factor     = 21
	BabyDuckParserRULE_parexpr    = 22
	BabyDuckParserRULE_factorsign = 23
	BabyDuckParserRULE_value      = 24
	BabyDuckParserRULE_funcs      = 25
	BabyDuckParserRULE_func       = 26
	BabyDuckParserRULE_param_list = 27
	BabyDuckParserRULE_param      = 28
	BabyDuckParserRULE_funcbody   = 29
	BabyDuckParserRULE_f_call     = 30
	BabyDuckParserRULE_arg_list   = 31
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Body() IBodyContext
	Vars() IVarsContext
	AllFuncs() []IFuncsContext
	Funcs(i int) IFuncsContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *ProgramContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *ProgramContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ProgramContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *ProgramContext) AllFuncs() []IFuncsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncsContext); ok {
			len++
		}
	}

	tst := make([]IFuncsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncsContext); ok {
			tst[i] = t.(IFuncsContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Funcs(i int) IFuncsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BabyDuckParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(BabyDuckParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__3 {
		{
			p.SetState(67)
			p.Vars()
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserVOID {
		{
			p.SetState(70)
			p.Funcs()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
		p.Match(BabyDuckParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Body()
	}
	{
		p.SetState(78)
		p.Match(BabyDuckParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_decl() []IVar_declContext
	Var_decl(i int) IVar_declContext

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_vars
	return p
}

func InitEmptyVarsContext(p *VarsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_vars
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) AllVar_decl() []IVar_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declContext); ok {
			len++
		}
	}

	tst := make([]IVar_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declContext); ok {
			tst[i] = t.(IVar_declContext)
			i++
		}
	}

	return tst
}

func (s *VarsContext) Var_decl(i int) IVar_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitVars(s)
	}
}

func (s *VarsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitVars(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BabyDuckParserRULE_vars)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(BabyDuckParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BabyDuckParserID {
		{
			p.SetState(81)
			p.Var_decl()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_list() IId_listContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SEMICOLON() antlr.TerminalNode

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_var_decl
	return p
}

func InitEmptyVar_declContext(p *Var_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_var_decl
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) Id_list() IId_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_listContext)
}

func (s *Var_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOLON, 0)
}

func (s *Var_declContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Var_declContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterVar_decl(s)
	}
}

func (s *Var_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitVar_decl(s)
	}
}

func (s *Var_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitVar_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BabyDuckParserRULE_var_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Id_list()
	}
	{
		p.SetState(87)
		p.Match(BabyDuckParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Type_()
	}
	{
		p.SetState(89)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IId_listContext is an interface to support dynamic dispatch.
type IId_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsId_listContext differentiates from other interfaces.
	IsId_listContext()
}

type Id_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_listContext() *Id_listContext {
	var p = new(Id_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_id_list
	return p
}

func InitEmptyId_listContext(p *Id_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_id_list
}

func (*Id_listContext) IsId_listContext() {}

func NewId_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_listContext {
	var p = new(Id_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_id_list

	return p
}

func (s *Id_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_listContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserID)
}

func (s *Id_listContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, i)
}

func (s *Id_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *Id_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *Id_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterId_list(s)
	}
}

func (s *Id_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitId_list(s)
	}
}

func (s *Id_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitId_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Id_list() (localctx IId_listContext) {
	localctx = NewId_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BabyDuckParserRULE_id_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserCOMMA {
		{
			p.SetState(92)
			p.Match(BabyDuckParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(BabyDuckParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTTYPE() antlr.TerminalNode
	FLOATTYPE() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INTTYPE() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserINTTYPE, 0)
}

func (s *TypeContext) FLOATTYPE() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserFLOATTYPE, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BabyDuckParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserINTTYPE || _la == BabyDuckParserFLOATTYPE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BabyDuckParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(BabyDuckParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073747200) != 0 {
		{
			p.SetState(102)
			p.Statement()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(108)
		p.Match(BabyDuckParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() IAssignContext
	Cycle() ICycleContext
	F_call() IF_callContext
	Print_stmt() IPrint_stmtContext
	Condition() IConditionContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatementContext) Cycle() ICycleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICycleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICycleContext)
}

func (s *StatementContext) F_call() IF_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IF_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IF_callContext)
}

func (s *StatementContext) Print_stmt() IPrint_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrint_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrint_stmtContext)
}

func (s *StatementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BabyDuckParserRULE_statement)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Cycle()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.F_call()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Print_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.Condition()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *AssignContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BabyDuckParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(BabyDuckParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Expression()
	}
	{
		p.SetState(120)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICycleContext is an interface to support dynamic dispatch.
type ICycleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Body() IBodyContext
	SEMICOLON() antlr.TerminalNode

	// IsCycleContext differentiates from other interfaces.
	IsCycleContext()
}

type CycleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCycleContext() *CycleContext {
	var p = new(CycleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_cycle
	return p
}

func InitEmptyCycleContext(p *CycleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_cycle
}

func (*CycleContext) IsCycleContext() {}

func NewCycleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CycleContext {
	var p = new(CycleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_cycle

	return p
}

func (s *CycleContext) GetParser() antlr.Parser { return s.parser }

func (s *CycleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *CycleContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CycleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *CycleContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *CycleContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *CycleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CycleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CycleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterCycle(s)
	}
}

func (s *CycleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitCycle(s)
	}
}

func (s *CycleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitCycle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Cycle() (localctx ICycleContext) {
	localctx = NewCycleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BabyDuckParserRULE_cycle)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(BabyDuckParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Expression()
	}
	{
		p.SetState(125)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(BabyDuckParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Body()
	}
	{
		p.SetState(128)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Body() IBodyContext
	Else_part() IElse_partContext
	SEMICOLON() antlr.TerminalNode

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *ConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *ConditionContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ConditionContext) Else_part() IElse_partContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_partContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_partContext)
}

func (s *ConditionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BabyDuckParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(BabyDuckParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Expression()
	}
	{
		p.SetState(133)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Body()
	}
	{
		p.SetState(135)
		p.Else_part()
	}
	{
		p.SetState(136)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_partContext is an interface to support dynamic dispatch.
type IElse_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Body() IBodyContext

	// IsElse_partContext differentiates from other interfaces.
	IsElse_partContext()
}

type Else_partContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_partContext() *Else_partContext {
	var p = new(Else_partContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_else_part
	return p
}

func InitEmptyElse_partContext(p *Else_partContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_else_part
}

func (*Else_partContext) IsElse_partContext() {}

func NewElse_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_partContext {
	var p = new(Else_partContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_else_part

	return p
}

func (s *Else_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_partContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *Else_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterElse_part(s)
	}
}

func (s *Else_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitElse_part(s)
	}
}

func (s *Else_partContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitElse_part(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Else_part() (localctx IElse_partContext) {
	localctx = NewElse_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BabyDuckParserRULE_else_part)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__10 {
		{
			p.SetState(138)
			p.Match(BabyDuckParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Body()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrint_stmtContext is an interface to support dynamic dispatch.
type IPrint_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllPrintexpr() []IPrintexprContext
	Printexpr(i int) IPrintexprContext
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPrint_stmtContext differentiates from other interfaces.
	IsPrint_stmtContext()
}

type Print_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_stmtContext() *Print_stmtContext {
	var p = new(Print_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_print_stmt
	return p
}

func InitEmptyPrint_stmtContext(p *Print_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_print_stmt
}

func (*Print_stmtContext) IsPrint_stmtContext() {}

func NewPrint_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_stmtContext {
	var p = new(Print_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_print_stmt

	return p
}

func (s *Print_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_stmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *Print_stmtContext) AllPrintexpr() []IPrintexprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintexprContext); ok {
			len++
		}
	}

	tst := make([]IPrintexprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintexprContext); ok {
			tst[i] = t.(IPrintexprContext)
			i++
		}
	}

	return tst
}

func (s *Print_stmtContext) Printexpr(i int) IPrintexprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintexprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintexprContext)
}

func (s *Print_stmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *Print_stmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *Print_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *Print_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *Print_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterPrint_stmt(s)
	}
}

func (s *Print_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitPrint_stmt(s)
	}
}

func (s *Print_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitPrint_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Print_stmt() (localctx IPrint_stmtContext) {
	localctx = NewPrint_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BabyDuckParserRULE_print_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(BabyDuckParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Printexpr()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserCOMMA {
		{
			p.SetState(145)
			p.Match(BabyDuckParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Printexpr()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintexprContext is an interface to support dynamic dispatch.
type IPrintexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	STRING() antlr.TerminalNode

	// IsPrintexprContext differentiates from other interfaces.
	IsPrintexprContext()
}

type PrintexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintexprContext() *PrintexprContext {
	var p = new(PrintexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printexpr
	return p
}

func InitEmptyPrintexprContext(p *PrintexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printexpr
}

func (*PrintexprContext) IsPrintexprContext() {}

func NewPrintexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintexprContext {
	var p = new(PrintexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_printexpr

	return p
}

func (s *PrintexprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintexprContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *PrintexprContext) STRING() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSTRING, 0)
}

func (s *PrintexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterPrintexpr(s)
	}
}

func (s *PrintexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitPrintexpr(s)
	}
}

func (s *PrintexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitPrintexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Printexpr() (localctx IPrintexprContext) {
	localctx = NewPrintexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BabyDuckParserRULE_printexpr)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserT__15, BabyDuckParserT__16, BabyDuckParserLPAREN, BabyDuckParserID, BabyDuckParserINT, BabyDuckParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Exp()
		}

	case BabyDuckParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(BabyDuckParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserINT, 0)
}

func (s *ConstantContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserFLOAT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BabyDuckParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserINT || _la == BabyDuckParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	Relational() IRelationalContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpressionContext) Relational() IRelationalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BabyDuckParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Exp()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57344) != 0 {
		{
			p.SetState(162)
			p.Relational()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalContext is an interface to support dynamic dispatch.
type IRelationalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Relop() IRelopContext
	Exp() IExpContext

	// IsRelationalContext differentiates from other interfaces.
	IsRelationalContext()
}

type RelationalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalContext() *RelationalContext {
	var p = new(RelationalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relational
	return p
}

func InitEmptyRelationalContext(p *RelationalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relational
}

func (*RelationalContext) IsRelationalContext() {}

func NewRelationalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalContext {
	var p = new(RelationalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_relational

	return p
}

func (s *RelationalContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *RelationalContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *RelationalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterRelational(s)
	}
}

func (s *RelationalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitRelational(s)
	}
}

func (s *RelationalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitRelational(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Relational() (localctx IRelationalContext) {
	localctx = NewRelationalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BabyDuckParserRULE_relational)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Relop()
	}
	{
		p.SetState(166)
		p.Exp()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relop
	return p
}

func InitEmptyRelopContext(p *RelopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relop
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }
func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (s *RelopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitRelop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BabyDuckParserRULE_relop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57344) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op rule contexts.
	GetOp() IAddopContext

	// SetOp sets the op rule contexts.
	SetOp(IAddopContext)

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllAddop() []IAddopContext
	Addop(i int) IAddopContext

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     IAddopContext
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetOp() IAddopContext { return s.op }

func (s *ExpContext) SetOp(v IAddopContext) { s.op = v }

func (s *ExpContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpContext) AllAddop() []IAddopContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAddopContext); ok {
			len++
		}
	}

	tst := make([]IAddopContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAddopContext); ok {
			tst[i] = t.(IAddopContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Addop(i int) IAddopContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddopContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddopContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitExp(s)
	}
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BabyDuckParserRULE_exp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Term()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__15 || _la == BabyDuckParserT__16 {
		{
			p.SetState(171)

			var _x = p.Addop()

			localctx.(*ExpContext).op = _x
		}
		{
			p.SetState(172)
			p.Term()
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddopContext is an interface to support dynamic dispatch.
type IAddopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAddopContext differentiates from other interfaces.
	IsAddopContext()
}

type AddopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddopContext() *AddopContext {
	var p = new(AddopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_addop
	return p
}

func InitEmptyAddopContext(p *AddopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_addop
}

func (*AddopContext) IsAddopContext() {}

func NewAddopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddopContext {
	var p = new(AddopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_addop

	return p
}

func (s *AddopContext) GetParser() antlr.Parser { return s.parser }
func (s *AddopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterAddop(s)
	}
}

func (s *AddopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitAddop(s)
	}
}

func (s *AddopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitAddop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Addop() (localctx IAddopContext) {
	localctx = NewAddopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BabyDuckParserRULE_addop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserT__15 || _la == BabyDuckParserT__16) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllMulop() []IMulopContext
	Mulop(i int) IMulopContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMulop() []IMulopContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMulopContext); ok {
			len++
		}
	}

	tst := make([]IMulopContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMulopContext); ok {
			tst[i] = t.(IMulopContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Mulop(i int) IMulopContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulopContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulopContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BabyDuckParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Factor()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserT__17 || _la == BabyDuckParserT__18 {
		{
			p.SetState(182)
			p.Mulop()
		}
		{
			p.SetState(183)
			p.Factor()
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMulopContext is an interface to support dynamic dispatch.
type IMulopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMulopContext differentiates from other interfaces.
	IsMulopContext()
}

type MulopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulopContext() *MulopContext {
	var p = new(MulopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_mulop
	return p
}

func InitEmptyMulopContext(p *MulopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_mulop
}

func (*MulopContext) IsMulopContext() {}

func NewMulopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulopContext {
	var p = new(MulopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_mulop

	return p
}

func (s *MulopContext) GetParser() antlr.Parser { return s.parser }
func (s *MulopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterMulop(s)
	}
}

func (s *MulopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitMulop(s)
	}
}

func (s *MulopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitMulop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Mulop() (localctx IMulopContext) {
	localctx = NewMulopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BabyDuckParserRULE_mulop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserT__17 || _la == BabyDuckParserT__18) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Parexpr() IParexprContext
	Factorsign() IFactorsignContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Parexpr() IParexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParexprContext)
}

func (s *FactorContext) Factorsign() IFactorsignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorsignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorsignContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BabyDuckParserRULE_factor)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Parexpr()
		}

	case BabyDuckParserT__15, BabyDuckParserT__16, BabyDuckParserID, BabyDuckParserINT, BabyDuckParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Factorsign()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParexprContext is an interface to support dynamic dispatch.
type IParexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsParexprContext differentiates from other interfaces.
	IsParexprContext()
}

type ParexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParexprContext() *ParexprContext {
	var p = new(ParexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_parexpr
	return p
}

func InitEmptyParexprContext(p *ParexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_parexpr
}

func (*ParexprContext) IsParexprContext() {}

func NewParexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParexprContext {
	var p = new(ParexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_parexpr

	return p
}

func (s *ParexprContext) GetParser() antlr.Parser { return s.parser }

func (s *ParexprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *ParexprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParexprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *ParexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterParexpr(s)
	}
}

func (s *ParexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitParexpr(s)
	}
}

func (s *ParexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitParexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Parexpr() (localctx IParexprContext) {
	localctx = NewParexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BabyDuckParserRULE_parexpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Expression()
	}
	{
		p.SetState(198)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorsignContext is an interface to support dynamic dispatch.
type IFactorsignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	Addop() IAddopContext

	// IsFactorsignContext differentiates from other interfaces.
	IsFactorsignContext()
}

type FactorsignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorsignContext() *FactorsignContext {
	var p = new(FactorsignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factorsign
	return p
}

func InitEmptyFactorsignContext(p *FactorsignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factorsign
}

func (*FactorsignContext) IsFactorsignContext() {}

func NewFactorsignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorsignContext {
	var p = new(FactorsignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_factorsign

	return p
}

func (s *FactorsignContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorsignContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FactorsignContext) Addop() IAddopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddopContext)
}

func (s *FactorsignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorsignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorsignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFactorsign(s)
	}
}

func (s *FactorsignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFactorsign(s)
	}
}

func (s *FactorsignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFactorsign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Factorsign() (localctx IFactorsignContext) {
	localctx = NewFactorsignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BabyDuckParserRULE_factorsign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__15 || _la == BabyDuckParserT__16 {
		{
			p.SetState(200)
			p.Addop()
		}

	}
	{
		p.SetState(203)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Constant() IConstantContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *ValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BabyDuckParserRULE_value)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(BabyDuckParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BabyDuckParserINT, BabyDuckParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Constant()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncsContext is an interface to support dynamic dispatch.
type IFuncsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Func_() IFuncContext

	// IsFuncsContext differentiates from other interfaces.
	IsFuncsContext()
}

type FuncsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncsContext() *FuncsContext {
	var p = new(FuncsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcs
	return p
}

func InitEmptyFuncsContext(p *FuncsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcs
}

func (*FuncsContext) IsFuncsContext() {}

func NewFuncsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncsContext {
	var p = new(FuncsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_funcs

	return p
}

func (s *FuncsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncsContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *FuncsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFuncs(s)
	}
}

func (s *FuncsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFuncs(s)
	}
}

func (s *FuncsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFuncs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Funcs() (localctx IFuncsContext) {
	localctx = NewFuncsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BabyDuckParserRULE_funcs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Func_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VOID() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Funcbody() IFuncbodyContext
	Param_list() IParam_listContext

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) VOID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserVOID, 0)
}

func (s *FuncContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *FuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *FuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *FuncContext) Funcbody() IFuncbodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncbodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FuncContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BabyDuckParserRULE_func)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(BabyDuckParserVOID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserID {
		{
			p.SetState(214)
			p.Param_list()
		}

	}
	{
		p.SetState(217)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Funcbody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param_list
	return p
}

func InitEmptyParam_listContext(p *Param_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param_list
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *Param_listContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Param_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *Param_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterParam_list(s)
	}
}

func (s *Param_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitParam_list(s)
	}
}

func (s *Param_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitParam_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Param_list() (localctx IParam_listContext) {
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BabyDuckParserRULE_param_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Param()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserCOMMA {
		{
			p.SetState(221)
			p.Match(BabyDuckParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Param()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *ParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOLON, 0)
}

func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BabyDuckParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(BabyDuckParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncbodyContext is an interface to support dynamic dispatch.
type IFuncbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	Body() IBodyContext
	RBRACKET() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Vars() IVarsContext

	// IsFuncbodyContext differentiates from other interfaces.
	IsFuncbodyContext()
}

type FuncbodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncbodyContext() *FuncbodyContext {
	var p = new(FuncbodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcbody
	return p
}

func InitEmptyFuncbodyContext(p *FuncbodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_funcbody
}

func (*FuncbodyContext) IsFuncbodyContext() {}

func NewFuncbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncbodyContext {
	var p = new(FuncbodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_funcbody

	return p
}

func (s *FuncbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncbodyContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLBRACKET, 0)
}

func (s *FuncbodyContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *FuncbodyContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRBRACKET, 0)
}

func (s *FuncbodyContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *FuncbodyContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *FuncbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFuncbody(s)
	}
}

func (s *FuncbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFuncbody(s)
	}
}

func (s *FuncbodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFuncbody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Funcbody() (localctx IFuncbodyContext) {
	localctx = NewFuncbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BabyDuckParserRULE_funcbody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(BabyDuckParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserT__3 {
		{
			p.SetState(233)
			p.Vars()
		}

	}
	{
		p.SetState(236)
		p.Body()
	}
	{
		p.SetState(237)
		p.Match(BabyDuckParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IF_callContext is an interface to support dynamic dispatch.
type IF_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Arg_list() IArg_listContext

	// IsF_callContext differentiates from other interfaces.
	IsF_callContext()
}

type F_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_callContext() *F_callContext {
	var p = new(F_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_f_call
	return p
}

func InitEmptyF_callContext(p *F_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_f_call
}

func (*F_callContext) IsF_callContext() {}

func NewF_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_callContext {
	var p = new(F_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_f_call

	return p
}

func (s *F_callContext) GetParser() antlr.Parser { return s.parser }

func (s *F_callContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *F_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *F_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *F_callContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *F_callContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *F_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterF_call(s)
	}
}

func (s *F_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitF_call(s)
	}
}

func (s *F_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitF_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) F_call() (localctx IF_callContext) {
	localctx = NewF_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BabyDuckParserRULE_f_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7524777984) != 0 {
		{
			p.SetState(242)
			p.Arg_list()
		}

	}
	{
		p.SetState(245)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Arg_listContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Arg_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *Arg_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterArg_list(s)
	}
}

func (s *Arg_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitArg_list(s)
	}
}

func (s *Arg_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitArg_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BabyDuckParserRULE_arg_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Expression()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserCOMMA {
		{
			p.SetState(249)
			p.Match(BabyDuckParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Expression()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
