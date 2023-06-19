/*
cmark-ast prints the nodes making up a markdown file

Usage:

	cmark-ast <markdown-file>
*/
package main

import (
	"fmt"
	"os"

	"github.com/matthewhughes934/go-cmark/go-cmark"
)

func main() {
	if err := cmarkAST(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cmarkAST(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: %s <markdown-file>", args[0])
	}

	content, err := os.ReadFile(args[1])
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
