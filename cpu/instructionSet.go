package cpu

import (
	"github.com/mrratatosk/oort-framework/memory"
	"github.com/mrratatosk/oort-framework/processor"
	"github.com/mrratatosk/oort-framework/tools"
)

type CPU = processor.ProcessorUnit[uint16, uint8]
type MEM = memory.Memory[uint16, uint8]

func instructionSet() map[uint]processor.Instruction[uint16, uint8] {
	return map[uint]processor.Instruction[uint16, uint8]{
		0x00: newIns("NOP", 4, 0, func(pu CPU, m MEM, params ...uint8) {}),
		0x01: newIns("LD BC, nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", params[1])
			pu.Registers.Set8("C", params[0])
		}),
		0x02: newIns("LD (BC),A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Combine8("B", "C").Value, pu.Registers.Get8("A").Value)
		}),
		0x03: newIns("INC BC", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			b, c := tools.Split8(pu.Registers.Combine8("B", "C").Inc().Value)
			pu.Registers.Set8("B", b)
			pu.Registers.Set8("C", c)
		}),
		0x04: newIns("INC B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.AddR8Val("B", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x05: newIns("DEC B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.SubR8Val("B", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x06: newIns("LD B,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", params[0])
		}),
		0x08: newIns("LD (nn),SP", 20, 1, func(pu CPU, m MEM, params ...uint8) {
			s, p := pu.Registers.SplitRL16("SP")
			addr := tools.Combine8(params[1], params[0])
			m.Write(addr, p)
			addr++
			m.Write(addr, s)
		}),
		0x09: newIns("ADD HL,BC", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			r, c, hc, _ := tools.Add16(pu.Registers.Combine8("H", "L").Value, pu.Registers.Combine8("B", "C").Value)
			h, l := tools.Split8(r)
			pu.Registers.Set8("H", h)
			pu.Registers.Set8("L", l)
			setFlags(pu, Flag{C, c}, Flag{N, false}, Flag{H, hc})
		}),
		0x10: newIns("STOP", 4, 1, func(pu CPU, m MEM, params ...uint8) {}),
		0x0A: newIns("LD A,(BC)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(pu.Registers.Combine8("B", "C").Value))
		}),
		0x0B: newIns("DEC BC", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			b, c := tools.Split8(pu.Registers.Combine8("B", "C").Dec().Value)
			pu.Registers.Set8("B", b)
			pu.Registers.Set8("C", c)
		}),
		0x0C: newIns("INC C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.AddR8Val("C", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x0D: newIns("DEC C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.SubR8Val("C", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x0E: newIns("LD C,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", params[0])
		}),
		0x11: newIns("LD DE, nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", params[1])
			pu.Registers.Set8("E", params[0])
		}),
		0x12: newIns("LD (DE),A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Combine8("D", "E").Value, pu.Registers.Get8("A").Value)
		}),
		0x13: newIns("INC DE", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			d, e := tools.Split8(pu.Registers.Combine8("D", "E").Inc().Value)
			pu.Registers.Set8("D", d)
			pu.Registers.Set8("E", e)
		}),
		0x14: newIns("INC D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.AddR8Val("D", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x15: newIns("DEC D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.SubR8Val("D", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x16: newIns("LD D,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", params[0])
		}),
		0x18: newIns("JP n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			steps := int8(params[0])
			if steps > 0 {
				pu.Registers.Get16("PC").Add(uint16(steps))
			} else {
				pu.Registers.Get16("PC").Sub(uint16(-steps))
			}
		}),
		0x19: newIns("ADD HL,DE", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			r, c, hc, _ := tools.Add16(pu.Registers.Combine8("H", "L").Value, pu.Registers.Combine8("D", "E").Value)
			h, l := tools.Split8(r)
			pu.Registers.Set8("H", h)
			pu.Registers.Set8("L", l)
			setFlags(pu, Flag{C, c}, Flag{N, false}, Flag{H, hc})
		}),
		0x1A: newIns("LD A,(DE)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(pu.Registers.Combine8("D", "E").Value))
		}),
		0x1B: newIns("DEC DE", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			d, e := tools.Split8(pu.Registers.Combine8("D", "E").Dec().Value)
			pu.Registers.Set8("D", d)
			pu.Registers.Set8("E", e)
		}),
		0x1C: newIns("INC E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.AddR8Val("E", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x1D: newIns("DEC E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.SubR8Val("E", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x1E: newIns("LD E,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", params[0])
		}),
		0x20: newIns("JR NZ,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, Z) {
				steps := int8(params[0])
				if steps > 0 {
					pu.Registers.Get16("PC").Add(uint16(steps))
				} else {
					pu.Registers.Get16("PC").Sub(uint16(-steps))
				}
			}
		}),
		0x21: newIns("LD HL, nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", params[1])
			pu.Registers.Set8("L", params[0])
		}),
		0x22: newIns("LDI (HL),A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Combine8("H", "L").Value, pu.Registers.Get8("A").Value)
			pu.Registers.SplitR16ToRL8(pu.Registers.Combine8("H", "L").Inc(), "H", "L")
		}),
		0x23: newIns("INC HL", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Combine8("H", "L").Inc().Value)
			pu.Registers.Set8("H", h)
			pu.Registers.Set8("L", l)
		}),
		0x24: newIns("INC H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.AddR8Val("H", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x25: newIns("DEC H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.SubR8Val("H", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x26: newIns("LD H,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", params[0])
		}),
		0x27: newIns("DAA", 4, 1, func(pu CPU, m MEM, params ...uint8) {
			n := getFlag(pu, N)
			cy := getFlag(pu, C)
			h := getFlag(pu, H)
			a := pu.Registers.Get8("A").Value

			if n {
				if cy {
					pu.Registers.Set8("A", a-0x60)
				}
				if h {
					pu.Registers.Set8("A", a-0x06)
				}
			} else {
				if cy || (a&0xFF) > 0x99 {
					pu.Registers.Set8("A", a+0x60)
					setFlags(pu, Flag{C, true})
				}
				if h || (a&0x0F) > 0x09 {
					pu.Registers.Set8("A", a+0x06)
				}
			}

			setFlags(pu, Flag{Z, pu.Registers.Get8("A").Value == 0}, Flag{H, false})
		}),
		0x28: newIns("JR Z,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, Z) {
				steps := int8(params[0])
				if steps > 0 {
					pu.Registers.Get16("PC").Add(uint16(steps))
				} else {
					pu.Registers.Get16("PC").Sub(uint16(-steps))
				}
			}
		}),
		0x29: newIns("ADD HL,HL", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			r, c, hc, _ := tools.Add16(pu.Registers.Combine8("H", "L").Value, pu.Registers.Combine8("H", "L").Value)
			h, l := tools.Split8(r)
			pu.Registers.Set8("H", h)
			pu.Registers.Set8("L", l)
			setFlags(pu, Flag{C, c}, Flag{N, false}, Flag{H, hc})
		}),
		0x2A: newIns("LDI A,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			pu.Registers.SplitR16ToRL8(pu.Registers.Combine8("H", "L").Inc(), "H", "L")
		}),
		0x2B: newIns("DEC HL", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Combine8("H", "L").Dec().Value)
			pu.Registers.Set8("H", h)
			pu.Registers.Set8("L", l)
		}),
		0x2C: newIns("INC L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.AddR8Val("L", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x2D: newIns("DEC L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.SubR8Val("L", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x2E: newIns("LD L,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", params[0])
		}),
		0x2F: newIns("CPL A", 4, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.NotR8("A")
			setFlags(pu, Flag{N, true}, Flag{H, true})
		}),
		0x30: newIns("JR NC,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, C) {
				steps := int8(params[0])
				if steps > 0 {
					pu.Registers.Get16("PC").Add(uint16(steps))
				} else {
					pu.Registers.Get16("PC").Sub(uint16(-steps))
				}
			}
		}),
		0x31: newIns("LD SP, nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set16("SP", tools.Combine8(params[1], params[0]))
		}),
		0x32: newIns("LDD (HL),A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Combine8("H", "L").Value, pu.Registers.Get8("A").Value)
			pu.Registers.SplitR16ToRL8(pu.Registers.Combine8("H", "L").Dec(), "H", "L")
		}),
		0x33: newIns("INC SP", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Get16("SP").Inc()
		}),
		0x34: newIns("INC (HL)", 12, 0, func(pu CPU, m MEM, params ...uint8) {
			addr := pu.Registers.Combine8("H", "L").Value
			r, _, hc, z := tools.Add8(m.Read(addr), 1)
			m.Write(addr, r)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x35: newIns("DEC (HL)", 12, 0, func(pu CPU, m MEM, params ...uint8) {
			addr := pu.Registers.Combine8("H", "L").Value
			r, _, hc, z := tools.Sub8(m.Read(addr), 1)
			m.Write(addr, r)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x36: newIns("LD (HL),n", 12, 1, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, params[0])
		}),
		0x37: newIns("SCF", 4, 1, func(pu CPU, m MEM, params ...uint8) {
			setFlags(pu, Flag{N, false}, Flag{H, false}, Flag{C, true})
		}),
		0x38: newIns("JR C,n", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, C) {
				steps := int8(params[0])
				if steps > 0 {
					pu.Registers.Get16("PC").Add(uint16(steps))
				} else {
					pu.Registers.Get16("PC").Sub(uint16(-steps))
				}
			}
		}),
		0x39: newIns("ADD HL,SP", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			r, c, hc, _ := tools.Add16(pu.Registers.Combine8("H", "L").Value, pu.Registers.Get16("SP").Value)
			h, l := tools.Split8(r)
			pu.Registers.Set8("H", h)
			pu.Registers.Set8("L", l)
			setFlags(pu, Flag{C, c}, Flag{N, false}, Flag{H, hc})
		}),
		0x3A: newIns("LDD A,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			pu.Registers.SplitR16ToRL8(pu.Registers.Combine8("H", "L").Dec(), "H", "L")
		}),
		0x3B: newIns("DEC SP", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Get16("SP").Dec()
		}),
		0x3C: newIns("INC A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.AddR8Val("A", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, false}, Flag{H, hc})
		}),
		0x3D: newIns("DEC A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			_, hc, z := pu.Registers.SubR8Val("A", 1)
			setFlags(pu, Flag{Z, z}, Flag{N, true}, Flag{H, hc})
		}),
		0x3E: newIns("LD A,#", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", params[0])
		}),
		0x3F: newIns("CCF", 4, 1, func(pu CPU, m MEM, params ...uint8) {
			cy := getFlag(pu, C)
			setFlags(pu, Flag{N, false}, Flag{H, false}, Flag{C, !cy})
		}),
		0x40: newIns("LD B,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", pu.Registers.Get8("B").Value)
		}),
		0x41: newIns("LD B,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", pu.Registers.Get8("C").Value)
		}),
		0x42: newIns("LD B,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", pu.Registers.Get8("D").Value)
		}),
		0x43: newIns("LD B,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", pu.Registers.Get8("E").Value)
		}),
		0x44: newIns("LD B,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", pu.Registers.Get8("H").Value)
		}),
		0x45: newIns("LD B,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", pu.Registers.Get8("L").Value)
		}),
		0x46: newIns("LD B,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", m.Read(pu.Registers.Get16("HL").Value))
		}),
		0x47: newIns("LD B,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("B", pu.Registers.Get8("A").Value)
		}),
		0x48: newIns("LD C,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", pu.Registers.Get8("B").Value)
		}),
		0x49: newIns("LD C,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", pu.Registers.Get8("C").Value)
		}),
		0x4A: newIns("LD C,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", pu.Registers.Get8("D").Value)
		}),
		0x4B: newIns("LD C,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", pu.Registers.Get8("E").Value)
		}),
		0x4C: newIns("LD C,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", pu.Registers.Get8("H").Value)
		}),
		0x4D: newIns("LD C,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", pu.Registers.Get8("L").Value)
		}),
		0x4E: newIns("LD C,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", m.Read(pu.Registers.Get16("HL").Value))
		}),
		0x4F: newIns("LD C,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", pu.Registers.Get8("A").Value)
		}),
		0x50: newIns("LD D,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", pu.Registers.Get8("B").Value)
		}),
		0x51: newIns("LD D,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", pu.Registers.Get8("C").Value)
		}),
		0x52: newIns("LD D,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", pu.Registers.Get8("D").Value)
		}),
		0x53: newIns("LD D,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", pu.Registers.Get8("E").Value)
		}),
		0x54: newIns("LD D,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", pu.Registers.Get8("H").Value)
		}),
		0x55: newIns("LD D,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", pu.Registers.Get8("L").Value)
		}),
		0x56: newIns("LD D,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", m.Read(pu.Registers.Get16("HL").Value))
		}),
		0x57: newIns("LD D,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("D", pu.Registers.Get8("A").Value)
		}),
		0x58: newIns("LD E,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", pu.Registers.Get8("B").Value)
		}),
		0x59: newIns("LD E,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", pu.Registers.Get8("C").Value)
		}),
		0x5A: newIns("LD E,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", pu.Registers.Get8("D").Value)
		}),
		0x5B: newIns("LD E,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", pu.Registers.Get8("E").Value)
		}),
		0x5C: newIns("LD E,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", pu.Registers.Get8("H").Value)
		}),
		0x5D: newIns("LD E,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", pu.Registers.Get8("L").Value)
		}),
		0x5E: newIns("LD E,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", m.Read(pu.Registers.Get16("HL").Value))
		}),
		0x5F: newIns("LD E,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", pu.Registers.Get8("A").Value)
		}),
		0x60: newIns("LD H,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", pu.Registers.Get8("B").Value)
		}),
		0x61: newIns("LD H,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", pu.Registers.Get8("C").Value)
		}),
		0x62: newIns("LD H,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", pu.Registers.Get8("D").Value)
		}),
		0x63: newIns("LD H,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", pu.Registers.Get8("E").Value)
		}),
		0x64: newIns("LD H,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", pu.Registers.Get8("H").Value)
		}),
		0x65: newIns("LD H,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", pu.Registers.Get8("L").Value)
		}),
		0x66: newIns("LD H,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", m.Read(pu.Registers.Get16("HL").Value))
		}),
		0x67: newIns("LD H,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("H", pu.Registers.Get8("A").Value)
		}),
		0x68: newIns("LD L,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", pu.Registers.Get8("B").Value)
		}),
		0x69: newIns("LD L,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", pu.Registers.Get8("C").Value)
		}),
		0x6A: newIns("LD L,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", pu.Registers.Get8("D").Value)
		}),
		0x6B: newIns("LD L,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", pu.Registers.Get8("E").Value)
		}),
		0x6C: newIns("LD L,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", pu.Registers.Get8("H").Value)
		}),
		0x6D: newIns("LD L,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", pu.Registers.Get8("L").Value)
		}),
		0x6E: newIns("LD L,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", m.Read(pu.Registers.Get16("HL").Value))
		}),
		0x6F: newIns("LD L,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", pu.Registers.Get8("A").Value)
		}),
		0x70: newIns("LD (HL),B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, pu.Registers.Get8("B").Value)
		}),
		0x71: newIns("LD (HL),C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, pu.Registers.Get8("C").Value)
		}),
		0x72: newIns("LD (HL),D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, pu.Registers.Get8("D").Value)
		}),
		0x73: newIns("LD (HL),E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, pu.Registers.Get8("E").Value)
		}),
		0x74: newIns("LD (HL),H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, pu.Registers.Get8("H").Value)
		}),
		0x75: newIns("LD (HL),L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, pu.Registers.Get8("L").Value)
		}),
		0x76: newIns("HALT", 4, 0, func(pu CPU, m MEM, params ...uint8) {}),
		0x77: newIns("LD (HL),A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("HL").Value, pu.Registers.Get8("A").Value)
		}),
		0x78: newIns("LD A,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", pu.Registers.Get8("B").Value)
		}),
		0x79: newIns("LD A,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", pu.Registers.Get8("C").Value)
		}),
		0x7A: newIns("LD A,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", pu.Registers.Get8("D").Value)
		}),
		0x7B: newIns("LD A,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", pu.Registers.Get8("E").Value)
		}),
		0x7C: newIns("LD A,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", pu.Registers.Get8("H").Value)
		}),
		0x7D: newIns("LD A,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", pu.Registers.Get8("L").Value)
		}),
		0x7E: newIns("LD A,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(pu.Registers.Get16("HL").Value))
		}),
		0x7F: newIns("LD A,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Get8("A").Set(pu.Registers.Get8("A").Value)
		}),
		0x80: newIns("ADD A,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "B")
			setAllFlags(pu, c, hc, z, false)
		}),
		0x81: newIns("ADD A,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "C")
			setAllFlags(pu, c, hc, z, false)
		}),
		0x82: newIns("ADD A,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "D")
			setAllFlags(pu, c, hc, z, false)
		}),
		0x83: newIns("ADD A,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "E")
			setAllFlags(pu, c, hc, z, false)
		}),
		0x84: newIns("ADD A,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "H")
			setAllFlags(pu, c, hc, z, false)
		}),
		0x85: newIns("ADD A,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "L")
			setAllFlags(pu, c, hc, z, false)
		}),
		0x86: newIns("ADD A,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			setAllFlags(pu, c, hc, z, false)
		}),
		0x87: newIns("ADD A,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "A")
			setAllFlags(pu, c, hc, z, false)
		}),
		0x88: newIns("ADC A,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "B")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x89: newIns("ADC A,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "C")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x8A: newIns("ADC A,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "D")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x8B: newIns("ADC A,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "E")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x8C: newIns("ADC A,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "H")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x8D: newIns("ADC A,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "L")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x8E: newIns("ADC A,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x8F: newIns("ADC A,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8R8("A", "A")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0x90: newIns("SUB A,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "B")
			setAllFlags(pu, c, hc, z, true)
		}),
		0x91: newIns("SUB A,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "C")
			setAllFlags(pu, c, hc, z, true)
		}),
		0x92: newIns("SUB A,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "D")
			setAllFlags(pu, c, hc, z, true)
		}),
		0x93: newIns("SUB A,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "E")
			setAllFlags(pu, c, hc, z, true)
		}),
		0x94: newIns("SUB A,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "H")
			setAllFlags(pu, c, hc, z, true)
		}),
		0x95: newIns("SUB A,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "L")
			setAllFlags(pu, c, hc, z, true)
		}),
		0x96: newIns("SUB A,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			setAllFlags(pu, c, hc, z, true)
		}),
		0x97: newIns("SUB A,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "A")
			setAllFlags(pu, c, hc, z, true)
		}),
		0x98: newIns("SBC A,B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "B")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0x99: newIns("SBC A,C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "C")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0x9A: newIns("SBC A,D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "D")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0x9B: newIns("SBC A,E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "E")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0x9C: newIns("SBC A,H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "H")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0x9D: newIns("SBC A,L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "L")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0x9E: newIns("SBC A,(HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0x9F: newIns("SBC A,A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8R8("A", "A")
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0xA0: newIns("AND B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8R8("A", "B")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA1: newIns("AND C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8R8("A", "C")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA2: newIns("AND D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8R8("A", "D")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA3: newIns("AND E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8R8("A", "E")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA4: newIns("AND H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8R8("A", "H")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA5: newIns("AND L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8R8("A", "L")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA6: newIns("AND (HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA7: newIns("AND A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8R8("A", "A")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA8: newIns("XOR B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8R8("A", "B")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xA9: newIns("XOR C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8R8("A", "C")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xAA: newIns("XOR D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8R8("A", "D")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xAB: newIns("XOR E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8R8("A", "E")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xAC: newIns("XOR H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8R8("A", "H")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xAD: newIns("XOR L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8R8("A", "L")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xAE: newIns("XOR (HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			setAllFlags(pu, c, hc, z, false)
		}),
		0xAF: newIns("XOR A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8R8("A", "A")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB0: newIns("OR B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8R8("A", "B")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB1: newIns("OR C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8R8("A", "C")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB2: newIns("OR D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8R8("A", "D")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB3: newIns("OR E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8R8("A", "E")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB4: newIns("OR H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8R8("A", "H")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB5: newIns("OR L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8R8("A", "L")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB6: newIns("OR (HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB7: newIns("OR A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8R8("A", "A")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB8: newIns("CP B", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8R8("A", "B")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xB9: newIns("CP C", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8R8("A", "C")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xBA: newIns("CP D", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8R8("A", "D")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xBB: newIns("CP E", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8R8("A", "E")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xBC: newIns("CP H", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8R8("A", "H")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xBD: newIns("CP L", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8R8("A", "L")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xBE: newIns("CP (HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8Val("A", m.Read(pu.Registers.Combine8("H", "L").Value))
			setAllFlags(pu, c, hc, z, false)
		}),
		0xBF: newIns("CP A", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8R8("A", "A")
			setAllFlags(pu, c, hc, z, false)
		}),
		0xC0: newIns("RET NZ", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, Z) {
				l := m.Read(pu.Registers.Get16("SP").Value)
				h := m.Read(pu.Registers.Get16("SP").Inc().Value)
				pu.Registers.Get16("SP").Inc()
				pu.Registers.Set16("PC", tools.Combine8(h, l))
			}
		}),
		0xC1: newIns("POP BC", 12, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("C", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
			pu.Registers.Set8("B", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
		}),
		0xC2: newIns("JP NZ,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, Z) {
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xC3: newIns("JP nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
		}),
		0xC4: newIns("CALL NZ,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, Z) {
				h, l := tools.Split8(pu.Registers.Get16("PC").Value)
				m.Write(pu.Registers.Get16("SP").Dec().Value, h)
				m.Write(pu.Registers.Get16("SP").Dec().Value, l)
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xC5: newIns("PUSH BC", 16, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("C").Value)
			pu.Registers.Get16("SP").Dec()
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("B").Value)
			pu.Registers.Get16("SP").Dec()
		}),
		0xC6: newIns("ADD A,#", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8Val("A", params[0])
			setAllFlags(pu, c, hc, z, false)
		}),
		0xC7: newIns("RST $00", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x00)
		}),
		0xC8: newIns("RET Z", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, Z) {
				l := m.Read(pu.Registers.Get16("SP").Value)
				h := m.Read(pu.Registers.Get16("SP").Inc().Value)
				pu.Registers.Get16("SP").Inc()
				pu.Registers.Set16("PC", tools.Combine8(h, l))
			}
		}),
		0xC9: newIns("RET", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			l := m.Read(pu.Registers.Get16("SP").Value)
			h := m.Read(pu.Registers.Get16("SP").Inc().Value)
			pu.Registers.Get16("SP").Inc()
			pu.Registers.Set16("PC", tools.Combine8(h, l))
		}),
		0xCA: newIns("JP Z,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, Z) {
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xCC: newIns("CALL Z,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, Z) {
				h, l := tools.Split8(pu.Registers.Get16("PC").Value + 1)
				m.Write(pu.Registers.Get16("SP").Dec().Value, h)
				m.Write(pu.Registers.Get16("SP").Dec().Value, l)
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xCD: newIns("CALL nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
		}),
		0xCE: newIns("ADC A,#", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8Val("A", params[0])
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, false)
		}),
		0xCF: newIns("RST $08", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x08)
		}),
		0xD0: newIns("RET NC", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, C) {
				l := m.Read(pu.Registers.Get16("SP").Value)
				h := m.Read(pu.Registers.Get16("SP").Inc().Value)
				pu.Registers.Get16("SP").Inc()
				pu.Registers.Set16("PC", tools.Combine8(h, l))
			}
		}),
		0xD1: newIns("POP DE", 12, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("E", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
			pu.Registers.Set8("D", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
		}),
		0xD2: newIns("JP NC,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, C) {
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xD4: newIns("CALL NC,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if !getFlag(pu, C) {
				h, l := tools.Split8(pu.Registers.Get16("PC").Value)
				m.Write(pu.Registers.Get16("SP").Dec().Value, h)
				m.Write(pu.Registers.Get16("SP").Dec().Value, l)
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xD5: newIns("PUSH DE", 16, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("E").Value)
			pu.Registers.Get16("SP").Dec()
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("D").Value)
			pu.Registers.Get16("SP").Dec()
		}),
		0xD6: newIns("SUB A,#", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AddR8Val("A", params[0])
			setAllFlags(pu, c, hc, z, true)
		}),
		0xD7: newIns("RST $10", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x10)
		}),
		0xD8: newIns("RET C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, C) {
				l := m.Read(pu.Registers.Get16("SP").Value)
				h := m.Read(pu.Registers.Get16("SP").Inc().Value)
				pu.Registers.Get16("SP").Inc()
				pu.Registers.Set16("PC", tools.Combine8(h, l))
			}
		}),
		0xD9: newIns("RETI", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			l := m.Read(pu.Registers.Get16("SP").Value)
			h := m.Read(pu.Registers.Get16("SP").Inc().Value)
			pu.Registers.Get16("SP").Inc()
			pu.Registers.Set16("PC", tools.Combine8(h, l))
		}),
		0xDA: newIns("JP C,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, C) {
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xDC: newIns("CALL C,nn", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			if getFlag(pu, C) {
				h, l := tools.Split8(pu.Registers.Get16("PC").Value)
				m.Write(pu.Registers.Get16("SP").Dec().Value, h)
				m.Write(pu.Registers.Get16("SP").Dec().Value, l)
				pu.Registers.Set16("PC", tools.Combine8(params[1], params[0]))
			}
		}),
		0xDE: newIns("SBC A,#", 8, 1, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.SubR8Val("A", params[0])
			if c {
				c, hc, z = pu.Registers.AddR8Val("A", 1)
			}
			setAllFlags(pu, c, hc, z, true)
		}),
		0xDF: newIns("RST $18", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x18)
		}),
		0xE0: newIns("LDH (n),A", 12, 1, func(pu CPU, m MEM, params ...uint8) {
			m.Write(0xFF00+uint16(params[0]), pu.Registers.Get8("A").Value)
		}),
		0xE1: newIns("POP HL", 12, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("L", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
			pu.Registers.Set8("H", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
		}),
		0xE2: newIns("LD (C),A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(0xFF00+uint16(pu.Registers.Get8("C").Value), pu.Registers.Get8("A").Value)
		}),
		0xE5: newIns("PUSH HL", 16, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("H").Value)
			pu.Registers.Get16("SP").Dec()
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("L").Value)
			pu.Registers.Get16("SP").Dec()
		}),
		0xE6: newIns("AND A,#", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.AndR8Val("A", params[0])
			setAllFlags(pu, c, hc, z, false)
		}),
		0xE7: newIns("RST $20", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x20)
		}),
		0xE8: newIns("ADD SP,#", 16, 0, func(pu CPU, m MEM, params ...uint8) {
			r, c, hc, _ := tools.Add16(pu.Registers.Get16("SP").Value, uint16(params[0]))
			pu.Registers.Set16("SP", r)
			setAllFlags(pu, c, hc, false, false)
		}),
		0xE9: newIns("JP (HL)", 4, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set16("PC", pu.Registers.Combine8("H", "L").Value)
		}),
		0xEA: newIns("LD (nn),A", 16, 2, func(pu CPU, m MEM, params ...uint8) {
			m.Write(tools.Combine8(params[1], params[0]), pu.Registers.Get8("A").Value)
		}),
		0xEE: newIns("XOR A,#", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.XorR8Val("A", params[0])
			setAllFlags(pu, c, hc, z, false)
		}),
		0xEF: newIns("RST $28", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x28)
		}),
		0xF0: newIns("LDH A,(n)", 12, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(0xFF00+uint16(params[0])))
		}),
		0xF1: newIns("POP AF", 12, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("F", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
			pu.Registers.Set8("A", m.Read(pu.Registers.Get16("SP").Value))
			pu.Registers.Get16("SP").Inc()
		}),
		0xF2: newIns("LD A,(C)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(0xFF00+uint16(pu.Registers.Get8("C").Value)))
		}),
		0xF3: newIns("DI", 4, 0, func(pu CPU, m MEM, params ...uint8) {

		}),
		0xF5: newIns("PUSH AF", 16, 0, func(pu CPU, m MEM, params ...uint8) {
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("A").Value)
			pu.Registers.Get16("SP").Dec()
			m.Write(pu.Registers.Get16("SP").Value, pu.Registers.Get8("F").Value)
			pu.Registers.Get16("SP").Dec()
		}),
		0xF6: newIns("OR A,#", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.OrR8Val("A", params[0])
			setAllFlags(pu, c, hc, z, false)
		}),
		0xF7: newIns("RST $30", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x30)
		}),
		0xF8: newIns("LD HL,SP+n", 12, 1, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Get16("SP").Add(uint16(params[0]))
			pu.Registers.SplitRL16ToRL8("SP", "H", "L")
		}),
		0xF9: newIns("LD SP,HL", 12, 2, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set16("SP", pu.Registers.Combine8("H", "L").Value)
		}),
		0xFA: newIns("LD A,(nn)", 16, 2, func(pu CPU, m MEM, params ...uint8) {
			pu.Registers.Set8("A", m.Read(tools.Combine8(params[1], params[0])))
		}),
		0xFB: newIns("DI", 4, 0, func(pu CPU, m MEM, params ...uint8) {

		}),
		0xFE: newIns("CP #", 8, 0, func(pu CPU, m MEM, params ...uint8) {
			c, hc, z := pu.Registers.CpR8Val("A", params[0])
			setAllFlags(pu, c, hc, z, false)
		}),
		0xFF: newIns("RST $38", 32, 0, func(pu CPU, m MEM, params ...uint8) {
			h, l := tools.Split8(pu.Registers.Get16("PC").Value)
			m.Write(pu.Registers.Get16("SP").Dec().Value, h)
			m.Write(pu.Registers.Get16("SP").Dec().Value, l)
			pu.Registers.Set16("PC", 0x38)
		}),
	}
}

func newIns(name string, cycle uint, params uint, fn func(CPU, MEM, ...uint8)) processor.Instruction[uint16, uint8] {
	return processor.Instruction[uint16, uint8]{
		Name:     name,
		Cycle:    cycle,
		Params:   params,
		Callback: fn,
	}
}
