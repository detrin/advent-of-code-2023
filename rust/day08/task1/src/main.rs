use std::collections::HashMap;
use std::io::{self, BufRead};

struct Node {
    left: String,
    right: String,
    name: String,
}

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();

    let instructions = lines.next().unwrap().unwrap();

    let _ = lines.next();

    let mut nodes: HashMap<String, Node> = HashMap::new();
    for line in lines {
        let line = line.unwrap();
        let parts: Vec<&str> = line.split(" = ").collect();
        let node_name = parts[0].to_string();
        let links: Vec<&str> = parts[1].trim_matches(|c| c == '(' || c == ')')
            .split(", ").collect();
        let node = Node {
            name: node_name.clone(),
            left: links[0].to_string(),
            right: links[1].to_string(),
        };
        nodes.insert(node_name, node);
    }

    let mut steps = 0;
    let mut curr_name = String::from("AAA");
    let mut instruction_pos = 0;
    let end_reached = false;

    while !end_reached {
        let curr_node = nodes.get(&curr_name).unwrap();
        let instruction = instructions.chars().nth(instruction_pos).unwrap();
        if instruction == 'R' {
            curr_name = curr_node.right.clone();
        }
        if instruction == 'L' {
            curr_name = curr_node.left.clone();
        }
        instruction_pos = (instruction_pos + 1) % instructions.len();
        steps += 1;
        if curr_name == "ZZZ" {
            break;
        }
    }

    println!("{}", steps);

    Ok(())
}
