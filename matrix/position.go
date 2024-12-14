package matrix

type Position struct{ X, Y int }

func (p *Position) SetMinX(i int) {
	if p.X < i {
		p.X = i
	}
}

func (p *Position) SetMaxX(i int) {
	if p.X > i {
		p.X = i
	}
}

func (p *Position) SetMinY(i int) {
	if p.Y < i {
		p.Y = i
	}
}

func (p *Position) SetMaxY(i int) {
	if p.Y > i {
		p.Y = i
	}
}

func (p *Position) Up() {
	p.X--
	p.SetMinX(0)
}

func (p *Position) Left() {
	p.Y--
	p.SetMinY(0)
}

func (p *Position) Down()  { p.X++ }
func (p *Position) Right() { p.Y++ }
