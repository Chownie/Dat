// parsing.go
package parsing

import (
	"fmt"
	"strings"
	"strconv"
)

type player struct {
	addr string
	x    int
	y    int
}

func NewPlayer(s string) *player {
	d := new(player)
	d.addr = s
	d.y = 0
	d.x = 0
	return d
}

func move(pl *player, s string) string {
	switch s {
	case "L":
		pl.x -= 1
	case "R":
		pl.x += 1
	case "D":
		pl.y -= 1
	case "U":
		pl.y += 1
	}
	fmt.Println("ARGS:", s)
	fmt.Println(pl.addr, "has moved! Now at", pl.x, pl.y)
	return "AT " + strconv.Itoa(pl.x) + " " + strconv.Itoa(pl.y)
}

func die(pl *player, s string) string {
	pl.x = 0
	pl.y = 0
	fmt.Println(pl.addr, "was reset")
	return "AT " + strconv.Itoa(pl.x) + " " + strconv.Itoa(pl.y)
}

func Parse(pl *player, s string, output chan string) []uint8 {
	var commands = map[string]func(*player, string) string{
		"MOVE": move,
		"DIE":  die}

	s = strings.TrimSpace(s)
	str := strings.SplitN(s, ",", 2)

	l := ""
	if cmd, ok := commands[str[0]]; ok {
		if len(str) < 2 {
			return []uint8("Incorrect number of args\n")
		}
		l = commands[str[0]](pl, str[1])
		fmt.Println(cmd, l)
	} else {
		l = "Illegal move"
		fmt.Println("No such command called", string(str[0]))
	}
	output <- (l + "\n")
	return []uint8(l + "\n")
}
