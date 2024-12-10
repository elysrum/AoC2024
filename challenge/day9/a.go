package day9

import (
	"container/list"
	"fmt"
	"io"

	"AoC2024/challenge"
	"AoC2024/util"

	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 9, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {

	type DiskBlock struct {
		blockNumber int
		blockType   int // File = 0, FS =
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

nextFreespace:
	for fs := blocks.Front(); fs != nil; fs = fs.Next() {

		fsBlock = fs.Value.(DiskBlock)

		if fsBlock.blockType == 1 && fsBlock.noBlocks > 0 {
			// Loop over files backwards
			for f := blocks.Back(); f != nil; f = f.Prev() {
				fBlock = f.Value.(DiskBlock)
				// Its a File
				if fBlock.blockType == 0 && fBlock.noBlocks > 0 {

					// FS has enough Space to hold file, split FS and delete file block
					if fsBlock.noBlocks >= fBlock.noBlocks {
						remainder := fsBlock.noBlocks - fBlock.noBlocks
						if remainder > 0 {
							// split the FS block
							newBlock := DiskBlock{blockType: 1, blockNumber: -1, noBlocks: remainder}
							blocks.InsertAfter(newBlock, fs)

						}
						fs.Value = fBlock
						blocks.Remove(f)

						// More FS available, on with next file

					} else { // FS does not have enough space, convert block to File, split file
						remainder := fBlock.noBlocks - fsBlock.noBlocks

						newFBlock := DiskBlock{blockType: fBlock.blockType, blockNumber: fBlock.blockNumber, noBlocks: remainder}
						newFSBlock := DiskBlock{blockType: fBlock.blockType, blockNumber: fBlock.blockNumber, noBlocks: fsBlock.noBlocks}

						//Convert fs to f
						fs.Value = newFSBlock
						f.Value = newFBlock

						// Filled FS, need to grab next one.
						continue nextFreespace
					}
				}

			}
		}

	}

	idx := 0
	for e := blocks.Front(); e != nil; e = e.Next() {
		block := e.Value.(DiskBlock)
		if block.blockType == 0 {
			blocks := block.noBlocks
			for i := 0; i < blocks; i++ {
				result = result + (idx * block.blockNumber)
				idx += 1
			}
		}
	}

	return result
}
