package point

type Point struct {
	X       uint
	Y       uint
	Val     uint
	Visited bool
	Parent  *Point
}

func (p *Point) IsFinish() bool {
	if p.Val == 3 {
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
