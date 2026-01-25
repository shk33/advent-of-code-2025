use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn solve(rotations: Vec<String>) -> i64 {
    let mut current_position: i64 = 50;
    let mut zero_count: i64 = 0;

    for line in rotations {
        if line.is_empty() {
            continue;
        }

        let direction = line.chars().next().unwrap();
        let distance: i64 = line[1..].parse().unwrap();

        if direction == 'R' {
            zero_count += (current_position + distance).div_euclid(100) - current_position.div_euclid(100);
            current_position += distance;
        } else if direction == 'L' {
            zero_count += (current_position - 1).div_euclid(100) - (current_position - distance - 1).div_euclid(100);
            current_position -= distance;
        }
    }

    zero_count
}

fn main() -> io::Result<()> {
    let path = Path::new("day1/part2/input1.txt");
    let file = File::open(&path)?;
    let lines = io::BufReader::new(file).lines();

    let rotations: Vec<String> = lines.map(|line| line.unwrap()).collect();
    
    let password = solve(rotations);
    println!("The password is: {}", password);

    Ok(())
}