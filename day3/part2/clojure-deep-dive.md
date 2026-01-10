# Day 3, Part 2: Clojure Deep Dive (Corrected Algorithm)

This document provides a detailed, line-by-line explanation of the **correct** Clojure solution for Day 3, Part 2. The initial approach was flawed because the problem requires finding the largest 12-digit **subsequence**, not a contiguous substring. This solution implements a greedy algorithm to solve the problem.

## The Full (Corrected) Clojure Code

```clojure
(ns solution
  (:require [clojure.string :as str]))

(defn find-max-subsequence [line k]
  (if (< (count line) k)
    (str/join (repeat k "0"))
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
            (str/join (concat result-strings (repeat remaining-to-find "0")))
            (let [best-digit-str (last (sort (map str search-str)))
                  best-digit-offset (.indexOf search-str best-digit-str)
                  best-digit-abs-index (+ current-start-index best-digit-offset)]
              (recur (inc i)
                     (inc best-digit-abs-index)
                     (conj result-strings best-digit-str)))))))))

(defn -main [& args]
  (let [input-path (str (.getParent (java.io.File. *file*)) "/input.txt")
        k 12
        lines (str/split-lines (slurp input-path))
        total-joltage (->> lines
                           (map str/trim)
                           (filter not-empty)
                           (map #(find-max-subsequence % k))
                           (map bigint)
                           (reduce +))]
    (println (str "The total output joltage is: " total-joltage))))

(-main)
```

---

## Line-by-Line Breakdown

### `find-max-subsequence` Function

This function implements the greedy algorithm to find the largest subsequence of length `k`. It's structured as a tail-recursive loop.

```clojure
(defn find-max-subsequence [line k]
```
- Defines the function taking a `line` string and subsequence length `k`.

```clojure
  (if (< (count line) k)
    (str/join (repeat k "0"))
```
- A guard clause: if the line is too short to form a k-digit number, it returns a string of `k` zeros.

```clojure
    (loop [i 0
           current-start-index 0
           result-strings []]
```
- `(loop ...)`: This starts a loop, which is Clojure's idiomatic construct for efficient recursion. It initializes three "loop variables":
    - `i`: A counter for which digit we are currently finding (from 0 to 11).
    - `current-start-index`: The starting position for our search in the `line`. This moves forward as we pick digits.
    - `result-strings`: A vector that accumulates the chosen digit strings.

```clojure
      (if (= i k)
        (str/join result-strings)
```
- **Base Case**: If our counter `i` equals `k`, we have found all 12 digits. `(str/join result-strings)` concatenates them into the final result string, and the loop terminates.

```clojure
        (let [remaining-to-find (- k i)
              search-end-index (- (count line) remaining-to-find)
              search-str (if (<= current-start-index search-end-index)
                           (subs line current-start-index (inc search-end-index))
                           "")]
```
- **Recursive Step**: This `let` block calculates the bounds for our search.
    - `remaining-to-find`: How many more digits we need to select.
    - `search-end-index`: The last possible index we can search for the current digit while still leaving enough characters to its right for the `remaining-to-find` digits.
    - `search-str`: The actual substring we will search within for our current best digit.

```clojure
          (if (empty? search-str)
            (str/join (concat result-strings (repeat remaining-to-find "0")))
```
- A safeguard. If the search string is somehow empty, it fills the rest of the result with zeros and terminates.

```clojure
            (let [best-digit-str (last (sort (map str search-str)))
```
- This is the core of the greedy choice. It robustly finds the largest digit in the search string.
    1. `(map str search-str)`: Converts each character in `search-str` to a single-character string (e.g., `"451"` -> `("4" "5" "1")`).
    2. `(sort ...)`: Sorts these strings lexicographically. For single digits, this is equivalent to a numerical sort (e.g., `("1" "4" "5")`).
    3. `(last ...)`: Takes the last element of the sorted list, which is the largest digit as a string (e.g., `"5"`).
    
```clojure
                  best-digit-offset (.indexOf search-str best-digit-str)
                  best-digit-abs-index (+ current-start-index best-digit-offset)]
```
- `best-digit-offset`: We find the *first* occurrence of our `best-digit-str` within the `search-str`. This is a Java interop call on the String object.
- `best-digit-abs-index`: We calculate the absolute index of our chosen digit relative to the original `line`.

```clojure
              (recur (inc i)
                     (inc best-digit-abs-index)
                     (conj result-strings best-digit-str)))))))))
```
- `(recur ...)`: This is the magic of the `loop`. It re-invokes the loop with updated values, effectively starting the next iteration.
    - `(inc i)`: We increment the digit counter.
    - `(inc best-digit-abs-index)`: The search for the *next* digit must begin immediately after the one we just found.
    - `(conj result-strings best-digit-str)`: We add our newly found `best-digit-str` to our vector of results. `conj` (conjoin) is the idiomatic way to add an element to a collection in Clojure.

### `-main` Function
The main function is largely the same, but it's worth noting how the pipeline works with the corrected `find-max-subsequence` function.

```clojure
(defn -main [& args]
  (let [...
        total-joltage (->> lines
                           ...
                           (map #(find-max-subsequence % k))
                           (map bigint)
                           (reduce +))]
    ...))
```
- `(map #(find-max-subsequence % k))`: This now maps each line to its largest 12-digit subsequence string.
- `(map bigint)`: This is crucial. It converts each 12-digit string result into a `BigInt`.
- `(reduce +)`: The `+` function automatically works with `BigInt`s, correctly summing them up into a final `BigInt` result without overflow.