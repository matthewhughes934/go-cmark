package gfm

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserAPI(t *testing.T) {
	parser := NewParser(NewParserOpts())
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

func TestParserOpts(t *testing.T) {
	for _, tc := range []struct {
		content  string
		opts     *ParserOpts
		expected string
	}{
		{
			"plain paragraph",
			NewParserOpts(),
			"plain paragraph",
		},
		{
			"bad UTF8: \xFF",
			NewParserOpts().WithValidateUTF8(),
			"bad UTF8: �",
		},
		{
			`text -- "with quotes" ---`,
			NewParserOpts().WithSmart(),
			"text – “with quotes” —",
		},
	} {
		t.Run(tc.content, func(t *testing.T) {
			parser := NewParser(tc.opts)
			document := parser.ParseDocument(tc.content)
			parsedContent := document.FirstChild().FirstChild().GetLiteral()
			t.Log(RenderHTML(document, NewRenderOpts(), nil))

			require.NotNil(t, parsedContent)
			require.Equal(t, tc.expected, *parsedContent)
		})
	}
}
