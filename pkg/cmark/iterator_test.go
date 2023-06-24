package cmark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterAPI(t *testing.T) {
	root := ParseDocument("# heading\n\nparagraph\n", OptDefault)
	iter := NewIter(root)

	// step into document
	require.Equal(t, EventTypeEnter, iter.Next())
	require.Equal(t, EventTypeEnter, iter.GetEventType())
	require.Equal(t, NodeTypeDocument, iter.GetNode().GetType())

	// step into heading
	require.Equal(t, EventTypeEnter, iter.Next())
	require.Equal(t, NodeTypeHeading, iter.GetNode().GetType())

	// step into heading text
	require.Equal(t, EventTypeEnter, iter.Next())
	require.Equal(t, NodeTypeText, iter.GetNode().GetType())

	// step out to heading
	require.Equal(t, EventTypeExit, iter.Next())
	require.Equal(t, NodeTypeHeading, iter.GetNode().GetType())

	// step into paragraph
	require.Equal(t, EventTypeEnter, iter.Next())
	require.Equal(t, NodeTypeParagraph, iter.GetNode().GetType())

	// step into paragraph text
	require.Equal(t, EventTypeEnter, iter.Next())
	require.Equal(t, NodeTypeText, iter.GetNode().GetType())

	// step out to paragraph
	require.Equal(t, EventTypeExit, iter.Next())
	require.Equal(t, NodeTypeParagraph, iter.GetNode().GetType())

	// step out to document
	require.Equal(t, EventTypeExit, iter.Next())
	require.Equal(t, NodeTypeDocument, iter.GetNode().GetType())

	// done
	require.Equal(t, EventTypeDone, iter.Next())
}
