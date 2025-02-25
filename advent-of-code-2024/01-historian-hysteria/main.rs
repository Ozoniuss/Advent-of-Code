use std::fs;
fn main() {
    let contents = fs::read_to_string("input.txt").expect("could not read file to string");

    let mut left: Vec<i32> = Vec::new();
    let mut right: Vec<i32> = Vec::new();

    for line in contents.lines() {
        let parts: Vec<&str> = line.split("   ").collect();

        let leftnr: i32 = parts[0].trim().parse().expect("failed to parse left");
        let rightnr: i32 = parts[1].trim().parse().expect("failed to parse left");

        left.push(leftnr);
        right.push(rightnr);
    }

    left.sort_by(|a, b| b.cmp(a));
    right.sort_by(|a, b| b.cmp(a));

    assert_eq!(left.len(), right.len(), "vecs should have the same length");
    let mut diff = 0;

    for (l, r) in left.iter().zip(right.iter()) {
        diff += (l - r).abs();
    }

    println!("{diff}");
    diff = 0;
    for l in left {
        let lb: i32 = right
            .partition_point(|&x| x >= l + 1)
            .try_into()
            .expect("too big");
        let ub: i32 = right
            .partition_point(|&x| x >= l)
            .try_into()
            .expect("TOO BIG");

        println!("{ub} {lb}");

        diff += (ub - lb).abs() * l;
    }
    println!("{diff}");
}
