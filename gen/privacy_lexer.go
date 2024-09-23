// Code generated from privacy.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type privacyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PrivacyLexerLexerStaticData struct {
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

func privacylexerLexerInit() {
	staticData := &PrivacyLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'ALLOW'", "'DENY'", "'TO'", "'WHEN'", "';'", "'share'", "'collect'",
		"'use'", "'third parties'", "'partners'", "'internal services'", "'time is within'",
		"'location is'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "TIME_RANGE",
		"LOCATION", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "TIME_RANGE", "LOCATION", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 169, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 4, 14, 157, 8, 14, 11, 14, 12, 14, 158, 1, 14, 1, 14, 1, 15, 4, 15,
		164, 8, 15, 11, 15, 12, 15, 165, 1, 15, 1, 15, 0, 0, 16, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 1, 0, 3, 1, 0, 48, 57, 5, 0, 32, 32, 44, 44,
		48, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 170, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1,
		33, 1, 0, 0, 0, 3, 39, 1, 0, 0, 0, 5, 44, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0,
		9, 52, 1, 0, 0, 0, 11, 54, 1, 0, 0, 0, 13, 60, 1, 0, 0, 0, 15, 68, 1, 0,
		0, 0, 17, 72, 1, 0, 0, 0, 19, 86, 1, 0, 0, 0, 21, 95, 1, 0, 0, 0, 23, 113,
		1, 0, 0, 0, 25, 128, 1, 0, 0, 0, 27, 140, 1, 0, 0, 0, 29, 154, 1, 0, 0,
		0, 31, 163, 1, 0, 0, 0, 33, 34, 5, 65, 0, 0, 34, 35, 5, 76, 0, 0, 35, 36,
		5, 76, 0, 0, 36, 37, 5, 79, 0, 0, 37, 38, 5, 87, 0, 0, 38, 2, 1, 0, 0,
		0, 39, 40, 5, 68, 0, 0, 40, 41, 5, 69, 0, 0, 41, 42, 5, 78, 0, 0, 42, 43,
		5, 89, 0, 0, 43, 4, 1, 0, 0, 0, 44, 45, 5, 84, 0, 0, 45, 46, 5, 79, 0,
		0, 46, 6, 1, 0, 0, 0, 47, 48, 5, 87, 0, 0, 48, 49, 5, 72, 0, 0, 49, 50,
		5, 69, 0, 0, 50, 51, 5, 78, 0, 0, 51, 8, 1, 0, 0, 0, 52, 53, 5, 59, 0,
		0, 53, 10, 1, 0, 0, 0, 54, 55, 5, 115, 0, 0, 55, 56, 5, 104, 0, 0, 56,
		57, 5, 97, 0, 0, 57, 58, 5, 114, 0, 0, 58, 59, 5, 101, 0, 0, 59, 12, 1,
		0, 0, 0, 60, 61, 5, 99, 0, 0, 61, 62, 5, 111, 0, 0, 62, 63, 5, 108, 0,
		0, 63, 64, 5, 108, 0, 0, 64, 65, 5, 101, 0, 0, 65, 66, 5, 99, 0, 0, 66,
		67, 5, 116, 0, 0, 67, 14, 1, 0, 0, 0, 68, 69, 5, 117, 0, 0, 69, 70, 5,
		115, 0, 0, 70, 71, 5, 101, 0, 0, 71, 16, 1, 0, 0, 0, 72, 73, 5, 116, 0,
		0, 73, 74, 5, 104, 0, 0, 74, 75, 5, 105, 0, 0, 75, 76, 5, 114, 0, 0, 76,
		77, 5, 100, 0, 0, 77, 78, 5, 32, 0, 0, 78, 79, 5, 112, 0, 0, 79, 80, 5,
		97, 0, 0, 80, 81, 5, 114, 0, 0, 81, 82, 5, 116, 0, 0, 82, 83, 5, 105, 0,
		0, 83, 84, 5, 101, 0, 0, 84, 85, 5, 115, 0, 0, 85, 18, 1, 0, 0, 0, 86,
		87, 5, 112, 0, 0, 87, 88, 5, 97, 0, 0, 88, 89, 5, 114, 0, 0, 89, 90, 5,
		116, 0, 0, 90, 91, 5, 110, 0, 0, 91, 92, 5, 101, 0, 0, 92, 93, 5, 114,
		0, 0, 93, 94, 5, 115, 0, 0, 94, 20, 1, 0, 0, 0, 95, 96, 5, 105, 0, 0, 96,
		97, 5, 110, 0, 0, 97, 98, 5, 116, 0, 0, 98, 99, 5, 101, 0, 0, 99, 100,
		5, 114, 0, 0, 100, 101, 5, 110, 0, 0, 101, 102, 5, 97, 0, 0, 102, 103,
		5, 108, 0, 0, 103, 104, 5, 32, 0, 0, 104, 105, 5, 115, 0, 0, 105, 106,
		5, 101, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108, 5, 118, 0, 0, 108, 109,
		5, 105, 0, 0, 109, 110, 5, 99, 0, 0, 110, 111, 5, 101, 0, 0, 111, 112,
		5, 115, 0, 0, 112, 22, 1, 0, 0, 0, 113, 114, 5, 116, 0, 0, 114, 115, 5,
		105, 0, 0, 115, 116, 5, 109, 0, 0, 116, 117, 5, 101, 0, 0, 117, 118, 5,
		32, 0, 0, 118, 119, 5, 105, 0, 0, 119, 120, 5, 115, 0, 0, 120, 121, 5,
		32, 0, 0, 121, 122, 5, 119, 0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5,
		116, 0, 0, 124, 125, 5, 104, 0, 0, 125, 126, 5, 105, 0, 0, 126, 127, 5,
		110, 0, 0, 127, 24, 1, 0, 0, 0, 128, 129, 5, 108, 0, 0, 129, 130, 5, 111,
		0, 0, 130, 131, 5, 99, 0, 0, 131, 132, 5, 97, 0, 0, 132, 133, 5, 116, 0,
		0, 133, 134, 5, 105, 0, 0, 134, 135, 5, 111, 0, 0, 135, 136, 5, 110, 0,
		0, 136, 137, 5, 32, 0, 0, 137, 138, 5, 105, 0, 0, 138, 139, 5, 115, 0,
		0, 139, 26, 1, 0, 0, 0, 140, 141, 5, 34, 0, 0, 141, 142, 7, 0, 0, 0, 142,
		143, 6, 13, 0, 0, 143, 144, 5, 58, 0, 0, 144, 145, 7, 0, 0, 0, 145, 146,
		6, 13, 1, 0, 146, 147, 5, 45, 0, 0, 147, 148, 7, 0, 0, 0, 148, 149, 6,
		13, 2, 0, 149, 150, 5, 58, 0, 0, 150, 151, 7, 0, 0, 0, 151, 152, 6, 13,
		3, 0, 152, 153, 5, 34, 0, 0, 153, 28, 1, 0, 0, 0, 154, 156, 5, 34, 0, 0,
		155, 157, 7, 1, 0, 0, 156, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161,
		5, 34, 0, 0, 161, 30, 1, 0, 0, 0, 162, 164, 7, 2, 0, 0, 163, 162, 1, 0,
		0, 0, 164, 165, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0,
		166, 167, 1, 0, 0, 0, 167, 168, 6, 15, 4, 0, 168, 32, 1, 0, 0, 0, 3, 0,
		158, 165, 5, 1, 13, 0, 1, 13, 1, 1, 13, 2, 1, 13, 3, 6, 0, 0,
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

// privacyLexerInit initializes any static state used to implement privacyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewprivacyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PrivacyLexerInit() {
	staticData := &PrivacyLexerLexerStaticData
	staticData.once.Do(privacylexerLexerInit)
}

// NewprivacyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewprivacyLexer(input antlr.CharStream) *privacyLexer {
	PrivacyLexerInit()
	l := new(privacyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PrivacyLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "privacy.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// privacyLexer tokens.
const (
	privacyLexerT__0       = 1
	privacyLexerT__1       = 2
	privacyLexerT__2       = 3
	privacyLexerT__3       = 4
	privacyLexerT__4       = 5
	privacyLexerT__5       = 6
	privacyLexerT__6       = 7
	privacyLexerT__7       = 8
	privacyLexerT__8       = 9
	privacyLexerT__9       = 10
	privacyLexerT__10      = 11
	privacyLexerT__11      = 12
	privacyLexerT__12      = 13
	privacyLexerTIME_RANGE = 14
	privacyLexerLOCATION   = 15
	privacyLexerWS         = 16
)

func (l *privacyLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 13:
		l.TIME_RANGE_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *privacyLexer) TIME_RANGE_Action(localctx antlr.RuleContext, actionIndex int) {
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
