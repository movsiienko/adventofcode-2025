# Advent of Code 2025

My solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Project Structure

```
.
├── cmd/
│   ├── day01/          # Day 1 solution
│   │   ├── main.go
│   │   └── input.txt
│   ├── day02/          # Day 2 solution
│   │   └── ...
│   └── ...
├── internal/
│   └── util/           # Shared utilities
│       └── input.go
├── go.mod
├── Makefile
└── README.md
```

## Usage

### Building

Build a specific day:
```bash
just build 01
```

Build all days:
```bash
just build-all
```

### Running

Run a specific day:
```bash
just run 01
```

Build and run in one command:
```bash
just solve 01
```

Or run the executable directly:
```bash
./bin/day01
```

### Adding a New Day

The easiest way is to use the built-in scaffold:
```bash
just new 02
```

This creates the directory structure and a template `main.go` file.

Or manually:
1. Create a new directory under `cmd/`:
   ```bash
   mkdir cmd/day02
   ```

2. Create `main.go` in the new directory:
   ```bash
   touch cmd/day02/main.go
   ```

3. Add your input file:
   ```bash
   touch cmd/day02/input.txt
   ```

4. Build and run:
   ```bash
   just solve 02
   ```

## Utilities

Common utilities are available in `internal/util/`:
- `ReadInput(filename)` - Read entire file as string
- `ReadLines(filename)` - Read file as slice of lines

Example usage:
```go
import "ovsiienko.xyz/advent-of-code/internal/util"

lines, err := util.ReadLines("input.txt")
```

## Other Commands

View all available commands:
```bash
just
```

Remove all built executables:
```bash
just clean
```

Run tests:
```bash
just test
```

Format code:
```bash
just fmt
```
