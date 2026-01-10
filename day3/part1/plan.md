# Day 3, Part 1: Lobby Escalator Power

## 1. Problem Understanding

The goal is to calculate a total "output joltage" from a series of battery banks. Each bank is a line of digits in the input file. For each bank, we must find the largest possible two-digit number (joltage) that can be formed by selecting two digits from the line, in their original order. The total output joltage is the sum of these maximum joltages from each bank.

**Example:**
- For the bank `818181911112111`, we can form numbers like `81`, `88`, `89`, `18`, `11`, `91`, `92`, etc.
- The largest number we can form is `92`.
- We do this for every line and sum the results.

## 2. Core Logic

The core task is to find the maximum two-digit number for a single line of digits.

A brute-force approach is feasible and clear:
1.  For a given line of digits (string `s`):
2.  Initialize a variable, `max_line_joltage`, to 0.
3.  Use nested loops to iterate through all possible ordered pairs of indices `(i, j)` such that `0 <= i < j < length(s)`.
4.  For each pair of indices, create a two-digit number by concatenating the digits `s[i]` and `s[j]`.
5.  Convert this two-digit string into an integer.
6.  Compare this integer with `max_line_joltage`. If it's larger, update `max_line_joltage`.
7.  After checking all pairs, `max_line_joltage` will hold the largest possible two-digit number for that line.

## 3. Overall Algorithm

1.  Initialize a variable, `total_joltage`, to 0.
2.  Read the input file line by line.
3.  For each line:
    a.  Apply the "Core Logic" described above to find the `max_line_joltage`.
    b.  Add this `max_line_joltage` to the `total_joltage`.
4.  After processing all lines, the `total_joltage` will be the final answer.

## 4. Implementation Details

-   All solutions will read from `input.txt` in the same directory.
-   The core logic function will take a string (a single line) as input and return an integer (the max joltage for that line).
-   The main part of each script will handle file I/O, call the core logic function for each line, and sum the results.

This plan is simple and directly models the problem description. Given the constraints on the length of the lines, a nested loop approach is computationally acceptable.
