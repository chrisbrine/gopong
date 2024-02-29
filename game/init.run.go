package pong

import (
	"time"

	"github.com/gdamore/tcell"
)

func (p *Pong) getInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := p.screen.PollEvent().(type) {
				case *tcell.EventKey:
					switch ev.Key() {
						case tcell.KeyEscape:
							inputChan <- "exit"
						case tcell.KeyUp:
							inputChan <- "up2"
						case tcell.KeyDown:
							inputChan <- "down2"
						case tcell.KeyCtrlC:
							inputChan <- "exit"
						case tcell.KeyRune:
							if ev.Rune() == 'b' || ev.Rune() == 'B' {
								inputChan <- "addBall"
							} else if ev.Rune() == 'n' || ev.Rune() == 'N' {
								inputChan <- "removeBall"
							} else if ev.Rune() == 'p' || ev.Rune() == 'P' {
								inputChan <- "increasePaddle"
							} else if ev.Rune() == 'o' || ev.Rune() == 'O' {
								inputChan <- "decreasePaddle"
							} else if (ev.Rune() == '1') {
								inputChan <- "toggleAI1"
							} else if (ev.Rune() == '2') {
								inputChan <- "toggleAI2"
							} else if ev.Rune() == 'q' || ev.Rune() == 'Q' {
								inputChan <- "exit"
							} else if ev.Rune() == 'w' || ev.Rune() == 'W' {
								inputChan <- "up1"
							} else if ev.Rune() == 's' || ev.Rune() == 'S' {
								inputChan <- "down1"
							} else if ev.Rune() == '=' || ev.Rune() == '+' {
								inputChan <- "increaseSpeed"
							} else if ev.Rune() == '-' || ev.Rune() == '_' {
								inputChan <- "decreaseSpeed"
							}
					}
			}
		}
	}()

	return inputChan
}

func (p *Pong) Run() {
	inputChan := p.getInput()
	var gameSpeed int
	for {
		// p.Draw()
		// get key event
		gameSpeed = 200 - (p.gameSpeed * 2)
		if gameSpeed < 1 {
			gameSpeed = 1
		} else if gameSpeed > 100 {
			gameSpeed = 100
		}
		time.Sleep(time.Duration(gameSpeed) * time.Millisecond)
		p.Draw()
		if p.gameStarted {
			p.MoveBalls()
			p.MoveAIPaddle()
		}
		p.CheckPaddles()
		p.CheckBalls()
		var option string
		select {
			case option = <-inputChan:
			default:
				option = ""
		}
		if option != "" {
			p.gameStarted = true
		}
		switch option {
			case "up1":
				p.MovePaddle1(-1)
			case "down1":
				p.MovePaddle1(1)
			case "up2":
				p.MovePaddle2(-1)
			case "down2":
				p.MovePaddle2(1)
			case "addBall":
				p.AddBall()
			case "removeBall":
				p.RemoveBall()
			case "increasePaddle":
				p.IncreasePaddle()
			case "decreasePaddle":
				p.DecreasePaddle()
			case "toggleAI1":
				p.paddles[0].ai = !p.paddles[0].ai
			case "toggleAI2":
				p.paddles[1].ai = !p.paddles[1].ai
			case "increaseSpeed":
				p.gameSpeed += 10
				if (p.gameSpeed > 100) {
					p.gameSpeed = 100
				}
			case "decreaseSpeed":
				p.gameSpeed -= 10
				if (p.gameSpeed < 0) {
					p.gameSpeed = 0
				}
			case "exit":
				p.screen.Fini()
				return
		}
		// go func () {
		// 	switch ev := p.screen.PollEvent().(type) {
		// 		case *tcell.EventResize:
		// 			p.screen.Sync()
		// 			p.Draw()
		// 		case *tcell.EventKey:
		// 			switch ev.Key() {
		// 				case tcell.KeyUp:
		// 					p.MovePaddle(-1)
		// 				case tcell.KeyDown:
		// 					p.MovePaddle(1)
		// 				case tcell.KeyCtrlC:
		// 					p.screen.Fini()
		// 					return
		// 				case tcell.KeyRune:
		// 					if ev.Rune() == 'b' || ev.Rune() == 'B' {
		// 						p.AddBall()
		// 					}
		// 					if ev.Rune() == 'n' || ev.Rune() == 'N' {
		// 						p.RemoveBall()
		// 					}
		// 					if ev.Rune() == 'p' || ev.Rune() == 'P' {
		// 						p.IncreasePaddle()
		// 					}
		// 					if ev.Rune() == 'o' || ev.Rune() == 'O' {
		// 						p.DecreasePaddle()
		// 					}
		// 					if (ev.Rune() == '1') {
		// 						p.paddles[0].ai = !p.paddles[0].ai
		// 					}
		// 					if (ev.Rune() == '2') {
		// 						p.paddles[1].ai = !p.paddles[1].ai
		// 					}
		// 					if ev.Rune() == 'q' || ev.Rune() == 'Q' {
		// 						p.screen.Fini()
		// 						return
		// 					}
		// 			}
		// 	}
		// }()
	}
}