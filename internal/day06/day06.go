package day06

import (
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type point struct{ r, c int }

type state struct {
	r, c, d int
}

var dirs = []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func normalize(input string) []string {
	s := strings.TrimSpace(input)
	if s == "" {
		return nil
	}
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.Split(s, "\n")
}

func findStart(grid []string) (int, int, int) {
	for r := range grid {
		for c := range grid[r] {
			switch grid[r][c] {
			case '^':
				return r, c, 0
			case '>':
				return r, c, 1
			case 'v':
				return r, c, 2
			case '<':
				return r, c, 3
			}
		}
	}
	return 0, 0, 0
}

func (s *Solver) Part1(input string) string {
	grid := normalize(input)
	if len(grid) == 0 {
		return "0"
	}
	sr, sc, dir := findStart(grid)

	visited := map[point]bool{}
	r, c := sr, sc
	rows := len(grid)
	cols := len(grid[0])

	for {
		visited[point{r, c}] = true
		nr := r + dirs[dir].r
		nc := c + dirs[dir].c
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			break
		}
		ch := grid[nr][nc]
		if ch == '#' {
			dir = (dir + 1) % 4
			continue
		}
		r, c = nr, nc
	}

	return strconv.Itoa(len(visited))
}

func simulateLoop(grid []string, sr, sc, dir int, block point) bool {
	rows := len(grid)
	cols := len(grid[0])
	seen := make(map[state]bool)
	r, c, d := sr, sc, dir

	for {
		st := state{r, c, d}
		if seen[st] {
			return true
		}
		seen[st] = true
		nr := r + dirs[d].r
		nc := c + dirs[d].c
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			return false
		}
		if (nr == block.r && nc == block.c) || grid[nr][nc] == '#' {
			d = (d + 1) % 4
			continue
		}
		r, c = nr, nc
	}
}

func (s *Solver) Part2(input string) string {
	grid := normalize(input)
	if len(grid) == 0 {
		return "0"
	}
	sr, sc, dir := findStart(grid)

	rows := len(grid)
	cols := len(grid[0])

	visited := make(map[point]bool)
	r, c, d := sr, sc, dir

	for {
		visited[point{r, c}] = true
		nr := r + dirs[d].r
		nc := c + dirs[d].c
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			break
		}
		if grid[nr][nc] == '#' {
			d = (d + 1) % 4
			continue
		}
		r, c = nr, nc
	}

	count := 0
	for p := range visited {
		if p.r == sr && p.c == sc {
			continue
		}
		if grid[p.r][p.c] == '#' {
			continue
		}
		if simulateLoop(grid, sr, sc, dir, p) {
			count++
		}
	}

	return strconv.Itoa(count)
}