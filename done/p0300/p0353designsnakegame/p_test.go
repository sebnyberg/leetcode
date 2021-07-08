package p0353designsnakegame

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnakeGame(t *testing.T) {
	g := Constructor(3, 2, [][]int{{1, 2}, {0, 1}})
	res := g.Move("R")
	require.Equal(t, 0, res)
	res = g.Move("D")
	require.Equal(t, 0, res)
	res = g.Move("R")
	require.Equal(t, 1, res)
	res = g.Move("U")
	require.Equal(t, 1, res)
	res = g.Move("L")
	require.Equal(t, 2, res)
	res = g.Move("U")
	require.Equal(t, -1, res)
}

type SnakeGame struct {
	width, height int
	food          [][]int
	foodPos       int
	body          [][]bool
	snake         *list.List
}

/** Initialize your data structure here.
  @param width - screen width
  @param height - screen height
  @param food - A list of food positions
  E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */
func Constructor(width int, height int, food [][]int) SnakeGame {
	snake := list.New()
	snake.PushFront(position{0, 0})
	body := make([][]bool, height)
	for i := range body {
		body[i] = make([]bool, width)
	}
	sg := SnakeGame{
		width:  width,
		height: height,
		food:   food,
		snake:  snake,
		body:   body,
	}
	sg.body[0][0] = true
	return sg
}

type position struct {
	i, j int
}

/** Moves the snake.
  @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
  @return The game's score after the move. Return -1 if game over.
  Game over when snake crosses the screen boundary or bites its body. */
func (this *SnakeGame) Move(direction string) int {
	head := this.snake.Front().Value.(position)
	switch direction {
	case "U":
		head.i--
	case "D":
		head.i++
	case "L":
		head.j--
	case "R":
		head.j++
	}
	// If head goes out of bounds, it does not matter whether it eats food or not
	if head.i < 0 || head.i >= this.height ||
		head.j < 0 || head.j >= this.width {
		return -1
	}
	// If head eats food, it may impact whether the head will hit the body
	// check whether it eats food first.
	eatsFood := this.foodPos < len(this.food) &&
		this.food[this.foodPos][0] == head.i &&
		this.food[this.foodPos][1] == head.j

	// If the head does not eat food, the tail will move to its next position
	if !eatsFood {
		tail := this.snake.Remove(this.snake.Back()).(position)
		this.body[tail.i][tail.j] = false

		// Check if head hit its own body
		if this.body[head.i][head.j] {
			return -1
		}
	} else {
		this.foodPos++
	}

	// Move head forward
	this.snake.PushFront(head)
	this.body[head.i][head.j] = true

	return this.foodPos
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
