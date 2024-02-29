package pong

import (
	"github.com/gdamore/tcell"
)

func (p *Pong) drawPaddles() {
	for i, paddle := range p.paddles {
		p.CheckPaddle(i)
		size := p.GetRealSize(paddle.size)
		for i := 0; i < size; i++ {
			style := tcell.StyleDefault.Background(paddle.color)
			p.screen.SetContent(paddle.x, paddle.y+i, ' ', nil, style)
		}
	}
}
