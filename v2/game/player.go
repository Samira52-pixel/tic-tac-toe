package game

type Player struct {
	figure BoardField
}

func NewPlayer() *Player {
	return &Player{
		figure: cross,
	}
}

func (p *Player) SwitchPlayer() {
	if p.figure == cross {
		p.figure = nought
	} else {
		p.figure = cross
	}
}

func (p *Player) Figure() BoardField {
	return p.figure
}

func (p *Player) GetSymbol() string {
	if p.figure == cross {
		return "X"
	}

	return "O"
}
