package main

import (
    "testing"
)

func TestSoma(t *testing.T) {
    got := soma(15,15)
    want := 30

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }	
	
}

