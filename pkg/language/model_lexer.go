// Code generated from ModelLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package language
import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 34, 273, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 6, 28, 207, 10, 28, 
	13, 28, 14, 28, 208, 3, 29, 3, 29, 5, 29, 213, 10, 29, 3, 30, 3, 30, 7, 
	30, 217, 10, 30, 12, 30, 14, 30, 220, 11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 
	7, 31, 226, 10, 31, 12, 31, 14, 31, 229, 11, 31, 3, 31, 3, 31, 3, 32, 3, 
	32, 7, 32, 235, 10, 32, 12, 32, 14, 32, 238, 11, 32, 3, 33, 3, 33, 3, 33, 
	3, 33, 7, 33, 244, 10, 33, 12, 33, 14, 33, 247, 11, 33, 3, 33, 3, 33, 3, 
	33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 257, 10, 34, 12, 34, 14, 
	34, 260, 11, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 6, 35, 268, 
	10, 35, 13, 35, 14, 35, 269, 3, 35, 3, 35, 5, 218, 227, 258, 2, 36, 3, 
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 2, 
	61, 2, 63, 31, 65, 32, 67, 33, 69, 34, 3, 2, 7, 3, 2, 50, 59, 5, 2, 67, 
	92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 12, 12, 
	15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2, 278, 2, 3, 3, 2, 2, 2, 2, 5, 3, 
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 63, 
	3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 3, 
	71, 3, 2, 2, 2, 5, 81, 3, 2, 2, 2, 7, 87, 3, 2, 2, 2, 9, 92, 3, 2, 2, 2, 
	11, 97, 3, 2, 2, 2, 13, 103, 3, 2, 2, 2, 15, 109, 3, 2, 2, 2, 17, 112, 
	3, 2, 2, 2, 19, 117, 3, 2, 2, 2, 21, 125, 3, 2, 2, 2, 23, 132, 3, 2, 2, 
	2, 25, 136, 3, 2, 2, 2, 27, 146, 3, 2, 2, 2, 29, 155, 3, 2, 2, 2, 31, 162, 
	3, 2, 2, 2, 33, 169, 3, 2, 2, 2, 35, 174, 3, 2, 2, 2, 37, 180, 3, 2, 2, 
	2, 39, 189, 3, 2, 2, 2, 41, 191, 3, 2, 2, 2, 43, 193, 3, 2, 2, 2, 45, 195, 
	3, 2, 2, 2, 47, 197, 3, 2, 2, 2, 49, 199, 3, 2, 2, 2, 51, 201, 3, 2, 2, 
	2, 53, 203, 3, 2, 2, 2, 55, 206, 3, 2, 2, 2, 57, 212, 3, 2, 2, 2, 59, 214, 
	3, 2, 2, 2, 61, 223, 3, 2, 2, 2, 63, 232, 3, 2, 2, 2, 65, 239, 3, 2, 2, 
	2, 67, 252, 3, 2, 2, 2, 69, 267, 3, 2, 2, 2, 71, 72, 7, 99, 2, 2, 72, 73, 
	7, 118, 2, 2, 73, 74, 7, 118, 2, 2, 74, 75, 7, 116, 2, 2, 75, 76, 7, 107, 
	2, 2, 76, 77, 7, 100, 2, 2, 77, 78, 7, 119, 2, 2, 78, 79, 7, 118, 2, 2, 
	79, 80, 7, 103, 2, 2, 80, 4, 3, 2, 2, 2, 81, 82, 7, 101, 2, 2, 82, 83, 
	7, 110, 2, 2, 83, 84, 7, 99, 2, 2, 84, 85, 7, 117, 2, 2, 85, 86, 7, 117, 
	2, 2, 86, 6, 3, 2, 2, 2, 87, 88, 7, 101, 2, 2, 88, 89, 7, 113, 2, 2, 89, 
	90, 7, 102, 2, 2, 90, 91, 7, 103, 2, 2, 91, 8, 3, 2, 2, 2, 92, 93, 7, 103, 
	2, 2, 93, 94, 7, 112, 2, 2, 94, 95, 7, 119, 2, 2, 95, 96, 7, 111, 2, 2, 
	96, 10, 3, 2, 2, 2, 97, 98, 7, 103, 2, 2, 98, 99, 7, 116, 2, 2, 99, 100, 
	7, 116, 2, 2, 100, 101, 7, 113, 2, 2, 101, 102, 7, 116, 2, 2, 102, 12, 
	3, 2, 2, 2, 103, 104, 7, 104, 2, 2, 104, 105, 7, 99, 2, 2, 105, 106, 7, 
	110, 2, 2, 106, 107, 7, 117, 2, 2, 107, 108, 7, 103, 2, 2, 108, 14, 3, 
	2, 2, 2, 109, 110, 7, 107, 2, 2, 110, 111, 7, 112, 2, 2, 111, 16, 3, 2, 
	2, 2, 112, 113, 7, 110, 2, 2, 113, 114, 7, 107, 2, 2, 114, 115, 7, 112, 
	2, 2, 115, 116, 7, 109, 2, 2, 116, 18, 3, 2, 2, 2, 117, 118, 7, 110, 2, 
	2, 118, 119, 7, 113, 2, 2, 119, 120, 7, 101, 2, 2, 120, 121, 7, 99, 2, 
	2, 121, 122, 7, 118, 2, 2, 122, 123, 7, 113, 2, 2, 123, 124, 7, 116, 2, 
	2, 124, 20, 3, 2, 2, 2, 125, 126, 7, 111, 2, 2, 126, 127, 7, 103, 2, 2, 
	127, 128, 7, 118, 2, 2, 128, 129, 7, 106, 2, 2, 129, 130, 7, 113, 2, 2, 
	130, 131, 7, 102, 2, 2, 131, 22, 3, 2, 2, 2, 132, 133, 7, 113, 2, 2, 133, 
	134, 7, 119, 2, 2, 134, 135, 7, 118, 2, 2, 135, 24, 3, 2, 2, 2, 136, 137, 
	7, 114, 2, 2, 137, 138, 7, 99, 2, 2, 138, 139, 7, 116, 2, 2, 139, 140, 
	7, 99, 2, 2, 140, 141, 7, 111, 2, 2, 141, 142, 7, 103, 2, 2, 142, 143, 
	7, 118, 2, 2, 143, 144, 7, 103, 2, 2, 144, 145, 7, 116, 2, 2, 145, 26, 
	3, 2, 2, 2, 146, 147, 7, 116, 2, 2, 147, 148, 7, 103, 2, 2, 148, 149, 7, 
	117, 2, 2, 149, 150, 7, 113, 2, 2, 150, 151, 7, 119, 2, 2, 151, 152, 7, 
	116, 2, 2, 152, 153, 7, 101, 2, 2, 153, 154, 7, 103, 2, 2, 154, 28, 3, 
	2, 2, 2, 155, 156, 7, 117, 2, 2, 156, 157, 7, 118, 2, 2, 157, 158, 7, 116, 
	2, 2, 158, 159, 7, 119, 2, 2, 159, 160, 7, 101, 2, 2, 160, 161, 7, 118, 
	2, 2, 161, 30, 3, 2, 2, 2, 162, 163, 7, 118, 2, 2, 163, 164, 7, 99, 2, 
	2, 164, 165, 7, 116, 2, 2, 165, 166, 7, 105, 2, 2, 166, 167, 7, 103, 2, 
	2, 167, 168, 7, 118, 2, 2, 168, 32, 3, 2, 2, 2, 169, 170, 7, 118, 2, 2, 
	170, 171, 7, 116, 2, 2, 171, 172, 7, 119, 2, 2, 172, 173, 7, 103, 2, 2, 
	173, 34, 3, 2, 2, 2, 174, 175, 7, 120, 2, 2, 175, 176, 7, 99, 2, 2, 176, 
	177, 7, 110, 2, 2, 177, 178, 7, 119, 2, 2, 178, 179, 7, 103, 2, 2, 179, 
	36, 3, 2, 2, 2, 180, 181, 7, 120, 2, 2, 181, 182, 7, 99, 2, 2, 182, 183, 
	7, 116, 2, 2, 183, 184, 7, 107, 2, 2, 184, 185, 7, 99, 2, 2, 185, 186, 
	7, 100, 2, 2, 186, 187, 7, 110, 2, 2, 187, 188, 7, 103, 2, 2, 188, 38, 
	3, 2, 2, 2, 189, 190, 7, 66, 2, 2, 190, 40, 3, 2, 2, 2, 191, 192, 7, 125, 
	2, 2, 192, 42, 3, 2, 2, 2, 193, 194, 7, 127, 2, 2, 194, 44, 3, 2, 2, 2, 
	195, 196, 7, 93, 2, 2, 196, 46, 3, 2, 2, 2, 197, 198, 7, 95, 2, 2, 198, 
	48, 3, 2, 2, 2, 199, 200, 7, 42, 2, 2, 200, 50, 3, 2, 2, 2, 201, 202, 7, 
	43, 2, 2, 202, 52, 3, 2, 2, 2, 203, 204, 7, 63, 2, 2, 204, 54, 3, 2, 2, 
	2, 205, 207, 9, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 
	206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 56, 3, 2, 2, 2, 210, 213, 5, 
	59, 30, 2, 211, 213, 5, 61, 31, 2, 212, 210, 3, 2, 2, 2, 212, 211, 3, 2, 
	2, 2, 213, 58, 3, 2, 2, 2, 214, 218, 7, 36, 2, 2, 215, 217, 11, 2, 2, 2, 
	216, 215, 3, 2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 218, 
	216, 3, 2, 2, 2, 219, 221, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 222, 
	7, 36, 2, 2, 222, 60, 3, 2, 2, 2, 223, 227, 7, 98, 2, 2, 224, 226, 11, 
	2, 2, 2, 225, 224, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 228, 3, 2, 2, 
	2, 227, 225, 3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 
	231, 7, 98, 2, 2, 231, 62, 3, 2, 2, 2, 232, 236, 9, 3, 2, 2, 233, 235, 
	9, 4, 2, 2, 234, 233, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 
	2, 2, 236, 237, 3, 2, 2, 2, 237, 64, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 
	239, 240, 7, 49, 2, 2, 240, 241, 7, 49, 2, 2, 241, 245, 3, 2, 2, 2, 242, 
	244, 10, 5, 2, 2, 243, 242, 3, 2, 2, 2, 244, 247, 3, 2, 2, 2, 245, 243, 
	3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 248, 3, 2, 2, 2, 247, 245, 3, 2, 
	2, 2, 248, 249, 8, 33, 2, 2, 249, 250, 3, 2, 2, 2, 250, 251, 8, 33, 3, 
	2, 251, 66, 3, 2, 2, 2, 252, 253, 7, 49, 2, 2, 253, 254, 7, 44, 2, 2, 254, 
	258, 3, 2, 2, 2, 255, 257, 11, 2, 2, 2, 256, 255, 3, 2, 2, 2, 257, 260, 
	3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 259, 261, 3, 2, 
	2, 2, 260, 258, 3, 2, 2, 2, 261, 262, 7, 44, 2, 2, 262, 263, 7, 49, 2, 
	2, 263, 264, 3, 2, 2, 2, 264, 265, 8, 34, 3, 2, 265, 68, 3, 2, 2, 2, 266, 
	268, 9, 6, 2, 2, 267, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 267, 
	3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 272, 8, 35, 
	3, 2, 272, 70, 3, 2, 2, 2, 11, 2, 208, 212, 218, 227, 236, 245, 258, 269, 
	4, 3, 33, 2, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'attribute'", "'class'", "'code'", "'enum'", "'error'", "'false'", 
	"'in'", "'link'", "'locator'", "'method'", "'out'", "'parameter'", "'resource'", 
	"'struct'", "'target'", "'true'", "'value'", "'variable'", "'@'", "'{'", 
	"'}'", "'['", "']'", "'('", "')'", "'='",
}

