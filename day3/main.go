package main

import (
	"fmt"
	"github.com/cquon/aoc-2020/inputreader"
)

/*
--- Day 3: Toboggan Trajectory ---
With the toboggan login problems resolved, you set off toward the airport. While travel by toboggan might be easy, it's certainly not safe: there's very minimal steering and the area is covered in trees. You'll need to see which angles will take you near the fewest trees.

Due to the local geology, trees in this area only grow on exact integer coordinates in a grid. You make a map (your puzzle input) of the open squares (.) and trees (#) you can see. For example:

..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
These aren't the only trees, though; due to something you read about once involving arboreal genetics and biome stability, the same pattern repeats to the right many times:

..##.........##.........##.........##.........##.........##.......  --->
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on your map).

The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers); start by counting all the trees you would encounter for the slope right 3, down 1:

From your starting position at the top-left, check the position that is right 3 and down 1. Then, check the position that is right 3 and down 1 from there, and so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with O where there was an open square and X where there was a tree:

..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........X.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...#X....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
In this example, traversing the map using this slope would cause you to encounter 7 trees.

Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?

--- Part Two ---
Time to check the rest of the slopes - you need to minimize the probability of a sudden arboreal stop, after all.

Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left corner and traverse the map all the way to the bottom:

Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively; multiplied together, these produce the answer 336.

What do you get if you multiply together the number of trees encountered on each of the listed slopes?
*/

type slopeLine struct {
	length        int
	treePositions []int
}

func (s *slopeLine) isTree(pos int) bool {
	if pos >= s.length {
		pos = pos % s.length
	}
	if len(s.treePositions) == 0 {
		return false
	}

	currTreeIdx := 0
	for s.treePositions[currTreeIdx] <= pos {
		if pos == s.treePositions[currTreeIdx] {
			return true
		}
		if currTreeIdx == (len(s.treePositions) - 1) {
			return false
		}
		currTreeIdx++
	}

	return false
}

func lineParser(b []byte) interface{} {
	sl := &slopeLine{
		length: len(b),
	}
	for i := range b {
		if b[i] == '#' {
			sl.treePositions = append(sl.treePositions, i)
		}
	}
	return sl
}

func part1() {
	ir := reader.NewInputReader("input.txt", lineParser)
	slopeLines := ir.ParseInput()
	currPos := 0
	treeCtr := 0
	for _, sl := range slopeLines {
		if sl.(*slopeLine).isTree(currPos) {
			treeCtr++
		}
		currPos += 3
	}
	fmt.Println(treeCtr)
}

func getTreesEncountered(slopeLines []interface{}, right int, down int) int {
	currCol := 0
	currRow := 0
	treeCtr := 0
	for currRow+down < len(slopeLines) {
		currCol += right
		currRow += down
		if slopeLines[currRow].(*slopeLine).isTree(currCol) {
			treeCtr++
		}
	}
	fmt.Println(treeCtr)
	return treeCtr
}

func part2() {
	ir := reader.NewInputReader("input.txt", lineParser)
	slopeLines := ir.ParseInput()
	var answer int64
	answer = int64(getTreesEncountered(slopeLines, 1, 1)) * int64(getTreesEncountered(slopeLines, 3, 1)) * int64(getTreesEncountered(slopeLines, 5, 1)) * int64(getTreesEncountered(slopeLines, 7, 1)) * int64(getTreesEncountered(slopeLines, 1, 2))
	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
