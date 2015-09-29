package main

import (
	"reflect"
	"testing"
)

func TestNewField(t *testing.T) {
	f := NewField()
	if reflect.TypeOf(f).String() != "*main.Field" {
		t.Errorf("Error NewField: return %q, want '*main.Field'", reflect.TypeOf(f))
	}
}

func TestNewFunction(t *testing.T) {
	f := NewFunction()
	if reflect.TypeOf(f).String() != "*main.Function" {
		t.Errorf("Error NewFunction: return %q, want '*main.Function'", reflect.TypeOf(f))
	}
}

func TestNewEntity(t *testing.T) {
	e := NewEntity()
	if reflect.TypeOf(e).String() != "*main.Entity" {
		t.Errorf("Error NewEntity: return %q, want '*main.Entity'", reflect.TypeOf(e))
	}
}
