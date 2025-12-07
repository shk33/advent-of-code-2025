import os

def solve(rotations):
    """
    Calculates the password based on a list of rotation instructions.
    """
    current_position = 50
    zero_count = 0

    for line in rotations:
        line = line.strip()
        if not line:
            continue

        direction = line[0]
        distance = int(line[1:])

        if direction == 'R':
            current_position += distance
        elif direction == 'L':
            current_position -= distance
        
        current_position %= 100

        if current_position == 0:
            zero_count += 1
    
    return zero_count

def main():
    """
    Reads the input file and prints the solution.
    """
    input_file_path = os.path.join(os.path.dirname(__file__), 'input1.txt')
    
    with open(input_file_path, 'r') as f:
        rotations = f.readlines()
        
    password = solve(rotations)
    print(f"The password is: {password}")

if __name__ == "__main__":
    main()
