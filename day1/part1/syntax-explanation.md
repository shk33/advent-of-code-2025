# In-Depth Syntax and Convention Explanations

This document provides a deep dive into the solutions written in Python, Ruby, Go, and Clojure. It's written for a programmer who already knows JavaScript and wants to understand the "how" and "why" behind these other languages.

A core concept in all solutions is the **separation of concerns**. The `solve` function in each file is a "pure function": its output depends only on its inputs, and it has no side effects (like reading files or printing to the console). The `main` function handles the "dirty work" of I/O, making the `solve` function easy to test and reuse.

---

## Python: `solution.py`

**Paradigm**: Python is a multi-paradigm language, but it's most often used in an object-oriented and imperative style. Its philosophy emphasizes readability and simplicity ("There should be one-- and preferably only one --obvious way to do it"). It is dynamically typed, like JavaScript.

```python
import os

# The 'solve' function is pure. It only performs calculations.
def solve(rotations):
    current_position = 50
    zero_count = 0
    # ... (rest of the logic)
    return zero_count

# The 'main' function handles side effects (I/O).
def main():
    # __file__ is a special variable holding the path to the current script.
    input_file_path = os.path.join(os.path.dirname(__file__), 'input1.txt')
    with open(input_file_path, 'r') as f:
        rotations = f.readlines()
    password = solve(rotations)
    print(f"The password is: {password}")

# This ensures 'main()' is only called when the script is run directly.
if __name__ == "__main__":
    main()
```

| Concept | Python | JavaScript Equivalent | Notes |
| :--- | :--- | :--- | :--- |
| **Variable Scope** | Function-scoped (or global) | `var` (ES5), or `let`/`const` (block-scoped) | Python does not have block-level scope for variables. |
| **Function Def** | `def name(arg):` | `function name(arg) { ... }` | Indentation (` `) replaces curly braces `{}` for defining blocks. |
| **File Reading** | `with open(...) as f:` | `fs.readFileSync(...)` | The `with` statement is a **Context Manager**. It guarantees that cleanup code (like closing the file) is run, even if errors occur. It's cleaner than a `try...finally` block. |
| **Entry Point** | `if __name__ == "__main__":` | `if (require.main === module)` in Node.js | In Python, modules have a special `__name__` variable. When a file is run directly, its `__name__` is `"__main__"`. If it's imported, `__name__` is the filename. This is the standard convention for creating runnable scripts. |
| **Math Modulo** | `-18 % 100` -> `82` | `(-18 % 100 + 100) % 100` | Python's `%` operator is a true mathematical modulo, which simplifies the logic for wrapping around the dial compared to JS. |

---

## Ruby: `solution.rb`

**Paradigm**: Ruby is a pure object-oriented language (everything is an object) with a focus on developer happiness and productivity. Its syntax is designed to be elegant and close to natural language. It is dynamically typed.

```ruby
# The 'solve' function is pure.
def solve(rotations)
  # ... (logic using an enumerable)
  zero_count
end

# The 'main' function handles side effects.
def main
  # __dir__ is a special variable for the current file's directory.
  input_file_path = File.join(__dir__, 'input1.txt')
  # Read all lines and call .strip on each one.
  rotations = File.readlines(input_file_path).map(&:strip)
  password = solve(rotations)
  puts "The password is: #{password}"
end

# Standard entry point check for a runnable script.
main if __FILE__ == $PROGRAM_NAME
```

| Concept | Ruby | JavaScript Equivalent | Notes |
| :--- | :--- | :--- | :--- |
| **Blocks** | `do ... end` or `{...}` | Anonymous functions / lambdas `() => {}` | Blocks are a fundamental concept, passed to methods like `each` or `map`. They are a core part of Ruby's expressiveness. `|line|` is the block's parameter. |
| **Iteration** | `rotations.each do ...` | `rotations.forEach(...)` | `each` is the idiomatic way to iterate. All enumerables (like arrays) have this method. |
| **`.map(&:strip)`** | `array.map { |i| i.strip }` | `array.map(i => i.trim())` | The `&:` syntax is a shortcut. It creates a block that calls the method named by the symbol (`:strip`) on each item in the enumerable. It's a common and beloved idiom. |
| **Return Value**| `zero_count` (last line) | `return zero_count;` | In Ruby, the last evaluated expression in a function is **implicitly returned**. You can use the `return` keyword, but it's often omitted by convention if it's on the last line. |
| **Entry Point** | `if __FILE__ == $PROGRAM_NAME` | `if (require.main === module)` | `__FILE__` and `$PROGRAM_NAME` are special global variables that let you check if the file is the main program being run. `$` indicates a global variable. |
| **String Interp.**| `puts "pass: #{password}"`| `` console.log(`pass: ${password}`) `` | Identical to JS template literals. `puts` adds a newline after the output. |

