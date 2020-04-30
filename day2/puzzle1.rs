use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    if let Ok(lines) = read_lines("./rust_input.txt") {
        for line in lines {
            if let Ok(intcode_line) = line {
                let intcodes: Vec<i32> = intcode_line
                    .split(',')
                    .map(|s| s.parse().expect("parse error"))
                    .collect();
                run_intcodes(intcodes);
            }
        }
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn run_intcodes(mut ic: std::vec::Vec<i32>){
    let mut pointer = 0;
    loop{
        if ic[pointer] == 99 {
            break;
        } else if ic[pointer] == 1 {
            let arg1 = ic[pointer+1] as usize;
            let arg2 = ic[pointer+2] as usize;
            let answer = ic[pointer+3] as usize;
            ic[answer] = ic[arg1] + ic[arg2];
            pointer += 4;
        } else if ic[pointer] == 2 {
            let arg1 = ic[pointer+1] as usize;
            let arg2 = ic[pointer+2] as usize;
            let answer = ic[pointer+3] as usize;
            ic[answer] = ic[arg1] * ic[arg2];
            pointer += 4;
        } else {
            println!("Invalid opcode -> {}", ic[pointer]);
            break;
        }
    }
    println!("Final array -> {:?}", ic);
}