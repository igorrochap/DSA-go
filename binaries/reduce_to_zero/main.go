package main

func numberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		steps++
		if num&1 == 1 { // checks if number is odd by it's parity bit
			num -= 1
			continue
		}
		num >>= 1 //right shift by one means divide by 2 one time
	}
	return steps
}

func main() {
	testcases := []int{14, 8, 123}
	for _, testcase := range testcases {
		println(numberOfSteps(testcase))
	}
}
