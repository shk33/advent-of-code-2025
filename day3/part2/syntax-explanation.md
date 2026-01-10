# Day 3, Part 2: Syntax Explanation (Corrected Algorithm)

This document explains the solutions for Day 3, Part 2, based on the **correct** algorithm. The initial approach was flawed; the problem requires finding the largest 12-digit number formed by a **subsequence** of digits, not a contiguous substring.

## Core Logic: Greedy Subsequence Algorithm

To build the lexicographically largest 12-digit number, we must pick the best possible digit for each position from left to right. The "best" digit is the largest one (`'9'` > `'8'`, etc.) that can be chosen while still leaving enough characters to its right to complete the 12-digit number.

**JavaScript Example of the Algorithm:**
```javascript
function findMaxSubsequence(line, k) {
    if (line.length < k) return "0".repeat(k);

    let resultChars = [];
    let currentStartIndex = 0;

    for (let i = 0; i < k; i++) {
        // How many digits are we still looking for?
        const remainingToFind = k - i;
        // Determine the search window. We can't pick a digit so far right
        // that there aren't enough remaining characters to finish the number.
        const searchEndIndex = line.length - remainingToFind;

        let bestDigit = -1;
        let bestDigitIndex = -1;

        // Find the largest digit (and its first occurrence) in the window
        for (let j = currentStartIndex; j <= searchEndIndex; j++) {
            const digit = parseInt(line[j], 10);
            if (digit > bestDigit) {
                bestDigit = digit;
                bestDigitIndex = j;
            }
        }
        
        resultChars.push(bestDigit.toString());
        // The next search must start *after* the digit we just picked.
        currentStartIndex = bestDigitIndex + 1;
    }

    return resultChars.join('');
}
```

---

## 1. Python Solution

```python
# From find_max_subsequence function
result_chars = []
current_start_index = 0

for i in range(k):
    remaining_to_find = k - i
    search_end_index = len(line) - remaining_to_find
    
    best_digit = -1
    best_digit_index = -1

    for j in range(current_start_index, search_end_index + 1):
        digit = int(line[j])
        if digit > best_digit:
            best_digit = digit
            best_digit_index = j
    
    result_chars.append(str(best_digit))
    current_start_index = best_digit_index + 1

# return "".join(result_chars)
```
- The Python code is a direct translation of the greedy algorithm. `range(start, end)` is used for loops, and standard list `append` and string `join` are used. Python's automatic handling of large integers simplifies the final summation.

---

## 2. Go Solution

```go
// From findMaxSubsequence function
var result strings.Builder
currentStartIndex := 0

for i := 0; i < k; i++ {
    remainingToFind := k - i
    searchEndIndex := len(line) - remainingToFind

    bestDigit := -1
    bestDigitIndex := -1

    for j := currentStartIndex; j <= searchEndIndex; j++ {
        digit, _ := strconv.Atoi(string(line[j]))
        if digit > bestDigit {
            bestDigit = digit
            bestDigitIndex = j
        }
    }
    
    result.WriteString(strconv.Itoa(bestDigit))
    currentStartIndex = bestDigitIndex + 1
}

// return result.String()
```
- Go's implementation is also a direct translation.
- `strings.Builder` is used for efficient string construction in a loop, which is better than repeated concatenation.
- `strconv.Atoi` converts a character (as a string) to an integer for comparison.
- **Large Integers**: For the final summation, Go requires the `math/big` package. `totalJoltage.Add(totalJoltage, currentLineJoltage)` is used for addition.

---

## 3. Ruby Solution

```ruby
# From find_max_subsequence method
result_chars = []
current_start_index = 0

(0...k).each do |i|
  remaining_to_find = k - i
  search_end_index = line.length - remaining_to_find

  best_digit = -1
  best_digit_index = -1
  
  (current_start_index..search_end_index).each do |j|
    digit = line[j].to_i
    if digit > best_digit
      best_digit = digit
      best_digit_index = j
    end
  end
  
  result_chars << best_digit.to_s
  current_start_index = best_digit_index + 1
end

# return result_chars.join
```
- Ruby's implementation uses ranges (`0...k` and `current_start_index..search_end_index`) and the `.each` iterator, which are idiomatic for loops.
- `line[j].to_i` converts a character to an integer.
- Like Python, Ruby's `Integer` type handles arbitrary size, so no special `BigInt` objects are needed.

---

## 4. Clojure Solution

```clojure
; find-max-subsequence function
(loop [i 0
       current-start-index 0
       result-strings []]
  (if (= i k)
    (str/join result-strings)
    (let [remaining-to-find (- k i)
          search-end-index (- (count line) remaining-to-find)
          search-str (if (<= current-start-index search-end-index)
                       (subs line current-start-index (inc search-end-index))
                       "")]
      (if (empty? search-str)
        ; ... safeguard ...
        (let [best-digit-str (last (sort (map str search-str)))
              best-digit-offset (.indexOf search-str best-digit-str)
              best-digit-abs-index (+ current-start-index best-digit-offset)]
          (recur (inc i)
                 (inc best-digit-abs-index)
                 (conj result-strings best-digit-str)))))))
```
-   **`loop/recur`**: Clojure uses `loop/recur` for efficient, stack-safe recursion, which is the idiomatic way to write complex loops that carry state from one iteration to the next (`current-start-index` and `result-strings`).
-   **State Management**: The loop's state is explicitly managed through the `recur` call, which re-invokes the `loop` with new values.
-   **Finding the Max Digit**: The expression `(last (sort (map str search-str)))` is a robust, functional way to find the largest digit character:
    1.  `search-str`: The substring to search within.
    2.  `(map str ...)`: Converts each character of the string into its own single-character string (e.g., `"451"` -> `("4" "5" "1")`).
    3.  `(sort ...)`: Sorts these strings lexicographically (e.g., `("1" "4" "5")`).
    4.  `(last ...)`: Takes the last element of the sorted sequence, which is the largest.
-   **`.indexOf`**: This is a Java interop call on the `String` object `search-str` to find the relative index of the `best-digit-str` we found.
-   **`bigint`**: For the final summation, `(map bigint ...)` is used to convert the resulting 12-digit strings into `BigInt`s before they are summed by `(reduce +)`.
