package pong

func (p *Pong) CheckBall(ball int) {
	width, height := p.screen.Size()
	if (p.balls[ball].y < (p.balls[ball].size - 1)) ||
		(p.balls[ball].y > height-(p.balls[ball].size-1)) {
		if p.balls[ball].y < (p.balls[ball].size - 1) {
			p.balls[ball].y = p.balls[ball].size
		} else {
			p.balls[ball].y = height - (p.balls[ball].size - 1)
		}
		p.balls[ball].direction.y = -p.balls[ball].direction.y
	}
	if (p.balls[ball].x < (p.balls[ball].size - 1)) ||
		(p.balls[ball].x > width-(p.balls[ball].size-1)) {
		if p.balls[ball].x < (p.balls[ball].size - 1) {
			p.balls[ball].x = p.balls[ball].size
		} else {
			p.balls[ball].x = width - (p.balls[ball].size - 1)
		}
		p.balls[ball].direction.x = -p.balls[ball].direction.x
	}

}

func (p *Pong) CheckBalls() {
	for ball := range p.balls {
		p.CheckBall(ball)
	}
}