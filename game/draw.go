package game

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

func SetCell(x, y int, c rune) {
	termbox.SetCell(x, y, c, termbox.ColorLightBlue, termbox.ColorDefault)
}

// 2048までには4マス必要

// 4*4で一つにしよう
func ConvertRune(z int) [4]rune {
	s := strconv.Itoa(z)
	// aは空白の数
	a := 4 - len(s)
	var ans [4]rune
	for i := 0; i < a; i++ {
		ans[3-i] = ' '
	}
	for i := 0; i < 4-a; i++ {
		ans[i] = rune(s[i])
	}
	return ans
}

func (field *Field) Draw() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a := ConvertRune(field[j][i])
			for k, c := range a {
				SetCell(4*i+k, j, c)
			}
		}
	}
}
