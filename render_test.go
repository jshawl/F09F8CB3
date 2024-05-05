package main

import (
	"testing"
)

func TestRenderEmptyDirectory(t *testing.T) {
	expected := `.

0 directories, 0 files
`
	node := Node{
		Name: ".",
		Type: Directory,
	}
	actual := render(node)
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestRenderDirectoryWithFiles(t *testing.T) {
	expected := `.
├── file1
└── file2

1 directory, 2 files
`
	file1 := Node{Name: "file1", Type: File}
	file2 := Node{Name: "file2", Type: File}
	node := Node{
		Name:     ".",
		Children: []*Node{&file1, &file2},
		Type:     Directory,
	}
	actual := render(node)
	if actual != expected {
		t.Errorf("Expected:\n%s, Actual:\n%s", expected, actual)
	}
}

func TestRenderDirectoryWithSubdirectories(t *testing.T) {
	expected := `.
├── directory1
│   ├── directory2
│   │   └── file4
│   └── file3
├── file1
└── file2

3 directories, 4 files
`
	file1 := Node{Name: "file1", Type: File}
	file2 := Node{Name: "file2", Type: File}
	file3 := Node{Name: "file3", Type: File}
	file4 := Node{Name: "file4", Type: File}
	directory2 := Node{Name: "directory2", Type: Directory, Children: []*Node{&file4}}
	directory1 := Node{Name: "directory1", Type: Directory, Children: []*Node{&file3, &directory2}}
	node := Node{
		Name:     ".",
		Children: []*Node{&file1, &file2, &directory1},
		Type:     Directory,
	}
	actual := render(node)
	if actual != expected {
		t.Errorf("Expected:\n%s, Actual:\n%s", expected, actual)
	}
}
