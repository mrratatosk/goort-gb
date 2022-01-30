package apu

import (
	"fmt"
	"sync"
)

type Apu struct {
}

func New() Apu {
	return Apu{}
}

func (apu Apu) ClockDivider() uint8 {
	return 2
}

func (apu Apu) Tick(wg *sync.WaitGroup) {
	fmt.Println("APU Tick")
	wg.Done()
}
