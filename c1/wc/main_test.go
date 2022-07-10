package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")
	exp := 3
	res, _ := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, bug got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("one\ntwo\nthree\n")
	exp := 3
	res, _ := count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d, but got %d instead", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("0123456789")
	exp := 10
	res, _ := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d, but got %d instead", exp, res)
	}
}

func TestExclusiveOptions(t *testing.T) {
	b := bytes.NewBufferString("")
	_, err := count(b, true, true)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
