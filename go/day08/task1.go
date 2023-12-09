package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Node struct {
	left  string
	right string
	name string
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
	 
	steps := 0
	curr_name := "AAA"
	curr_node := nodes[curr_name]
	instruction_pos := 0
	end_reached := false
	for !end_reached {
		curr_node = nodes[curr_name]
		instruction := instructions[instruction_pos]
		if instruction == 'R' {
			curr_name = curr_node.right
		}
		if instruction == 'L' {
			curr_name = curr_node.left
		}
		instruction_pos = (instruction_pos + 1) % len(instructions)
		steps++
		if curr_name == "ZZZ" {
			end_reached = true
		}
	}
	fmt.Println(steps)
}