use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    if let Ok(lines) = read_lines("./rust_input.txt") {
        let mut total_mass = 0;
        for line in lines {
            if let Ok(module) = line {
                total_mass = total_mass + (module.parse::<i32>().unwrap() / 3) - 2;
            }
        }
        println!("Total mass -> {}", total_mass);
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}