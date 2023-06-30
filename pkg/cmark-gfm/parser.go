package gfm

/*
#include "cmark-gfm.h"
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"
)

type Parser struct {
	parser *C.cmark_parser
}

// free wraps cmark_parser_free
func (parser *Parser) free() { //go-cov:skip
	C.cmark_parser_free(parser.parser)
}

// NewParser wraps cmark_parser_new
// Creates a new parser object.
func NewParser(options ParserOpt) *Parser {
	parser := &Parser{parser: C.cmark_parser_new(C.int(options))}
	runtime.SetFinalizer(parser, (*Parser).free)

	return parser
}

// Feed wraps cmark_parser_feed
// Feeds a string from 'text' to 'parser'.
func (parser *Parser) Feed(text string) {
	Cstr := C.CString(text)
	defer C.free(unsafe.Pointer(Cstr))

	C.cmark_parser_feed(parser.parser, Cstr, C.size_t(len(text)))
}

// Finish wraps cmark_parser_finish
// Finish parsing and return a pointer to a tree of nodes.
func (parser *Parser) Finish() *Node {
	return &Node{
		node: C.cmark_parser_finish(parser.parser),
	}
}
