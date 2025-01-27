package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type MeteorTimer struct {
	currentTicks int
	targetTicks  int
}

func NewTimer(d time.Duration) *MeteorTimer {
	return &MeteorTimer{
		currentTicks: 0,
		targetTicks:  int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *MeteorTimer) update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *MeteorTimer) reset() {
	t.currentTicks = 0
}

func (t *MeteorTimer) isReady() bool {
	return t.currentTicks >= t.targetTicks
}
