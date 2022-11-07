package main

import "testing"

func TestCat(t *testing.T) {
	got := Cat()
	if got != "miao~~~~~~~" {
		t.Errorf("mao cant %s", got)
	}
}
