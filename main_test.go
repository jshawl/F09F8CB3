package main

import (
	"testing"
)
func TestRender(t *testing.T){
	expected := "./"
	actual := render()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}