---

## Go: `solution.go` (Special Emphasis)

**Paradigm**: Go is a **statically-typed, compiled language** focused on simplicity, reliability, and efficiency. It feels like a simple, modern version of C. Its biggest selling point is its built-in support for **concurrency** (handling many tasks at once), although that's not used in this example.

Unlike JS, which is interpreted at runtime, Go code is compiled into a single executable file with no external dependencies.

```go
package main // Defines an executable program.

import (
	"bufio"  // Buffered I/O
	"fmt"    // Formatting text
	"os"     // Operating system functions
	"strconv"// String conversions
)

// A pure function. Note the explicit type declarations.
func solve(rotations []string) int {
	// ... (logic)
	return zeroCount
}

func main() {
	inputFilePath := "day1/part1/input1.txt"
	file, err := os.Open(inputFilePath)
	// Explicit error handling is a core Go convention.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	// 'defer' guarantees this runs just before 'main' exits.
	defer file.Close()

    // ... (file scanning logic) ...
	password := solve(rotations)
	fmt.Printf("The password is: %d\n", password)
}
```

#### Deeper Concepts vs. JavaScript

1.  **Static Typing**: This is the biggest difference. In Go, every variable's type must be known when the code is compiled.
    - `var name string = "Eduardo"` is the full declaration.
    - `name := "Eduardo"` is shorthand that infers the type is a `string`.
    - **Benefit**: This catches a huge category of bugs before the program even runs. `let x = 5; x = "hello";` is valid in JS but would be a compile-time error in Go.
    - **Analogy**: It's like using TypeScript everywhere, but it's built into the language.

2.  **Explicit Error Handling**: Go rejects the `try...catch` model of JavaScript.
    - **Convention**: Functions that can fail return their result *and* an `error` value (e.g., `value, err := myFunction()`).
    - **Philosophy**: In Go, errors are just values. You are forced to consciously handle them with `if err != nil`. This makes the code more robust and readable, as error paths are explicit and not hidden by an invisible `try...catch` block somewhere up the call stack.

3.  **Slices vs. Arrays**:
    - In JS, `Array` is your all-purpose list.
    - In Go, `[10]int` is a fixed-size **array**. Its size is part of its type.
    - `[]int` is a **slice**. It's a lightweight, flexible view onto a section of an underlying array. In practice, it feels very much like a JS `Array` and is what you use most of the time. The `append` function is used to add items, similar to `push`.

4.  **`defer` Statement**:
    - JS has `try...finally`. Go has `defer`.
    - `defer file.Close()` schedules `file.Close()` to be run at the end of the current function's execution. It's a simple, elegant way to ensure resources are always cleaned up. You write the cleanup code right next to the allocation code.

---

## Clojure: `solution.clj` (Special Emphasis)

**Paradigm**: Clojure is a modern **Lisp dialect** and a **functional programming language**. It runs on the Java Virtual Machine (JVM). Its philosophy is centered on **simplicity**, which it achieves through **immutability** and a "code is data" approach.

```clojure
(ns solution
  (:require [clojure.string :as str]))

;; A pure function operating on immutable data.
(defn solve [input]
  (let [lines (str/split-lines input)]
    (loop [rotations lines
           position 50
           zero-count 0]
      (if (empty? rotations)
        zero-count ; Base case: return final count
        (let [line (first rotations)
              direction (first line)
              distance (Integer/parseInt (subs line 1))
              new-position (if (= direction \R)
                             (+ position distance)
                             (- position distance))
              wrapped-position (mod new-position 100)]
          ;; 'recur' re-invokes the loop with new, immutable values.
          (recur (rest rotations)
                 wrapped-position
                 (if (zero? wrapped-position)
                   (inc zero-count)
                   zero-count)))))))

;; The main function, handling side effects.
(defn -main [& args]
  (let [input (slurp "day1/part1/input1.txt")]
    (println (str "The password is: " (solve input)))))

(-main)
```

#### Deeper Concepts vs. JavaScript

1.  **Lisp Syntax & "Code is Data"**:
    - Everything is a list wrapped in parentheses: `(function argument1 argument2)`. This is called an S-expression (Symbolic Expression).
    - **This is the most alien concept**. In JS, you have statements, expressions, operators, etc. In Clojure, there's just one rule.
    - **Code is Data**: A line of Clojure code like `(+ 1 2)` is a `List` data structure containing a symbol `+` and two numbers. Because the code itself *is* a data structure, you can write programs that manipulate and generate code, a powerful feature called **macros**.

