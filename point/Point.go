package point

type Point struct {
	X   uint
	Y   uint
	Val uint
}

func (p *Point) isFinish() bool {
	if p.Val == 3 {
		return true
	}

	return false
}

func (p *Point) isWall() bool {
	if p.Val == 1 {
		return true
	}

	return false
}
