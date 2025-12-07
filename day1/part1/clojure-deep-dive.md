# Clojure Deep Dive: From JavaScript to a Functional Lisp

This document will walk you through the Clojure solution line-by-line, explaining the core concepts of the language and its functional paradigm. The mental model is different from JavaScript, so we'll build it from the ground up.

### Core Concept 1: Everything is an S-expression (The Parentheses)

In JavaScript, you have different syntax for declaring variables, calling functions, and using operators:
`let x = 1 + 2;`
`console.log("Hello");`

In Clojure (and other Lisps), there is only **one** syntax rule: a list in parentheses `()`, where the **first item is the function** and the rest are arguments.

- `1 + 2` in JS becomes `(+ 1 2)` in Clojure.
- `console.log("Hello")` in JS becomes `(println "Hello")` in Clojure.

This "code is data" principle means the code you write is itself a simple Clojure list, which is an incredibly powerful concept for advanced metaprogramming. For now, just remember: **(function arg1 arg2...)**.

### Core Concept 2: Immutability (Data Never Changes)

This is the most important and most different concept. In JavaScript, `const` prevents re-assigning a variable, but the data it points to can still be changed (mutated).

```javascript
// JavaScript
const person = { name: "John" };
const another = person; // Both point to the same object
another.name = "Jane";
console.log(person.name); // Prints "Jane". The original object was changed!
```

In Clojure, this **cannot happen**. All data structures are **immutable**. You can't change them. Ever. When you want to "change" something, Clojure efficiently creates a *new* data structure with the updated value.

```clojure
;; Clojure
(let [person {:name "John"}
      another (assoc person :name "Jane")] ; 'assoc' returns a NEW map
  (println (:name person)) ; Prints "John". The original is unchanged.
  (println (:name another))) ; Prints "Jane".
```

This prevents an entire class of bugs related to shared state and makes code, especially concurrent code, far easier to reason about.

---

## Line-by-Line Code Walkthrough

Here is the full code. We will break it down piece by piece.

```clojure
(ns solution
  (:require [clojure.string :as str]))

(defn solve [input]
  (let [lines (str/split-lines input)]
    (loop [rotations lines
           position 50
           zero-count 0]
      (if (empty? rotations)
        zero-count
        (let [line (first rotations)
              direction (first line)
              distance (Integer/parseInt (subs line 1))
              new-position (if (= direction \R)
                             (+ position distance)
                             (- position distance))
              wrapped-position (mod new-position 100)]
          (recur (rest rotations)
                 wrapped-position
                 (if (zero? wrapped-position)
                   (inc zero-count)
                   zero-count)))))))

(defn -main [& args]
  (let [input (slurp "day1/part1/input1.txt")]
    (println (str "The password is: " (solve input)))))

(-main)
```

---

### Line 1: `(ns solution ...)`

```clojure
(ns solution
  (:require [clojure.string :as str]))
```

- **`(ns solution ...)`**: This declares a **n**ame**s**pace, which is Clojure's equivalent of a module in JavaScript or a package in Java. It prevents naming conflicts.
- **`(:require ...)`**: This clause is like JavaScript's `import`.
- **`[clojure.string :as str]`**: We are importing the built-in `clojure.string` library and aliasing it as `str`. This is almost identical to `import * as str from 'some-string-library';` in JS. The `[...]` denotes a **vector**, which is Clojure's version of a JS Array. `:require` and `:as` are **keywords**. Think of them as special, efficient strings used for keys in maps (like JS objects).

### Line 4: `(defn solve [input])`

```clojure
(defn solve [input]
   ; ... body of the function
)
```

- **`defn`**: "define function". This is how you create a new function.
- **`solve`**: The name of the function.
- **`[input]`**: A vector containing the function's parameters. This is like `(input)` in a JS function declaration.

### Line 5: `(let [lines (str/split-lines input)] ...)`

```clojure
(let [lines (str/split-lines input)]
  ; ... rest of the 'solve' function
)
```