var lexerSymbolicNames = []string{
	"", "ATTRIBUTE", "CLASS", "CODE", "ENUM", "ERROR", "FALSE", "IN", "LINK", 
	"LOCATOR", "METHOD", "OUT", "PARAMETER", "RESOURCE", "STRUCT", "TARGET", 
	"TRUE", "VALUE", "VARIABLE", "AT_SIGN", "LEFT_CURLY_BRACKET", "RIGHT_CURLY_BRACKET", 
	"LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET", "LEFT_PARENTHESIS", "RIGHT_PARENTHESIS", 
	"EQUALS_SIGN", "INTEGER_LITERAL", "STRING_LITERAL", "IDENTIFIER", "LINE_COMMENT", 
	"BLOCK_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"ATTRIBUTE", "CLASS", "CODE", "ENUM", "ERROR", "FALSE", "IN", "LINK", "LOCATOR", 
	"METHOD", "OUT", "PARAMETER", "RESOURCE", "STRUCT", "TARGET", "TRUE", "VALUE", 
	"VARIABLE", "AT_SIGN", "LEFT_CURLY_BRACKET", "RIGHT_CURLY_BRACKET", "LEFT_SQUARE_BRACKET", 
	"RIGHT_SQUARE_BRACKET", "LEFT_PARENTHESIS", "RIGHT_PARENTHESIS", "EQUALS_SIGN", 
	"INTEGER_LITERAL", "STRING_LITERAL", "SHORT_STRING", "LONG_STRING", "IDENTIFIER", 
	"LINE_COMMENT", "BLOCK_COMMENT", "WS",
}
type ModelLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewModelLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ModelLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewModelLexer(input antlr.CharStream) *ModelLexer {
	l := new(ModelLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ModelLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ModelLexer tokens.
const (
	ModelLexerATTRIBUTE = 1
	ModelLexerCLASS = 2
	ModelLexerCODE = 3
	ModelLexerENUM = 4
	ModelLexerERROR = 5
	ModelLexerFALSE = 6
	ModelLexerIN = 7
	ModelLexerLINK = 8
	ModelLexerLOCATOR = 9
	ModelLexerMETHOD = 10
	ModelLexerOUT = 11
	ModelLexerPARAMETER = 12
	ModelLexerRESOURCE = 13
	ModelLexerSTRUCT = 14
	ModelLexerTARGET = 15
	ModelLexerTRUE = 16
	ModelLexerVALUE = 17
	ModelLexerVARIABLE = 18
	ModelLexerAT_SIGN = 19
	ModelLexerLEFT_CURLY_BRACKET = 20
	ModelLexerRIGHT_CURLY_BRACKET = 21
	ModelLexerLEFT_SQUARE_BRACKET = 22
	ModelLexerRIGHT_SQUARE_BRACKET = 23
	ModelLexerLEFT_PARENTHESIS = 24
	ModelLexerRIGHT_PARENTHESIS = 25
	ModelLexerEQUALS_SIGN = 26
	ModelLexerINTEGER_LITERAL = 27
	ModelLexerSTRING_LITERAL = 28
	ModelLexerIDENTIFIER = 29
	ModelLexerLINE_COMMENT = 30
	ModelLexerBLOCK_COMMENT = 31
	ModelLexerWS = 32
)


  func (l *ModelLexer) comment() {
    comments[l.GetLine()] = l.GetText()
  }



func (l *ModelLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 31:
			l.LINE_COMMENT_Action(localctx, actionIndex)


	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *ModelLexer) LINE_COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	this := l
	_ = this

	switch actionIndex {
	case 0:
			 l.comment() 


	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}

