package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetField(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := &Person{Name: "Alice", Age: 30}
	name := GetField[string](p, "Name")
	fmt.Println(name)
}

func TestSetField(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := &Person{Name: "Alice", Age: 30}
	err := SetField(p, "Name", "Bob")
	if err != nil {
		t.Errorf("setField error: %s", err.Error())
	}
	if p.Name != "Bob" {
		t.Error("setField failed")
	}
	assert.Panics(t, func() {
		SetField(p, "Age", "0")
	}, "set string to int should panic")

	fmt.Println(p)
}
