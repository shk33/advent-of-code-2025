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
    let mut base: u64 = 1;
    loop {
        let s_base = base.to_string();
        let s_invalid = s_base.clone() + &s_base;
        let n_invalid: u64 = match s_invalid.parse() {
            Ok(n) => n,
            Err(_) => break, // Should not happen with this logic
        };

        if n_invalid > max_id {
            if s_base.len() > (base -1).to_string().len() && base > 1 {
                 break;
            }
        }
        
        invalid_ids.push(n_invalid);
        base += 1;
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
    let input_str = fs::read_to_string("day2/part1/input.txt")?;
    let result = solve(&input_str);
    println!("The sum of all invalid IDs is: {}", result);
    Ok(())
}