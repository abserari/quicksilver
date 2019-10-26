package main

import "log"

func main() {
	var s string = "(())))(())()()"

	n := longestValidParentheses(s)
	log.Println(n)
}

func longestValidParentheses(s string) int {
	stack := []int{-1}
	temp := 0
	longest := 0

	for pos, char := range s {
		if char == '(' {
			stack = append(stack, pos)
		} else {
			stack = stack[0 : len(stack)-1]
			if len(stack) > 0 {
				temp = pos - stack[len(stack)-1]
				log.Println(pos, "-", +stack[len(stack)-1], "=", temp)
				if temp > longest {
					longest = temp
				}
			} else {
				stack = append(stack, pos)
				log.Println(stack)
			}
		}
	}
	return longest
}
