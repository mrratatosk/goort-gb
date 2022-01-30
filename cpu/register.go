package cpu

import "github.com/mrratatosk/oort-framework/register"

func registers() register.RegisterList {
	rl := register.NewRegisterList()

	rl.AddRegister8("A")
	rl.AddRegister8("B")
	rl.AddRegister8("C")
	rl.AddRegister8("D")
	rl.AddRegister8("E")
	rl.AddRegister8("F")
	rl.AddRegister8("H")
	rl.AddRegister8("L")

	rl.AddRegister16("PC")
	rl.AddRegister16("SP")

	return rl
}

type Flags uint8

const (
	Z Flags = 4
	N       = 5
	H       = 6
	C       = 7
)
