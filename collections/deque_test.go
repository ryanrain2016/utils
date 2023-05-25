package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque1(t *testing.T) {
	// test NewDeque
	dq := NewDeque(3, 1, 2, 3)
	assert.Equal(t, 3, dq.Len())
	dq1 := NewDeque[int](0)
	assert.Equal(t, -1, dq1.MaxLen)

	// test Append
	dq.Append(4)
	dq.Append(5)
	v, err := dq.Pop()
	assert.Equal(t, 5, v)
	assert.Nil(t, err, "normal pop should return no error")
	v, err = dq.Pop()
	assert.Equal(t, 4, v)
	assert.Nil(t, err, "normal pop should return no error")

	// test AppendLeft
	dq.AppendLeft(0)
	dq.AppendLeft(-1)
	v, err = dq.PopLeft()
	assert.Equal(t, -1, v)
	assert.Nil(t, err, "normal pop should return no error")
	v, err = dq.PopLeft()
	assert.Equal(t, 0, v)
	assert.Nil(t, err, "normal pop should return no error")

	// test Clear
	dq.Clear()
	assert.Equal(t, 0, dq.Len())

	// test Copy
	dq = NewDeque(0, 1, 2, 3)
	dq1 = dq.Copy().(*Deque[int])
	dq1.Append(4)
	assert.NotEqual(t, dq.Len(), dq1.Len())

	// test Count
	dq = NewDeque(0, 1, 2, 2, 3)
	assert.Equal(t, 2, dq.Count(2))

	// test Extend
	dq = NewDeque(0, 1, 2, 3)
	dq1 = NewDeque(0, 4, 5, 6)
	dq.Extend([]int{7, 8})
	dq.Extend(dq1)
	assert.Equal(t, 8, dq.Len())

	// test ExtendLeft
	dq = NewDeque(0, 1, 2, 3)
	dq1 = NewDeque(0, 4, 5, 6)
	dq.ExtendLeft([]int{7, 8})
	dq.ExtendLeft(dq1)
	assert.Equal(t, 8, dq.Len())

	// test Index
	dq = NewDeque(0, 1, 2, 3, 2)
	i, err := dq.Index(2)
	assert.Nil(t, err)
	assert.Equal(t, 1, i)
	i, err = dq.Index(2, 2)
	assert.Nil(t, err)
	assert.Equal(t, 3, i)

	// test Insert
	dq = NewDeque(0, 1, 2, 3)
	dq.Insert(1, 4)
	v, err = dq.GetItem(1)
	assert.Equal(t, 4, v)
	assert.Nil(t, err, "normal Get should return no error")
	dq.Insert(-1, 5)
	v, err = dq.GetItem(-2)
	assert.Equal(t, 5, v)
	assert.Nil(t, err, "normal Get should return no error")

	// test Pop
	dq = NewDeque(0, 1, 2, 3)
	x, err := dq.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 3, x)
	x, err = dq.PopLeft()
	assert.Nil(t, err)
	assert.Equal(t, 1, x)

	// test Delete
	dq = NewDeque(0, 1, 2, 3, 4)
	err = dq.Delete(2)
	assert.Nil(t, err)
	assert.Equal(t, 3, dq.Len())
	err = dq.Delete(-1)
	assert.Nil(t, err)
	assert.Equal(t, 2, dq.Len())

	// test Remove
	dq = NewDeque(0, 1, 2, 2, 3)
	err = dq.Remove(2)
	assert.Nil(t, err)
	assert.Equal(t, 3, dq.Len())

	// test Reverse
	dq = NewDeque(0, 1, 2, 3)
	dq.Reverse()
	v, _ = dq.Pop()
	assert.Equal(t, 1, v)
	v, _ = dq.Pop()
	assert.Equal(t, 2, v)

	// test Rotate
	dq = NewDeque(0, 1, 2, 3, 4)
	dq.Rotate()
	v, err = dq.PopLeft()
	assert.Nil(t, err)
	assert.Equal(t, 4, v)
	dq.Rotate(-1)
	v, err = dq.PopLeft()
	assert.Nil(t, err)
	assert.Equal(t, 2, v)

	// testGet and Set
	dq = NewDeque(0, 1, 2, 3)
	x, err = dq.GetItem(1)
	assert.Nil(t, err)
	assert.Equal(t, 2, x)
	err = dq.SetItem(1, 5)
	assert.Nil(t, err)
	x, err = dq.GetItem(1)
	assert.Equal(t, 5, x)
	assert.Nil(t, err)
}
