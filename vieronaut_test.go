package vieronaut

import "testing"

func TestQuest(t *testing.T) {
	v := Vieronaut{Name: "Vieronaut", Health: 100}
	_, err := Quest(v)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
