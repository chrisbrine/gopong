package pong

func (p *Pong) GetRealSize(size int) int {
	_, height := p.screen.Size()
  newSize := height / (10 / size)
	return newSize
}

func (p *Pong) CheckPaddle(paddle int) {
	width, height := p.screen.Size()
	size := p.GetRealSize(p.paddles[paddle].size)
	if p.paddles[paddle].x > width / 2 {
		p.paddles[paddle].x = width - 1
	} else {
		p.paddles[paddle].x = 0
	}
	if p.paddles[paddle].y < 1 {
		p.paddles[paddle].y = 0
	} else if p.paddles[paddle].y >= (height - size){
		p.paddles[paddle].y = height - size
	}
}

func (p *Pong) CheckPaddles() {
	for paddle := range p.paddles {
		p.CheckPaddle(paddle)
	}
}