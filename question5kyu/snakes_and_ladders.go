package question5kyu

import "fmt"

//ref: https://www.codewars.com/kata/587136ba2eefcb92a9000027

var Ladders = map[int]int{
	2:  38,
	7:  14,
	8:  31,
	15: 26,
	21: 42,
	28: 84,
	36: 44,
	51: 67,
	71: 91,
	78: 98,
	87: 94,
}

var Snakes = map[int]int{
	16: 6,
	46: 25,
	49: 11,
	62: 19,
	64: 60,
	74: 53,
	89: 68,
	92: 88,
	95: 75,
	99: 80,
}

type SnakesLadders struct {
	Player1Pos int
	Player2Pos int
	Turn       bool
}

func (sl *SnakesLadders) NewGame() {
	sl.Player1Pos = 0
	sl.Player2Pos = 0
	sl.Turn = true // true = Player 1, false = Player 2
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
	if sl.Player1Pos == 100 || sl.Player2Pos == 100 {
		return "Game over!"
	}

	sum := die1 + die2
	if sl.Turn {
		sl.Player1Pos = walking(sl.Player1Pos, sum)
	}
	if !sl.Turn {
		sl.Player2Pos = walking(sl.Player2Pos, sum)
	}

	if sl.Player1Pos == 100 {
		return "Player 1 Wins!"
	}
	if sl.Player2Pos == 100 {
		return "Player 2 Wins!"
	}

	return sl.endTurn(die1, die2)
}

func walking(pos int, dice int) int {
	newPos := pos + dice
	if newPos > 100 {
		newPos = 100 - (newPos - 100)
	}
	pos = newPos
	pos = useLadders(pos)
	pos = useSnake(pos)

	return pos
}

func useLadders(pos int) int {
	if Ladders[pos] != 0 {
		return Ladders[pos]
	}

	return pos
}

func useSnake(pos int) int {
	if Snakes[pos] != 0 {
		return Snakes[pos]
	}

	return pos
}

func (sl *SnakesLadders) endTurn(die1, die2 int) string {
	result := ""
	if sl.Turn {
		result = fmt.Sprintf("Player 1 is on square %d", sl.Player1Pos)
	} else {
		result = fmt.Sprintf("Player 2 is on square %d", sl.Player2Pos)
	}

	if die1 != die2 {
		sl.Turn = !sl.Turn
	}

	return result
}
