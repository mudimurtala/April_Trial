# Go Basics Training Guide for Complete Beginners

## Purpose of this guide

This material is designed for absolute beginners who feel lost in Go training.
It is written so you can teach in small sessions and continue over time.

This guide focuses on:
- Core Go basics in the right order
- Clear examples
- Hands-on exercises after every concept
- Practical style used in your current training where printing is often done with z01.PrintRune

## How to use this guide in a live session

Recommended session structure:
1. Explain one concept in simple words
2. Run a tiny example
3. Let learners solve a short exercise
4. Review common mistakes
5. Move to next concept

Recommended pace:
- Session 1: What is Go, program structure, data types
- Session 2: Variables and constants
- Session 3: Functions
- Session 4: Conditionals
- Session 5: Loops
- Session 6: Rune and PrintRune deep practice
- Session 7: Mixed exercises and checkpoints

------------------------------------------------------------

## 1. What is Go programming language

Go (Golang) is a programming language created to be:
- Simple to read
- Fast to run
- Good for real software projects

Why beginners like Go:
- Clean syntax
- Strong typing helps catch mistakes early
- One clear loop style (for loop)
- Good standard library

Your learners should remember this first:
- A Go program is made of packages, functions, and statements.
- Every runnable program starts at main function.

Example:

```go
package main

import "github.com/01-edu/z01"

func main() {
    z01.PrintRune('H')
    z01.PrintRune('i')
    z01.PrintRune('\n')
}
```

Exercise 1:
- Print the word Go and then a newline using only z01.PrintRune.

------------------------------------------------------------

## 2. Program structure in Go

Most beginner programs have:
- package declaration
- imports
- functions

Template:

```go
package main

import "github.com/01-edu/z01"

func main() {
    // your code here
}
```

Key rule:
- main function is the entry point for runnable programs.

Common mistake:
- Writing main as Main (capital M). Go is case-sensitive.

Exercise 2:
- Create a file that prints A, B, C on the same line, then newline.

------------------------------------------------------------

## 3. Data types in Go (very important)

A data type tells Go what kind of value is stored.

### 3.1 Integer types

Used for whole numbers.
Examples:
- int
- int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64

Most beginner tasks use int.

```go
var age int = 25
var score = 90 // inferred as int
```

### 3.2 Floating-point types

Used for decimal numbers.
Examples:
- float32
- float64

```go
var price float64 = 99.95
```

### 3.3 Boolean type

Used for true or false.

```go
var isReady bool = true
```

### 3.4 String type

A sequence of characters.

```go
var name string = "Musa"
```

### 3.5 Byte and Rune (critical for your training)

This is where many beginners get confused.

- byte represents raw byte data (alias of uint8)
- rune represents a Unicode code point (alias of int32)

Why rune is special:
- Many languages hide this detail from beginners.
- In Go, rune is explicit and very useful when working with characters.
- z01.PrintRune expects a rune, not a string.

Examples:

```go
var letter rune = 'A'
var digit rune = '7'
var symbol rune = '#'
```

Important:
- Single quotes produce a rune literal: 'A'
- Double quotes produce a string: "A"

If you use z01.PrintRune, this is correct:

```go
z01.PrintRune('A')
```

This is wrong:

```go
// z01.PrintRune("A") // wrong type
```

### 3.6 Rune vs string iteration

If you iterate over a string by index, you get bytes.
If you iterate with range, you get runes.

```go
s := "Go"

for i := 0; i < len(s); i++ {
    // s[i] is a byte
}

for _, r := range s {
    // r is a rune
    _ = r
}
```

Exercise 3:
- Create rune variables for X, y, and 5 and print each with z01.PrintRune.

Exercise 4:
- Explain in your own words: difference between 'A' and "A".

------------------------------------------------------------

## 4. Variables and constants

### 4.1 Declaring variables

You can declare with type:

```go
var count int = 10
```

Or let Go infer type:

```go
var count = 10
```

Short declaration (inside functions only):

```go
count := 10
```

### 4.2 Updating variables

```go
count := 10
count = count + 5
count++
```

### 4.3 Constants

Constant values do not change.

