package main

import (
	"termboxPractice/game"

	"github.com/nsf/termbox-go"
)

// 簡単なタイマー
// 1秒おきに、時間を表示する
func main() {
	termbox.Init()
	game.Start()
}
