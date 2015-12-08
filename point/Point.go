package point

import (
	"github.com/mattn/go-gtk/gtk"
)

const (
	WAY    uint = 0
	WALL   uint = 1
	START  uint = 2
	FINISH uint = 3
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
	if p.Val == FINISH {
		return true
	}

	return false
}

func (p *Point) IsStart() bool {
	if p.Val == START {
		return true
	}

	return false
}

func (p *Point) IsWall() bool {
	if p.Val == WALL {
		return true
	}

	return false
}

func (p *Point) ColorVisitedPoint(tab *gtk.Table) {
	p.Img = gtk.NewImageFromFile("./images/reddot.png")
	tab.Attach(p.Img, p.X, p.X+1, p.Y, p.Y+1, gtk.FILL, gtk.FILL, 0, 0)
	p.Img.Show()
}
