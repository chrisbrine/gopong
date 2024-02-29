package pong

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

func NewGame() Pong {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	paddleHeight := 2
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	width, height := screen.Size()
	pong := Pong{
		paddles: [2]Paddle{
			NewPaddle(0, height / 2, tcell.ColorBlue, paddleHeight, false),
			NewPaddle(width, height / 2, tcell.ColorRed, paddleHeight, true),
		},
		balls: []Ball{NewBall(tcell.ColorWhite, 0, 1, screen, rng)},
		screen: screen,
		initialPaddleSize: paddleHeight,
		initialBallSize: 1,
		rand: rng,
		gameSpeed: 50, // as a percentage
		gameStarted: false,
	}
	return pong
}
