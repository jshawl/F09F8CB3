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

func render(node Node) string {
	numFiles := len(node.Children)
	directories := ""
	if numFiles > 0 {
		directories = fmt.Sprintf("%d directory", 1)
	} else {
		directories = fmt.Sprintf("%d directories", 0)
	}
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
	return fmt.Sprintf("%s\n%s\n%s, %d files\n", name, list.String(), directories, numFiles)
}
