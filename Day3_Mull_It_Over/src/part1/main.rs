use regex::Regex;
use std::fs;
use std::error::Error;
use shared::mul;

fn main() -> Result<(), Box<dyn Error>> {
    let re = Regex::new(r"mul\((\d+),(\d+)\)")?;
    let hay = fs::read_to_string("../../input.txt")?;

    let mut sum = 0;
    for cap in re.captures_iter(&hay) {
        if let (Some(x), Some(y)) = (cap.get(1), cap.get(2)) {
            let x: i32 = x.as_str().parse()?;
            let y: i32 = y.as_str().parse()?;
            let result = mul(x, y);
            sum += result;
        }
    }
    
    println!("{}", sum);
    Ok(())
}
