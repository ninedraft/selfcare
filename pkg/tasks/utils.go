package tasks

import (
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = newRnd()
}

func newRnd() *rand.Rand {
	var seed = time.Now().UnixNano()
	var source = rand.NewSource(seed)
	return rand.New(source)
}

func DefaultLabels() []string {
	return []string{
		"💚", "❤", "♥", "🧡", "💜", "💛", "💙", "💗",
		"💖", "✨", "✔", "🥇", "🍫", "🧁",
		"🍬", "🌼", "🌻", "🌞", "☀",
	}
}
