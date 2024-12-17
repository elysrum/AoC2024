package day17

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"slices"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 17, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %vd\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

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

	instructions := re.FindAllString(dataPart[1], -1)

	for _, i := range instructions {

		program = append(program, util.MustAtoI(i))
	}

	startA := 0
	best := 0
	val := 0
	for {

		// brute force it - calc chunks of numbers at a time - thanks to JPaulson

		startA += 1
		//val := startA
		val = startA*int(math.Pow(float64(8), float64(10))) + 0o2756025052
		// val := startA*2097152 + 0o44025052

		output := runSim(val, program)
		if slices.Equal(output, program) {
			break
		} else if len(output) > best {
			fmt.Printf("A=%d, oct(A)=%o, best=%d, len=%d\n", val, val, best, len(program))
			best = len(output)
		}

	}
	fmt.Printf("Result: %d", val)
	return val
}

func runSim(startA int, program []int) []int {

	A = startA
	B = 0
	C = 0
	output := make([]int, 0)
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
			if output[len(output)-1] != program[len(output)-1] {

				return output
			}
			instrIdx += 2
		// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored
		// in the B register. (The numerator is still read from the A register.)
		case 6:
			numerator := A

			res := numerator >> evalCombo(operand)
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
	return output

}
