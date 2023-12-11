use std::collections::HashMap;
use std::collections::HashSet;
use std::io::{self, BufRead};

#[derive(Hash, Eq, PartialEq, Debug, Copy, Clone)]
struct Pos {
    r: i32,
    c: i32,
}

#[derive(Default)]
struct Graph {
    edges: HashMap<Pos, HashSet<Pos>>,
}

impl Graph {
    fn new() -> Self {
        Self::default()
    }

    fn add_edge(&mut self, u: Pos, v: Pos) {
        self.edges.entry(u).or_default().insert(v);
    }

    fn dfs(&mut self, current: Pos, start: Pos, visited: &mut HashSet<Pos>, main_loop: &mut HashSet<Pos>, depth: usize) -> usize {
        visited.insert(current);
        let mut max_depth = depth;
        if let Some(neighbours) = self.edges.get(&current).cloned() {
            for &nbr in &neighbours {
                if nbr == start && max_depth > 4 {
                    main_loop.clear();
                    main_loop.extend(visited.iter());
                    return depth + 1;
                }
                if !visited.contains(&nbr) {
                    max_depth = max_depth.max(self.dfs(nbr, start, visited, main_loop, depth + 1));
                }
            }
        }
        visited.remove(&current);
        max_depth
    }
    
}

fn new_pos(r: i32, c: i32) -> Pos {
    Pos { r, c }
}

fn is_inside_increment(pos: Pos, loop_: &HashSet<Pos>, g: &Graph, m: i32) -> f64 {
    let mut crosses = 0.0;
    if loop_.contains(&pos) {
        let is_below = g.edges.get(&pos).map(|e| e.contains(&new_pos(pos.r - 1, pos.c))).unwrap_or(false);
        let is_above = g.edges.get(&pos).map(|e| e.contains(&new_pos(pos.r + 1, pos.c))).unwrap_or(false);
        let is_prev = g.edges.get(&pos).map(|e| e.contains(&new_pos(pos.r, pos.c - 1))).unwrap_or(false);
        let is_next = g.edges.get(&pos).map(|e| e.contains(&new_pos(pos.r, pos.c + 1))).unwrap_or(false);
        if is_below && is_above {
            crosses += 1.0;
        } else if is_below && is_next {
            crosses += 0.5;
        } else if is_below && is_prev {
            crosses -= 0.5;
        } else if is_above && is_next {
            crosses -= 0.5;
        } else if is_above && is_prev {
            crosses += 0.5;
        }
    }
    crosses
}

fn main() {
    let stdin = io::stdin();
    let buffer: Vec<String> = stdin.lock().lines().filter_map(Result::ok).collect();
    let m = buffer[0].len() as i32;
    let n = buffer.len() as i32;

    let mut g = Graph::new();
    let mut animal = Pos { r: 0, c: 0 };
    for (i, line) in buffer.iter().enumerate() {
        for (j, char) in line.chars().enumerate() {
            let pos = new_pos(i as i32, j as i32);
            match char {
                'S' => animal = pos,
                '|' => {
                    g.add_edge(pos, new_pos(i as i32 - 1, j as i32));
                    g.add_edge(pos, new_pos(i as i32 + 1, j as i32));
                }
                '-' => {
                    g.add_edge(pos, new_pos(i as i32, j as i32 - 1));
                    g.add_edge(pos, new_pos(i as i32, j as i32 + 1));
                }
                'L' => {
                    g.add_edge(pos, new_pos(i as i32 - 1, j as i32));
                    g.add_edge(pos, new_pos(i as i32, j as i32 + 1));
                }
                'J' => {
                    g.add_edge(pos, new_pos(i as i32 - 1, j as i32));
                    g.add_edge(pos, new_pos(i as i32, j as i32 - 1));
                }
                '7' => {
                    g.add_edge(pos, new_pos(i as i32, j as i32 - 1));
                    g.add_edge(pos, new_pos(i as i32 + 1, j as i32));
                }
                'F' => {
                    g.add_edge(pos, new_pos(i as i32, j as i32 + 1));
                    g.add_edge(pos, new_pos(i as i32 + 1, j as i32));
                }
                _ => {}
            }
        }
    }

    let deltas = vec![new_pos(-1, 0), new_pos(1, 0), new_pos(0, -1), new_pos(0, 1)];
    for &delta in &deltas {
        let nbr = new_pos(animal.r + delta.r, animal.c + delta.c);
        if g.edges.contains_key(&nbr) && g.edges[&nbr].contains(&animal) {
            g.add_edge(animal, nbr);
        }
    }

    let mut visited = HashSet::new();
    let mut longest_loop = HashSet::new();
    g.dfs(animal, animal, &mut visited, &mut longest_loop, 0);

    let mut total = 0;
    for i in 0..n {
        let mut cross_value = 1000.0;
        for j in 0..m {
            let pos = new_pos(i, j);
            let increment = is_inside_increment(pos, &longest_loop, &g, m);
            cross_value += increment;
            if !longest_loop.contains(&pos) && (cross_value as i32) % 2 == 1 {
                total += 1;
            }
        }
    }

    println!("{}", total);
}


