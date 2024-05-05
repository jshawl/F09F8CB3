package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type NodeType string

type Node struct {
	Children []*Node
	Name     string
	Type     NodeType
}

type Options struct {
	all bool
}

const (
	File      NodeType = "file"
	Directory NodeType = "directory"
)

func walk(path string, node *Node, options Options) Node {
	entries, _ := os.ReadDir(path)
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") && !options.all {
			continue
		}
		current := Node{Type: File, Name: e.Name(), Children: []*Node{}}
		if e.IsDir() {
			current.Type = Directory
			nextPath := fmt.Sprintf("%s/%s", path, e.Name())
			current = walk(nextPath, &current, options)
		}
		node.Children = append(node.Children, &current)
	}
	return *node
}

func entrypoint(path string, options Options) string {
	node := Node{Name: "test", Type: Directory, Children: []*Node{}}
	return render(walk(path, &node, options))
}

func main() {
	all := flag.Bool("a", false, "")
	flag.Parse()

	options := Options{all: *all}
	fmt.Print(entrypoint(flag.Args()[0], options))
}
