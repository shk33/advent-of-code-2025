# Advent of Code 2025 Solutions

This repository contains solutions for the Advent of Code 2025 challenges, implemented in multiple programming languages to explore different paradigms and approaches.

## Project Structure

Each day's puzzles are organized into separate directories (e.g., `day1/`). Within each day, individual parts of the problem (typically two per day) are further organized into `partX/` subdirectories (e.g., `day1/part1/`, `day1/part2/`).

For every part of a problem, you will find the following files:

*   **`problemX.md`**: The original problem description for that part.
*   **`inputX.txt`**: The specific input data for solving the puzzle.
*   **`planX.md`**: A language-agnostic plan outlining the core logic and algorithm to solve the problem.
*   **`solution.<lang_ext>`**: The implemented solution in each of the chosen programming languages. These solutions are designed to separate the core problem-solving logic from file I/O operations for better testability and reusability.
*   **`syntax-explanation.md`**: A detailed explanation of the code for all implemented languages, tailored for a programmer familiar with JavaScript. It highlights key syntax, common conventions, and comparisons to JavaScript concepts.
*   **`clojure-deep-dive.md`**: A focused, line-by-line explanation of the Clojure solution, delving into its unique functional programming paradigms, immutable data structures, and Lisp syntax.

## Programming Languages Used

We are solving each problem part using the following languages:

*   **Python 3**
*   **JavaScript (Node.js)**
*   **Go**
*   **Ruby**
*   **Clojure**

## Methodology

For each part of an Advent of Code problem, we follow a consistent methodology:

1.  **Understand the Problem**: Read `problemX.md` and analyze `inputX.txt`.
2.  **Develop a Language-Agnostic Plan**: Create `planX.md` outlining the general algorithm.
3.  **Implement Solutions**: Write the `solution.<lang_ext>` files in each of the target languages.
4.  **Verify Solutions**: Run each solution against the `inputX.txt` to ensure correctness.
5.  **Document Explanations**:
    *   Create `syntax-explanation.md` to help developers familiar with JavaScript understand the solutions in other languages.
    *   Create `clojure-deep-dive.md` to provide an in-depth understanding of the Clojure solution, focusing on its unique functional programming aspects.

This approach allows for a comprehensive exploration of problem-solving techniques across different programming paradigms.

## How to Run the Solutions

All commands should be run from the root of the project. The examples below are for `day1/part1/`, but the commands are applicable to any day and part.

### Python

Requires Python 3.

```bash
python3 day1/part1/solution.py
```

### JavaScript

Requires Node.js.

```bash
node day1/part1/solution.js
```

### Go

Requires the Go toolchain.

```bash
go run day1/part1/solution.go
```

### Ruby

Requires Ruby.

```bash
ruby day1/part1/solution.rb
```

### Clojure

Requires a Clojure installation (e.g., the `clj` command-line tool).

```bash
clj -M day1/part1/solution.clj
```
