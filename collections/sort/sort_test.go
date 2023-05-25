package sort

import (
	"reflect"
	"testing"

	"github.com/ryanrain2016/utils/collections"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name     string
	Fullname string
	Age      int
}

func TestSortSortBy(t *testing.T) {
	l := collections.NewList[*Person](0)
	l.Append(&Person{
		"bob", "bob bar", 27,
	})
	l.Append(&Person{
		"bob", "bob foo", 29,
	})
	l.Append(&Person{
		"bob", "bob bar", 28,
	})
	l.Append(&Person{
		"alice", "alice foo", 18,
	})
	sort := SortBy(func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}, func(p1, p2 *Person) bool {
		return p1.Fullname < p2.Fullname
	}, func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	})
	sort.Sort(l)
	sl := l.ToSlice()
	expected := []*Person{
		{
			"alice", "alice foo", 18,
		},
		{
			"bob", "bob bar", 27,
		},
		{
			"bob", "bob bar", 28,
		},
		{
			"bob", "bob foo", 29,
		},
	}
	if !reflect.DeepEqual(sl, expected) {
		t.Errorf("Sort failed.")
	}
}

func TestSortSortByKey(t *testing.T) {
	l := collections.NewList[*Person](0)
	l.Append(&Person{
		"bob", "bob bar", 27,
	})
	l.Append(&Person{
		"bob", "bob foo", 29,
	})
	l.Append(&Person{
		"bob", "bob bar", 28,
	})
	l.Append(&Person{
		"alice", "alice foo", 18,
	})
	sort := SortByKey[*Person](func(p1 *Person) string {
		return p1.Name
	}, func(p1 *Person) string {
		return p1.Fullname
	}, func(p1 *Person) int {
		return p1.Age
	})
	sort.Sort(l)
	sl := l.ToSlice()
	expected := []*Person{
		{
			"alice", "alice foo", 18,
		},
		{
			"bob", "bob bar", 27,
		},
		{
			"bob", "bob bar", 28,
		},
		{
			"bob", "bob foo", 29,
		},
	}
	if !reflect.DeepEqual(sl, expected) {
		t.Errorf("Sort failed.")
	}
}

func TestSortSortByKeyWithWrongKeyFunc(t *testing.T) {
	assert.Panics(t, func() {
		SortByKey[*Person](func(p1 *Person) *Person {
			return p1
		})
	})
}
