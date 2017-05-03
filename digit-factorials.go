// This is a solution to the digit factorial problem listed at:
// https://projecteuler.net/problem=34. This has been limited to 10^8, simply
// because it takes ~ 30 seconds for my laptop to complete this using that
// value.
package main

import "fmt"

func main() {
	var factorial_sum int
	for n := 3; n < 100000000; n++ {
		if digit_factorial(n) {
			factorial_sum += n
		}
	}
	fmt.Println(factorial_sum, "is the sum of all positive numbers, less than 10^8, which are equal to the sum of the factorial of their digits.")
}

func digit_factorial(num int) bool {
	// Using bool as a return value as we want a yes/no answer if the argument
	// being passed in equals the sum of the factorial of each of its digits.
	var result int
	n := num

	// Would prefer to convert num to []int, but I technically haven't
	// gotten that far yet in Todd's course, so using division to modify the
	// control variable (n), and modulus to get the next digit in num.
	for n > 0 {
		result += factorial(n % 10)
		n /= 10
	}

	if result == num {
		return true
	}

	return false
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
