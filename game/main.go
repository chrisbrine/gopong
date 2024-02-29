package pong

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

type Paddle struct {
	x int
	y int
	score int
	color tcell.Color
	size int
	ai bool
}

type Direction struct {
	x int
	y int
}

type Ball struct {
	x int
	y int
	direction Direction
	color tcell.Color
	size int
}

type Pong struct {
	paddles [2]Paddle
	balls []Ball
	screen tcell.Screen
	initialPaddleSize int
	initialBallSize int
	gameSpeed int
	rand *rand.Rand
	gameStarted bool
}
