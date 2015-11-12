package point

import (
	"testing"
)

func TestIsWall(t *testing.T) {
	var r Point = Point{1, 1, 3, false, nil} //meta
	var w Point = Point{1, 1, 1, false, nil} //ściana
	var p Point = Point{1, 1, 0, false, nil} //droga

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
	var r Point = Point{1, 1, 3, false, nil} //meta
	var w Point = Point{1, 1, 1, false, nil} //ściana
	var p Point = Point{1, 1, 0, false, nil} //droga

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
