
import os
import math

def solve(rotations):
    """
    Calculates the password based on the number of times the dial clicks on 0.
    """
    current_position = 50
    total_zero_count = 0

    for line in rotations:
        line = line.strip()
        if not line:
            continue

        direction = line[0]
        distance = int(line[1:])
        
        zeros_this_turn = 0
        if direction == 'R':
            # In Python, // is floor division, which works correctly for positives and negatives.
            zeros_this_turn = (current_position + distance) // 100 - current_position // 100
        elif direction == 'L':
            zeros_this_turn = (current_position - 1) // 100 - (current_position - distance - 1) // 100
        
        total_zero_count += zeros_this_turn
        
        # Update position for the next turn
        if direction == 'R':
            current_position += distance
        elif direction == 'L':
            current_position -= distance

    # The final position state is not part of the password, only the running count of clicks.
    # The modulo for the next turn's start is implicitly handled by the number line logic.
    # For clarity, let's keep track of the final dial position separately.
    # The problem can be solved without ever using a modulo operator.
    
    return total_zero_count

def main():
    """
    Reads the input file and prints the solution.
    """
    # Path is relative to this script, going up one level and into part1
    input_file_path = os.path.join(os.path.dirname(__file__), '../part1/input1.txt')
    
    with open(input_file_path, 'r') as f:
        rotations = f.readlines()
        
    password = solve(rotations)
    print(f"The password is: {password}")

if __name__ == "__main__":
    main()
