package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestNoArguments(t *testing.T) {
	version, _ := exec.Command("tree", "--version").Output()
	fmt.Println(string(version))
	out, _ := exec.Command("tree", "test").Output()
	expected := strings.ReplaceAll(string(out), "\u00a0", " ")
	actual := entrypoint()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
