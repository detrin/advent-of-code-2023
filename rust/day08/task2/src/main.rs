use std::collections::HashMap;  
use std::io::{self, BufRead};  
use std::rc::Rc;  
use std::cell::RefCell;  
  
#[derive(Clone)]  
struct Node {  
    left: String,  
    right: String,  
    name: String,  
}  
  
struct Position {  
    name: String,  
    node: Rc<RefCell<Node>>,  
}  

impl Position {  
    fn move_position(&mut self, direction: char, nodes: &HashMap<String, Rc<RefCell<Node>>>) {  
        if direction == 'R' {  
            self.name = self.node.borrow().right.clone();  
        }  
        if direction == 'L' {  
            self.name = self.node.borrow().left.clone();  
        }  
        self.node = Rc::clone(&nodes[&self.name]);  
    }  
} 

fn gcd(a: usize, b: usize) -> usize {
    let mut a = a;
    let mut b = b;
    while b != 0 {
        let t = b;
        b = a % b;
        a = t;
    }
    a
}

fn lcm(a: usize, b: usize) -> usize {
    a * b / gcd(a, b)
}

fn lcm_of_list(nums: &[usize]) -> usize {
    let mut result = nums[0];
    for &num in &nums[1..] {
        result = lcm(result, num);
    }
    result
}

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();

    let instructions = lines.next().unwrap().unwrap();

    let _ = lines.next();

    let mut nodes: HashMap<String, Rc<RefCell<Node>>> = HashMap::new();  
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
        nodes.insert(node_name, Rc::new(RefCell::new(node)));  
    }  

    let mut positions: Vec<Position> = Vec::new();  
    for node in nodes.values() {  
        if node.borrow().name.ends_with("A") {  
            positions.push(Position {  
                node: Rc::clone(node),  
                name: node.borrow().name.clone(),  
            });  
        }  
    } 

    let mut cycle_lengths = vec![0; positions.len()];
    for i in 0..positions.len() {
        let mut instruction_pos = 0;
        let mut end_reached = false;
        while !end_reached {
            let instruction = instructions.chars().nth(instruction_pos).unwrap();
            positions[i].move_position(instruction, &nodes);
            instruction_pos = (instruction_pos + 1) % instructions.len();
            cycle_lengths[i] += 1;
            if positions[i].name.ends_with("Z") {
                end_reached = true;
            }
        }
    }

    println!("{}", lcm_of_list(&cycle_lengths));

    Ok(())
}