- **`let`**: This creates local bindings (like `const` or `let` in JS). These bindings are immutable and only exist within the scope of the `let` block.
- **`[lines (str/split-lines input)]`**: This is the bindings vector. We are binding the name `lines` to the result of `(str/split-lines input)`. This call is like `input.split('\n')` in JS.

### Line 6: `(loop [...] ...)` - The Heart of the Logic

This is the most complex part. Since data is immutable, you can't have a traditional `for` loop with a counter that you change (`i++`). A `loop` provides the mechanism for efficient, recursive-style iteration.

```clojure
(loop [rotations lines
       position 50
       zero-count 0]
   ; ... loop body ...
)
```

- **`loop`**: This establishes a "recursion target."
- **`[...]`**: These are the **loop's state variables**. Think of this like initializing your variables right before a `while` loop in JS:
  - `rotations lines` is like `let rotations = lines;`
  - `position 50` is like `let position = 50;`
  - `zero-count 0` is like `let zeroCount = 0;`

### Line 10: `(if (empty? rotations) ...)` - The Exit Condition

```clojure
(if (empty? rotations)
  zero-count ; "then" block: if the list is empty, we're done.
  ; ... "else" block: if the list is not empty, keep processing.
)
```

- **`if`**: The standard conditional. It has the form `(if condition then-expression else-expression)`.
- **`(empty? rotations)`**: The condition. This is a function call. The `?` at the end is a common Clojure convention for functions that return a boolean (`true`/`false`). It's equivalent to `rotations.length === 0`.
- **`zero-count`**: If the condition is true, this is what the entire `loop` expression evaluates to. It's the final result, like `return zeroCount;` at the end of a JS function.

### Line 12: `(let [...] ...)` - Inside the Loop

This inner `let` block is for processing a single line.

```clojure
(let [line (first rotations)
      direction (first line)
      distance (Integer/parseInt (subs line 1))
      ; ... and so on
```

- **`(first rotations)`**: Gets the first item from the `rotations` collection. Equivalent to `rotations[0]`.
- **`(Integer/parseInt (subs line 1))`**: An example of **Java Interoperability**. Since Clojure runs on the JVM, it can call any Java code. This line calls the static method `parseInt` from the Java `Integer` class. `(subs line 1)` is like `line.substring(1)`.
- **`(mod new-position 100)`**: Calls the `mod` function. Unlike JS's `%` remainder operator, `mod` in Clojure is a true mathematical modulo, so `(mod -10 100)` correctly returns `90`.

### Line 20: `(recur ...)` - The Next Iteration

This is the key to making the loop continue.

```clojure
(recur (rest rotations) ; The list, minus its first item
       wrapped-position ; The new position
       (if (zero? wrapped-position) ; The new count
         (inc zero-count)
         zero-count))
```

- **`recur`**: This keyword causes execution to jump back to the `loop` statement. It **must** be in the "tail position" (the very last thing that happens). It's a low-level instruction that re-uses the current stack frame, so it's as fast as a `while` loop and won't cause a stack overflow.
- **The arguments to `recur` provide the new values for the loop variables for the next iteration**:
  1. **`(rest rotations)`**: The new value for `rotations`. `rest` returns the entire collection *except* for the first item. This is how we move to the next line, like `rotations.slice(1)`.
  2. **`wrapped-position`**: The new value for `position`.
  3. **`(if (zero? ...))`**: The new value for `zero-count`. `(inc zero-count)` is a function that returns `zero-count + 1`. It doesn't *change* `zero-count`; it returns a new number.

Think of one full iteration as consuming the first line of the list and then calling `recur` with the *rest* of the list and the *newly calculated* state. This continues until the list is empty.

### Lines 24-28: `(-main)` - The Entry Point

```clojure
(defn -main [& args]
  (let [input (slurp "day1/part1/input1.txt")]
    (println (str "The password is: " (solve input)))))

(-main)
```
- **`defn -main`**: By convention, `-main` is the function that runs when you execute a file as a script.
- **`(slurp ...)`**: A convenience function that reads an entire file into one big string.
- **`(println (str ...))`**: `str` concatenates its arguments into a string, and `println` prints it to the console.
- **`(-main)`**: This final line actually calls the main function, starting the program.
