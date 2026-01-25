use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn solve_line(line: &str) -> u32 {
    let mut max_line_joltage = 0;
    let chars: Vec<char> = line.chars().collect();
    for i in 0..chars.len() {
        for j in (i + 1)..chars.len() {
            let joltage_str = format!("{}{}", chars[i], chars[j]);
            let joltage: u32 = joltage_str.parse().unwrap();
            if joltage > max_line_joltage {
                max_line_joltage = joltage;
            }
        }
    }
    max_line_joltage
}

fn main() -> io::Result<()> {
    let path = Path::new("day3/part1/input.txt");
    let file = File::open(&path)?;
    let lines = io::BufReader::new(file).lines();

    let mut total_joltage = 0;
    for line in lines {
        if let Ok(ip) = line {
            total_joltage += solve_line(&ip);
        }
    }

    println!("The total output joltage is: {}", total_joltage);

    Ok(())
}