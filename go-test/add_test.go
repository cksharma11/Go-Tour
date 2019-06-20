package test

import (
	"testing"
)

func TestAdd(t *testing.T){
	got := Add(1,2)
	if got != 3 {
		t.Errorf("Failed")
	}
}