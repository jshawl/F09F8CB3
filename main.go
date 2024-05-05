package main

import (
	"fmt"
	"os"
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

func walk(path string, node *Node) Node {
	entries, _ := os.ReadDir(path)
	for _, e := range entries {
		current := Node{Type: File, Name: e.Name(), Children: []*Node{}}
		if e.IsDir() {
			current.Type = Directory
			nextPath := fmt.Sprintf("%s/%s", path, e.Name())
			current = walk(nextPath, &current)
		}
		node.Children = append(node.Children, &current)
	}
	return *node
}

func entrypoint() string {
	node := Node{Name: "test", Type: Directory, Children: []*Node{}}
	return render(walk("./test", &node))
}

func main() {
	fmt.Print(entrypoint())
}
