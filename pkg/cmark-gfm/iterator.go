package gfm

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
