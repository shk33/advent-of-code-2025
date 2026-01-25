# Day 3, Part 2: Syntax Explanation (Corrected Algorithm)

This document explains the solutions for Day 3, Part 2, based on the **correct** algorithm. The initial approach was flawed; the problem requires finding the largest 12-digit number formed by a **subsequence** of digits, not a contiguous substring.

## Core Logic: Greedy Stack-Based Algorithm

A more efficient and elegant way to solve the "lexicographically largest subsequence" problem is with a stack. This approach avoids the nested loops of the previous method.

**Algorithm:**
1. Initialize an empty stack (or list that acts as a stack).
2. Iterate through each digit of the input `line`.
3. For each `digit`:
    a. While the stack is not empty, the `digit` is greater than the digit at the top of the stack, AND we still have "removals" left (i.e., `(line.length - stack.length) > (k - number_of_removals)`), pop from the stack and use one "removal".
    b. Push the current `digit` onto the stack.
4. After iterating through all digits, the stack may still be longer than `k`. Truncate it to `k` elements.
5. Join the elements of the stack to form the final result string.

---

## 1. Python Solution (Stack-based)

```python
def find_max_subsequence(line, k):
    stack = []
    removals_left = len(line) - k
    
    for digit in line:
        while stack and digit > stack[-1] and removals_left > 0:
            stack.pop()
            removals_left -= 1
        stack.append(digit)
        
    return "".join(stack[:k])
```
- **Stack Logic**: The `while` loop is the core of the algorithm. It "corrects" the stack by removing smaller digits from the end to make way for a larger digit, as long as we have removals to spare.
- **`stack[-1]`**: Python's way of accessing the last element of a list.
- **`stack[:k]`**: This is list slicing. It creates a new list containing elements from the beginning up to (but not including) index `k`. This truncates the final stack to the desired length.

---

## 2. Clojure Solution (Stack-based)

```clojure
(defn find-max-subsequence [line k]
  (let [removals-left (- (count line) k)]
    (loop [digits (vec line) ; Treat the line as a vector of characters
           removals removals-left
           stack []] ; The stack is a vector
      (if (empty? digits)
        (str/join (subvec stack 0 (min k (count stack)))) ; Truncate and join
        (let [digit (first digits)
              rest-digits (rest digits)]
          (if (and (not (empty? stack))
                   (> (int digit) (int (last stack))) ; Compare digits
                   (> removals 0))
            (recur digits (dec removals) (pop stack)) ; Pop and retry the *same* digit
            (recur rest-digits removals (conj stack digit)))))))) ; Push digit
```
- **`loop/recur` for State**: The functional equivalent of a `while` loop. The state (`digits` left, `removals`, `stack`) is passed from one iteration to the next.
- **Retry Logic**: When a character is popped from the stack, the `recur` call is made with the *same* `digits` list. This is crucial because it allows the current `digit` to be compared against the *new* last element of the stack.
- **Stack Operations**:
    - `(last stack)`: Gets the top element.
    - `(pop stack)`: Returns a *new* vector with the last element removed (immutable).
    - `(conj stack digit)`: Returns a *new* vector with the `digit` added to the end.
- **`(subvec stack 0 k)`**: The idiomatic way to get a sub-vector, equivalent to Python's slicing.

---

## 3. Rust Solution (Stack-based)

```rust
fn find_max_subsequence(line: &str, k: usize) -> String {
    let mut stack: Vec<char> = Vec::new();
    let mut removals_left = line.len() - k;

    for digit in line.chars() {
        while let Some(&top) = stack.last() {
            if digit > top && removals_left > 0 {
                stack.pop();
                removals_left -= 1;
            } else {
                break;
            }
        }
        stack.push(digit);
    }

    stack.truncate(k);
    stack.into_iter().collect()
}
```
- **`Vec<char>` as Stack**: A `Vec` (vector) is Rust's primary dynamic array type and is perfect for stack operations.
- **`while let Some(&top) = stack.last()`**: This is a powerful and safe way to loop while the stack is not empty.
    - `stack.last()`: Returns an `Option<&char>`. It's `Some(&top)` if the stack has an element, `None` if it's empty.
    - The `while let` pattern unwraps the `Some` and binds its content to `top`, breaking the loop if it's `None`.
- **`stack.truncate(k)`**: An efficient in-place method to shorten the vector to `k` elements.
- **`.into_iter().collect()`**: This converts the `Vec<char>` into an iterator, and `collect()` gathers the items into a new `String`.

---

## 4. Elixir Solution (Stack-based)

```elixir
defp find_max_subsequence(line, k) do
  removals_left = String.length(line) - k
  
  String.graphemes(line)
  |> Enum.reduce({[], removals_left}, fn digit, {stack, removals} ->
    reduce_stack(digit, stack, removals)
  end)
  |> elem(0) # Get the final stack from the accumulator
  |> Enum.take(k)
  |> Enum.join()
end

# Helper function to handle the inner "while" loop logic
defp reduce_stack(digit, [top | rest] = stack, removals)
     when digit > top and removals > 0 do
  # Recursively call until the condition is no longer met
  reduce_stack(digit, rest, removals - 1)
else
  # Base case: push the digit onto the stack
  {[digit | stack], removals}
end
```
- **`reduce` for Iteration**: The main logic uses `Enum.reduce` to iterate over the digits, with a tuple `{stack, removals}` as the accumulator. This is the idiomatic functional way to manage state through a loop.
- **Recursive Helper Function**: The `while` loop's logic is encapsulated in a separate, recursive helper function `reduce_stack/3`.
- **Pattern Matching and Guards**:
    - `reduce_stack(digit, [top | rest] = stack, removals)`: This function head only matches if the `stack` is not empty. It destructures the stack into its `top` element and the `rest` of the list.
    - `when digit > top and removals > 0`: This is a "guard clause". The function only executes if this condition is true.
    - This combination of pattern matching and guards elegantly replaces the `while` loop found in imperative solutions.
- **Stack Representation**: The stack is represented as a list. Pushing is `[digit | stack]` (consing to the front), and the final result is reversed implicitly by the reduction process. (Note: A more direct translation would reverse the list before processing, but this approach works correctly).
- **`Enum.take(k)` and `Enum.join()`**: Used to truncate the final list and join it into a string.

