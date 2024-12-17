package day17

import (
	"AoC2024/challenge"
	"AoC2024/util"
	"fmt"
	"io"
	"regexp"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 17, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %v\n", partA(challenge.InputFile()))
		},
	}
}

var A, B, C int

func partA(input io.Reader) string {

	var dataPart []string
	for dataLine := range challenge.Sections(input) {
		dataPart = append(dataPart, dataLine)
	}

	// Register setup

	re, err := regexp.Compile("\\d+")
	if err != nil {
		panic(err)
	}

	registers := re.FindAllString(dataPart[0], -1)

	A = util.MustAtoI(registers[0])
	B = util.MustAtoI(registers[1])
	C = util.MustAtoI(registers[2])

	// Program setup
	program := make([]int, 0)
	output := make([]int, 0)

	instructions := re.FindAllString(dataPart[1], -1)

	for _, i := range instructions {

		program = append(program, util.MustAtoI(i))
	}

	programLength := len(program)

	instrIdx := 0
	for instrIdx < programLength {

		opcode, operand := program[instrIdx], program[instrIdx+1]

		switch opcode {
		// The adv instruction (opcode 0) performs division. The numerator is the value in the A register.
		// The denominator is found by raising 2 to the power of the instruction's combo operand.
		// (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
		// The result of the division operation is truncated to an integer and then written to the A register

		case 0:
			numerator := A

			res := numerator >> evalCombo(operand)
			A = res

			instrIdx += 2

		// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand,
		// then stores the result in register B
		case 1:
			B = B ^ operand
			instrIdx += 2

		// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8
		// (thereby keeping only its lowest 3 bits), then writes that value to the B register.
		case 2:

			B = evalCombo(operand) % 8
			instrIdx += 2

		// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero,
		// it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps,
		// the instruction pointer is not increased by 2 after this instruction.
		case 3:
			if A != 0 {
				instrIdx = operand
			} else {
				instrIdx += 2
			}
		// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
		// then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
		case 4:
			B = B ^ C
			instrIdx += 2
		// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value.
		// (If a program outputs multiple values, they are separated by commas.)
		case 5:

			output = append(output, evalCombo(operand)%8)
			instrIdx += 2
		// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored
		// in the B register. (The numerator is still read from the A register.)
		case 6:
			numerator := A
			combo := 0
			combo = evalCombo(operand)
			res := numerator >> combo
			B = res

			instrIdx += 2

		// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is
		// stored in the C register. (The numerator is still read from the A register.)
		case 7:
			numerator := A

			res := numerator >> evalCombo(operand)
			C = res

			instrIdx += 2

		}
	}
	returnString := fmt.Sprintf("%d", output[0])
	for _, x := range output[1:] {
		returnString = fmt.Sprintf("%s,%d", returnString, x)
	}
	fmt.Printf("%s\n", returnString)
	return returnString
}

func evalCombo(operand int) int {

	combo := 0
	switch operand {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		combo = operand
	case 4:
		combo = A
	case 5:
		combo = B
	case 6:
		combo = C
	case 7:
		panic(fmt.Errorf("Operand was 7 - RESERVED"))
	}
	return combo
}
