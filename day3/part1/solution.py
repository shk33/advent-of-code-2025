import os

def solve():
    """
    Solves the Day 3, Part 1 puzzle.
    """
    input_path = os.path.join(os.path.dirname(__file__), 'input.txt')
    with open(input_path, 'r') as f:
        lines = f.readlines()

    total_joltage = 0
    for line in lines:
        line = line.strip()
        if not line:
            continue
        
        max_line_joltage = 0
        for i in range(len(line)):
            for j in range(i + 1, len(line)):
                joltage_str = line[i] + line[j]
                joltage = int(joltage_str)
                if joltage > max_line_joltage:
                    max_line_joltage = joltage
        total_joltage += max_line_joltage
        
    print(f"The total output joltage is: {total_joltage}")

if __name__ == "__main__":
    solve()
