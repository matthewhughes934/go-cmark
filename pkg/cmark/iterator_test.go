package cmark

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterAPI(t *testing.T) {
	root := ParseDocument("# heading\n\nparagraph\n", NewParserOpts())
	iter := NewIter(root)

	// step into document
	require.Equal(t, EventTypeEnter, iter.Next())
	require.Equal(t, EventTypeEnter, iter.GetEventType())
	rootNode := iter.GetNode()
	require.Equal(t, NodeTypeDocument, rootNode.GetType())

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

	require.Equal(t, iter.GetRoot(), rootNode)
}

func ExampleIter() {
	root := ParseDocument("# Document\n\nsome text\n", NewParserOpts())
	iter := NewIter(root)

	evType := iter.Next()
	for ; evType != EventTypeDone; evType = iter.Next() {
		cur := iter.GetNode()
		fmt.Print(cur.GetTypeString())
		switch evType {
		case EventTypeNone:
			fmt.Println(" NONE")
		case EventTypeDone:
			fmt.Println(" DONE")
		case EventTypeEnter:
			fmt.Println(" ENTER")
		case EventTypeExit:
			fmt.Println(" EXIT")
		}
	}

	// Output:
	// document ENTER
	// heading ENTER
	// text ENTER
	// heading EXIT
	// paragraph ENTER
	// text ENTER
	// paragraph EXIT
	// document EXIT
}
