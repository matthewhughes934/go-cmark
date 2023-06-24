/*
cmark-ast prints the nodes making up a markdown file

Usage:

	cmark-ast [ markdown-file ]

If no file is provided, reads from "/dev/stdin"
*/
package main

import (
	"fmt"
	"os"

	"github.com/matthewhughes934/go-cmark/pkg/cmark"
)

func main() { //go-cov:skip
	if err := cmarkAST(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cmarkAST(args []string) error {
	var filename string
	if len(args) != 2 {
		filename = "/dev/stdin"
	} else {
		filename = args[1]
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error reading file %s: %v", args[1], err)
	}

	dumpTree(cmark.ParseDocument(string(content), cmark.OptDefault), 0)
	return nil
}

func dumpTree(node *cmark.Node, indentLevel int) {
	var indentStr string
	for i := 0; i < indentLevel; i++ {
		indentStr += "\t"
	}
	fmt.Print(indentStr)

	if nodeContent := node.GetLiteral(); nodeContent == nil {
		fmt.Printf("(%s)", node.GetTypeString())
	} else {
		fmt.Printf("(%s: '%s')", node.GetTypeString(), *nodeContent)
	}

	fmt.Println()
	if node.FirstChild() != nil {
		dumpTree(node.FirstChild(), indentLevel+1)
	}
	if node.Next() != nil {
		dumpTree(node.Next(), indentLevel)
	}
}