2.  **Immutability**:
    - In JS, you can have a `const` array, but you can still change its contents: `const arr = [1]; arr.push(2);`.
    - In Clojure, all data structures are **immutable**. You cannot change them.
    - When you "add" an item to a vector, Clojure efficiently creates and returns a *new* vector with the added item. The original is untouched.
    - **Benefit**: This eliminates a massive source of bugs related to shared state and unpredictable mutations. It's a key reason Clojure is excellent for concurrent programming.

3.  **`loop`/`recur` for Iteration**:
    - Since variables are immutable, you can't have a traditional `for` loop that re-assigns `i` (e.g., `i++`).
    - `loop` establishes a recursion point with a set of initial bindings (your "loop variables").
    - `recur` performs a tail-call jump back to the `loop` with a new set of values. It is a low-level, highly efficient primitive that doesn't consume stack space like a normal recursive function call.
    - This is the functional equivalent of a `while` or `for` loop.

4.  **Functional Idiom: `reduce`**
    - While `loop/recur` works, a more common functional pattern for "reducing" a list to a single value is the `reduce` function. The `solve` logic could be rewritten more idiomatically like this:

    ```clojure
    (defn solve-with-reduce [input]
      (let [lines (str/split-lines input)]
        (:zero-count ; 4. Return the final zero-count from the map
          (reduce
            (fn [acc line] ; 2. Reducer function for each line
              (let [{:keys [position]} acc ; Destructuring the map
                    direction (first line)
                    ; ... (parsing logic) ...
                    new-acc (assoc acc :position wrapped-position)]
                (if (zero? wrapped-position)
                  (update new-acc :zero-count inc) ; "update" the count
                  new-acc)))
            {:position 50 :zero-count 0} ; 1. Initial value for the accumulator
            lines))))) ; 3. The collection to iterate over
    ```
    - This is similar to `Array.prototype.reduce` in JavaScript. It takes a function, an initial value (`{:position 50, :zero-count 0}`), and a collection. It's often preferred as it's a higher-level abstraction than `loop/recur`.

5.  **The JVM and Java Interop**:
    - Clojure runs on the JVM, giving it two major benefits: it's battle-tested, high-performance, and has access to the entire universe of Java libraries.
    - `(Integer/parseInt (subs line 1))` is a direct example of this. It's calling the static `parseInt` method on the `java.lang.Integer` class. The syntax is `(JavaClassName/staticMethod args...)`.

---

## Rust: `solution.rs` (Special Emphasis)

**Paradigm**: Rust is a **statically-typed, compiled language** focused on **performance, reliability, and memory safety**. Its key innovation is the **ownership and borrowing** system, which guarantees memory safety without needing a garbage collector like in JS, Go, or Java. This gives it the performance of C++ with much stronger safety guarantees.

```rust
use std::fs;

// A pure function. Types are explicit.
fn solve(rotations: &str) -> i32 {
    let mut current_position = 50;
    let mut zero_count = 0;

    for line in rotations.lines() {
        let direction = line.chars().next().unwrap();
        let distance: i32 = line[1..].parse().unwrap();

        current_position += if direction == 'R' { distance } else { -distance };
        current_position = current_position.rem_euclid(100);

        if current_position == 0 {
            zero_count += 1;
        }
    }
    zero_count
}

fn main() {
    let input = fs::read_to_string("day1/part1/input1.txt").expect("Error reading file");
    let password = solve(&input);
    println!("The password is: {}", password);
}
```

#### Deeper Concepts vs. JavaScript

1.  **Static Typing & Type Inference**:
    - Like Go, all types must be known at compile time.
    - `let distance: i32 = ...` is an explicit type annotation for a 32-bit signed integer.
    - Often, the compiler can infer the type, so you can write `let password = solve(&input);` without an explicit type, just like in Go.

2.  **Ownership and Borrowing**: This is Rust's most unique feature.
    - **Ownership**: Each value in Rust has a single "owner." When the owner goes out of scope, the value is dropped (memory is freed). `let s1 = String::from("hello"); let s2 = s1;` **moves** ownership. `s1` can no longer be used. This prevents "double free" errors.
    - **Borrowing**: Instead of moving ownership, you can create "references" (borrows) to a value. This is what the `&` symbol does. `solve(&input)` passes a reference to the `input` string, so `solve` can read it without taking ownership and destroying it.
    - **Benefit**: This system provides C++-level performance (no garbage collector pauses) while making it impossible to have memory bugs like dangling pointers or data races at compile time.

