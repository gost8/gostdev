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

func TestFieldSetters(t *testing.T) {
	f := NewField().
		SetDescription("Test description").
		SetName("testFieldName").
		SetType("string").
		SetLength(255).
		SetMinval(1).
		SetMaxval(2.5)

	if f.Description != "Test description" {
		t.Errorf("Error SetDescription: return %q, want 'Test description'", f.Description)
	}
	if f.Name != "testFieldName" {
		t.Errorf("Error SetName: return %q, want 'testFieldName'", f.Name)
	}
	if f.Type != "string" {
		t.Errorf("Error SetType: return %q, want 'string'", f.Type)
	}
	if f.Length != 255 {
		t.Errorf("Error SetLength: return %q, want '255'", f.Length)
	}
	if f.Minval != 1 {
		t.Errorf("Error SetMinval: return %q, want '1'", f.Minval)
	}
	if f.Maxval != 2.5 {
		t.Errorf("Error SetMaxval: return %q, want '2.5'", f.Maxval)
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
