package pong

func (p *Pong) MoveAIPaddle() {
	// first find the closet ball to the paddle that is going in its direction
	// then move the paddle to the ball
	width, _ := p.screen.Size()
	for i, paddle := range p.paddles {
		if paddle.ai {
			closestBall := 0
			closestDistance := 1000
			for i, ball := range p.balls {
				if (ball.direction.x > 0 && paddle.x > (width / 2)) ||
				(ball.direction.x < 0 && paddle.x < (width / 2)) {
					distance := ball.x - paddle.x
					if distance < 0 {
						distance = -distance
					}
					if distance < closestDistance {
						closestDistance = distance
						closestBall = i
					}
				}
			}
			if p.balls[closestBall].y < paddle.y {
				p.paddles[i].y -= 1
			} else if p.balls[closestBall].y > paddle.y {
				p.paddles[i].y += 1
			}
		}
	}
}
