# Clojure Deep Dive (Part 2): The Functional Number Line

Welcome to the deep dive for Part 2. This document builds on the concepts from the Part 1 dive (Lisp syntax, immutability, `loop/recur`), so please make sure you're familiar with those.

The challenge in Part 2 is to track the dial's position on an "infinite" number line, counting every time it crosses a multiple of 100. Let's see how this is handled in a functional, immutable way.

## The Core Problem: Tracking State Without Changing It

In an imperative language, you'd have a single `current_position` variable and just add or subtract from it in a loop.

```javascript
// Imperative JavaScript
let currentPosition = 50;
for (const line of lines) {
    // ...
    currentPosition += distance; // Mutating the variable
}
```

In Clojure, we can't do this. The "variables" in our `loop` are immutable bindings. The solution is the same as before: to "change" the state, we pass a *new* value for it in the `recur` call. For Part 2, we simply treat `position` as the ever-growing number on the infinite line and thread it through the loop just like `rotations` and the `total-zero-count`.

---

## Line-by-Line Code Walkthrough

Here is the Part 2 `solve` function and its new helper.

```clojure
(defn floor-div [a b]
  (int (Math/floor (/ (double a) (double b)))))

(defn solve [input]
  (let [lines (str/split-lines input)]
    (loop [rotations lines
           position 50
           total-zero-count 0]
      (if (empty? rotations)
        total-zero-count
        (let [line (first rotations)
              direction (first line)
              distance (Integer/parseInt (subs line 1))
              zeros-this-turn (if (= direction \R)
                                (- (floor-div (+ position distance) 100) (floor-div position 100))
                                (- (floor-div (- position 1) 100) (floor-div (- position distance 1) 100)))
              new-position (if (= direction \R)
                             (+ position distance)
                             (- position distance))]
          (recur (rest rotations)
                 new-position
                 (+ total-zero-count zeros-this-turn)))))))
```

---

### The `floor-div` Helper Function

```clojure
(defn floor-div [a b]
  (int (Math/floor (/ (double a) (double b)))))
```

- **Why is this needed?** Clojure's standard division function `/` is very precise. `(/ 10 3)` produces a rational number `10/3`. `(/ -11 100)` produces `-11/100`. We need a classic integer floor division, like Python's `//`.
- **`(double a)`**: This is a **type hint** or **casting function**. It converts the integer `a` into a floating-point number (a "double"). We do this because Java's `Math/floor` operates on doubles.
- **`(/ ...)`**: Now that the arguments are doubles, `/` performs floating-point division.
- **`(Math/floor ...)`**: This is the Java interop call to `java.lang.Math.floor()`. It gives us the correct floor value.
- **`(int ...)`**: This casts the resulting double (e.g., `5.0`) back into an integer (`5`), which is the type we need for our calculations.

### The `solve` function

The structure is the same `loop/recur` as before, but the bindings and calculations inside have changed.

**Line 8: The Loop Bindings**

```clojure
(loop [rotations lines
       position 50
       total-zero-count 0]
```

- `position` is still our starting point, but now we will let it grow or shrink indefinitely, treating it as a point on the infinite number line.
- `total-zero-count` is our accumulator that will hold the final answer.

**Line 13: `zeros-this-turn`**

```clojure
(let [ ...
      zeros-this-turn (if (= direction \R)
                        (- (floor-div (+ position distance) 100) (floor-div position 100))
                        (- (floor-div (- position 1) 100) (floor-div (- position distance 1) 100)))
      ...
```

- This is a direct translation of the logic from our `plan2.md`.
- `(if (= direction \R) ... ...)`: We check the direction. Note that `\R` is the syntax for a single `R` character, distinct from the string `"R"`.
- `(- ...)`: The subtraction function. `(- a b)` is `a - b`.
- The `then` and `else` blocks of the `if` call our new `floor-div` helper to calculate how many multiples of 100 were crossed.

**Line 18: `new-position`**

```clojure
      new-position (if (= direction \R)
                     (+ position distance)
                     (- position distance))]
```
- This calculates the next value for `position` on the **unwrapped, infinite number line**. We are *not* using `mod` here, because we need the real value for the next turn's `floor-div` calculation.

**Line 21: `(recur ...)`**

```clojure
(recur (rest rotations)
       new-position
       (+ total-zero-count zeros-this-turn))
```
- The state of our loop is passed to the next iteration.
- `(rest rotations)`: The new value for the `rotations` binding.
- `new-position`: The new value for the `position` binding.
- `(+ total-zero-count zeros-this-turn)`: The new value for the `total-zero-count` binding. We add the zeros from this turn to our running total.

This flow perfectly demonstrates the functional approach: state isn't mutated. It is calculated and passed as an argument to the next step in a pipeline, which in this case is the next iteration of the loop.
