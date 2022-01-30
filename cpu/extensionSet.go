package cpu

import (
	"github.com/mrratatosk/oort-framework/processor"
	"github.com/mrratatosk/oort-framework/tools"
)

func extensionSet() map[uint]map[uint]processor.Instruction[uint16, uint8] {
	return map[uint]map[uint]processor.Instruction[uint16, uint8]{
		0xCB: {
			0x00: newIns("RLC B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("B")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x01: newIns("RLC C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("C")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x02: newIns("RLC D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("D")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x03: newIns("RLC E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("E")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x04: newIns("RLC H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("H")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x05: newIns("RLC L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("L")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x06: newIns("RLC (HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				r, cy := tools.RotateL8(m.Read(addr), 1)
				m.Write(addr, r)
				setAllFlags(pu, cy, false, r == 0, false)
			}),
			0x07: newIns("RLC A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("A")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x08: newIns("RRC B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.RotateRR8("B")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x09: newIns("RRC C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.RotateRR8("C")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x0A: newIns("RRC D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.RotateRR8("D")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x0B: newIns("RRC E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.RotateRR8("E")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x0C: newIns("RRC H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.RotateRR8("H")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x0D: newIns("RRC L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateRR8("L")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x0E: newIns("RRC (HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				r, cy := tools.RotateR8(m.Read(addr), 1)
				m.Write(addr, r)
				setAllFlags(pu, cy, false, r == 0, false)
			}),
			0x0F: newIns("RRC A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.RotateRR8("A")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x10: newIns("RL B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				msb, z := pu.Registers.RotateLR8("B")
				if cy {
					pu.Registers.Get8("B").BitSet(0)
				} else {
					pu.Registers.Get8("B").BitClear(0)
				}
				setAllFlags(pu, msb, false, z, false)
			}),
			0x11: newIns("RL C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				msb, z := pu.Registers.RotateLR8("C")
				if cy {
					pu.Registers.Get8("C").BitSet(0)
				} else {
					pu.Registers.Get8("C").BitClear(0)
				}
				setAllFlags(pu, msb, false, z, false)
			}),
			0x12: newIns("RL D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				msb, z := pu.Registers.RotateLR8("D")
				if cy {
					pu.Registers.Get8("D").BitSet(0)
				} else {
					pu.Registers.Get8("D").BitClear(0)
				}
				setAllFlags(pu, msb, false, z, false)
			}),
			0x13: newIns("RL E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				msb, z := pu.Registers.RotateLR8("E")
				if cy {
					pu.Registers.Get8("E").BitSet(0)
				} else {
					pu.Registers.Get8("E").BitClear(0)
				}
				setAllFlags(pu, msb, false, z, false)
			}),
			0x14: newIns("RL H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				msb, z := pu.Registers.RotateLR8("H")
				if cy {
					pu.Registers.Get8("H").BitSet(0)
				} else {
					pu.Registers.Get8("H").BitClear(0)
				}
				setAllFlags(pu, msb, false, z, false)
			}),
			0x15: newIns("RL L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				msb, z := pu.Registers.RotateLR8("L")
				if cy {
					pu.Registers.Get8("L").BitSet(0)
				} else {
					pu.Registers.Get8("L").BitClear(0)
				}
				setAllFlags(pu, msb, false, z, false)
			}),
			0x16: newIns("RL (HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				cy := getFlag(pu, C)
				r, msb := tools.RotateL8(m.Read(addr), 1)

				if cy {
					r = tools.Set8(r, 0)
				} else {
					r = tools.Clear8(r, 0)
				}

				m.Write(addr, r)
				setAllFlags(pu, msb, false, r == 0, false)
			}),
			0x17: newIns("RL A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.RotateLR8("A")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x18: newIns("RR B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				lsb, z := pu.Registers.RotateRR8("B")
				if cy {
					pu.Registers.Get8("B").BitSet(7)
				} else {
					pu.Registers.Get8("B").BitClear(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x19: newIns("RR C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				lsb, z := pu.Registers.RotateRR8("C")
				if cy {
					pu.Registers.Get8("C").BitSet(7)
				} else {
					pu.Registers.Get8("C").BitClear(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x1A: newIns("RR D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				lsb, z := pu.Registers.RotateRR8("D")
				if cy {
					pu.Registers.Get8("D").BitSet(7)
				} else {
					pu.Registers.Get8("D").BitClear(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x1B: newIns("RR E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				lsb, z := pu.Registers.RotateRR8("E")
				if cy {
					pu.Registers.Get8("E").BitSet(7)
				} else {
					pu.Registers.Get8("E").BitClear(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x1C: newIns("RR H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				lsb, z := pu.Registers.RotateRR8("H")
				if cy {
					pu.Registers.Get8("H").BitSet(7)
				} else {
					pu.Registers.Get8("H").BitClear(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x1D: newIns("RR L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				cy := getFlag(pu, C)
				lsb, z := pu.Registers.RotateRR8("L")
				if cy {
					pu.Registers.Get8("L").BitSet(7)
				} else {
					pu.Registers.Get8("L").BitClear(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x1E: newIns("RR (HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				cy := getFlag(pu, C)
				r, lsb := tools.RotateR8(m.Read(addr), 1)

				if cy {
					r = tools.Set8(r, 7)
				} else {
					r = tools.Clear8(r, 7)
				}

				m.Write(addr, r)
				setAllFlags(pu, lsb, false, r == 0, false)
			}),
			0x1F: newIns("RR A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.RotateRR8("A")
				cy := getFlag(pu, C)
				if cy {
					pu.Registers.Get8("A").BitSet(7)
				} else {
					pu.Registers.Get8("A").BitClear(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x20: newIns("SLA B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.ShiftLR8("B")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x21: newIns("SLA C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.ShiftLR8("C")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x22: newIns("SLA D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.ShiftLR8("D")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x23: newIns("SLA E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.ShiftLR8("E")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x24: newIns("SLA H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.ShiftLR8("H")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x25: newIns("SLA L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.ShiftLR8("L")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x26: newIns("SLA (HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				r, msb := tools.ShiftL8(m.Read(addr), 1)

				m.Write(addr, r)
				setAllFlags(pu, msb, false, r == 0, false)
			}),
			0x27: newIns("SLA A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb, z := pu.Registers.ShiftLR8("A")
				setAllFlags(pu, msb, false, z, false)
			}),
			0x28: newIns("SRA B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb := pu.Registers.Get8("B").Bit(7)
				lsb, z := pu.Registers.ShiftRR8("B")
				if msb {
					pu.Registers.Get8("B").BitSet(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x29: newIns("SRA C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb := pu.Registers.Get8("C").Bit(7)
				lsb, z := pu.Registers.ShiftRR8("C")
				if msb {
					pu.Registers.Get8("C").BitSet(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x2A: newIns("SRA D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb := pu.Registers.Get8("D").Bit(7)
				lsb, z := pu.Registers.ShiftRR8("D")
				if msb {
					pu.Registers.Get8("D").BitSet(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x2B: newIns("SRA E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb := pu.Registers.Get8("E").Bit(7)
				lsb, z := pu.Registers.ShiftRR8("E")
				if msb {
					pu.Registers.Get8("E").BitSet(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x2C: newIns("SRA H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb := pu.Registers.Get8("H").Bit(7)
				lsb, z := pu.Registers.ShiftRR8("H")
				if msb {
					pu.Registers.Get8("H").BitSet(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x2D: newIns("SRA L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb := pu.Registers.Get8("L").Bit(7)
				lsb, z := pu.Registers.ShiftRR8("L")
				if msb {
					pu.Registers.Get8("L").BitSet(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x2E: newIns("SRA (HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				val := m.Read(addr)
				msb := tools.Bit8(val, 7)
				r, lsb := tools.ShiftR8(val, 1)

				if msb {
					r = tools.Set8(r, 7)
				}

				m.Write(addr, r)
				setAllFlags(pu, lsb, false, r == 0, false)
			}),
			0x2F: newIns("SRA A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				msb := pu.Registers.Get8("A").Bit(7)
				lsb, z := pu.Registers.ShiftRR8("A")
				if msb {
					pu.Registers.Get8("A").BitSet(7)
				}
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x30: newIns("SWAP B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				z := pu.Registers.SwapR8("B")
				setAllFlags(pu, false, false, z, false)
			}),
			0x31: newIns("SWAP C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				z := pu.Registers.SwapR8("C")
				setAllFlags(pu, false, false, z, false)
			}),
			0x32: newIns("SWAP D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				z := pu.Registers.SwapR8("D")
				setAllFlags(pu, false, false, z, false)
			}),
			0x33: newIns("SWAP E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				z := pu.Registers.SwapR8("E")
				setAllFlags(pu, false, false, z, false)
			}),
			0x34: newIns("SWAP H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				z := pu.Registers.SwapR8("H")
				setAllFlags(pu, false, false, z, false)
			}),
			0x35: newIns("SWAP L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				z := pu.Registers.SwapR8("L")
				setAllFlags(pu, false, false, z, false)
			}),
			0x36: newIns("SWAP (HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				r := tools.Swap8(m.Read(addr))
				m.Write(addr, r)
				setAllFlags(pu, false, false, r == 0, false)
			}),
			0x37: newIns("SWAP A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				z := pu.Registers.SwapR8("A")
				setAllFlags(pu, false, false, z, false)
			}),
			0x38: newIns("SRL B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.ShiftRR8("B")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x39: newIns("SRL C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.ShiftRR8("C")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x3A: newIns("SRL D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.ShiftRR8("D")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x3B: newIns("SRL E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.ShiftRR8("E")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x3C: newIns("SRL H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.ShiftRR8("H")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x3D: newIns("SRL L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.ShiftRR8("L")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x3E: newIns("SRL (HL)", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				val := m.Read(addr)
				r, lsb := tools.ShiftR8(val, 1)
				m.Write(addr, r)
				setAllFlags(pu, lsb, false, r == 0, false)
			}),
			0x3F: newIns("SRL A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				lsb, z := pu.Registers.ShiftRR8("A")
				setAllFlags(pu, lsb, false, z, false)
			}),
			0x40: newIns("BIT 0,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(0))})
			}),
			0x41: newIns("BIT 0,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(0))})
			}),
			0x42: newIns("BIT 0,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(0))})
			}),
			0x43: newIns("BIT 0,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(0))})
			}),
			0x44: newIns("BIT 0,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(0))})
			}),
			0x45: newIns("BIT 0,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(0))})
			}),
			0x46: newIns("BIT 0,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 0)})
			}),
			0x47: newIns("BIT 0,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(0))})
			}),
			0x48: newIns("BIT 1,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(1))})
			}),
			0x49: newIns("BIT 1,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(1))})
			}),
			0x4A: newIns("BIT 1,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(1))})
			}),
			0x4B: newIns("BIT 1,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(1))})
			}),
			0x4C: newIns("BIT 1,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(1))})
			}),
			0x4D: newIns("BIT 1,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(1))})
			}),
			0x4E: newIns("BIT 1,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 1)})
			}),
			0x4F: newIns("BIT 1,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(1))})
			}),
			0x50: newIns("BIT 2,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(2))})
			}),
			0x51: newIns("BIT 2,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(2))})
			}),
			0x52: newIns("BIT 2,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(2))})
			}),
			0x53: newIns("BIT 2,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(2))})
			}),
			0x54: newIns("BIT 2,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(2))})
			}),
			0x55: newIns("BIT 2,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(2))})
			}),
			0x56: newIns("BIT 2,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 2)})
			}),
			0x57: newIns("BIT 2,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(2))})
			}),
			0x58: newIns("BIT 3,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(3))})
			}),
			0x59: newIns("BIT 3,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(3))})
			}),
			0x5A: newIns("BIT 3,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(3))})
			}),
			0x5B: newIns("BIT 3,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(3))})
			}),
			0x5C: newIns("BIT 3,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(3))})
			}),
			0x5D: newIns("BIT 3,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(3))})
			}),
			0x5E: newIns("BIT 3,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 3)})
			}),
			0x5F: newIns("BIT 3,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(3))})
			}),
			0x60: newIns("BIT 4,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(4))})
			}),
			0x61: newIns("BIT 4,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(4))})
			}),
			0x62: newIns("BIT 4,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(4))})
			}),
			0x63: newIns("BIT 4,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(4))})
			}),
			0x64: newIns("BIT 4,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(4))})
			}),
			0x65: newIns("BIT 4,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(4))})
			}),
			0x66: newIns("BIT 4,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 4)})
			}),
			0x67: newIns("BIT 4,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(4))})
			}),
			0x68: newIns("BIT 5,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(5))})
			}),
			0x69: newIns("BIT 5,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(5))})
			}),
			0x6A: newIns("BIT 5,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(5))})
			}),
			0x6B: newIns("BIT 5,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(5))})
			}),
			0x6C: newIns("BIT 5,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(5))})
			}),
			0x6D: newIns("BIT 5,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(5))})
			}),
			0x6E: newIns("BIT 5,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 5)})
			}),
			0x6F: newIns("BIT 5,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(5))})
			}),
			0x70: newIns("BIT 6,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(6))})
			}),
			0x71: newIns("BIT 6,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(6))})
			}),
			0x72: newIns("BIT 6,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(6))})
			}),
			0x73: newIns("BIT 6,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(6))})
			}),
			0x74: newIns("BIT 6,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(6))})
			}),
			0x75: newIns("BIT 6,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(6))})
			}),
			0x76: newIns("BIT 6,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 6)})
			}),
			0x77: newIns("BIT 6,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(6))})
			}),
			0x78: newIns("BIT 7,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("B").Bit(7))})
			}),
			0x79: newIns("BIT 7,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("C").Bit(7))})
			}),
			0x7A: newIns("BIT 7,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("D").Bit(7))})
			}),
			0x7B: newIns("BIT 7,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("E").Bit(7))})
			}),
			0x7C: newIns("BIT 7,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("H").Bit(7))})
			}),
			0x7D: newIns("BIT 7,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("L").Bit(7))})
			}),
			0x7E: newIns("BIT 7,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !tools.Bit8(m.Read(pu.Registers.Combine8("H", "L").Value), 7)})
			}),
			0x7F: newIns("BIT 7,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				setFlags(pu, Flag{N, false}, Flag{H, true}, Flag{Z, !(pu.Registers.Get8("A").Bit(7))})
			}),
			0x80: newIns("RES 0,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(0)
			}),
			0x81: newIns("RES 0,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(0)
			}),
			0x82: newIns("RES 0,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(0)
			}),
			0x83: newIns("RES 0,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(0)
			}),
			0x84: newIns("RES 0,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(0)
			}),
			0x85: newIns("RES 0,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(0)
			}),
			0x86: newIns("RES 0,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 0))
			}),
			0x87: newIns("RES 0,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(0)
			}),
			0x88: newIns("RES 1,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(1)
			}),
			0x89: newIns("RES 1,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(1)
			}),
			0x8A: newIns("RES 1,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(1)
			}),
			0x8B: newIns("RES 1,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(1)
			}),
			0x8C: newIns("RES 1,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(1)
			}),
			0x8D: newIns("RES 1,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(1)
			}),
			0x8E: newIns("RES 1,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 1))
			}),
			0x8F: newIns("RES 1,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(1)
			}),
			0x90: newIns("RES 2,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(2)
			}),
			0x91: newIns("RES 2,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(2)
			}),
			0x92: newIns("RES 2,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(2)
			}),
			0x93: newIns("RES 2,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(2)
			}),
			0x94: newIns("RES 2,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(2)
			}),
			0x95: newIns("RES 2,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(2)
			}),
			0x96: newIns("RES 2,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 2))
			}),
			0x97: newIns("RES 2,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(2)
			}),
			0x98: newIns("RES 3,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(3)
			}),
			0x99: newIns("RES 3,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(3)
			}),
			0x9A: newIns("RES 3,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(3)
			}),
			0x9B: newIns("RES 3,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(3)
			}),
			0x9C: newIns("RES 3,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(3)
			}),
			0x9D: newIns("RES 3,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(3)
			}),
			0x9E: newIns("RES 3,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 3))
			}),
			0x9F: newIns("RES 3,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(3)
			}),
			0xA0: newIns("RES 4,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(4)
			}),
			0xA1: newIns("RES 4,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(4)
			}),
			0xA2: newIns("RES 4,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(4)
			}),
			0xA3: newIns("RES 4,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(4)
			}),
			0xA4: newIns("RES 4,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(4)
			}),
			0xA5: newIns("RES 4,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(4)
			}),
			0xA6: newIns("RES 4,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 4))
			}),
			0xA7: newIns("RES 4,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(4)
			}),
			0xA8: newIns("RES 5,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(5)
			}),
			0xA9: newIns("RES 5,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(5)
			}),
			0xAA: newIns("RES 5,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(5)
			}),
			0xAB: newIns("RES 5,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(5)
			}),
			0xAC: newIns("RES 5,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(5)
			}),
			0xAD: newIns("RES 5,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(5)
			}),
			0xAE: newIns("RES 5,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 5))
			}),
			0xAF: newIns("RES 5,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(5)
			}),
			0xB0: newIns("RES 6,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(6)
			}),
			0xB1: newIns("RES 6,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(6)
			}),
			0xB2: newIns("RES 6,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(6)
			}),
			0xB3: newIns("RES 6,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(6)
			}),
			0xB4: newIns("RES 6,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(6)
			}),
			0xB5: newIns("RES 6,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(6)
			}),
			0xB6: newIns("RES 6,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 6))
			}),
			0xB7: newIns("RES 6,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(6)
			}),
			0xB8: newIns("RES 7,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitClear(7)
			}),
			0xB9: newIns("RES 7,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitClear(7)
			}),
			0xBA: newIns("RES 7,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitClear(7)
			}),
			0xBB: newIns("RES 7,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitClear(7)
			}),
			0xBC: newIns("RES 7,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitClear(7)
			}),
			0xBD: newIns("RES 7,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitClear(7)
			}),
			0xBE: newIns("RES 7,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Clear8(m.Read(addr), 7))
			}),
			0xBF: newIns("RES 7,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitClear(7)
			}),
			0xC0: newIns("SET 0,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(0)
			}),
			0xC1: newIns("SET 0,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(0)
			}),
			0xC2: newIns("SET 0,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(0)
			}),
			0xC3: newIns("SET 0,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(0)
			}),
			0xC4: newIns("SET 0,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(0)
			}),
			0xC5: newIns("SET 0,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(0)
			}),
			0xC6: newIns("SET 0,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 0))
			}),
			0xC7: newIns("SET 0,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(0)
			}),
			0xC8: newIns("SET 1,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(1)
			}),
			0xC9: newIns("SET 1,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(1)
			}),
			0xCA: newIns("SET 1,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(1)
			}),
			0xCB: newIns("SET 1,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(1)
			}),
			0xCC: newIns("SET 1,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(1)
			}),
			0xCD: newIns("SET 1,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(1)
			}),
			0xCE: newIns("SET 1,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 1))
			}),
			0xCF: newIns("SET 1,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(1)
			}),
			0xD0: newIns("SET 2,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(2)
			}),
			0xD1: newIns("SET 2,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(2)
			}),
			0xD2: newIns("SET 2,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(2)
			}),
			0xD3: newIns("SET 2,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(2)
			}),
			0xD4: newIns("SET 2,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(2)
			}),
			0xD5: newIns("SET 2,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(2)
			}),
			0xD6: newIns("SET 2,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 2))
			}),
			0xD7: newIns("SET 2,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(2)
			}),
			0xD8: newIns("SET 3,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(3)
			}),
			0xD9: newIns("SET 3,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(3)
			}),
			0xDA: newIns("SET 3,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(3)
			}),
			0xDB: newIns("SET 3,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(3)
			}),
			0xDC: newIns("SET 3,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(3)
			}),
			0xDD: newIns("SET 3,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(3)
			}),
			0xDE: newIns("SET 3,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 3))
			}),
			0xDF: newIns("SET 3,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(3)
			}),
			0xE0: newIns("SET 4,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(4)
			}),
			0xE1: newIns("SET 4,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(4)
			}),
			0xE2: newIns("SET 4,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(4)
			}),
			0xE3: newIns("SET 4,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(4)
			}),
			0xE4: newIns("SET 4,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(4)
			}),
			0xE5: newIns("SET 4,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(4)
			}),
			0xE6: newIns("SET 4,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 4))
			}),
			0xE7: newIns("SET 4,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(4)
			}),
			0xE8: newIns("SET 5,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(5)
			}),
			0xE9: newIns("SET 5,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(5)
			}),
			0xEA: newIns("SET 5,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(5)
			}),
			0xEB: newIns("SET 5,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(5)
			}),
			0xEC: newIns("SET 5,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(5)
			}),
			0xED: newIns("SET 5,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(5)
			}),
			0xEE: newIns("SET 5,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 5))
			}),
			0xEF: newIns("SET 5,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(5)
			}),
			0xF0: newIns("SET 6,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(6)
			}),
			0xF1: newIns("SET 6,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(6)
			}),
			0xF2: newIns("SET 6,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(6)
			}),
			0xF3: newIns("SET 6,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(6)
			}),
			0xF4: newIns("SET 6,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(6)
			}),
			0xF5: newIns("SET 6,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(6)
			}),
			0xF6: newIns("SET 6,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 6))
			}),
			0xF7: newIns("SET 6,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(6)
			}),
			0xF8: newIns("SET 7,B", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("B").BitSet(7)
			}),
			0xF9: newIns("SET 7,C", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("C").BitSet(7)
			}),
			0xFA: newIns("SET 7,D", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("D").BitSet(7)
			}),
			0xFB: newIns("SET 7,E", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("E").BitSet(7)
			}),
			0xFC: newIns("SET 7,H", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("H").BitSet(7)
			}),
			0xFD: newIns("SET 7,L", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("L").BitSet(7)
			}),
			0xFE: newIns("SET 7,(HL)", 16, 0, func(pu CPU, m MEM, params ...uint8) {
				addr := pu.Registers.Combine8("H", "L").Value
				m.Write(addr, tools.Set8(m.Read(addr), 7))
			}),
			0xFF: newIns("SET 7,A", 8, 0, func(pu CPU, m MEM, params ...uint8) {
				pu.Registers.Get8("A").BitSet(7)
			}),
		},
	}
}
