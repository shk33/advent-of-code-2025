import os

def solve_line(line):
    """
    Finds the largest 12-digit number in a line.
    """
    n = len(line)
    k = 12
    to_remove = n - k
    res = []

    for digit in line:
        while res and digit > res[-1] and to_remove > 0:
            res.pop()
            to_remove -= 1
        res.append(digit)

    while to_remove > 0:
        res.pop()
        to_remove -= 1
        
    return int("".join(res))

def solve():
    """
    Solves the Day 3, Part 2 puzzle.
    """
    input_path = os.path.join(os.path.dirname(__file__), 'input.txt')
    with open(input_path, 'r') as f:
        lines = f.readlines()

    total_joltage = 0
    for line in lines:
        line = line.strip()
        if not line:
            continue
        
        total_joltage += solve_line(line)
        
    print(f"The new total output joltage is: {total_joltage}")

if __name__ == "__main__":
    solve()