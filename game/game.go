package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Event int

const (
	up Event = iota
	down
	left
	right
	quit
	red
)

func Wait() Event {
	event := termbox.PollEvent()
	if event.Key == termbox.KeyCtrlC {
		return quit
	}
	if event.Key == termbox.KeyArrowUp {
		return up
	}
	if event.Key == termbox.KeyArrowDown {
		return down
	}
	if event.Key == termbox.KeyArrowLeft {
		return left
	}
	if event.Key == termbox.KeyArrowRight {
		return right
	}
	return red
}

// 開きますは0にする
// ロジック用のfield
type Field [4][4]int

func Init() *Field {
	return &Field{
		{2, 2, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
		{2, 0, 0, 0},
	}
}

// ハンドラ
func Start() {
	field := Init()
	for {
		field.Draw()
		termbox.Flush()
		event := Wait()
		if event == quit {
			termbox.Close()
			return
		}
		if event == down {
			field.down()
			field.downPlus()
		}
		if event == up {
			field.up()
			field.upPlus()
		}
		if event == right {
			field.right()
			field.rightPlus()
		}
		if event == left {
			field.left()
			field.leftPlus()
		}
		field.createNewBlock()
	}
}
func (field *Field) createNewBlock() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if field[i][j] == 2048 {
				fmt.Println("GAME Clear")
				fmt.Println("CTR-Cを押してください")
			}
		}
	}
	for cnt := 0; cnt < 10; cnt++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if random.Intn(2) == 0 && field[i][j] == 0 {
					field[i][j] = 2
					return
				}
			}
		}
	}
	fmt.Println("GAMEOVER")
	fmt.Println("CTR-Cを押してください")
}
