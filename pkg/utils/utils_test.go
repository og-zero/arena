package utils

import (
	"os"
	"testing"
)

func TestWorkDir(t *testing.T) {
	got := WorkDir()

	want, err := os.Getwd()
	if err != nil {
		t.Errorf("can't get os.Getwd() to compare: %s", err)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
