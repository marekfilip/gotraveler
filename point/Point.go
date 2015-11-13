package point

import (
	"github.com/mattn/go-gtk/gtk"
)

type Point struct {
	X       uint
	Y       uint
	Val     uint
	Visited bool
	Parent  *Point
	Img     *gtk.Image
}

func New(x, y, val uint) *Point {
	return &Point{x, y, val, false, nil, nil}
}

func (p *Point) IsFinish() bool {
	if p.Val == 3 {
		return true
	}

	return false
}

func (p *Point) IsStart() bool {
	if p.Val == 2 {
		return true
	}

	return false
}

func (p *Point) IsWall() bool {
	if p.Val == 1 {
		return true
	}

	return false
}
