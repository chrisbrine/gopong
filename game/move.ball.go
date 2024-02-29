package pong

func (p *Pong) MoveBall(ball int) {
	// p.balls[ball].x += p.balls[ball].direction.x
	// p.balls[ball].y += p.balls[ball].direction.y
	if p.balls[ball].direction.x < 0 {
		p.balls[ball].x -= 1
	} else {
		p.balls[ball].x += 1
	}
	if p.balls[ball].direction.y < 0 {
		p.balls[ball].y -= 1
	} else {
		p.balls[ball].y += 1
	}
}

func (p *Pong) MoveBalls() {
	for ball := range p.balls {
		p.MoveBall(ball)
		p.CheckBall(ball)
	}
	p.CheckCollision()
}
