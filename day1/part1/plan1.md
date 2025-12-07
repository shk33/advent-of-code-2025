# Plan to Solve Day 1, Problem 1

This document outlines the language-agnostic logic for solving the Advent of Code 2025 Day 1, Problem 1.

### 1. Initialization
- Create a variable to store the dial's current position, let's call it `current_position`. Initialize it to `50`, as per the problem description.
- Create a counter variable to track how many times the dial lands on zero, let's call it `zero_count`. Initialize it to `0`.

### 2. Data Input
- Read the entire `input1.txt` file.
- Process the file line by line, as each line represents a single rotation instruction.

### 3. Process Rotations
- For each line of input:
    a. **Parse the Instruction:** Separate the line into two parts: 
        - The **direction** (the first character, 'L' or 'R').
        - The **distance** (the remaining part of the string, converted to an integer).
    b. **Calculate New Position:**
        - If the direction is 'R' (right), **add** the distance to `current_position`.
        - If the direction is 'L' (left), **subtract** the distance from `current_position`.
    c. **Handle Dial Wrap-Around:** The dial is circular with 100 positions (0-99). The new position must be calculated using modulo 100 arithmetic to ensure it stays within this range.
        - `current_position = calculated_position % 100`
        - **Note:** Special care must be taken for negative results from left rotations. The result must wrap around correctly (e.g., a raw value of -1 should become 99). The specific implementation may vary by language, but a common robust formula is `(raw_value % 100 + 100) % 100`.
    d. **Check for Zero:** After each rotation is complete and `current_position` is updated, check if its new value is exactly `0`.
    e. **Increment Counter:** If `current_position` is `0`, increment `zero_count` by one.

### 4. Final Output
- After iterating through all the lines in the input file, the final value of `zero_count` is the answer. This value should be presented as the solution.
