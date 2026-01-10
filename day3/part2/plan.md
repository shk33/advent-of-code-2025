# Day 3, Part 2: Lobby Escalator Power (Enhanced)

## 1. Problem Understanding

The core problem remains the same as Part 1: calculate a total "output joltage" from a series of battery banks. Each bank is a line of digits. The key difference is how the joltage from each bank is determined.

**Change from Part 1:** Instead of selecting exactly *two* batteries (digits) to form a two-digit number, we now need to select exactly *twelve* batteries (digits) from the line to form a twelve-digit number. The order of the selected digits must be maintained, and we are looking for the *largest possible numerical value* that can be formed.

The total output joltage is the sum of these maximum twelve-digit numbers from all banks.

**Example Revisited (from problem description):**
- In `987654321111111`, the largest 12-digit joltage is `987654321111`.
- In `811111111111119`, the largest 12-digit joltage is `811111111119`.

## 2. Core Logic: Finding the Max Twelve-Digit Joltage in a Line

The initial assumption that the 12 digits must be a contiguous substring was incorrect. The problem requires finding the largest number formed by a **subsequence** of 12 digits, meaning the selected digits must maintain their relative order but do not need to be adjacent in the original string.

A greedy algorithm is well-suited for this. To construct the lexicographically largest result, we should pick the best possible digit for each position from left to right.

**Greedy Algorithm:**
1.  Initialize an empty result string, `result_joltage`.
2.  Initialize a search starting position, `start_index = 0`.
3.  We need to select `k=12` digits. Loop from `i = 0` to `11` (for each of the 12 positions in our final number):
    a.  Determine the search range for the current digit. We need to leave enough digits for the rest of the number. If we have `k - i` digits left to pick, the last possible starting index for the current digit is `len(s) - (k - i)`. So, the search range is from `start_index` up to `len(s) - (k - i)`.
    b.  Within this search range, find the largest digit.
    c.  If there are multiple occurrences of this largest digit, we must pick the one that appears earliest (i.e., at the lowest index) to maximize the chances for the subsequent digits.
    d.  Find the index of this chosen digit. Let's call it `best_digit_index`.
    e.  Append the chosen digit (`s[best_digit_index]`) to `result_joltage`.
    f.  Update the search starting position for the *next* iteration: `start_index = best_digit_index + 1`.
4.  After the loop completes, `result_joltage` will be the string representation of the largest possible 12-digit subsequence.

**Example:** Find the largest 4-digit subsequence in `s = "1548293"` (`k=4`).
- **Digit 1:** Need 3 more. Search range `s[0:len(s)-3]` -> `"1548"`. Best digit is '8' at index 3.
  - `result = "8"`. `start_index = 3 + 1 = 4`.
- **Digit 2:** Need 2 more. Search range `s[4:len(s)-2]` -> `"29"`. Best digit is '9' at index 5.
  - `result = "89"`. `start_index = 5 + 1 = 6`.
- **Digit 3:** Need 1 more. Search range `s[6:len(s)-1]` -> `"3"`. Best digit is '3' at index 6.
  - `result = "893"`. `start_index = 6 + 1 = 7`.
- **Digit 4:** Need 0 more. Search range `s[7:len(s)-0]` -> `""`. This is wrong. The end of the search range is inclusive. It should be `len(s) - (k-i)`.
Let's fix the range: `s[start_index : len(s) - (k-i-1)]`.

**Corrected Greedy Algorithm:**
1.  Initialize `result = ""` and `search_start_pos = 0`.
2.  For `i` from `0` to `11`:
    a. Let `num_digits_to_find = 12 - i`.
    b. `search_end_pos = len(s) - num_digits_to_find`.
    c. Find the max digit in `s` from `search_start_pos` up to and including `search_end_pos`.
    d. Find the index of its first occurrence in that range. Call it `best_index`.
    e. Append `s[best_index]` to `result`.
    f. Update `search_start_pos = best_index + 1`.
3. Return `result`.

## 3. Overall Algorithm

1.  Initialize a variable, `total_joltage`, to 0 (or a suitable large integer type).
2.  Read the input file line by line.
3.  For each line:
    a.  Apply the "Core Logic" described above to find the `max_line_joltage_str`.
    b.  Convert `max_line_joltage_str` to an integer.
    c.  Add this integer to `total_joltage`.
4.  After processing all lines, the `total_joltage` will be the final answer.

## 4. Implementation Details

-   All solutions will read from `input.txt` in the same directory.
-   The core logic function will take a string (a single line) as input and return a string (the max 12-digit joltage).
-   The main part of each script will handle file I/O, call the core logic function for each line, perform the string-to-integer conversion for the 12-digit number, and sum the results. Due to the potentially large size of the 12-digit numbers, languages like Python and Ruby handle large integers automatically. Go and JavaScript will need to use specific types (`int64` in Go, `BigInt` in JavaScript) to avoid overflow.

This plan focuses on efficiently finding the largest 12-digit number per line and correctly summing them, accounting for potential large number handling.
