# Day 3, Part 1: Syntax Explanation

This document explains the solutions for Day 3, Part 1 in Python, Go, Ruby, and Clojure, assuming a background in JavaScript. The core logic involves iterating through each line of input, finding the maximum two-digit number that can be formed from its digits, and summing these maximums.

## Core Logic: Finding the Max Joltage in a Line

The main task is to generate all possible two-digit numbers from a string of digits and find the largest one.

**JavaScript Example:**
```javascript
let maxLineJoltage = 0;
for (let i = 0; i < line.length; i++) {
    for (let j = i + 1; j < line.length; j++) {
        const joltageStr = line[i] + line[j];
        const joltage = parseInt(joltageStr, 10);
        if (joltage > maxLineJoltage) {
            maxLineJoltage = joltage;
        }
    }
}
```

This nested loop structure is a direct way to generate all ordered pairs and is mirrored in the other languages.

---

## 1. Python Solution

```python
# solution.py

max_line_joltage = 0
for i in range(len(line)):
    for j in range(i + 1, len(line)):
        joltage_str = line[i] + line[j]
        joltage = int(joltage_str)
        if joltage > max_line_joltage:
            max_line_joltage = joltage
```

### Key Differences & Concepts

-   **Looping**: Python's `for i in range(len(line))` is equivalent to `for (let i = 0; i < line.length; i++)` in JavaScript.
-   **String Concatenation**: `line[i] + line[j]` works just like in JavaScript.
-   **Type Conversion**: `int(joltage_str)` is Python's version of `parseInt(joltage_str, 10)`.

---

## 2. Go Solution

```go
// solution.go

maxLineJoltage := 0
for i := 0; i < len(line); i++ {
    for j := i + 1; j < len(line); j++ {
        joltageStr := string(line[i]) + string(line[j])
        joltage, _ := strconv.Atoi(joltageStr)
        if joltage > maxLineJoltage {
            maxLineJoltage = joltage
        }
    }
}
```

### Key Differences & Concepts

-   **Variable Declaration**: `maxLineJoltage := 0` is Go's shorthand for declaring and initializing a variable (`var maxLineJoltage = 0`).
-   **Looping**: Go's `for` loop syntax is very similar to JavaScript's, but without the parentheses.
-   **String and Character Handling**: In Go, iterating over a string gives you `rune`s (a character type). To concatenate them as strings, you must first cast them back to `string`, like `string(line[i])`.
-   **Type Conversion**: `strconv.Atoi(joltageStr)` (ASCII to Integer) is Go's equivalent of `parseInt`. It returns two values: the result and an error. The underscore `_` is used to discard the error value, a common practice when you don't expect an error (as we are building the string ourselves).

---

## 3. Ruby Solution

```ruby
# solution.rb

max_line_joltage = 0
(0...line.length).each do |i|
  (i + 1...line.length).each do |j|
    joltage_str = line[i] + line[j]
    joltage = joltage_str.to_i
    if joltage > max_line_joltage
      max_line_joltage = joltage
    end
  end
end
```

### Key Differences & Concepts

-   **Ranges and Iteration**: Ruby uses ranges like `(0...line.length)` to define a sequence of numbers (from 0 up to, but not including, the length). The `.each` method then iterates over this range, which is analogous to a `for` loop.
-   **Blocks**: The `do |variable| ... end` syntax defines a block, which is similar to a callback function passed to an iterator in JavaScript (e.g., `array.forEach(item => { ... })`).
-   **String Indexing**: `line[i]` works just like in JavaScript.
-   **Type Conversion**: The `.to_i` method is called on the string object to convert it to an integer, similar to `parseInt`.

---

## 4. Clojure Solution

```clojure
; solution.clj

(defn- get-max-joltage-for-line [line]
  (let [digits (vec line)]
    (->> (for [i (range (count digits))
               j (range (inc i) (count digits))]
           (str (get digits i) (get digits j)))
         (map #(Integer/parseInt %))
         (apply max 0))))
```

