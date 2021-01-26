package datastructure

func Calculate(s string) int {
	return cal(buildRPN(s))
}
func pop(a *[]int) int {
	if len(*a) == 0 {
		return 0
	}
	top := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return top
}
func cal(rpn []string) int {
	var nums []int
	for i := 0; i < len(rpn); i++ {
		switch rpn[i] {
		case "+":
			two, one := pop(&nums), pop(&nums)
			nums = append(nums, one+two)
		case "-":
			two, one := pop(&nums), pop(&nums)
			nums = append(nums, one-two)
		case "*":
			two, one := pop(&nums), pop(&nums)
			nums = append(nums, one*two)
		case "/":
			two, one := pop(&nums), pop(&nums)
			nums = append(nums, one/two)
		default:
			a := 0
			for j := 0; j < len(rpn[i]); j++ {
				a = a*10 + int(rpn[i][j]-'0')
			}
			nums = append(nums, a)
		}
	}
	return nums[0]
}

// Reverse Polish Notation
func buildRPN(s string) []string {
	var signed []byte
	var outer []string
	level := map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		current := ""
		for ; i < len(s) && s[i] <= '9' && s[i] >= '0'; i++ {
			current += string(s[i])
		}
		if len(current) != 0 {
			outer = append(outer, current)
		}
		if i >= len(s) {
			break
		}
		switch s[i] {
		case ' ':
			continue
		case '(':
			signed = append(signed, '(')
		case ')':
			for len(signed) != 0 && signed[len(signed)-1] != '(' {
				outer = append(outer, string(signed[len(signed)-1]))
				signed = signed[:len(signed)-1]
			}
			if len(signed) != 0 && signed[len(signed)-1] == '(' {
				signed = signed[:len(signed)-1]
			}
		default:
			for len(signed) != 0 && level[signed[len(signed)-1]] >= level[s[i]] {
				outer = append(outer, string(signed[len(signed)-1]))
				signed = signed[:len(signed)-1]
			}
			signed = append(signed, s[i])
		}
	}
	for len(signed) != 0 {
		outer = append(outer, string(signed[len(signed)-1]))
		signed = signed[:len(signed)-1]
	}
	return outer
}
