package queue

import (
	//"fmt"
	"gtktest/point"
)

type Queue []point.Point

func (q *Queue) IsEmpty() bool {
	if len(*q) > 0 {
		return false
	}

	return true
}

/*
Ściąga z końca
*/
func (q *Queue) Pop() *point.Point {
	if len(*q) == 0 {
		return nil
	}
	var x point.Point
	x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]
	return &x
}

/*
Dokłada na koniec
*/
func (q *Queue) Push(p point.Point) {
	*q = append(*q, p)
}

/*
Ściąga z początku
*/
func (q *Queue) Shift() *point.Point {
	if len((*q)) == 0 {
		return nil
	}
	var x point.Point
	x, *q = (*q)[0], (*q)[1:]

	return &x
}

/*
Dokłada na początek
*/
func (q *Queue) Unshift(p point.Point) {
	*q = append(Queue{p}, (*q)...)
}
