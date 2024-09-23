// Code generated from antlr/authz.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type authzLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AuthzLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func authzlexerLexerInit() {
	staticData := &AuthzLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'ALLOW'", "'DENY'", "'FROM'", "'TO'", "'WHEN'", "';'", "'user'",
		"'service'", "'share'", "'collect'", "'use'", "'data'", "'feature'",
		"'time_range'", "'='", "'location'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"TIME_RANGE", "LOCATION", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "TIME_RANGE",
		"LOCATION", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 160, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 4, 17, 148, 8, 17, 11, 17, 12, 17, 149, 1, 17,
		1, 17, 1, 18, 4, 18, 155, 8, 18, 11, 18, 12, 18, 156, 1, 18, 1, 18, 0,
		0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		1, 0, 3, 1, 0, 48, 57, 5, 0, 32, 32, 44, 44, 48, 57, 65, 90, 97, 122, 3,
		0, 9, 10, 13, 13, 32, 32, 161, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 50, 1, 0,
		0, 0, 7, 55, 1, 0, 0, 0, 9, 58, 1, 0, 0, 0, 11, 63, 1, 0, 0, 0, 13, 65,
		1, 0, 0, 0, 15, 70, 1, 0, 0, 0, 17, 78, 1, 0, 0, 0, 19, 84, 1, 0, 0, 0,
		21, 92, 1, 0, 0, 0, 23, 96, 1, 0, 0, 0, 25, 101, 1, 0, 0, 0, 27, 109, 1,
		0, 0, 0, 29, 120, 1, 0, 0, 0, 31, 122, 1, 0, 0, 0, 33, 131, 1, 0, 0, 0,
		35, 145, 1, 0, 0, 0, 37, 154, 1, 0, 0, 0, 39, 40, 5, 65, 0, 0, 40, 41,
		5, 76, 0, 0, 41, 42, 5, 76, 0, 0, 42, 43, 5, 79, 0, 0, 43, 44, 5, 87, 0,
		0, 44, 2, 1, 0, 0, 0, 45, 46, 5, 68, 0, 0, 46, 47, 5, 69, 0, 0, 47, 48,
		5, 78, 0, 0, 48, 49, 5, 89, 0, 0, 49, 4, 1, 0, 0, 0, 50, 51, 5, 70, 0,
		0, 51, 52, 5, 82, 0, 0, 52, 53, 5, 79, 0, 0, 53, 54, 5, 77, 0, 0, 54, 6,
		1, 0, 0, 0, 55, 56, 5, 84, 0, 0, 56, 57, 5, 79, 0, 0, 57, 8, 1, 0, 0, 0,
		58, 59, 5, 87, 0, 0, 59, 60, 5, 72, 0, 0, 60, 61, 5, 69, 0, 0, 61, 62,
		5, 78, 0, 0, 62, 10, 1, 0, 0, 0, 63, 64, 5, 59, 0, 0, 64, 12, 1, 0, 0,
		0, 65, 66, 5, 117, 0, 0, 66, 67, 5, 115, 0, 0, 67, 68, 5, 101, 0, 0, 68,
		69, 5, 114, 0, 0, 69, 14, 1, 0, 0, 0, 70, 71, 5, 115, 0, 0, 71, 72, 5,
		101, 0, 0, 72, 73, 5, 114, 0, 0, 73, 74, 5, 118, 0, 0, 74, 75, 5, 105,
		0, 0, 75, 76, 5, 99, 0, 0, 76, 77, 5, 101, 0, 0, 77, 16, 1, 0, 0, 0, 78,
		79, 5, 115, 0, 0, 79, 80, 5, 104, 0, 0, 80, 81, 5, 97, 0, 0, 81, 82, 5,
		114, 0, 0, 82, 83, 5, 101, 0, 0, 83, 18, 1, 0, 0, 0, 84, 85, 5, 99, 0,
		0, 85, 86, 5, 111, 0, 0, 86, 87, 5, 108, 0, 0, 87, 88, 5, 108, 0, 0, 88,
		89, 5, 101, 0, 0, 89, 90, 5, 99, 0, 0, 90, 91, 5, 116, 0, 0, 91, 20, 1,
		0, 0, 0, 92, 93, 5, 117, 0, 0, 93, 94, 5, 115, 0, 0, 94, 95, 5, 101, 0,
		0, 95, 22, 1, 0, 0, 0, 96, 97, 5, 100, 0, 0, 97, 98, 5, 97, 0, 0, 98, 99,
		5, 116, 0, 0, 99, 100, 5, 97, 0, 0, 100, 24, 1, 0, 0, 0, 101, 102, 5, 102,
		0, 0, 102, 103, 5, 101, 0, 0, 103, 104, 5, 97, 0, 0, 104, 105, 5, 116,
		0, 0, 105, 106, 5, 117, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108, 5, 101,
		0, 0, 108, 26, 1, 0, 0, 0, 109, 110, 5, 116, 0, 0, 110, 111, 5, 105, 0,
		0, 111, 112, 5, 109, 0, 0, 112, 113, 5, 101, 0, 0, 113, 114, 5, 95, 0,
		0, 114, 115, 5, 114, 0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 5, 110, 0,
		0, 117, 118, 5, 103, 0, 0, 118, 119, 5, 101, 0, 0, 119, 28, 1, 0, 0, 0,
		120, 121, 5, 61, 0, 0, 121, 30, 1, 0, 0, 0, 122, 123, 5, 108, 0, 0, 123,
		124, 5, 111, 0, 0, 124, 125, 5, 99, 0, 0, 125, 126, 5, 97, 0, 0, 126, 127,
		5, 116, 0, 0, 127, 128, 5, 105, 0, 0, 128, 129, 5, 111, 0, 0, 129, 130,
		5, 110, 0, 0, 130, 32, 1, 0, 0, 0, 131, 132, 5, 34, 0, 0, 132, 133, 7,
		0, 0, 0, 133, 134, 6, 16, 0, 0, 134, 135, 5, 58, 0, 0, 135, 136, 7, 0,
		0, 0, 136, 137, 6, 16, 1, 0, 137, 138, 5, 45, 0, 0, 138, 139, 7, 0, 0,
		0, 139, 140, 6, 16, 2, 0, 140, 141, 5, 58, 0, 0, 141, 142, 7, 0, 0, 0,
		142, 143, 6, 16, 3, 0, 143, 144, 5, 34, 0, 0, 144, 34, 1, 0, 0, 0, 145,
		147, 5, 34, 0, 0, 146, 148, 7, 1, 0, 0, 147, 146, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0,
		0, 0, 151, 152, 5, 34, 0, 0, 152, 36, 1, 0, 0, 0, 153, 155, 7, 2, 0, 0,
		154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156,
		157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 6, 18, 4, 0, 159, 38,
		1, 0, 0, 0, 3, 0, 149, 156, 5, 1, 16, 0, 1, 16, 1, 1, 16, 2, 1, 16, 3,
		6, 0, 0,
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

