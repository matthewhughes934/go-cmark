package cmark

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserAPI(t *testing.T) {
	parser := NewParser(OptDefault)
	document := "# heading\n\nparagraph here\n"

	scanner := bufio.NewScanner(strings.NewReader(document))

	for scanner.Scan() {
		parser.Feed(scanner.Text() + "\n")
	}
	require.NoError(t, scanner.Err())

	root := parser.Finish()
	firstChild := root.FirstChild()
	require.NotNil(t, firstChild)
	nextNode := firstChild.Next()
	require.NotNil(t, nextNode)

	require.Equal(t, NodeTypeHeading, firstChild.GetType())
	require.Equal(t, NodeTypeParagraph, nextNode.GetType())
}
