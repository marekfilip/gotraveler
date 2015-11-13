package queue

import (
	"gotraveler/point"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var q Queue = Queue{}

	if q.IsEmpty() == false {
		t.Error("Queue is empty returned false, expected true")
	}

	q = append(q, &point.Point{})

	if q.IsEmpty() == true {
		t.Error("Queue is not empty returned true, expected false")
	}
}

func TestPop(t *testing.T) {
	var q Queue = Queue{
		point.New(1, 1, 1),
		point.New(2, 2, 2),
	}
	var poped *point.Point

	poped = q.Pop()
	if poped.X != 2 || poped.Y != 2 || poped.Val != 2 {
		t.Error("Bad point poped:", poped, "expected all fields to be 2")
	}

	poped = q.Pop()

	if poped.X != 1 || poped.Y != 1 || poped.Val != 1 {
		t.Error("Bad point poped:", poped, "expected all fields to be 1")
	}

	poped = q.Pop()

	if poped != nil {
		t.Error("Bad point poped:", poped, "expected nil")
	}
}

func TestShift(t *testing.T) {
	var q Queue = Queue{
		point.New(1, 1, 1),
		point.New(2, 2, 2),
	}
	var shifted *point.Point

	shifted = q.Shift()

	if shifted.X != 1 || shifted.Y != 1 || shifted.Val != 1 {
		t.Error("Bad point shifted:", shifted, "expected all fields to be 1")
	}

	shifted = q.Shift()
	if shifted.X != 2 || shifted.Y != 2 || shifted.Val != 2 {
		t.Error("Bad point shifted:", shifted, "expected all fields to be 2")
	}

	shifted = q.Shift()
	if shifted != nil {
		t.Error("Bad point shifted:", shifted, "expected nil")
	}
}

func TestPush(t *testing.T) {
	var q Queue = Queue{
		point.New(1, 1, 1),
		point.New(2, 2, 2),
	}

	q.Push(point.New(3, 3, 3))

	if q[2].X != 3 || q[2].Y != 3 || q[2].Val != 3 {
		t.Error("Point didnt pop:", q, "expected last point to be all 3")
	}
}

func TestUnshift(t *testing.T) {
	var q Queue = Queue{
		point.New(1, 1, 1),
		point.New(2, 2, 2),
	}

	q.Unshift(point.New(3, 3, 3))

	if q[0].X != 3 || q[0].Y != 3 || q[0].Val != 3 {
		t.Error("Point didnt unshift:", q, "expected first point to be all 3")
	}
}
