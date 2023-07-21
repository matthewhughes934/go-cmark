package gfm

/*
#include <stdlib.h>
#include "cmark-gfm.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// ParserOpts provides options for configuring parsing behaviour
type ParserOpts struct {
	o C.int
}

func NewParserOpts() *ParserOpts {
	return &ParserOpts{o: C.CMARK_OPT_DEFAULT}
}

// WithValidateUTF8 maps to C.CMARK_OPT_VALIDATE_UTF8.
// Validates UTF-8 in the input before parsing, replacing illegal
// sequences with the replacement character U+FFFD.
func (o *ParserOpts) WithValidateUTF8() *ParserOpts {
	o.o |= C.CMARK_OPT_VALIDATE_UTF8
	return o
}

// WithSmart maps to C.CMARK_OPT_SMART.
// Converts straight quotes to curly, --- to em dashes, -- to en dashes.
func (o *ParserOpts) WithSmart() *ParserOpts {
	o.o |= C.CMARK_OPT_SMART
	return o
}

// WithFoonotes maps to c.CMARK_OPT_FOOTNOTES
// Parses footnotes
func (o *ParserOpts) WithFoonotes() *ParserOpts {
	o.o |= C.CMARK_OPT_FOOTNOTES
	return o
}

type Parser struct {
	parser *C.cmark_parser
}

// free wraps cmark_parser_free
func (parser *Parser) free() { //go-cov:skip
	C.cmark_parser_free(parser.parser)
}

// NewParser wraps cmark_parser_new.
// Creates a new parser object.
func NewParser(opts *ParserOpts) *Parser {
	parser := &Parser{parser: C.cmark_parser_new(opts.o)}
	runtime.SetFinalizer(parser, (*Parser).free)

	return parser
}

// Feed wraps cmark_parser_feed.
// Feeds a string from 'text' to 'parser'.
func (parser *Parser) Feed(text string) {
	Cstr := C.CString(text)
	defer C.free(unsafe.Pointer(Cstr))

	C.cmark_parser_feed(parser.parser, Cstr, C.size_t(len(text)))
}

// Finish wraps cmark_parser_finish.
// Finish parsing and return a pointer to a tree of nodes.
func (parser *Parser) Finish() *Node {
	return &Node{
		node: C.cmark_parser_finish(parser.parser),
	}
}

// ParseDocument parses a CommonMark document in 'document' and returns a pointer to a tree of nodes.
// The returned [cmark.Node] has a finalizer set that will call
// `cmark_node_free` which will free the memory allocated for the node and any
// of its children
func (parser *Parser) ParseDocument(document string) *Node {
	parser.Feed(document)
	node := parser.Finish()
	runtime.SetFinalizer(node, (*Node).free)

	return node
}
