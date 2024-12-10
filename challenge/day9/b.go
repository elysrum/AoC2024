package day9

import (
	"container/list"
	"fmt"
	"io"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {

	type DiskBlock struct {
		blockNumber int
		blockType   int // File = 0, FS = 1
		noBlocks    int
	}

	data := challenge.Lines(input)

	result := 0
	blocks := list.New()

	// Create the initial blocks array
	for inputLine := range data {

		blockIndex := 0
		for idx, cell := range inputLine {
			val := util.MustAtoI(string(cell))
			blockType := (idx % 2)
			var block DiskBlock
			if blockType == 0 {
				block = DiskBlock{blockType: blockType, blockNumber: blockIndex, noBlocks: val}
				blockIndex += 1
			} else {
				block = DiskBlock{blockType: blockType, blockNumber: -1, noBlocks: val}

			}
			blocks.PushBack(block)

		}
	}

	// Loop over free space forwards
	var fsBlock DiskBlock
	var fBlock DiskBlock
nextFile:
	for f := blocks.Back(); f != nil; f = f.Prev() {
		fBlock = f.Value.(DiskBlock)
		// Its a File
		if fBlock.blockType == 0 && fBlock.noBlocks > 0 {
			// Loop over files backwards
			for fs := blocks.Front(); fs != nil; fs = fs.Next() {

				if fs == f {
					continue nextFile
				}
				fsBlock = fs.Value.(DiskBlock)

				if fsBlock.blockType == 1 && fsBlock.noBlocks > 0 {

					// FS has enough Space to hold file, split FS and delete file block
					if fsBlock.noBlocks >= fBlock.noBlocks {
						remainder := fsBlock.noBlocks - fBlock.noBlocks
						if remainder > 0 {
							// split the FS block
							newBlock := DiskBlock{blockType: 1, blockNumber: -1, noBlocks: remainder}
							blocks.InsertAfter(newBlock, fs)

						}
						fs.Value = fBlock
						f.Value = DiskBlock{blockType: 1, blockNumber: -1, noBlocks: fBlock.noBlocks}

						continue nextFile

					}

				}

			}
		}

	}

	idx := 0
	for e := blocks.Front(); e != nil; e = e.Next() {
		block := e.Value.(DiskBlock)
		blocks := block.noBlocks
		for i := 0; i < blocks; i++ {
			if block.blockType == 0 {
				result = result + (idx * block.blockNumber)
			}
			idx += 1

		}
	}

	return result
}
