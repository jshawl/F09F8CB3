package main

import (
	"fmt"
	"strings"
)

type NodeType string

type Node struct {
  Children []*Node
  Name string
  Type NodeType
}

const (
	File   NodeType = "file"
	Directory NodeType = "directory"
)

func main() {
	file1 := Node{Name: "file1"}
	file2 := Node{Name: "file2"}
	node := Node{
		Name: ".",
		Children: []*Node{&file1, &file2},
	}
	fmt.Print(render(node))
}

func summary(numFiles int, numDirectories int) string {
	directories := ""
	files := fmt.Sprintf("%d files", numFiles)
	if numFiles > 0 {
		directories = fmt.Sprintf("%d directory", 1)
	} else {
		directories = fmt.Sprintf("%d directories", numDirectories)
	}
	return fmt.Sprintf("%s, %s", directories, files)
}

func render(node Node) string {
	numFiles := len(node.Children)
	numDirectories := len(node.Children)
	var list strings.Builder
    for i := 0; i < numFiles; i++ {
		if i == numFiles - 1 {
			list.WriteString("└── ")
		} else {
			list.WriteString("├── ")
		}
        list.WriteString(node.Children[i].Name)
		list.WriteString("\n")
    }
	name := node.Name
	return fmt.Sprintf(
		"%s\n%s\n%s\n", 
		name, 
		list.String(), 
		summary(numFiles, numDirectories),
	)
}
