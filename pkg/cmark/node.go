package cmark

/*
#include <stdlib.h>
#include "cmark.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// NodeType is cmark_node_type
type NodeType int

const (
	/* Error status */
	NodeTypeNone NodeType = C.CMARK_NODE_NONE

	/* Block */
	NodeTypeDocument      NodeType = C.CMARK_NODE_DOCUMENT
	NodeTypeBlockQuote    NodeType = C.CMARK_NODE_BLOCK_QUOTE
	NodeTypeList          NodeType = C.CMARK_NODE_LIST
	NodeTypeItem          NodeType = C.CMARK_NODE_ITEM
	NodeTypeCodeBlock     NodeType = C.CMARK_NODE_CODE_BLOCK
	NodeTypeHTMLBlock     NodeType = C.CMARK_NODE_HTML_BLOCK
	NodeTypeCustomBlock   NodeType = C.CMARK_NODE_CUSTOM_BLOCK
	NodeTypeParagraph     NodeType = C.CMARK_NODE_PARAGRAPH
	NodeTypeHeading       NodeType = C.CMARK_NODE_HEADING
	NodeTypeThematicBreak NodeType = C.CMARK_NODE_THEMATIC_BREAK

	NodeTypeFirstBlock NodeType = C.CMARK_NODE_DOCUMENT
	NodeTypeLastBlock  NodeType = C.CMARK_NODE_THEMATIC_BREAK

	/* Inline */
	NodeTypeText         NodeType = C.CMARK_NODE_TEXT
	NodeTypeSoftbreak    NodeType = C.CMARK_NODE_SOFTBREAK
	NodeTypeLinebreak    NodeType = C.CMARK_NODE_LINEBREAK
	NodeTypeCode         NodeType = C.CMARK_NODE_CODE
	NodeTypeHTMLInline   NodeType = C.CMARK_NODE_HTML_INLINE
	NodeTypeCustomInline NodeType = C.CMARK_NODE_CUSTOM_INLINE
	NodeTypeEmph         NodeType = C.CMARK_NODE_EMPH
	NodeTypeStrong       NodeType = C.CMARK_NODE_STRONG
	NodeTypeLink         NodeType = C.CMARK_NODE_LINK
	NodeTypeImage        NodeType = C.CMARK_NODE_IMAGE

	NodeTypeFirstInline NodeType = C.CMARK_NODE_TEXT
	NodeTypeLastInline  NodeType = C.CMARK_NODE_IMAGE
)

// ListType is cmark_list_type
type ListType int

const (
	TypeNoList      ListType = C.CMARK_NO_LIST
	TypeBulletList  ListType = C.CMARK_BULLET_LIST
	TypeOrderedList ListType = C.CMARK_ORDERED_LIST
)

// DelimType is cmark_delim_type
type DelimType int

const (
	TypeNoDelim     DelimType = C.CMARK_NO_DELIM
	TypePeriodDelim DelimType = C.CMARK_PERIOD_DELIM
	TypeParentDelim DelimType = C.CMARK_PAREN_DELIM
)

type ParserOpt int

const (
	// Default options.
	ParserOptDefault ParserOpt = C.CMARK_OPT_DEFAULT

	// Include a `data-sourcepos` attribute on all block elements.
	ParserOptSourcePos ParserOpt = C.CMARK_OPT_SOURCEPOS

	// Render `softbreak` elements as hard line breaks.
	ParserOptHardBreaks ParserOpt = C.CMARK_OPT_HARDBREAKS

	//  Render raw HTML and unsafe links (`javascript:`, `vbscript:`,
	// `file:`, and `data:`, except for `image/png`, `image/gif`,
	// `image/jpeg`, or `image/webp` mime types).  By default,
	// raw HTML is replaced by a placeholder HTML comment. Unsafe
	// links are replaced by empty strings.
	ParserOptUnsafe ParserOpt = C.CMARK_OPT_UNSAFE

	// Render `softbreak` elements as spaces.
	ParserOptNoBreaks ParserOpt = C.CMARK_OPT_NOBREAKS

	// Legacy option (no effect).
	ParserOptNormalize ParserOpt = C.CMARK_OPT_NORMALIZE

	//  Validate UTF-8 in the input before parsing, replacing illegal
	// sequences with the replacement character U+FFFD.
	ParserOptValidateUTF8 ParserOpt = C.CMARK_OPT_VALIDATE_UTF8

	// Convert straight quotes to curly, --- to em dashes, -- to en dashes.
	ParserOptSmart ParserOpt = C.CMARK_OPT_SMART
)

type Node struct {
	node *C.struct_cmark_node
}

// NewNode wraps cmark_node_new
// Creates a new node of type 'type'
func NewNode(nt NodeType) *Node {
	node := &Node{C.cmark_node_new(C.cmark_node_type(nt))}
	runtime.SetFinalizer(node, (*Node).free)
	return node
}

// Free wraps cmark_node_free
// Frees the memory allocated for a node and any children.
func (node *Node) free() { //go-cov:skip
	C.cmark_node_free(node.node)
}