3.  **Mutability**:
    - Variables are **immutable by default**.
    - `let x = 5; x = 6;` is a compile-time error.
    - To make a variable mutable, you must explicitly use the `mut` keyword: `let mut zero_count = 0;`. This makes the code's intent clearer.

4.  **Error Handling: `Result` and `Option`**:
    - Like Go, Rust rejects `try...catch`. It uses two special generic enums (sum types) for handling operations that might fail.
    - **`Option<T>`**: Represents a value that can be something (`Some(T)`) or nothing (`None`). It's used when a value might not exist. `line.chars().next()` returns an `Option<char>`. This is like a nullable type that you are forced to handle.
    - **`Result<T, E>`**: Represents a value that can be success (`Ok(T)`) or failure (`Err(E)`). `fs::read_to_string` returns a `Result<String, io::Error>`.
    - **`.unwrap()` and `.expect()`**: These are shortcuts to get the value out of an `Option` or `Result`. If the value is `None` or `Err`, the program will **panic** (crash). This is fine for quick scripts but in robust applications, you would use `match` or `if let` to handle the error cases gracefully. `expect("message")` is like `unwrap()` but with a custom error message.

5.  **Expressions vs. Statements**:
    - In Rust, nearly everything is an **expression**, meaning it returns a value.
    - `if`, `match`, and even code blocks `{}` are expressions.
    - This allows for very concise code: `let value = if condition { 1 } else { 0 };`. This is similar to a ternary operator in JS (`condition ? 1 : 0`).

---

## Elixir: `solution.exs` (Special Emphasis)

**Paradigm**: Elixir is a **dynamic, functional programming language** built on the **Erlang VM (BEAM)**. It is designed for building highly scalable, fault-tolerant, and concurrent applications. Its syntax is heavily inspired by Ruby.

```elixir
defmodule Solution do
  # The main 'solve' function is pure.
  def solve(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.reduce({50, 0}, fn line, {position, zero_count} ->
      direction = String.first(line)
      distance = String.slice(line, 1..-1) |> String.to_integer()

      new_position =
        if direction == "R" do
          position + distance
        else
          position - distance
        end

      wrapped_position = rem(new_position, 100)

      new_zero_count =
        if wrapped_position == 0 do
          zero_count + 1
        else
          zero_count
        end

      {wrapped_position, new_zero_count}
    end)
    |> elem(1) # Get the second element (zero_count) from the final tuple.
  end
end

input = File.read!("day1/part1/input1.txt")
password = Solution.solve(input)
IO.puts("The password is: #{password}")
```

#### Deeper Concepts vs. JavaScript

1.  **Immutability and Functional Purity**:
    - Like Clojure, all data in Elixir is **immutable**. You don't "change" data, you create new data through transformations.
    - This makes it a functional language at its core. Functions are predictable and easy to test.

2.  **The Pipe Operator `|>`**:
    - This is the most iconic feature of Elixir. It takes the result of the expression on its left and passes it as the **first argument** to the function on its right.
    - `c(b(a))` becomes `a |> b |> c`.
    - It makes data transformation pipelines extremely readable, from top to bottom.

3.  **Pattern Matching**:
    - The `=` sign in Elixir is not assignment, it's a **match operator**.
    - `x = 1` asserts "I expect `x` to be equal to 1." If `x` was already `2`, this would throw an error.
    - It's most powerful with data structures. ` {position, zero_count} = {50, 0} ` destructures the tuple into two variables. This is like JS destructuring: `const [position, zero_count] = [50, 0];`.
    - It's used heavily in function heads to define different function clauses for different input patterns, which is a powerful alternative to `if` or `switch` statements.

4.  **Modules and Functions**:
    - Code is organized into **modules** (`defmodule`).
    - Functions are defined with `def` (public) or `defp` (private).
    - There is no `return` keyword; the last expression's value is implicitly returned, just like in Ruby.

5.  **The Erlang VM (BEAM) and Concurrency**:
    - The true power of Elixir comes from its runtime environment, the BEAM, which has been used for decades to build ultra-reliable telecom systems.
    - It's famous for its **lightweight "processes"** (not OS processes). You can have millions of them running concurrently. They are isolated, communicate via messages, and if one crashes, it doesn't take down the others.
    - This makes Elixir an incredible choice for web servers, real-time applications, and distributed systems. While not used in this simple script, it's the primary reason to choose Elixir.
```
