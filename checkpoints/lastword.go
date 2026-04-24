// ## lastword
// ### Instructions
// Write a function `LastWord` that takes a `string` and returns its last word 
// followed by a `\n`.
// - A word is a section of `string` delimited by spaces or by the start/end 
// of the `string`.
// ### Expected function
// ```go
// func LastWord(s string) string{
// }
// ```
// ### Usage
// Here is a possible program to test your function :
// ```go
// package main
// import (
// 	"fmt"
// 	"piscine"
// )
// func main() {
// 	fmt.Print(piscine.LastWord("this        ...       is sparta, then again, maybe    not"))
// 	fmt.Print(piscine.LastWord(" lorem,ipsum "))
// 	fmt.Print(piscine.LastWord(" "))
// }
// ```
// And its output :
// ```console
// $ go run . | cat -e
// not$
// lorem,ipsum$
// $
// $
// ```
// ### Notions
// - [01-edu/z01](https://github.com/01-edu/z01)



package main
import (
	"fmt"
)

func LastWord(s string) string {
	result := ""

	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		result = string(s[i]) + result
		i--
	}

	return result + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
