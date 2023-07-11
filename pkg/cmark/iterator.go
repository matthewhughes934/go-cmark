package cmark

/*
#include "cmark.h"
*/
import "C"

import (
	"runtime"
)

type EventType int

const (
	EventTypeNone  EventType = C.CMARK_EVENT_NONE
	EventTypeDone  EventType = C.CMARK_EVENT_DONE
	EventTypeEnter EventType = C.CMARK_EVENT_ENTER
	EventTypeExit  EventType = C.CMARK_EVENT_EXIT
)

/*
Iter wraps cmark_iter

An iterator will walk through a tree of nodes, starting from a root node,
returning one node at a time, together with information about whether the node
is being entered or exited. The iterator will first descend to a child node, if
there is one. When there is no child, the iterator will go to the next sibling.
When there is no next sibling, the iterator will return to the parent (but with
a [EventType] of [EventTypeExit]). The iterator will return [EventTypeDone]
when it reaches the root node again.  One natural application is an HTML
renderer, where an [EventTypeEnter] event outputs an open tag and an
[EventTypeExit] event outputs a close tag. An iterator might also be used to
transform an AST in some systematic way, for example, turning all level-3
headings into regular paragraphs.

Iterators will never return [EventTypeExit] events for leaf noes, which are
nodes of type:

	* [NodeTypeHTMLBlock]
	* [NodeTypeThematicBreak]
	* [NodeTypeCodeBlock]
	* [NodeTypeText]
	* [NodeTypeSoftbreak]
	* [NodeTypeLinebreak]
	* [NodeTypeCode]
	* [NodeTypeHTMLInline]

Nodes must only be modified after an [EvenTypeExit] event or an
[EventTypeEnter] for leaf nodes
*/
type Iter struct {
	iter *C.cmark_iter
}

// NewIter wraps cmark_iter_new
// Creates a new iterator starting at 'root'.  The current node and event
// type are undefined until 'cmark_iter_next' is called for the first time.
func NewIter(root *Node) *Iter {
	iter := &Iter{iter: C.cmark_iter_new(root.node)}
	runtime.SetFinalizer(iter, (*Iter).free)
	return iter
}

// free wraps cmark_iter_free
func (iter *Iter) free() { //go-cov:skip
	C.cmark_iter_free(iter.iter)
}

// Next wraps cmark_iter_next
// Advances to the next node and returns the event type
func (iter *Iter) Next() EventType {
	return EventType(C.cmark_iter_next(iter.iter))
}

// GetNode wraps cmark_iter_get_node
// Returns the current node.
func (iter *Iter) GetNode() *Node {
	return &Node{node: C.cmark_iter_get_node(iter.iter)}
}

// GetEventType wraps cmark_iter_get_event_type
// Returns the current event type.
func (iter *Iter) GetEventType() EventType {
	return EventType(C.cmark_iter_get_event_type(iter.iter))
}

// GetRoot  wraps cmark_iter_get_root
// Returns the root node.
func (iter *Iter) GetRoot() *Node {
	return &Node{node: C.cmark_iter_get_root(iter.iter)}
}
