# Plan to Solve Day 1, Part 2

This document outlines the language-agnostic logic for solving the second part of the Advent of Code 2025 Day 1 problem.

## Problem Change Analysis

The core change from Part 1 is the definition of the "password". Instead of counting only the times a rotation *ends* at 0, we must now count every time the dial *clicks* on 0. This includes passing through 0 during a long rotation.

The logic can be modeled by imagining the dial's positions on an infinite number line. A right turn adds to the position, and a left turn subtracts. A "click" on 0 occurs every time a multiple of 100 is crossed on this number line.

- **Example (Right Turn):** Starting at position 80, a rotation of `R30` moves the position from 80 to 110 on the number line. In doing so, it crosses 100 exactly once. This corresponds to one click on 0.
- **Example (Left Turn):** Starting at position 20, a rotation of `L30` moves the position from 20 to -10. In doing so, it crosses 0 exactly once.

The number of times a multiple of 100 is crossed can be calculated using floor division.

## Language-Agnostic Plan

1.  **Initialization:**
    *   Create a variable `current_position` and initialize it to `50`.
    *   Create a counter variable `total_zero_count` and initialize it to `0`.

2.  **Data Input:**
    *   Read the input file (`day1/part1/input1.txt`) line by line.

3.  **Process Rotations (Loop):**
    *   For each line of input:
        a.  **Parse Instruction:** Extract the `direction` ('L' or 'R') and `distance` (as an integer).

        b.  **Calculate Zeros This Turn:** Calculate how many times 0 was clicked during this single rotation.
            *   If `direction` is 'R':
                `zeros_this_turn = floor((current_position + distance) / 100) - floor(current_position / 100)`
            *   If `direction` is 'L':
                `zeros_this_turn = floor((current_position - 1) / 100) - floor((current_position - distance - 1) / 100)`
            *   **Note:** This requires true mathematical `floor` division, especially for negative numbers. The formulas are designed to count the number of multiples of 100 crossed in the respective interval on the number line.

        c.  **Update Total Count:**
            *   Add `zeros_this_turn` to `total_zero_count`.

        d.  **Update Position for Next Turn:**
            *   Calculate the new `current_position` for the circular dial using modulo arithmetic, just like in Part 1.
            *   `current_position = (current_position +/- distance)`
            *   `current_position` must be wrapped to the `0-99` range, ensuring negative results are handled correctly (e.g., -1 becomes 99). The formula `(value % 100 + 100) % 100` is a robust way to do this.

4.  **Final Output:**
    *   After iterating through all the lines, the final value of `total_zero_count` is the answer.
