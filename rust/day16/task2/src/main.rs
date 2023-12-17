use std::collections::HashMap;
use std::io::{self, BufRead};

#[derive(Clone, Copy, Hash, Eq, PartialEq)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

#[derive(Clone, Copy)]
struct Node {
    value: char,
    energized: bool,
}

#[derive(Hash, Eq, PartialEq)]
struct NodeKey {
    x: usize,
    y: usize,
    dir: Direction,
}

struct Graph {
    nodes: Vec<Vec<Node>>,
    visited: HashMap<NodeKey, bool>,
}

impl Graph {
    fn new(lines: &[String]) -> Self {
        let nodes: Vec<Vec<_>> = lines
            .iter()
            .map(|line| {
                line.chars()
                    .map(|c| Node {
                        value: c,
                        energized: false,
                    })
                    .collect()
            })
            .collect();
        Self {
            nodes,
            visited: HashMap::new(),
        }
    }

    fn energize(&mut self, x: usize, y: usize, dir: Direction) {
        if x >= self.nodes[0].len() || y >= self.nodes.len() {
            return;
        }
        let key = NodeKey { x, y, dir };
        if self.visited.contains_key(&key) {
            return;
        }
        self.visited.insert(key, true);
        let node = &mut self.nodes[y][x];
        node.energized = true;
        match node.value {
            '.' => self.move_straight(x, y, dir),
            '/' => self.move_reflected(x, y, dir, '/'),
            '\\' => self.move_reflected(x, y, dir, '\\'),
            '|' => self.split_beam(x, y, dir, '|'),
            '-' => self.split_beam(x, y, dir, '-'),
            _ => (),
        }
    }

    fn move_straight(&mut self, x: usize, y: usize, dir: Direction) {
        match dir {
            Direction::Up => self.energize(x, y.checked_sub(1).unwrap_or(0), Direction::Up),
            Direction::Down => self.energize(x, y.saturating_add(1), Direction::Down),
            Direction::Left => self.energize(x.checked_sub(1).unwrap_or(0), y, Direction::Left),
            Direction::Right => self.energize(x.saturating_add(1), y, Direction::Right),
        }
    }

    fn move_reflected(&mut self, x: usize, y: usize, dir: Direction, mirror: char) {
        match dir {
            Direction::Up => {
                if mirror == '/' {
                    self.energize(x.saturating_add(1), y, Direction::Right);
                } else {
                    self.energize(x.checked_sub(1).unwrap_or(0), y, Direction::Left);
                }
            }
            Direction::Down => {
                if mirror == '/' {
                    self.energize(x.checked_sub(1).unwrap_or(0), y, Direction::Left);
                } else {
                    self.energize(x.saturating_add(1), y, Direction::Right);
                }
            }
            Direction::Left => {
                if mirror == '/' {
                    self.energize(x, y.saturating_add(1), Direction::Down);
                } else {
                    self.energize(x, y.checked_sub(1).unwrap_or(0), Direction::Up);
                }
            }
            Direction::Right => {
                if mirror == '/' {
                    self.energize(x, y.checked_sub(1).unwrap_or(0), Direction::Up);
                } else {
                    self.energize(x, y.saturating_add(1), Direction::Down);
                }
            }
        }
    }

    fn split_beam(&mut self, x: usize, y: usize, dir: Direction, splitter: char) {
        match dir {
            Direction::Up => {
                if splitter == '-' {
                    self.energize(x.saturating_add(1), y, Direction::Right);
                    self.energize(x.checked_sub(1).unwrap_or(0), y, Direction::Left);
                } else {
                    self.energize(x, y.checked_sub(1).unwrap_or(0), Direction::Up);
                }
            }
            Direction::Down => {
                if splitter == '-' {
                    self.energize(x.saturating_add(1), y, Direction::Right);
                    self.energize(x.checked_sub(1).unwrap_or(0), y, Direction::Left);
                } else {
                    self.energize(x, y.saturating_add(1), Direction::Down);
                }
            }
            Direction::Left => {
                if splitter == '|' {
                    self.energize(x, y.saturating_add(1), Direction::Down);
                    self.energize(x, y.checked_sub(1).unwrap_or(0), Direction::Up);
                } else {
                    self.energize(x.checked_sub(1).unwrap_or(0), y, Direction::Left);
                }
            }
            Direction::Right => {
                if splitter == '|' {
                    self.energize(x, y.saturating_add(1), Direction::Down);
                    self.energize(x, y.checked_sub(1).unwrap_or(0), Direction::Up);
                } else {
                    self.energize(x.saturating_add(1), y, Direction::Right);
                }
            }
        }
    }

    fn count_energized(&self) -> usize {
        self.nodes.iter().flatten().filter(|n| n.energized).count()
    }
}


fn main() {
    let stdin = io::stdin();
    let lines: Vec<String> = stdin
        .lock()
        .lines()
        .map(|line| line.unwrap())
        .collect();

    let size_x = lines[0].len();
    let size_y = lines.len();

    let mut energized_max = 0;
    for x in 0..size_x {
        let mut graph = Graph::new(&lines);
        graph.energize(x, 0, Direction::Down);
        let energized = graph.count_energized();
        energized_max = energized_max.max(energized);

        let mut graph = Graph::new(&lines);
        graph.energize(x, size_y - 1, Direction::Up);
        let energized = graph.count_energized();
        energized_max = energized_max.max(energized);
    }

    for y in 0..size_y {
        let mut graph = Graph::new(&lines);
        graph.energize(0, y, Direction::Right);
        let energized = graph.count_energized();
        energized_max = energized_max.max(energized);

        let mut graph = Graph::new(&lines);
        graph.energize(size_x - 1, y, Direction::Left);
        let energized = graph.count_energized();
        energized_max = energized_max.max(energized);
    }

    println!("{}", energized_max);
}
