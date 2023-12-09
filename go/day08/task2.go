package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
  
func gcd(a, b int) int {  
	for b != 0 {  
		t := b  
		b = a % b  
		a = t  
	}  
	return a  
}  
    
func lcm(a, b int) int {  
	return a * b / gcd(a, b)  
}  
   
func lcmOfList(nums []int) int {  
	result := nums[0]  
	for _, num := range nums[1:] {  
		result = lcm(result, num)  
	}  
	return result  
}  

type Node struct {
	left  string
	right string
	name string
}

type Position struct {
	name string
	node *Node
}

func (p *Position) move(direction string, nodes map[string]*Node) {
	if direction == "R" {
		p.name = p.node.right
	}
	if direction == "L" {
		p.name = p.node.left
	}
	p.node = nodes[p.name]
}

func main() {
	nodes := make(map[string]*Node)  
	var instructions string  
	
	scanner := bufio.NewScanner(os.Stdin)   
	scanner.Scan()  
	instructions = scanner.Text()  
	scanner.Scan()  
	for scanner.Scan() {  
		line := scanner.Text()  
		parts := strings.Split(line, " = ")  
		nodeName := parts[0]  
		links := strings.Trim(parts[1], "()")  
		linkNames := strings.Split(links, ", ")  
	
		node := &Node{name: nodeName}  
		nodes[nodeName] = node  
		node.left = linkNames[0]
		node.right = linkNames[1]
	}  
	
	if err := scanner.Err(); err != nil {  
		fmt.Println("Error reading input:", err)  
	}  
	 
	positions := make([]Position, 0)
	for _, node := range nodes {
		if strings.HasSuffix(node.name, "A") {
			positions = append(positions, Position{node: node, name: node.name})
		}
	}

	cycle_lengths := make([]int, len(positions))
	for i := 0; i < len(positions); i++ {
		instruction_pos := 0
		end_reached := false
		for !end_reached {
			instruction := instructions[instruction_pos]
			positions[i].move(string(instruction), nodes)	
			instruction_pos = (instruction_pos + 1) % len(instructions)
			cycle_lengths[i]++
			if strings.HasSuffix(positions[i].name, "Z") {
				end_reached = true
			}
		}
	}
	fmt.Println(lcmOfList(cycle_lengths))
}