# Day 3, Part 1: Clojure Deep Dive

This document provides a detailed, line-by-line explanation of the Clojure solution for Day 3, Part 1. The goal is to find the maximum two-digit number that can be formed from the digits in each line of input, and then sum these maximums.

## The Full Clojure Code

```clojure
(ns solution
  (:require [clojure.string :as str]))

(defn- get-max-joltage-for-line [line]
  (let [digits (vec line)]
    (->> (for [i (range (count digits))
               j (range (inc i) (count digits))]
           (str (get digits i) (get digits j)))
         (map #(Integer/parseInt %))
         (apply max 0))))

(defn -main [& args]
  (let [input-path (str (.getParent (java.io.File. *file*)) "/input.txt")
        lines (str/split-lines (slurp input-path))
        total-joltage (->> lines
                           (map str/trim)
                           (filter not-empty)
                           (map get-max-joltage-for-line)
                           (reduce +))]
    (println (str "The total output joltage is: " total-joltage))))

(-main)
```

---

## Line-by-Line Breakdown

### `get-max-joltage-for-line` Function

This is the core function that processes a single line. It's a great example of Clojure's functional, data-oriented approach.

```clojure
(defn- get-max-joltage-for-line [line]
```

-   `(defn- ...)`: Defines a private function named `get-max-joltage-for-line`. It's private (`-`) because it's only intended for use within this namespace.
-   `[line]`: The function takes one argument, `line`, which is a string.

```clojure
  (let [digits (vec line)]
```

-   `(let [...] ...)`: This is a `let` binding, which creates local variables.
-   `digits (vec line)`: We create a local variable `digits`. `(vec line)` converts the input string `line` into a vector of its characters. For example, `"1928"` becomes `[\1 \9 \2 \8]`. Using a vector allows for efficient, indexed access with the `get` function.

```clojure
    (->> ;; Start of the thread-last macro
```

-   `->>` (thread-last): This powerful macro lets us write a sequence of operations in a clean, top-to-bottom style. The result of the first form is passed as the *last* argument to the second form, and so on.

```clojure
      (for [i (range (count digits))
            j (range (inc i) (count digits))]
        (str (get digits i) (get digits j)))
```

-   This is a list comprehension, which generates a sequence. It's not a loop in the imperative sense.
-   `i (range (count digits))`: This binding iterates `i` from `0` to one less than the number of digits.
-   `j (range (inc i) (count digits))`: This is a nested binding. `(inc i)` is `i + 1`, so `j` iterates from `i + 1` to the end. This setup ensures that we only get ordered pairs `(i, j)` where `i < j`.
-   `(str (get digits i) (get digits j))`: For each pair of `i` and `j`, this form gets the character at each index from the `digits` vector and concatenates them into a two-digit string.
-   **Result**: This `for` form produces a lazy sequence of all possible two-digit strings (e.g., `"98"`, `"97"`, `"96"`, ... `"21"`).

```clojure
      (map #(Integer/parseInt %))
```

-   The sequence of strings from the `for` comprehension is now "threaded" into this `map` function.
-   `(map ...)`: Applies a function to every item in a sequence and returns a new sequence of the results.
-   `#(Integer/parseInt %)`: This is a shorthand for an anonymous function. It's equivalent to `(fn [s] (Integer/parseInt s))`. The `%` represents the argument (each two-digit string).
-   **Result**: This step transforms the sequence of strings into a sequence of integers (e.g., `(98 97 96 ... 21)`).

```clojure
      (apply max 0))))
```

-   The sequence of integers is now threaded into `apply`.
-   `(apply max ...)`: The `apply` function takes a function (`max` in this case) and a sequence, and "applies" the items of the sequence as if they were individual arguments to the function. So, `(apply max '(98 97 96))` is equivalent to `(max 98 97 96)`.
-   `0`: We include a `0` at the end. If the input line has fewer than two digits, the sequence of numbers would be empty. Calling `(apply max '())` would cause an error. By adding `0`, we ensure `max` has at least one argument, returning `0` in that edge case.
-   **Result**: This final step finds the largest number in the sequence, which is the maximum joltage for the line.

### `-main` Function

This function orchestrates the whole process, from reading the file to printing the final result.

```clojure
(defn -main [& args]
```
- Defines the main entry point of the script.

```clojure
  (let [input-path (str (.getParent (java.io.File. *file*)) "/input.txt")
        lines (str/split-lines (slurp input-path))
```
- **File I/O**:
  - `*file*`: A special variable that holds the path to the current file.
  - `(java.io.File. *file*)`: Creates a Java File object.
  - `(.getParent ...)`: A Java interop call to get the parent directory.
  - `(slurp input-path)`: Reads the entire content of the input file into a single string.
  - `(str/split-lines ...)`: Splits the file content into a sequence of lines.

```clojure
        total-joltage (->> lines
                           (map str/trim)
                           (filter not-empty)
                           (map get-max-joltage-for-line)
                           (reduce +))]
```
- **Data Pipeline**: Another `->>` macro is used to process the sequence of lines.
  1. `lines`: The initial sequence of strings from the file.
  2. `(map str/trim)`: Trims whitespace from both ends of every line.
  3. `(filter not-empty)`: Removes any blank lines from the sequence.
  4. `(map get-max-joltage-for-line)`: Applies our core logic function to each line, transforming the sequence of lines into a sequence of their corresponding max joltage numbers.
  5. `(reduce +)`: The `reduce` function collapses a sequence into a single value. `(reduce +)` starts with the first two numbers, adds them, then adds the next number to the result, and so on, effectively summing the entire sequence.

```clojure
    (println (str "The total output joltage is: " total-joltage))))
```
- Finally, the calculated `total-joltage` is concatenated with a descriptive string and printed to the console.
