package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type NodeType string

type Node struct {
	Children []*Node
	Name     string
	Type     NodeType
}

const (
	File      NodeType = "file"
	Directory NodeType = "directory"
)

type Counter struct {
	Files       int
	Directories int
}

func main() {

}

func summary(numDirectories int, numFiles int) string {
	directories := ""
	files := fmt.Sprintf("%d files", numFiles)

	if numFiles == 0 {
		numDirectories = 0
	}

	if numDirectories == 1 {
		directories = fmt.Sprintf("%d directory", numDirectories)
	} else {
		directories = fmt.Sprintf("%d directories", numDirectories)
	}
	return fmt.Sprintf("%s, %s", directories, files)
}

func indent(s string, depth int) string {
	var line strings.Builder
	for i := 0; i < depth; i++ {
		line.WriteString("│   ")
	}
	line.WriteString(s)
	return line.String()
}

func tree(node Node, depth int, counter *Counter) string {
	numFiles := len(node.Children)

	slices.SortFunc(node.Children, func(a, b *Node) int {
		return cmp.Compare(a.Type, b.Type)
	})

	var list strings.Builder
	for i := 0; i < numFiles; i++ {
		if i == numFiles-1 {
			list.WriteString(indent("└── ", depth))
		} else {
			list.WriteString(indent("├── ", depth))
		}
		list.WriteString(node.Children[i].Name)
		list.WriteString("\n")
		if node.Children[i].Type == Directory {
			counter.Directories++
			list.WriteString(tree(*node.Children[i], depth+1, counter))
		} else {
			counter.Files++
		}
	}
	if depth == 0 {
		list.WriteString("\n")
		list.WriteString(summary(counter.Directories, counter.Files))
		list.WriteString("\n")
	}
	return list.String()
}

func render(node Node) string {
	counter := &Counter{Files: 0, Directories: 1}

	return fmt.Sprintf(
		"%s\n%s",
		node.Name,
		tree(node, 0, counter),
	)
}
