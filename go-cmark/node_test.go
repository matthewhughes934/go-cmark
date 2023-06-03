package cmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSimpleDocument(t *testing.T) {
	document := "# My Document\n\nthis is a great document!\n"

	got := ParseDocument(document, OptDefault)
	require.NotNil(t, got)

	assert.Equal(t, NodeTypeDocument, got.GetType())
}

func TestNodeType(t *testing.T) {
	for _, tc := range []struct {
		nodeType     NodeType
		expectedName string
	}{
		{NodeTypeNone, "none"},
		{NodeTypeDocument, "document"},
		{NodeTypeBlockQuote, "block_quote"},
		{NodeTypeList, "list"},
		{NodeTypeItem, "item"},
		{NodeTypeCodeBlock, "code_block"},
		{NodeTypeHTMLBlock, "html_block"},
		{NodeTypeCustomBlock, "custom_block"},
		{NodeTypeParagraph, "paragraph"},
		{NodeTypeHeading, "heading"},
		{NodeTypeThematicBreak, "thematic_break"},
		{NodeTypeFirstBlock, "document"},
		{NodeTypeLastBlock, "thematic_break"},
		{NodeTypeText, "text"},
		{NodeTypeSoftbreak, "softbreak"},
		{NodeTypeLinebreak, "linebreak"},
		{NodeTypeCode, "code"},
		{NodeTypeHTMLInline, "html_inline"},
		{NodeTypeCustomInline, "custom_inline"},
		{NodeTypeEmph, "emph"},
		{NodeTypeStrong, "strong"},
		{NodeTypeLink, "link"},
		{NodeTypeImage, "image"},
		{NodeTypeFirstInline, "text"},
		{NodeTypeLastInline, "image"},
	} {
		t.Run(tc.expectedName, func(t *testing.T) {
			got := NewNode(tc.nodeType)

			require.NotNil(t, got)
			assert.Equal(t, tc.nodeType, got.GetType())
			assert.Equal(t, tc.expectedName, got.GetTypeString())
		})
	}
}

func TestGetLiteralNoContent(t *testing.T) {
	node := NewNode(NodeTypeNone)

	assert.Nil(t, node.GetLiteral())
}

func TestGetLiteralWithContent(t *testing.T) {
	content := "# heading\n"
	document := ParseDocument(content, OptDefault)
	require.NotNil(t, document)
	heading := document.FirstChild()
	require.NotNil(t, heading)

	headingText := heading.FirstChild()
	require.NotNil(t, headingText)
	assert.Equal(t, "heading", *headingText.GetLiteral())
}

func TestNodeFirstChildNoChild(t *testing.T) {
	got := ParseDocument("", OptDefault)
	require.NotNil(t, got)

	assert.Nil(t, got.FirstChild())
}

func TestNodeFirstChild(t *testing.T) {
	got := ParseDocument("# heading\n", OptDefault)
	require.NotNil(t, got)

	firstChild := got.FirstChild()
	assert.NotNil(t, firstChild)
	assert.Equal(t, NodeTypeHeading, firstChild.GetType())
}

func TestLastChildNoChild(t *testing.T) {
	document := ParseDocument("", OptDefault)

	assert.Nil(t, document.LastChild())
}

func TestLastChild(t *testing.T) {
	document := ParseDocument("# heading\n\nparagraph\n", OptDefault)

	lastChild := document.LastChild()
	require.NotNil(t, lastChild)
	require.Equal(t, NodeTypeParagraph, lastChild.GetType())
}

func TestNodeNextNoNext(t *testing.T) {
	document := ParseDocument("", OptDefault)
	require.NotNil(t, document)

	assert.Nil(t, document.Next())
}

func TestNodeNext(t *testing.T) {
	document := ParseDocument("first paragraph\n\nsecond paragraph\n", OptDefault)

	firstParagraph := document.FirstChild()
	require.Equal(t, NodeTypeParagraph, firstParagraph.GetType())

	secondParagraph := firstParagraph.Next()
	require.NotNil(t, secondParagraph)
	require.Equal(t, NodeTypeParagraph, secondParagraph.GetType())
	assert.NotSame(t, firstParagraph, secondParagraph)
}

func TestNodePreviousNoPrevious(t *testing.T) {
	document := ParseDocument("", OptDefault)
	require.NotNil(t, document)

	assert.Nil(t, document.Previous())
}

func TestNodePrevious(t *testing.T) {
	document := ParseDocument("first paragraph\n\nsecond paragraph\n", OptDefault)

	secondParagraph := document.FirstChild().Next()
	require.Equal(t, NodeTypeParagraph, secondParagraph.GetType())

	firstParagraph := secondParagraph.Previous()
	require.NotNil(t, firstParagraph)
	require.Equal(t, NodeTypeParagraph, firstParagraph.GetType())
	assert.NotSame(t, firstParagraph, secondParagraph)
}

