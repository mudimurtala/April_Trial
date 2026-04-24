package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = i
	}

	for {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(digits[i]) + '0')
		}

		last := true
		for i := 0; i < n; i++ {
			if digits[i] != 10-n+i {
				last = false
				break
			}
		}

		if last {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')

		i := n - 1
		for i >= 0 && digits[i] == 10-n+i {
			i--
		}

		digits[i]++
		for j := i + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}
	}

	z01.PrintRune('\n')
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}