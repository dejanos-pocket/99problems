package main

import (
	"advent23/util/assert"
	"advent23/util/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file, lines := input.LineByLine("day2/input.txt")
	defer file.Close()

	fmt.Println("part_1", part1(lines)) // 2617
	fmt.Println("part_2", part2(lines)) // 59795
}

type Game struct {
	id    int
	power int
	sets  []Set
}
type Set struct {
	cubes map[string]int
}

func part1(lines []string) int {
	constrains := map[string]int{"red": 12, "green": 13, "blue": 14}
	games := newGames(lines)
	sum := 0
	for _, game := range games {
		possible := true
		for _, set := range game.sets {
			for color, cnt := range set.cubes {
				if cnt > constrains[color] {
					possible = false
				}
			}
		}
		if possible {
			sum += game.id
		}
	}
	return sum
}

func part2(lines []string) int {
	games := newGames(lines)
	sum := 0
	for i := range games {
		game := games[i]
		maxcnts := map[string]int{}
		for _, set := range game.sets {
			for color, cnt := range set.cubes {
				if cnt > maxcnts[color] {
					maxcnts[color] = cnt
				}
			}
		}
		for _, v := range maxcnts {
			game.power = game.power * v
		}
		sum += game.power
	}
	return sum
}

func newGames(lines []string) []*Game {
	var games []*Game
	for _, line := range lines {
		games = append(games, newGame(line))
	}
	return games
}

func newGame(line string) *Game {
	gameInfo := strings.Split(line, ":")
	gameID, err := strconv.Atoi(strings.Split(gameInfo[0], " ")[1])
	assert.Empty(err)
	game := &Game{id: gameID, power: 1}
	for _, set := range strings.Split(gameInfo[1], ";") {
		s := Set{cubes: map[string]int{}}
		for _, cube := range strings.Split(set, ", ") {
			cntByColor := strings.Split(strings.TrimSpace(cube), " ")
			cnt, err := strconv.Atoi(cntByColor[0])
			assert.Empty(err)
			color := cntByColor[1]
			s.cubes[color] = cnt
		}
		game.sets = append(game.sets, s)
	}
	return game
}
