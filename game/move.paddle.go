package pong

func (p *Pong) MovePaddle1(directionY int) {
	// check if both paddles are not AI, if so just move the first paddle, otherwise trigger MovePaddle()
	if !p.paddles[0].ai && !p.paddles[1].ai {
		p.paddles[0].y += directionY
	} else {
		p.MovePaddle(directionY)
	}
}

func (p *Pong) MovePaddle2(directionY int) {
	// check if both paddles are not AI, if so just move the first paddle, otherwise trigger MovePaddle()
	if !p.paddles[0].ai && !p.paddles[1].ai {
		p.paddles[1].y += directionY
	} else {
		p.MovePaddle(directionY)
	}
}

func (p *Pong) MovePaddle(directionY int) {
	// find non AI paddle and move it
	for i, paddle := range p.paddles {
		if !paddle.ai {
			p.paddles[i].y += directionY
		}
	}
}

func (p *Pong) IncreasePaddle() {
	// find non AI paddle and increase its size
	for i, paddle := range p.paddles {
		if !paddle.ai {
			p.paddles[i].size++
			if (p.paddles[i].size > 10) {
				p.paddles[i].size = 10
			}
		}
	}
}

func (p *Pong) DecreasePaddle() {
	// find non AI paddle and decrease its size
	for i, paddle := range p.paddles {
		if !paddle.ai {
			p.paddles[i].size--
			if (p.paddles[i].size < 1) {
				p.paddles[i].size = 1
			}
		}
	}
}
