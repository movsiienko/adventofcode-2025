# Advent of Code 2025

# Default recipe - show help
default:
    @just --list

# Build executable for specific day
build DAY:
    @echo "Building day {{DAY}}..."
    @go build -o bin/day{{DAY}} ./cmd/day{{DAY}}

# Run executable for specific day
run DAY:
    @if [ ! -f bin/day{{DAY}} ]; then \
        echo "Building day {{DAY}}..."; \
        go build -o bin/day{{DAY}} ./cmd/day{{DAY}}; \
    fi
    @echo "Running day {{DAY}}..."
    @cd cmd/day{{DAY}} && ../../bin/day{{DAY}}

# Build and run in one command
solve DAY:
    @just build {{DAY}}
    @just run {{DAY}}

# Build all day executables
build-all:
    @echo "Building all days..."
    @for dir in cmd/day*; do \
        day=$$(basename $$dir | sed 's/day//'); \
        echo "Building day $$day..."; \
        go build -o bin/day$$day ./$$dir; \
    done

# Remove all built executables
clean:
    @echo "Cleaning..."
    @rm -rf bin/

# Run all tests
test:
    @echo "Running tests..."
    @go test ./...

# Run tests with coverage
test-coverage:
    @echo "Running tests with coverage..."
    @go test -cover ./...

# Format all Go code
fmt:
    @echo "Formatting code..."
    @go fmt ./...

# Run linter
lint:
    @echo "Running linter..."
    @golangci-lint run ./...

# Create a new day's directory structure
new DAY:
    @echo "Creating day {{DAY}}..."
    @mkdir -p cmd/day{{DAY}}
    @touch cmd/day{{DAY}}/input.txt
    @echo 'package main\n\nimport (\n\t"fmt"\n\t"ovsiienko.xyz/advent-of-code/internal/util"\n)\n\nfunc main() {\n\tlines, err := util.ReadLines("input.txt")\n\tif err != nil {\n\t\tfmt.Printf("Error reading input: %v\\n", err)\n\t\treturn\n\t}\n\n\t// Part 1\n\tfmt.Println("Part 1:", part1(lines))\n\n\t// Part 2\n\tfmt.Println("Part 2:", part2(lines))\n}\n\nfunc part1(lines []string) int {\n\t// TODO: Implement part 1\n\treturn 0\n}\n\nfunc part2(lines []string) int {\n\t// TODO: Implement part 2\n\treturn 0\n}\n' > cmd/day{{DAY}}/main.go
    @echo "Day {{DAY}} created! Edit cmd/day{{DAY}}/main.go and add your input to cmd/day{{DAY}}/input.txt"
