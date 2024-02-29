package pong

import (
	"fmt"

	"github.com/gdamore/tcell"
)

func (p *Pong) Write(x int, y int, text string) {
	p.ShowText(tcell.ColorWhite, tcell.ColorBlack, x, y, text)
}

func createScoreString(player int, score int, ai bool) string {
	scoreString := "Player " + fmt.Sprint(player) + ": " + fmt.Sprint(score)
	if ai {
		scoreString += " (AI)"
	}
	return scoreString
}

func (p *Pong) Draw() {
	width, height := p.screen.Size()
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	p.screen.SetStyle(style)
	p.screen.Clear()
	score1String := createScoreString(1, p.paddles[0].score, p.paddles[0].ai)
	score2String := createScoreString(2, p.paddles[1].score, p.paddles[1].ai)
	p.Write(width / 2 - 3, 0, "GoPONG")
	p.Write(2, 1, score1String)
	p.Write(width - len(score2String) - 2, 1, score2String)
	if !p.gameStarted {
		p.Write(width / 2 - 11, height / 2, "Press any movement key to start!")
		if (!p.paddles[0].ai) {
			if (!p.paddles[1].ai) {
				p.Write(2, height - 1, "W: Up, S: Down")
			} else {
				p.Write(2, height - 1, "/\\: Up, \\/: Down")
			}
		}
		if (!p.paddles[1].ai) {
			p.Write(width - 16 - 2, height - 1, "/\\: Up, \\/: Down")
		}
		p.Write(width / 2 - 11, height - 2, "Press Ctrl+C to Quit!")
		}
	p.drawPaddles()
	p.drawBall()
	p.screen.Show()
}
