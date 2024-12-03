use regex::Regex;
use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    let re = Regex::new(r"(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))")?;
    let hay = fs::read_to_string("../../input.txt");

    let mut mul_enabled = true; 
    let mut result_sum = 0;

    for cap in re.captures_iter(&hay) {
        if let Some(do_match) = cap.get(4) {
            mul_enabled = true;
        } else if let Some(dont_match) = cap.get(5) {
            mul_enabled = false;
        } else if let (Some(_), Some(x), Some(y)) = (cap.get(1), cap.get(2), cap.get(3)) {
            if mul_enabled {
                let x: i32 = x.as_str().parse()?;
                let y: i32 = y.as_str().parse()?;
                let result = mul(x, y);
                result_sum += result;
            }
        }
    }

    println!("Sum of enabled mul results: {}", result_sum);
    Ok(())
}
