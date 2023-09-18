package gfm

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup/teardown of extensions for entire test run
	CoreExtensionsEnsureRegistered()
	defer ReleasePlugins()

	os.Exit(m.Run())
}

func Example() {
	document := "# My great document\n\nWhat a great read!\n"
	root := NewParser().ParseDocument(document)

	heading := root.FirstChild()
	headingContent := heading.FirstChild()

	paragraph := heading.Next()
	paragraphContent := paragraph.FirstChild()

	fmt.Println(root.GetTypeString())
	fmt.Println(heading.GetTypeString())
	fmt.Println(headingContent.GetType() == NodeTypeText)
	fmt.Println(headingContent.GetLiteral())
	fmt.Println(paragraph.GetTypeString())
	fmt.Println(paragraphContent.GetType() == NodeTypeText)
	fmt.Println(paragraphContent.GetLiteral())

	// Output:
	// document
	// heading
	// true
	// My great document
	// paragraph
	// true
	// What a great read!
}

func Example_extensions() {
	// commented out since these are handled separately in the test suite
	// CoreExtensionsEnsureRegistered()
	// defer ReleasePlugins()
	document := "# My document\nWith ~~no~~ an extension\n"
	parser := NewParser().WithSyntaxExtension(FindSyntaxExtension(Strikethrough))

	root := parser.ParseDocument(document)

	fmt.Print(RenderHTML(root, NewRenderOpts(), parser.SyntaxExtensions()))

	// Output:
	// <h1>My document</h1>
	// <p>With <del>no</del> an extension</p>
}
