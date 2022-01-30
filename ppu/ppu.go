package ppu

import (
	"fmt"
	"sync"
)

type Ppu struct {
}

func New() Ppu {
	return Ppu{}
}

func (ppu Ppu) ClockDivider() uint8 {
	return 2
}

func (ppu Ppu) Tick(wg *sync.WaitGroup) {
	fmt.Println("PPU Tick")
	wg.Done()
}
