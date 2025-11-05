package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strconv"

    "github.com/bigsapper/advent-of-code-2024-go/internal/day06"
    "github.com/bigsapper/advent-of-code-2024-go/internal/runner"
)

type solver interface {
    Part1(input string) string
    Part2(input string) string
}

func getSolver(day int) (solver, error) {
    switch day {
    case 6:
        return day06.New(), nil
    default:
        return nil, fmt.Errorf("day %d not implemented", day)
    }
}

func main() {
    dayFlag := flag.Int("day", 6, "day number")
    partFlag := flag.Int("part", 1, "part number (1 or 2)")
    inputFlag := flag.String("input", "", "path to input file (defaults to inputs/dayXX.txt)")
    flag.Parse()

    s, err := getSolver(*dayFlag)
    if err != nil {
        log.Fatal(err)
    }

    inputPath := *inputFlag
    if inputPath == "" {
        dd := fmt.Sprintf("%02s", strconv.Itoa(*dayFlag))
        inputPath = filepath.Join("inputs", "day"+dd+".txt")
    }

    data, err := runner.ReadAll(inputPath)
    if err != nil {
        log.Fatal(err)
    }

    var out string
    switch *partFlag {
    case 1:
        out = s.Part1(data)
    case 2:
        out = s.Part2(data)
    default:
        log.Fatal("part must be 1 or 2")
    }

    _, _ = fmt.Fprintln(os.Stdout, out)
}