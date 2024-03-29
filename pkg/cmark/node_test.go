package cmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSimpleDocument(t *testing.T) {
	document := "# My Document\n\nthis is a great document!\n"

	got := NewParser().ParseDocument(document)
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

	assert.Equal(t, "", node.GetLiteral())
}

func TestGetLiteralWithContent(t *testing.T) {
	content := "# heading\n"
	document := NewParser().ParseDocument(content)
	require.NotNil(t, document)
	heading := document.FirstChild()
	require.NotNil(t, heading)

	headingText := heading.FirstChild()
	require.NotNil(t, headingText)
	assert.Equal(t, "heading", headingText.GetLiteral())
}

func TestNodeFirstChildNoChild(t *testing.T) {
	got := NewParser().ParseDocument("")
	require.NotNil(t, got)

	assert.Nil(t, got.FirstChild())
}

func TestNodeFirstChild(t *testing.T) {
	got := NewParser().ParseDocument("# heading\n")
	require.NotNil(t, got)

	firstChild := got.FirstChild()
	assert.NotNil(t, firstChild)
	assert.Equal(t, NodeTypeHeading, firstChild.GetType())
}

func TestLastChildNoChild(t *testing.T) {
	document := NewParser().ParseDocument("")

	assert.Nil(t, document.LastChild())
}

func TestLastChild(t *testing.T) {
	document := NewParser().ParseDocument("# heading\n\nparagraph\n")

	lastChild := document.LastChild()
	require.NotNil(t, lastChild)
	require.Equal(t, NodeTypeParagraph, lastChild.GetType())
}

func TestNodeNextNoNext(t *testing.T) {
	document := NewParser().ParseDocument("")
	require.NotNil(t, document)

	assert.Nil(t, document.Next())
}

func TestNodeNext(t *testing.T) {
	document := NewParser().ParseDocument("first paragraph\n\nsecond paragraph\n")

	firstParagraph := document.FirstChild()
	require.Equal(t, NodeTypeParagraph, firstParagraph.GetType())

	secondParagraph := firstParagraph.Next()
	require.NotNil(t, secondParagraph)
	require.Equal(t, NodeTypeParagraph, secondParagraph.GetType())
	assert.NotSame(t, firstParagraph, secondParagraph)
}

func TestNodePreviousNoPrevious(t *testing.T) {
	document := NewParser().ParseDocument("")
	require.NotNil(t, document)

	assert.Nil(t, document.Previous())
}

func TestNodePrevious(t *testing.T) {
	document := NewParser().ParseDocument("first paragraph\n\nsecond paragraph\n")

	secondParagraph := document.FirstChild().Next()
	require.Equal(t, NodeTypeParagraph, secondParagraph.GetType())

	firstParagraph := secondParagraph.Previous()
	require.NotNil(t, firstParagraph)
	require.Equal(t, NodeTypeParagraph, firstParagraph.GetType())
	assert.NotSame(t, firstParagraph, secondParagraph)
}

func TestNodeParentNoParent(t *testing.T) {
	document := NewParser().ParseDocument("")
	require.NotNil(t, document)

	assert.Nil(t, document.Parent())
}

func TestNodeParent(t *testing.T) {
	document := NewParser().ParseDocument("# heading\n")

	heading := document.FirstChild()
	require.NotNil(t, heading)

	parent := heading.Parent()
	require.NotNil(t, parent)
	require.Equal(t, NodeTypeDocument, parent.GetType())
}

func TestGetHeadingLevelNotHeading(t *testing.T) {
	document := NewParser().ParseDocument("")

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
			document := NewParser().ParseDocument(tc.heading + "\n")

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
			document := NewParser().ParseDocument(tc.content + "\n")

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
			document := NewParser().ParseDocument(tc.content + "\n")

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
			document := NewParser().ParseDocument(tc.content + "\n")

			list := document.FirstChild()
			require.NotNil(t, list)

			require.Equal(t, tc.expected, list.IsTightList())
		})
	}
}

func TestGetFenceInfoNoInfo(t *testing.T) {
	document := NewParser().ParseDocument("No code fence\n")

	content := document.FirstChild()
	require.NotNil(t, content)

	require.Equal(t, "", content.GetFenceInfo())
}

func TestGetFenceInfo(t *testing.T) {
	for _, tc := range []struct {
		content  string
		expected string
	}{
		{"```bash\necho 'hello'\n```\n", "bash"},
		{"~~~python\nprint('hello')\n~~~\n", "python"},
	} {
		t.Run(tc.content, func(t *testing.T) {
			document := NewParser().ParseDocument(tc.content)
			fenceNode := document.FirstChild()

			require.Equal(t, fenceNode.GetFenceInfo(), tc.expected)
		})
	}
}

func TestGetUrlNoUrl(t *testing.T) {
	document := NewParser().ParseDocument("No URL here\n")

	content := document.FirstChild()
	require.NotNil(t, content)

	require.Equal(t, "", content.GetUrl())
}

func TestGetURL(t *testing.T) {
	for _, tc := range []struct {
		content  string
		expected string
	}{
		{"<https://example.com>", "https://example.com"},
		{"[link here](https://example.com)", "https://example.com"},
		{"[link here]()", ""},
		{`![alt-name](https://example.com/image.png "some-title")`, "https://example.com/image.png"},
	} {
		t.Run(tc.content, func(t *testing.T) {
			document := NewParser().ParseDocument(tc.content)

			linkNode := document.FirstChild().FirstChild()

			require.Equal(t, tc.expected, linkNode.GetUrl())
		})
	}
}

func TestGetTitleNoTitle(t *testing.T) {
	for _, content := range []string{
		"paragraph",
		"# heading",
		"* list point",
	} {
		t.Run(content, func(t *testing.T) {
			document := NewParser().ParseDocument(content)

			linkNode := document.FirstChild().FirstChild()

			require.Equal(t, "", linkNode.GetTitle())
		})
	}
}

func TestGetTitle(t *testing.T) {
	for _, tc := range []struct {
		content  string
		expected string
	}{
		{"<https://example.com>", ""},
		{`[link text](https://example.com "link here")`, "link here"},
		{`[](/link "hmm")`, "hmm"},
		{`![alt-name](https://example.com/image.png "some title")`, "some title"},
	} {
		t.Run(tc.content, func(t *testing.T) {
			document := NewParser().ParseDocument(tc.content)

			linkNode := document.FirstChild().FirstChild()

			require.Equal(t, tc.expected, linkNode.GetTitle())
		})
	}
}

func TestNodePositionFunctions(t *testing.T) {
	content := "# heading\nparagraph that\nlasts\nseveral\nlines\n"

	document := NewParser().ParseDocument(content)

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
