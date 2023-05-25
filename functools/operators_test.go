package functools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myBooler bool

func (b myBooler) Bool() bool {
	return bool(b)
}

// 测试实现了 lener 接口的类型
type myLener []int

func (l myLener) Len() int {
	return len(l)
}

// 测试实现了 stringer 接口的类型
type myStringer int

func (i myStringer) String() string {
	return fmt.Sprintf("%d", i)
}

func TestTruth(t *testing.T) {
	// 测试常见类型的真值
	if !Truth(1) {
		t.Error("1 should be true")
	}
	if Truth(0) {
		t.Error("0 should be false")
	}
	if !Truth(3.14) {
		t.Error("3.14 should be true")
	}
	if Truth(0.0) {
		t.Error("0.0 should be false")
	}
	if !Truth("hello") {
		t.Error(`"hello" should be true`)
	}
	if Truth("") {
		t.Error(`"" should be false`)
	}

	// 测试实现了 booler 接口的类型

	if !Truth(myBooler(true)) {
		t.Error("myBooler(true) should be true")
	}
	if Truth(myBooler(false)) {
		t.Error("myBooler(false) should be false")
	}

	if !Truth(myLener{1, 2, 3}) {
		t.Error("myLener{1, 2, 3} should be true")
	}
	if Truth(myLener{}) {
		t.Error("myLener{} should be false")
	}

	if !Truth(myStringer(42)) {
		t.Error("myStringer(42) should be true")
	}

	if Truth(myStringer(0)) {
		t.Error("myStringer(0) should be false")
	}

	// 测试指针类型
	var p *int
	if Truth(p) {
		t.Error("nil pointer should be false")
	}
	var i int = 42
	p = &i
	if !Truth(p) {
		t.Error("non-nil pointer should be true")
	}
	*p = 0
	if Truth(p) {
		t.Error("non-nil pointer to 0 should be false")
	}

	// 测试其他类型
	type myStruct struct {
		x int
	}
	s := myStruct{}
	if !Truth(s) {
		t.Error("myStruct{} should be true")
	}
	m := map[string]int{"a": 1}
	if !Truth(m) {
		t.Error(`map[string]int{"a": 1} should be true`)
	}

	m = map[string]int{}
	if Truth(m) {
		t.Error(`empty map should be false`)
	}

	c := make(chan int)
	if Truth(c) {
		t.Error("empty channel should be false")
	}
	close(c)
	c = make(chan int, 1)
	c <- 1
	if !Truth(c) {
		t.Error("non-empty channel should be true")
	}
	var x any
	if Truth(x) {
		t.Error("uninitialed empty iterface should be false")
	}
}

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
