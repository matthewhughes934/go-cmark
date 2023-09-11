package cmark

import (
	"fmt"
)

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