```go
const school = "01"
const maxStudents = 30
```

Exercise 5:
- Create variables name, age, and isStudent.
- Change age by adding 1.
- Print a single letter based on isStudent (S for true, N for false) using PrintRune.

------------------------------------------------------------

## 5. Operators you need now

Arithmetic:
- +, -, *, /, %

Comparison:
- ==, !=, <, >, <=, >=

Logical:
- &&, ||, !

Example:

```go
a := 10
b := 3

sum := a + b      // 13
mod := a % b      // 1
ok := a > b       // true
both := ok && b > 0
_ = sum
_ = mod
_ = both
```

Exercise 6:
- Check if a number is even using n % 2 == 0.
- Print E for even, O for odd.

------------------------------------------------------------

## 6. Functions

Functions let you group reusable logic.

### 6.1 Basic function

```go
func SayA() {
    z01.PrintRune('A')
    z01.PrintRune('\n')
}
```

### 6.2 Function with parameter

```go
func PrintLetter(r rune) {
    z01.PrintRune(r)
    z01.PrintRune('\n')
}
```

### 6.3 Function returning value

```go
func Add(a int, b int) int {
    return a + b
}
```

### 6.4 Multiple return values

```go
func DivideAndRemainder(a int, b int) (int, int) {
    return a / b, a % b
}
```

### 6.5 Why this matters for your training

Your tasks like printalphabet, isnegative, printnbr all depend on function thinking.

Exercise 7:
- Write IsPositive(n int) that prints T if n > 0, otherwise F.

Exercise 8:
- Write a function Double(n int) int and test it with 4, 9, 0.

------------------------------------------------------------

## 7. Conditionals (if, else if, else)

Conditionals control decision making.

```go
if n < 0 {
    z01.PrintRune('N')
} else if n == 0 {
    z01.PrintRune('Z')
} else {
    z01.PrintRune('P')
}
```

Common beginner mistake:
- Using assignment = instead of comparison == in conditions.

### 7.1 switch

Useful when checking many exact values.

```go
switch day {
case 1:
    // Monday
case 2:
    // Tuesday
default:
    // unknown
}
```

Exercise 9:
- Write a function Grade(mark int) that returns:
  - A if mark >= 70
  - B if mark >= 60
  - C if mark >= 50
  - F otherwise

Exercise 10:
- Write IsNegative style function and test with -1, 0, 2.

------------------------------------------------------------

## 8. Loops (the engine of many Go tasks)

Go has one loop keyword: for.

### 8.1 Basic counting loop

```go
for i := 0; i < 5; i++ {
    // repeat
}
```

### 8.2 While-style loop

```go
i := 0
for i < 5 {
    i++
}
```

### 8.3 Infinite loop

```go
for {
    break
}
```

### 8.4 range loop

```go
for _, r := range "Go" {
    _ = r
}
```

### 8.5 Loop patterns from your current exercises

Pattern A: print alphabet

```go
for r := 'a'; r <= 'z'; r++ {
    z01.PrintRune(r)
}
z01.PrintRune('\n')
```

Pattern B: reverse alphabet

```go
for r := 'z'; r >= 'a'; r-- {
    z01.PrintRune(r)
}
z01.PrintRune('\n')
```

Pattern C: digits

```go
for r := '0'; r <= '9'; r++ {
    z01.PrintRune(r)
}
z01.PrintRune('\n')
```

Exercise 11:
- Print letters from a to f.

Exercise 12:
- Print odd digits only: 1, 3, 5, 7, 9 (no commas, one line).

------------------------------------------------------------

## 9. Deep focus: Rune and z01.PrintRune

Since your learners are not using normal print functions, they must be excellent at PrintRune.

### 9.1 Printing a string manually with PrintRune

```go
func PrintWord(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}
```

### 9.2 Printing an int manually (idea behind printnbr)

Core idea:
- Handle sign
- Print higher digits first (often recursion)
- Convert each digit to rune using + '0'

