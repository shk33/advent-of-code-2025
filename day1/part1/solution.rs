use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn solve(rotations: Vec<String>) -> i32 {
    let mut current_position = 50;
    let mut zero_count = 0;

    for line in rotations {
        if line.is_empty() {
            continue;
        }

        let direction = line.chars().next().unwrap();
        let distance: i32 = line[1..].parse().unwrap();

        if direction == 'R' {
            current_position += distance;
        } else if direction == 'L' {
            current_position -= distance;
        }

        current_position = (current_position % 100 + 100) % 100;

        if current_position == 0 {
            zero_count += 1;
        }
    }

    zero_count
}

fn main() -> io::Result<()> {
    let path = Path::new("day1/part1/input1.txt");
    let file = File::open(&path)?;
    let lines = io::BufReader::new(file).lines();

    let rotations: Vec<String> = lines.map(|line| line.unwrap()).collect();
    
    let password = solve(rotations);
    println!("The password is: {}", password);

    Ok(())
}