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

func main() {
	file1 := Node{Name: "file1", Type: File}
	file2 := Node{Name: "file2", Type: File}
	directory1 := Node{Name: "directory1", Type: Directory}
	node := Node{
		Name:     ".",
		Children: []*Node{&file1, &file2, &directory1},
		Type:     Directory,
	}
	fmt.Print(render(node))
}

func summary(numFiles int, numDirectories int) string {
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

func tree(node Node) string {
	numFiles := len(node.Children)

	slices.SortFunc(node.Children, func(a, b *Node) int {
		return cmp.Compare(a.Type, b.Type)
	})

	var list strings.Builder
	for i := 0; i < numFiles; i++ {
		if i == numFiles-1 {
			list.WriteString("└── ")
		} else {
			list.WriteString("├── ")
		}
		list.WriteString(node.Children[i].Name)
		list.WriteString("\n")
	}
	return list.String()
}

func render(node Node) string {
	numFiles := 0
	numDirectories := 1
	for i := 0; i < len(node.Children); i++ {
		if node.Children[i].Type == File {
			numFiles++
		}
		if node.Children[i].Type == Directory {
			numDirectories++
		}
	}

	return fmt.Sprintf(
		"%s\n%s\n%s\n",
		node.Name,
		tree(node),
		summary(numFiles, numDirectories),
	)
}
