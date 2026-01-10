import os

def find_max_subsequence(line, k):
    """
    Finds the lexicographically largest subsequence of length k.
    """
    if len(line) < k or k == 0:
        return "0" * k

    result_chars = []
    current_start_index = 0
    
    for i in range(k):
        remaining_to_find = k - i
        search_end_index = len(line) - remaining_to_find
        
        best_digit = -1
        best_digit_index = -1

        # Find the best digit in the current search window
        for j in range(current_start_index, search_end_index + 1):
            digit = int(line[j])
            if digit > best_digit:
                best_digit = digit
                best_digit_index = j
        
        result_chars.append(str(best_digit))
        current_start_index = best_digit_index + 1

    return "".join(result_chars)

def solve():
    """
    Solves the Day 3, Part 2 puzzle.
    """
    input_path = os.path.join(os.path.dirname(__file__), 'input.txt')
    with open(input_path, 'r') as f:
        lines = f.readlines()

    total_joltage = 0
    k = 12 # Number of digits to select

    for line in lines:
        line = line.strip()
        if not line:
            continue
        
        max_line_joltage_str = find_max_subsequence(line, k)
        total_joltage += int(max_line_joltage_str)
        
    print(f"The total output joltage is: {total_joltage}")

if __name__ == "__main__":
    solve()
