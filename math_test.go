package main

import (
    "fmt"
    "testing"
)

func TestSoma(t *testing.T) {
    got := Soma(15,15)
    want := 30

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }	
	
}

