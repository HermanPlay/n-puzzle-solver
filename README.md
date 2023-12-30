# N-Puzzle Solver in Go

This repository contains a simple N-Puzzle solver implemented in Go. The N-Puzzle problem is a classic sliding puzzle game where a board of tiles must be rearranged to reach a goal state.

## Implementation details

For searching [A* algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm) is used together with [Manhattan Distance](https://en.wikipedia.org/wiki/Taxicab_geometry) + [Linear Conflicts](https://medium.com/swlh/looking-into-k-puzzle-heuristics-6189318eaca2#:~:text=Linear%20Conflicts,reach%20their%20final%20goal%20position.) as heuristics.

Board states are stored in 1D array of ints. To store states in a min-heap states are converted to strings as such: 
```
Node{[1, 2, 3, 4, 5, 6, 7, 8, 0]} -> "1 2 3 4 5 6 7 8 0"
```

## Usage

1. Ensure you have Go installed on your machine.
2. Clone this repository:

   ```bash
   git clone https://github.com/HermanPlay/n-puzzle-solver.git
   ```
3. Navigate to project directory:
   ```bash
   cd n-puzzle-solver
   ```
4. To run predefined tests:
   ```bash
   go test .\n-puzzle
   ```
5. To run custom board, build the solution:
   ```bash
   go build
   ```
## Flags
```
Usage of n-puzzle:
  -gen
        (OPTIONAL) generates puzzle in the given path file
  -path string
        (REQUIRED) path to a file to save generated puzzle, requires size
  -size int
        (REQUIRED) size of one of the sides of the puzzle
  -v    (OPTIONAL) output verbose solution path
```

## License
This project is licensed under the MIT License - see the LICENSE file for details.
