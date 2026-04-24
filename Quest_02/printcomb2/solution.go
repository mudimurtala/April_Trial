package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for first := 0; first <= 98; first++ {
		for second := first + 1; second <= 99; second++ {
			z01.PrintRune(rune(first/10) + '0')
			z01.PrintRune(rune(first%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(second/10) + '0')
			z01.PrintRune(rune(second%10) + '0')

			if !(first == 98 && second == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

func main() {
	PrintComb2()
}
