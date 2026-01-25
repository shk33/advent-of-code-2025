use std::collections::HashSet;
use std::fs;
use std::io;

fn parse_ranges(input_str: &str) -> (Vec<(u64, u64)>, u64) {
    let mut ranges = Vec::new();
    let mut max_id = 0;
    for r_str in input_str.trim().split(',') {
        if r_str.is_empty() {
            continue;
        }
        let parts: Vec<&str> = r_str.split('-').collect();
        let start = parts[0].parse().unwrap();
        let end = parts[1].parse().unwrap();
        ranges.push((start, end));
        if end > max_id {
            max_id = end;
        }
    }
    (ranges, max_id)
}

fn generate_invalid_ids(max_id: u64) -> Vec<u64> {
    let mut invalid_ids = Vec::new();
    let mut base_num: u64 = 1;

    loop {
        let base_s = base_num.to_string();
        
        let first_repetition_s = base_s.clone() + &base_s;
        let first_repetition_val: u64 = match first_repetition_s.parse() {
            Ok(n) => n,
            Err(_) => break,
        };

        if first_repetition_val > max_id {
            break;
        }
        
        let mut current_repeated_s = base_s.clone();
        
        loop {
            current_repeated_s += &base_s;
            
            let n_invalid: u64 = match current_repeated_s.parse() {
                Ok(n) => n,
                Err(_) => break,
            };

            if n_invalid > max_id {
                break;
            }
            
            invalid_ids.push(n_invalid);
        }
        
        base_num += 1;
    }
    invalid_ids
}

fn solve(input_str: &str) -> u64 {
    let (ranges, max_id) = parse_ranges(input_str);
    let potential_ids = generate_invalid_ids(max_id);
    
    let mut found_invalid_ids = HashSet::new();

    for &invalid_id in &potential_ids {
        for &(start, end) in &ranges {
            if invalid_id >= start && invalid_id <= end {
                found_invalid_ids.insert(invalid_id);
                break; 
            }
        }
    }
    
    found_invalid_ids.iter().sum()
}

fn main() -> io::Result<()> {
    let input_str = fs::read_to_string("day2/part2/input.txt")?;
    let result = solve(&input_str);
    println!("The sum of all invalid IDs is: {}", result);
    Ok(())
}