// authzLexerInit initializes any static state used to implement authzLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewauthzLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AuthzLexerInit() {
	staticData := &AuthzLexerLexerStaticData
	staticData.once.Do(authzlexerLexerInit)
}

// NewauthzLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewauthzLexer(input antlr.CharStream) *authzLexer {
	AuthzLexerInit()
	l := new(authzLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AuthzLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "authz.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// authzLexer tokens.
const (
	authzLexerT__0       = 1
	authzLexerT__1       = 2
	authzLexerT__2       = 3
	authzLexerT__3       = 4
	authzLexerT__4       = 5
	authzLexerT__5       = 6
	authzLexerT__6       = 7
	authzLexerT__7       = 8
	authzLexerT__8       = 9
	authzLexerT__9       = 10
	authzLexerT__10      = 11
	authzLexerT__11      = 12
	authzLexerT__12      = 13
	authzLexerT__13      = 14
	authzLexerT__14      = 15
	authzLexerT__15      = 16
	authzLexerTIME_RANGE = 17
	authzLexerLOCATION   = 18
	authzLexerWS         = 19
)

func (l *authzLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 16:
		l.TIME_RANGE_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *authzLexer) TIME_RANGE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		2

	case 1:
		2

	case 2:
		2

	case 3:
		2

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
