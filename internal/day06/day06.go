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

var dirs = []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (s *Solver) Part1(input string) string {
    sIn := strings.TrimSpace(input)
    if sIn == "" {
        return "0"
    }
    sIn = strings.ReplaceAll(sIn, "\r\n", "\n")
    lines := strings.Split(sIn, "\n")

    grid := lines

    var sr, sc, dir int
    found := false
    for r := range grid {
        for c := range grid[r] {
            switch grid[r][c] {
            case '^':
                sr, sc, dir = r, c, 0
                found = true
            case '>':
                sr, sc, dir = r, c, 1
                found = true
            case 'v':
                sr, sc, dir = r, c, 2
                found = true
            case '<':
                sr, sc, dir = r, c, 3
                found = true
            }
            if found {
                break
            }
        }
        if found {
            break
        }
    }

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

func (s *Solver) Part2(input string) string {
    return ""
}