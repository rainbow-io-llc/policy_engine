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
		"'use'", "'third parties'", "'partners'", "'internal services'", "'AND'",
		"'OR'", "'time_range'", "'='", "'location'",
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
		4, 0, 19, 177, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 4, 17, 165,
		8, 17, 11, 17, 12, 17, 166, 1, 17, 1, 17, 1, 18, 4, 18, 172, 8, 18, 11,
		18, 12, 18, 173, 1, 18, 1, 18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 3, 1, 0, 48, 57, 5, 0, 32, 32,
		44, 44, 48, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 178, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0,
		0, 0, 3, 45, 1, 0, 0, 0, 5, 50, 1, 0, 0, 0, 7, 53, 1, 0, 0, 0, 9, 58, 1,
		0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 66, 1, 0, 0, 0, 15, 74, 1, 0, 0, 0, 17,
		78, 1, 0, 0, 0, 19, 92, 1, 0, 0, 0, 21, 101, 1, 0, 0, 0, 23, 119, 1, 0,
		0, 0, 25, 123, 1, 0, 0, 0, 27, 126, 1, 0, 0, 0, 29, 137, 1, 0, 0, 0, 31,
		139, 1, 0, 0, 0, 33, 148, 1, 0, 0, 0, 35, 162, 1, 0, 0, 0, 37, 171, 1,
		0, 0, 0, 39, 40, 5, 65, 0, 0, 40, 41, 5, 76, 0, 0, 41, 42, 5, 76, 0, 0,
		42, 43, 5, 79, 0, 0, 43, 44, 5, 87, 0, 0, 44, 2, 1, 0, 0, 0, 45, 46, 5,
		68, 0, 0, 46, 47, 5, 69, 0, 0, 47, 48, 5, 78, 0, 0, 48, 49, 5, 89, 0, 0,
		49, 4, 1, 0, 0, 0, 50, 51, 5, 84, 0, 0, 51, 52, 5, 79, 0, 0, 52, 6, 1,
		0, 0, 0, 53, 54, 5, 87, 0, 0, 54, 55, 5, 72, 0, 0, 55, 56, 5, 69, 0, 0,
		56, 57, 5, 78, 0, 0, 57, 8, 1, 0, 0, 0, 58, 59, 5, 59, 0, 0, 59, 10, 1,
		0, 0, 0, 60, 61, 5, 115, 0, 0, 61, 62, 5, 104, 0, 0, 62, 63, 5, 97, 0,
		0, 63, 64, 5, 114, 0, 0, 64, 65, 5, 101, 0, 0, 65, 12, 1, 0, 0, 0, 66,
		67, 5, 99, 0, 0, 67, 68, 5, 111, 0, 0, 68, 69, 5, 108, 0, 0, 69, 70, 5,
		108, 0, 0, 70, 71, 5, 101, 0, 0, 71, 72, 5, 99, 0, 0, 72, 73, 5, 116, 0,
		0, 73, 14, 1, 0, 0, 0, 74, 75, 5, 117, 0, 0, 75, 76, 5, 115, 0, 0, 76,
		77, 5, 101, 0, 0, 77, 16, 1, 0, 0, 0, 78, 79, 5, 116, 0, 0, 79, 80, 5,
		104, 0, 0, 80, 81, 5, 105, 0, 0, 81, 82, 5, 114, 0, 0, 82, 83, 5, 100,
		0, 0, 83, 84, 5, 32, 0, 0, 84, 85, 5, 112, 0, 0, 85, 86, 5, 97, 0, 0, 86,
		87, 5, 114, 0, 0, 87, 88, 5, 116, 0, 0, 88, 89, 5, 105, 0, 0, 89, 90, 5,
		101, 0, 0, 90, 91, 5, 115, 0, 0, 91, 18, 1, 0, 0, 0, 92, 93, 5, 112, 0,
		0, 93, 94, 5, 97, 0, 0, 94, 95, 5, 114, 0, 0, 95, 96, 5, 116, 0, 0, 96,
		97, 5, 110, 0, 0, 97, 98, 5, 101, 0, 0, 98, 99, 5, 114, 0, 0, 99, 100,
		5, 115, 0, 0, 100, 20, 1, 0, 0, 0, 101, 102, 5, 105, 0, 0, 102, 103, 5,
		110, 0, 0, 103, 104, 5, 116, 0, 0, 104, 105, 5, 101, 0, 0, 105, 106, 5,
		114, 0, 0, 106, 107, 5, 110, 0, 0, 107, 108, 5, 97, 0, 0, 108, 109, 5,
		108, 0, 0, 109, 110, 5, 32, 0, 0, 110, 111, 5, 115, 0, 0, 111, 112, 5,
		101, 0, 0, 112, 113, 5, 114, 0, 0, 113, 114, 5, 118, 0, 0, 114, 115, 5,
		105, 0, 0, 115, 116, 5, 99, 0, 0, 116, 117, 5, 101, 0, 0, 117, 118, 5,
		115, 0, 0, 118, 22, 1, 0, 0, 0, 119, 120, 5, 65, 0, 0, 120, 121, 5, 78,
		0, 0, 121, 122, 5, 68, 0, 0, 122, 24, 1, 0, 0, 0, 123, 124, 5, 79, 0, 0,
		124, 125, 5, 82, 0, 0, 125, 26, 1, 0, 0, 0, 126, 127, 5, 116, 0, 0, 127,
		128, 5, 105, 0, 0, 128, 129, 5, 109, 0, 0, 129, 130, 5, 101, 0, 0, 130,
		131, 5, 95, 0, 0, 131, 132, 5, 114, 0, 0, 132, 133, 5, 97, 0, 0, 133, 134,
		5, 110, 0, 0, 134, 135, 5, 103, 0, 0, 135, 136, 5, 101, 0, 0, 136, 28,
		1, 0, 0, 0, 137, 138, 5, 61, 0, 0, 138, 30, 1, 0, 0, 0, 139, 140, 5, 108,
		0, 0, 140, 141, 5, 111, 0, 0, 141, 142, 5, 99, 0, 0, 142, 143, 5, 97, 0,
		0, 143, 144, 5, 116, 0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5, 111, 0,
		0, 146, 147, 5, 110, 0, 0, 147, 32, 1, 0, 0, 0, 148, 149, 5, 34, 0, 0,
		149, 150, 7, 0, 0, 0, 150, 151, 7, 0, 0, 0, 151, 152, 5, 58, 0, 0, 152,
		153, 7, 0, 0, 0, 153, 154, 7, 0, 0, 0, 154, 155, 5, 45, 0, 0, 155, 156,
		7, 0, 0, 0, 156, 157, 7, 0, 0, 0, 157, 158, 5, 58, 0, 0, 158, 159, 7, 0,
		0, 0, 159, 160, 7, 0, 0, 0, 160, 161, 5, 34, 0, 0, 161, 34, 1, 0, 0, 0,
		162, 164, 5, 34, 0, 0, 163, 165, 7, 1, 0, 0, 164, 163, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168,
		1, 0, 0, 0, 168, 169, 5, 34, 0, 0, 169, 36, 1, 0, 0, 0, 170, 172, 7, 2,
		0, 0, 171, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 6, 18, 0, 0, 176,
		38, 1, 0, 0, 0, 3, 0, 166, 173, 1, 6, 0, 0,
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
	privacyLexerT__13      = 14
	privacyLexerT__14      = 15
	privacyLexerT__15      = 16
	privacyLexerTIME_RANGE = 17
	privacyLexerLOCATION   = 18
	privacyLexerWS         = 19
)
