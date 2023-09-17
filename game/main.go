package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	width  = 20
	height = 10
)

var (
	playerX, playerY int
	foodX, foodY     int
	player           = "🐛"
	food             = '🍏'
	gameOver         bool
	score            = 0
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initGame()
	defer keyboard.Close()

	for !gameOver {
		printGame()
		handleInput()
		updateGame()
		time.Sleep(100 * time.Millisecond)
		clearScreen()
	}
	fmt.Println("Game Over!")
}

func initGame() {
	playerX = width / 2
	playerY = height / 2
	placeFood()
	keyboard.Open()
}

func printGame() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == playerX && y == playerY {
				fmt.Print(string(player))
			} else if x == foodX && y == foodY {
				fmt.Print(string(food))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func handleInput() {
	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	if score == 10 {
		gameOver = true
	}
	switch key {
	case keyboard.KeyArrowUp:
		playerY--
	case keyboard.KeyArrowDown:
		playerY++
	case keyboard.KeyArrowLeft:
		playerX--
	case keyboard.KeyArrowRight:
		playerX++
	case keyboard.KeyEsc:	
		gameOver = true
	}
}

func updateGame() {
	checkCollision()
	if playerX < 0 {
		playerX = 0
	}
	if playerX >= width {
		playerX = width - 1
	}
	if playerY < 0 {
		playerY = 0
	}
	if playerY >= height {
		playerY = height - 1
	}
}

func checkCollision() {
	if playerX == foodX && playerY == foodY {
		player += "🐛"
		score++
		placeFood()
	}
}

func placeFood() {
	foodX = rand.Intn(width)
	foodY = rand.Intn(height)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(score)
}
