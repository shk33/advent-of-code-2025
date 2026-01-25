use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn solve_line(line: &str) -> u128 {
    let n = line.len();
    let k = 12;
    let mut to_remove = n - k;
    let mut res: Vec<char> = Vec::new();

    for c in line.chars() {
        while !res.is_empty() && c > *res.last().unwrap() && to_remove > 0 {
            res.pop();
            to_remove -= 1;
        }
        res.push(c);
    }

    while to_remove > 0 {
        res.pop();
        to_remove -= 1;
    }
    
    let res_str: String = res.iter().collect();
    res_str.parse().unwrap()
}

fn main() -> io::Result<()> {
    let path = Path::new("day3/part2/input.txt");
    let file = File::open(&path)?;
    let lines = io::BufReader::new(file).lines();

    let mut total_joltage: u128 = 0;
    for line in lines {
        if let Ok(ip) = line {
            if !ip.is_empty() {
                total_joltage += solve_line(&ip);
            }
        }
    }

    println!("The new total output joltage is: {}", total_joltage);

    Ok(())
}