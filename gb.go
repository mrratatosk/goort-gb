package oortgb

import (
	"os"
	"sync"

	oortframework "github.com/mrratatosk/oort-framework"
	"github.com/mrratatosk/oort-framework/memory"
	"github.com/mrratatosk/oort-framework/processor"
	"github.com/mrratatosk/oort-framework/tools"
	"github.com/mrratatosk/oort-gb/cpu"
)

type GbEmulator struct {
	oortframework.Emulator[uint16, uint8]
}

func New(biosPath string) GbEmulator {
	mem := memory.NewMemory[uint16, uint8](0x10000)

	gb := GbEmulator{
		oortframework.Emulator[uint16, uint8]{
			Memory: mem,
			Bios:   biosPath,
			Units: []processor.ITicker{
				//ppu.New(),
				//apu.New(),
				cpu.New(mem),
			},
		},
	}

	return gb
}

func (gb GbEmulator) loadBios() {
	dat, err := os.ReadFile(gb.Bios)
	tools.Check(err)

	gb.Memory.WriteRange(0x0, dat)
}

func (gb GbEmulator) Start() {
	gb.loadBios()

	var wg sync.WaitGroup
	cycles, stop := uint8(0), uint8(0)

	for stop < 10 {
		if cycles >= 4 {
			cycles = 0
			stop++
		}

		for _, unit := range gb.Units {
			if cycles%unit.ClockDivider() == 0 {
				wg.Add(1)
				go unit.Tick(&wg)
			}
		}

		wg.Wait()
		cycles++
	}
}
