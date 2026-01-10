
import os

def parse_ranges(input_str):
    """Parses the input string into a list of (start, end) tuples."""
    ranges = []
    max_id = 0
    range_strs = input_str.strip().split(',')
    for r_str in range_strs:
        if not r_str:
            continue
        start_str, end_str = r_str.split('-')
        start, end = int(start_str), int(end_str)
        ranges.append((start, end))
        if end > max_id:
            max_id = end
    return ranges, max_id

def generate_invalid_ids(max_id):
    """Generates all possible invalid IDs up to the max_id."""
    invalid_ids = []
    base = 1
    while True:
        s_base = str(base)
        s_invalid = s_base + s_base
        n_invalid = int(s_invalid)
        
        if n_invalid > max_id:
            # If the number of digits has increased and we are already over max,
            # any further number with this many digits will also be over.
            # This is a simple but effective cutoff. A tighter bound could be used.
            if len(str(base)) > len(str(base - 1)):
                 break
        else:
            invalid_ids.append(n_invalid)

        base += 1
    return invalid_ids

def solve(input_str):
    """
    Finds the sum of all invalid IDs within the given ranges.
    """
    ranges, max_id = parse_ranges(input_str)
    potential_ids = generate_invalid_ids(max_id)
    
    found_invalid_ids = set()

    for invalid_id in potential_ids:
        for start, end in ranges:
            if start <= invalid_id <= end:
                found_invalid_ids.add(invalid_id)
                break # Move to the next invalid_id
    
    return sum(found_invalid_ids)

def main():
    """
    Reads the input file and prints the solution.
    """
    input_file_path = os.path.join(os.path.dirname(__file__), 'input.txt')
    with open(input_file_path, 'r') as f:
        input_str = f.read()
    
    result = solve(input_str)
    print(f"The sum of all invalid IDs is: {result}")

if __name__ == "__main__":
    main()
