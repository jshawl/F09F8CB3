package main

import (
	"testing"
)

func TestRenderEmptyDirectory(t *testing.T){
	expected := `.

0 directories, 0 files
`
	actual := render(Tree{})
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