func TestNodeParentNoParent(t *testing.T) {
	document := ParseDocument("", OptDefault)
	require.NotNil(t, document)

	assert.Nil(t, document.Parent())
}

func TestNodeParent(t *testing.T) {
	document := ParseDocument("# heading\n", OptDefault)

	heading := document.FirstChild()
	require.NotNil(t, heading)

	parent := heading.Parent()
	require.NotNil(t, parent)
	require.Equal(t, NodeTypeDocument, parent.GetType())
}

func TestGetHeadingLevelNotHeading(t *testing.T) {
	document := ParseDocument("", OptDefault)

	require.Equal(t, 0, document.GetHeadingLevel())
}

func TestGetHeadingLevel(t *testing.T) {
	for _, tc := range []struct {
		heading       string
		expectedLevel int
	}{
		{"# heading", 1},
		{"## heading", 2},
	} {
		t.Run(tc.heading, func(t *testing.T) {
			document := ParseDocument(tc.heading+"\n", OptDefault)

			heading := document.FirstChild()
			require.NotNil(t, heading)

			require.Equal(t, tc.expectedLevel, heading.GetHeadingLevel())
		})
	}
}

func TestGetListType(t *testing.T) {
	for _, tc := range []struct {
		content          string
		expectedListType ListType
	}{
		{"# Heading is not a list", TypeNoList},
		{"* foo\n* bar", TypeBulletList},
		{"1. foo\n2. bar", TypeOrderedList},
	} {
		t.Run(tc.content, func(t *testing.T) {
			document := ParseDocument(tc.content+"\n", OptDefault)

			list := document.FirstChild()
			require.NotNil(t, list)

			require.Equal(t, tc.expectedListType, list.GetListType())
		})
	}
}

func TestGetListStart(t *testing.T) {
	for _, tc := range []struct {
		content           string
		expectedListStart int
	}{
		{"# Heading is not a list", 0},
		{"* foo\n* bar", 0},
		{"1. foo\n2. bar", 1},
		{"2. foo\n3. bar", 2},
	} {
		t.Run(tc.content, func(t *testing.T) {
			document := ParseDocument(tc.content+"\n", OptDefault)

			list := document.FirstChild()
			require.NotNil(t, list)

			require.Equal(t, tc.expectedListStart, list.GetListStart())
		})
	}
}

func TestIsTightList(t *testing.T) {
	for _, tc := range []struct {
		content  string
		expected bool
	}{
		{"Not a list", false},
		{"* not a tight\n\n* list", false},
		{"* tight\n*list", true},
	} {
		t.Run(tc.content, func(t *testing.T) {
			document := ParseDocument(tc.content+"\n", OptDefault)

			list := document.FirstChild()
			require.NotNil(t, list)

			require.Equal(t, tc.expected, list.IsTightList())
		})
	}
}

func TestGetUrlNoUrl(t *testing.T) {
	document := ParseDocument("No URL here\n", OptDefault)

	content := document.FirstChild()
	require.NotNil(t, content)

	require.Nil(t, content.GetUrl())
}

func TestGetURL(t *testing.T) {
	for _, tc := range []struct {
		content  string
		expected string
	}{
		{"<https://example.com>", "https://example.com"},
		{"[link here](https://example.com)", "https://example.com"},
		{"[link here]()", ""},
	} {
		t.Run(tc.content, func(t *testing.T) {
			document := ParseDocument(tc.content, OptDefault)

			linkNode := document.FirstChild().FirstChild()

			require.Equal(t, tc.expected, *linkNode.GetUrl())
		})
	}
}

func TestNodePositionFunctions(t *testing.T) {
	content := "# heading\nparagraph that\nlasts\nseveral\nlines\n"

	document := ParseDocument(content, OptDefault)

	headingNode := document.FirstChild()
	require.NotNil(t, headingNode)
	paragraphNode := headingNode.Next()
	require.NotNil(t, paragraphNode)

	assert.Equal(t, 1, headingNode.GetStartLine())
	assert.Equal(t, 1, headingNode.GetEndLine())
	assert.Equal(t, 1, headingNode.GetStartColumn())
	assert.Equal(t, 9, headingNode.GetEndColumn())

	assert.Equal(t, 2, paragraphNode.GetStartLine())
	assert.Equal(t, 5, paragraphNode.GetEndLine())
	assert.Equal(t, 1, paragraphNode.GetStartColumn())
	assert.Equal(t, 5, paragraphNode.GetEndColumn())
}
