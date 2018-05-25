package main

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	input := "group_chat"

	output := GetCamelCase(input)

	if output != "GroupChat" {
		t.Fatal("Fail")
	}
}
