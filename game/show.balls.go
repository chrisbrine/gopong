package pong

import (
	"github.com/gdamore/tcell"
)

func (p *Pong) drawBall() {
	for i, ball := range p.balls {
		p.CheckBall(i)
		style := tcell.StyleDefault.Background(ball.color)
		p.screen.SetContent(ball.x, ball.y, ' ', nil, style)
	}
}