// GetType returns the type of 'node' or [NodeTypeNone] on failure
func (node *Node) GetType() NodeType {
	return NodeType(C.cmark_node_get_type(node.node))
}

// GetTypeString is like GetType but returns a string representation of the
// type, or `"<unknown>"`
func (node *Node) GetTypeString() string {
	return C.GoString(C.cmark_node_get_type_string(node.node))
}

// GetLiteral wraps cmark_node_get_literal
// Returns a pointer to the string contents of 'node', or an empty
// string if none is set. Returns nil if called on a
// node that does not have string content.
func (node *Node) GetLiteral() *string {
	literal := C.cmark_node_get_literal(node.node)
	if literal == nil {
		return nil
	}

	str := C.GoString(literal)
	return &str
}

// GetHeadingLevel wraps cmark_node_get_heading_level
// Returns the heading level of 'node', or 0 if 'node' is not a heading.
func (node *Node) GetHeadingLevel() int {
	return int(C.cmark_node_get_heading_level(node.node))
}

// GetListType wraps cmark_node_get_list_type
// Returns the list delimiter type of 'node', or [TypeNoList] if 'node'
// is not a list
func (node *Node) GetListType() ListType {
	return ListType(C.cmark_node_get_list_type(node.node))
}

// GetListStart wraps cmark_node_get_list_start
// Returns starting number of 'node', if it is an ordered list, otherwise 0.
func (node *Node) GetListStart() int {
	return int(C.cmark_node_get_list_start(node.node))
}

// IsTightList wraps cmark_node_get_list_tight
// Returns whether 'node' is a tight list
func (node *Node) IsTightList() bool {
	return C.cmark_node_get_list_tight(node.node) == C.int(1)
}

// GetUrl wraps cmark_node_get_url
// Returns the URL of a link or image 'node', or an empty string
// if no URL is set.  Returns NULL if called on a node that is
// not a link or image.
func (node *Node) GetUrl() *string {
	literal := C.cmark_node_get_url(node.node)
	if literal == nil {
		return nil
	}

	str := C.GoString(literal)
	return &str
}

// GetTitle wraps cmark_node_get_title
// returns the title of a link or image 'node', or an empty string
// if no URL is set. Returns nil if called on a node that is not a
// link or image
func (node *Node) GetTitle() *string {
	literal := C.cmark_node_get_title(node.node)
	if literal == nil {
		return nil
	}

	str := C.GoString(literal)
	return &str
}

// GetStartLine wraps cmark_node_get_start_line
// Returns the line at which 'node' begins.
func (node *Node) GetStartLine() int {
	return int(C.cmark_node_get_start_line(node.node))
}

// GetEndLine wraps cmark_node_get_end_line
// Returns the line at which 'node' begins.
func (node *Node) GetEndLine() int {
	return int(C.cmark_node_get_end_line(node.node))
}

// GetStartColumn wraps cmark_node_get_start_column
// Returns the column at which 'node' begins.
func (node *Node) GetStartColumn() int {
	return int(C.cmark_node_get_start_column(node.node))
}

// GetEndColumn wraps cmark_node_get_end_column
// Returns the column at which 'node' begins.
func (node *Node) GetEndColumn() int {
	return int(C.cmark_node_get_end_column(node.node))
}

// Next wraps cmark_node_next
// Returns the next node in the sequence after 'node', or nil if
// there is none
func (node *Node) Next() *Node {
	next := C.cmark_node_next(node.node)
	if next == nil {
		return nil
	}
	return &Node{node: next}
}

// Previous wraps cmark_node_previous
// Returns the previous node in the sequence after 'node', or none if
// there is none
func (node *Node) Previous() *Node {
	previous := C.cmark_node_previous(node.node)
	if previous == nil {
		return nil
	}
	return &Node{node: previous}
}

// Parent wraps cmark_node_parent
// Returns the parent of 'node', or nil if there is none.
func (node *Node) Parent() *Node {
	parent := C.cmark_node_parent(node.node)
	if parent == nil {
		return nil
	}
	return &Node{node: parent}
}

// FirstChild wraps cmark_node_first_child
// Returns the first child of 'node', or nil if 'node' has no children.
func (node *Node) FirstChild() *Node {
	firstChild := C.cmark_node_first_child(node.node)
	if firstChild == nil {
		return nil
	}
	return &Node{node: firstChild}
}

// LastChild wraps cmark_node_last_child
// Returns the last child of 'node', or nil if 'node' has no children.
func (node *Node) LastChild() *Node {
	lastChild := C.cmark_node_last_child(node.node)
	if lastChild == nil {
		return nil
	}

	return &Node{node: lastChild}
}

// ParseDocument wraps cmark_parse_document
// Parse a CommonMark document in 'document'
// Returns a pointer to a tree of nodes.
func ParseDocument(document string, options ParserOpt) *Node {
	str := C.CString(document)
	defer C.free(unsafe.Pointer(str))

	node := &Node{
		node: C.cmark_parse_document(str, C.size_t(len(document)), C.int(options)),
	}
	runtime.SetFinalizer(node, (*Node).free)

	return node
}
