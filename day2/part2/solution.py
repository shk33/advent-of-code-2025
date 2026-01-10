
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
    """Generates all possible invalid IDs up to the max_id (repeated at least twice)."""
    invalid_ids = []
    base_num = 1

    while True:
        base_s = str(base_num)
        
        # Optimization: if the smallest possible repetition (base_s + base_s) is too big, stop.
        first_repetition_val = int(base_s + base_s)
        if first_repetition_val > max_id:
            break
        
        current_repeated_s = base_s
        
        # Generate repetitions: base_s + base_s, base_s + base_s + base_s, etc.
        while True:
            current_repeated_s += base_s # Append base_s again (e.g., "1"->"11", "11"->"111")
            
            # The problem example "1111111 (1 seven times)" implies base_s itself can be repeated.
            # But "123123123 (123 three times)" implies (base_s repeated) + base_s.
            # My current logic:
            # base_s = "1", current_repeated_s starts as "1". First += base_s makes it "11".
            # Second += base_s makes it "111". This is correct.
            # base_s = "123", current_repeated_s starts as "123". First += base_s makes it "123123".
            # Second += base_s makes it "123123123". This is correct.

            n_invalid = int(current_repeated_s)

            if n_invalid > max_id:
                break # This base_s repetition is too long
            
            invalid_ids.append(n_invalid)
        
        base_num += 1
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
