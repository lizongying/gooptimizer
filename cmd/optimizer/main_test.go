package main

import (
	"github.com/lizongying/gooptimizer"
	"testing"
)

func TestMain_T1(t *testing.T) {
	result := gooptimizer.StructAlign(new(T1))
	if !result {
		t.Error("Expected true, but got false")
	}
}

func TestMain_Alignment(t *testing.T) {
	result := gooptimizer.StructAlign(new(gooptimizer.Alignment))
	if !result {
		t.Error("Expected true, but got false")
	}
}
