package pong

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

func NewPaddle(x int, y int, color tcell.Color, size int, ai bool) Paddle {
	return Paddle{x: x, y: y, score: 0, color: color, size: size, ai: ai}
}

func restartPaddle(paddle Paddle, height int, width int) Paddle {
	paddle.x = height
	paddle.y = width
	return paddle
}

func randomCount(positiveOrNegative int, count int, rng *rand.Rand) int {
	return (rng.Intn(count) + 1) * positiveOrNegative
}

func randomDirection(positiveOrNegativeX int, positiveOrNegativeY int, rng *rand.Rand) Direction {
	return Direction{randomCount(positiveOrNegativeX, 5, rng), randomCount(positiveOrNegativeY, 5, rng)}
}

func randomNegOrPositive(current int, rng *rand.Rand) int {
	if current == 0 {
		current = rng.Intn(2)
		if current != 1 {
			return -1
		}
	} else if current > 1 {
		return 1
	
	} else if current < -1 {
		return -1
	}
	return current

}

func NewBall(color tcell.Color, directionX int, size int, screen tcell.Screen, rng *rand.Rand) Ball {
	direction := randomDirection(randomNegOrPositive(directionX, rng), randomNegOrPositive(0, rng), rng)
	width, height := screen.Size()
	return Ball{
		x: (width / 2) + (rng.Intn(10) - 5),
		y: (height / 2) + (rng.Intn(10) - 5),
		direction: direction,
		color: color,
		size: size,
	}
}

func (p *Pong) InitBalls(directionX int) {
	for i, ball := range p.balls {
		color := ball.color
		size := ball.size
		rng := p.rand
		p.balls[i] = NewBall(color, directionX, size, p.screen, rng)
	}
}

func (p *Pong) InitPaddles() {
	width, height := p.screen.Size()
	for i, paddle := range p.paddles {
		p.paddles[i] = restartPaddle(paddle, height, width)
	}
}

func (p *Pong) AddBall() {
	color := tcell.Color(p.rand.Intn(255) + 1)
	size := p.initialBallSize
	rng := p.rand
	if len(p.balls) < 10 {
		p.balls = append(p.balls, NewBall(color, 1, size, p.screen, rng))
	}
}

func (p *Pong) RemoveBall() {
	if len(p.balls) > 1 {
		p.balls = p.balls[:len(p.balls)-1]
	}
}
