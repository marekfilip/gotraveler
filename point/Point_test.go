package point

import (
	"testing"
)

func TestIsWall(t *testing.T) {
	var r *Point = New(1, 1, 3) //meta
	var w *Point = New(1, 1, 1) //ściana
	var p *Point = New(1, 1, 0) //droga

	if r.IsWall() == true {
		t.Error("Point", p, " is not wall, expected false, recived true")
	}

	if w.IsWall() == false {
		t.Error("Point", p, "is wall, expected true, recived false")
	}

	if p.IsWall() == true {
		t.Error("Point", p, " is not wall, expected false, recived true")
	}
}

func TestIsFinish(t *testing.T) {
	var r *Point = New(1, 1, 3) //meta
	var w *Point = New(1, 1, 1) //ściana
	var p *Point = New(1, 1, 0) //droga

	if r.IsFinish() == false {
		t.Error("Point", p, " is finish, expected false, recived true")
	}

	if w.IsFinish() == true {
		t.Error("Point", p, " is not finish, expected true, recived false")
	}

	if p.IsFinish() == true {
		t.Error("Point", p, " is not finish, expected false, recived true")
	}
}
