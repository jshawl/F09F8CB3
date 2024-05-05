package main

import "fmt"

type Tree struct {
  Directories []Tree
  Files []string
}

func main() {
	tree := Tree {}
	fmt.Print(render(tree))
}

func render(tree Tree) string {
	directories := len(tree.Directories)
	files := len(tree.Files)
	return fmt.Sprintf(".\n\n%d directories, %d files\n", directories, files)
}
