package pong

import (
	"github.com/gdamore/tcell"
)

func (p *Pong) ShowText(color tcell.Color, bg tcell.Color, x int, y int, text string) {
	style := tcell.StyleDefault.Foreground(color).Background(bg)
	for i, c := range text {
		p.screen.SetContent(x+i, y, c, nil, style)
	}
}