```go
func PrintNbr(n int) {
    if n < 0 {
        z01.PrintRune('-')
    }

    if n <= -10 || n >= 10 {
        if n < 0 {
            PrintNbr(-(n / 10))
        } else {
            PrintNbr(n / 10)
        }
    }

    d := n % 10
    if d < 0 {
        d = -d
    }
    z01.PrintRune(rune(d) + '0')
}
```

Exercise 13:
- Write PrintYes that prints YES using only PrintRune.

Exercise 14:
- Write PrintLine(s string) using range and PrintRune.

------------------------------------------------------------

## 10. Guided practice from your local tasks

Use these local files as practice references.

### 10.1 Very first print practice
- checkpoints/onlya.go
- Task: change it to print z instead.

### 10.2 Alphabet loop practice
- Quest_02/printalphabet/solution.go
- Task: print only lowercase vowels.

### 10.3 Conditionals practice
- Quest_02/isnegative/solution.go
- Task: create IsZero that prints T if zero else F.

### 10.4 Combination logic practice
- Quest_02/printcomb/solution.go
- Quest_02/printcomb2/solution.go
- Task: explain why second loop starts at first+1.

### 10.5 Character processing practice
- checkpoints/repeatalpha.go
- checkpoints/countalpha.go
- Task: update countalpha to use range over runes.

------------------------------------------------------------

## 11. Frequent beginner errors and fixes

1. Confusing rune and string
- Fix: use single quotes for rune, double quotes for string.

2. Off-by-one loop errors
- Fix: read condition carefully: < versus <=.

3. Forgetting newline
- Fix: add z01.PrintRune('\n') when output format requires it.

4. Using = in condition
- Fix: use == for equality checks.

5. Wrong order in combinations
- Fix: ensure first < second (and third for 3-digit combinations).

6. Trailing comma issues
- Fix: print separator only if not last element.

7. Shadowing variables accidentally
- Fix: use = to update existing variable, not :=.

------------------------------------------------------------

## 12. Simple checkpoint ladder (for teaching)

Use this order in class:
1. Print one character
2. Print a fixed word
3. Print alphabet
4. Print reverse alphabet
5. Print digits
6. IsNegative
7. PrintNbr
8. PrintComb
9. PrintComb2
10. PrintCombN

At each step ask learners:
- What is the input?
- What is the output format?
- Which loop or condition is needed?
- Where can mistakes happen?

------------------------------------------------------------

## 13. Homework set (hands-on)

### Homework A (very basic)
1. Print your initials with PrintRune.
2. Print digits 0 to 5.
3. Print letters c to h.

### Homework B (conditionals)
1. Print T if number is even, else F.
2. Print P, N, or Z for positive/negative/zero.

### Homework C (functions)
1. Write Max(a, b int) int.
2. Write Min(a, b int) int.
3. Write Abs(n int) int.

### Homework D (mixed)
1. Count alphabetic characters in a string.
2. Repeat alphabetic characters by index (like repeatalpha idea).
3. Print all two-digit pairs where first is smaller than second.

------------------------------------------------------------

## 14. Suggested teaching script for your first meeting

Use this script flow:

Part 1 (15 min)
- What is Go and why we use it
- Structure of a Go file

Part 2 (20 min)
- Data types
- Focus on rune and why it matters

Part 3 (20 min)
- Variables and constants
- Tiny tasks

Part 4 (25 min)
- If/else and loops
- IsNegative style exercise

Part 5 (20 min)
- PrintRune-only output tasks
- Alphabet and digits drills

Part 6 (20 min)
- Q and A
- Assign homework A and B

------------------------------------------------------------

## 15. Quick reference sheet

- main entry point: func main()
- rune literal: 'A'
- string literal: "A"
- print one rune: z01.PrintRune(r)
- digit to rune: rune(d) + '0'
- newline: z01.PrintRune('\n')
- for loop: for i := 0; i < n; i++
- condition: if ... else
- equality check: ==

------------------------------------------------------------

## 16. Final encouragement for your learners

Go gets easier with repetition.
The students are not slow. They only need more guided practice.

Best rule:
- Read problem
- Predict output
- Write small code
- Test
- Fix one error at a time

If they keep practicing these basics, the harder topics later will become much easier.

End of guide.
