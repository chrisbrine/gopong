package pong

func (p *Pong) CheckCollision() {
	// check if the ball collides with the paddles, if so then change the direction of the ball
	width, _ := p.screen.Size()
	for ball := range p.balls {
		for paddle := range p.paddles {
			size := p.GetRealSize(p.paddles[paddle].size)
			if p.balls[ball].x == p.paddles[paddle].x ||
			   p.balls[ball].x == p.paddles[paddle].x+1 ||
				 p.balls[ball].x == p.paddles[paddle].x-1 {
				if p.balls[ball].y >= p.paddles[paddle].y &&
				p.balls[ball].y <= p.paddles[paddle].y+size {
					p.balls[ball].direction.x = -p.balls[ball].direction.x
					// check where the ball hits the paddle and modify the y direction based on it
					if p.balls[ball].y < p.paddles[paddle].y+(size/2) {
						// the ball hits the upper part of the paddle
						// the y should be negative, but modify angle based on the distance from the center of the paddle
						p.balls[ball].direction.y = -1 - (p.paddles[paddle].y + (size / 2) - p.balls[ball].y)
					} else if p.balls[ball].y > p.paddles[paddle].y+(size/2) {
						// the ball hits the lower part of the paddle
						// the y should be positive, but modify angle based on the distance from the center of the paddle
						p.balls[ball].direction.y = 1 + (p.balls[ball].y - (p.paddles[paddle].y + (size / 2)))
					}
				// also check if hits the wall if outside of the paddle
				} else {
					// give point to other paddles, and reset the ball
					for i := range p.paddles {
						if i != paddle {
							p.paddles[i].score++
						}
					}
					// get a directionX for the ball toward where the winning paddle is
					var directionX int = -1
					if p.paddles[paddle].x < width/2 {
						directionX = 1
					}
					p.InitBalls(directionX)
					// p.InitPaddles()
					p.gameStarted = false
				}
			}
		}
	}
	p.CheckBalls()
	p.CheckPaddles()
	p.Draw()
}