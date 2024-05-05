package main

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

}