### Key Differences & Concepts

-   **Functional Approach**: Instead of loops with mutable variables, the Clojure solution uses a functional, data-transformation pipeline.
-   **`for` Comprehension**: `(for [...] ...)` is not a traditional loop. It's a "list comprehension" that generates a sequence of values based on the iterations. Here, it generates all the two-digit `joltage_str` values.
-   **`->>` (Thread-Last Macro)**: This macro takes the result of one expression and "threads" it as the *last* argument to the next function. It makes the code easier to read, top-to-bottom.
-   **Data Transformation**:
    1.  The `for` comprehension creates a lazy sequence of all two-digit strings.
    2.  `(map #(Integer/parseInt %))` takes that sequence and applies the `Integer/parseInt` function to each string, producing a sequence of numbers. The `#(...)` is a shorthand for an anonymous function, where `%` is the argument.
    3.  `(apply max 0)` takes the sequence of numbers and "applies" them as arguments to the `max` function. `max` finds the largest of its arguments. We include `0` in case the input sequence is empty.
-   **Immutability**: No variables are reassigned. Each step transforms the data from the previous step into a new sequence.

---

## 5. Rust Solution

```rust
// solution.rs

fn get_max_joltage_for_line(line: &str) -> i32 {
    let mut max_line_joltage = 0;
    let chars: Vec<char> = line.chars().collect();
    for i in 0..chars.len() {
        for j in i + 1..chars.len() {
            let joltage_str = format!("{}{}", chars[i], chars[j]);
            let joltage: i32 = joltage_str.parse().unwrap();
            if joltage > max_line_joltage {
                max_line_joltage = joltage;
            }
        }
    }
    max_line_joltage
}
```

### Key Differences & Concepts

-   **Mutability**: The `max_line_joltage` variable must be declared with `mut` to allow it to be updated.
-   **Character Handling**: In Rust, a `&str` is a UTF-8 encoded string slice. To access characters by index, it's idiomatic to first convert it into a `Vec<char>` (a vector of characters). `line.chars().collect()` achieves this.
-   **Looping**: `for i in 0..chars.len()` is Rust's syntax for a `for` loop over a range of numbers, similar to Python.
-   **String Formatting**: The `format!` macro is used for building strings. It's a powerful and type-safe way to handle string interpolation and concatenation.
-   **Type Conversion**: `.parse()` is used to convert the string to an integer. The type annotation `: i32` helps the compiler infer what type to parse into. `.unwrap()` is used to get the value out of the `Result` returned by `parse`, crashing if the parse fails.

---

## 6. Elixir Solution

```elixir
# solution.exs

defp get_max_joltage_for_line(line) do
  digits = String.graphemes(line)
  
  for i <- 0..(length(digits) - 1),
      j <- (i + 1)..(length(digits) - 1),
      reduce: 0 do
    max_joltage ->
      joltage_str = Enum.at(digits, i) <> Enum.at(digits, j)
      joltage = String.to_integer(joltage_str)
      max(joltage, max_joltage)
  end
end
```

### Key Differences & Concepts

-   **`for` Comprehension with `reduce`**: Elixir's `for` is a powerful comprehension that can both generate and reduce data. Here, it's used to iterate through all pairs of indices (`i` and `j`).
-   **`reduce: 0`**: This option transforms the `for` from a generator into a reducer. `0` is the initial value of the accumulator (`max_joltage`).
-   **Accumulator**: In each iteration, the block `do max_joltage -> ... end` is executed. The value returned by this block becomes the new value of `max_joltage` for the next iteration. This is a purely functional way to achieve the same result as a mutable `max_line_joltage` variable.
-   **`max/2`**: The `max` function is used to find the greater of the two values, which is more concise than an `if` statement.
-   **String/Character Handling**: `String.graphemes` splits the string into a list of its characters (grapheme clusters). `Enum.at` is used for indexed access, and `<>` is the string concatenation operator.
