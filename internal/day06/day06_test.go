package day06

import (
	"strings"
	"testing"
)

const exampleInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1_Example(t *testing.T) {
	s := New()
	got := s.Part1(exampleInput)
	want := "41"
	if got != want {
		t.Fatalf("Part1(example) = %s, want %s", got, want)
	}
}

func TestPart2_Example(t *testing.T) {
	s := New()
	got := s.Part2(exampleInput)
	want := "6"
	if got != want {
		t.Fatalf("Part2(example) = %s, want %s", got, want)
	}
}

func Test_CRLF_Handling(t *testing.T) {
	s := New()
	crlf := strings.ReplaceAll(exampleInput, "\n", "\r\n")
	if got := s.Part1(crlf); got != "41" {
		t.Fatalf("Part1(CRLF) = %s, want %s", got, "41")
	}
}