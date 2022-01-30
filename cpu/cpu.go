package cpu

import (
	"sync"

	"github.com/mrratatosk/oort-framework/memory"
	"github.com/mrratatosk/oort-framework/processor"
)

type Cpu struct {
	mem          *MEM
	currentCycle uint
	processor.ProcessorUnit[uint16, uint8]
}

func New(mem *memory.Memory[uint16, uint8]) *Cpu {
	return &Cpu{
		mem:          mem,
		currentCycle: 0,
		ProcessorUnit: processor.ProcessorUnit[uint16, uint8]{
			InstructionSet: instructionSet(),
			ExtensionSet:   extensionSet(),
			Registers:      registers(),
		},
	}
}

func (c Cpu) ClockDivider() uint8 {
	return 4
}

func (c *Cpu) Tick(wg *sync.WaitGroup) {
	if c.currentCycle == 0 {
		opcode := c.fetch()
		ins, params := c.decode(opcode)
		c.currentCycle = c.execute(ins, params)
	} else {
		c.currentCycle--
	}

	wg.Done()
}

func (c Cpu) fetch() uint8 {
	return c.mem.Read(c.Registers.Get16("PC").Value)
}

func (c Cpu) decode(opCode uint8) (processor.Instruction[uint16, uint8], []uint8) {
	set := c.InstructionSet
	code := opCode
	if val, ok := c.ExtensionSet[uint(code)]; ok {
		set = val
		code = c.mem.Read(c.Registers.Get16("PC").Inc().Value)
	}

	nextIns := set[uint(code)]

	params := make([]uint8, nextIns.Params)
	p := uint(0)
	for p < nextIns.Params {
		params[0] = c.mem.Read(c.Registers.Get16("PC").Inc().Value)
		p++
	}

	c.Registers.Get16("PC").Inc()
	return nextIns, params
}

func (c Cpu) execute(ins processor.Instruction[uint16, uint8], params []uint8) uint {
	ins.Callback(c.ProcessorUnit, *c.mem, params...)
	return ins.Cycle / 4
}

type Flag struct {
	name  Flags
	value bool
}

func getFlag(c CPU, flag Flags) bool {
	return c.Registers.Get8("F").Bit(uint8(flag))
}

func setFlag(c CPU, flag Flags, value bool) {
	if value {
		c.Registers.Get8("F").BitSet(uint8(flag))
	} else {
		c.Registers.Get8("F").BitClear(uint8(flag))
	}
}

func setFlags(c CPU, flags ...Flag) {
	for _, flag := range flags {
		setFlag(c, flag.name, flag.value)
	}
}

func setAllFlags(c CPU, cy bool, hc bool, z bool, n bool) {
	setFlags(c, Flag{C, cy}, Flag{H, hc}, Flag{Z, z}, Flag{N, n})
}
