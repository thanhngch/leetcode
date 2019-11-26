package main

import "testing"

func TestPrefix(t *testing.T) {
	in := []string{"trinh", "truong", "troi"}
	want := "tr"

	got := longestCommonPrefix(in)
	if got != want {
		t.Errorf("prefix(%v) == %v, want %v ", in, got, want)
	}
}
