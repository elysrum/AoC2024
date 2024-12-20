// Code generated by 'go run ./gen'; DO NOT EDIT

package cmd

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"AoC2024/challenge/day1"
	"AoC2024/challenge/day2"
	"AoC2024/challenge/day3"
	"AoC2024/challenge/day4"
	"AoC2024/challenge/day5"
	"AoC2024/challenge/day6"
	"AoC2024/challenge/day7"
	"AoC2024/challenge/day8"
	"AoC2024/challenge/day9"
	"AoC2024/challenge/day10"
	"AoC2024/challenge/day11"
	"AoC2024/challenge/day12"
	"AoC2024/challenge/day13"
	"AoC2024/challenge/day14"
	"AoC2024/challenge/day15"
	"AoC2024/challenge/day16"
	"AoC2024/challenge/day17"
	"AoC2024/challenge/day18"
	// "AoC2024/challenge/day19"
	// "AoC2024/challenge/day20"
	// "AoC2024/challenge/day21"
	// "AoC2024/challenge/day22"
	// "AoC2024/challenge/day23"
	// "AoC2024/challenge/day24"
	// "AoC2024/challenge/day25"
	"AoC2024/challenge/example"
)

func addDays(root *cobra.Command) {
	example.AddCommandsTo(root)
	day1.AddCommandsTo(root)
	day2.AddCommandsTo(root)
	day3.AddCommandsTo(root)
	day4.AddCommandsTo(root)
	day5.AddCommandsTo(root)
	day6.AddCommandsTo(root)
	day7.AddCommandsTo(root)
	day8.AddCommandsTo(root)
	day9.AddCommandsTo(root)
	day10.AddCommandsTo(root)
	day11.AddCommandsTo(root)
	day12.AddCommandsTo(root)
	day13.AddCommandsTo(root)
	day14.AddCommandsTo(root)
	day15.AddCommandsTo(root)
	day16.AddCommandsTo(root)
	day17.AddCommandsTo(root)
	day18.AddCommandsTo(root)
	// day19.AddCommandsTo(root)
	// day20.AddCommandsTo(root)
	// day21.AddCommandsTo(root)
	// day22.AddCommandsTo(root)
	// day23.AddCommandsTo(root)
	// day24.AddCommandsTo(root)
	// day25.AddCommandsTo(root)
}

type prof interface {
	Stop()
}

func NewRootCommand() *cobra.Command {
	var (
		start    time.Time
		profiler prof
	)

	result := &cobra.Command{
		Use:     "aoc2024",
		Short:   "Advent of Code 2024 Solutions",
		Long:    "Golang implementations for the 2024 Advent of Code problems",
		Example: "go run main.go 1 a -i ./challenge/day1/input.txt",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if viper.GetBool("profile") {
				profiler = profile.Start()
			}

			start = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			if profiler != nil {
				profiler.Stop()
			}

			fmt.Println("Took", time.Since(start))
		},
	}

	addDays(result)

	flags := result.PersistentFlags()

	flags.StringP("input", "i", "", "Input File to read. If not specified, assumes ./challenge/dayN/input.txt for the currently running challenge")
	flags.Bool("profile", false, "Profile implementation performance")

	_ = viper.BindPFlags(flags)

	return result
}
