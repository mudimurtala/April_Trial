# IsNegative: Checking if a Number is Negative

## What is This Exercise About?

In this exercise, you'll write your **first function** — a reusable block of code that performs a specific task. Your function will check whether a number is negative or not.

**The Mission:**
Write a function called `IsNegative` that:
- Takes an integer as input
- Prints the character `'T'` (meaning "true") if the number is **negative** (less than 0)
- Prints the character `'F'` (meaning "false") if the number is **not negative** (zero or positive)

## Key Concepts You'll Learn

### 1. What are Functions?

A **function** is a reusable block of code that performs a specific task. Instead of writing the same code repeatedly, you define it once as a function and call it whenever you need it.

**Why use functions?**
- Code reusability: Write once, use many times
- Organization: Keep your code organized and readable
- Modularity: Break complex problems into smaller, manageable pieces

**Syntax:**
```go
func FunctionName(parameter1 type1, parameter2 type2) {
	// Code to execute
}
```

### 2. Understanding Negative Numbers

**Negative numbers** are values less than zero. In mathematics and programming:
- `-5` is negative (less than 0)
- `-1` is negative (less than 0)
- `0` is **not** negative (it equals zero, not less than zero)
- `1` is **not** negative (greater than 0)

### 3. Conditional Statements: `if` and `else`

To check whether a number is negative, you use an **`if` statement** — the most basic decision-making tool in programming.

**Syntax:**
```go
if condition {
	// Code runs if condition is TRUE
} else {
	// Code runs if condition is FALSE
}
```

**Example:**
```go
if age >= 18 {
	fmt.Println("You are an adult")
} else {
	fmt.Println("You are not an adult")
}
```

### 4. Comparison Operators

Comparison operators let you compare two values and get a `true` or `false` result:

| Operator | Meaning | Example |
|----------|---------|---------|
| `<` | Less than | `3 < 5` is `true` |
| `>` | Greater than | `10 > 5` is `true` |
| `<=` | Less than or equal to | `5 <= 5` is `true` |
| `>=` | Greater than or equal to | `5 >= 3` is `true` |
| `==` | Equal to | `5 == 5` is `true` |
| `!=` | Not equal to | `5 != 3` is `true` |

**For this exercise:**
- We use `< 0` to check if a number is negative

## Step-by-Step Solution

### Step 1: Understand the Function Signature

```go
func IsNegative(nb int) {
	// nb is the parameter (the number we're checking)
	// The function takes one integer parameter
}
```

### Step 2: Write the Condition

Inside the function, use an `if` statement to check if the number is negative:

```go
if nb < 0 {
	// nb is negative, print 'T'
} else {
	// nb is not negative, print 'F'
}
```

### Step 3: Print the Result

Use `z01.PrintRune()` to print either `'T'` or `'F'`:

```go
if nb < 0 {
	z01.PrintRune('T')
} else {
	z01.PrintRune('F')
}
```

## Complete Function Code

```go
package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}

func main() {
	IsNegative(1)   // Output: F (1 is not negative)
	IsNegative(0)   // Output: F (0 is not negative)
	IsNegative(-1)  // Output: T (-1 is negative)
}
```

## How It Works: Line by Line

1. **`if nb < 0`** — Check if the parameter `nb` is less than 0
2. **`z01.PrintRune('T')`** — If true, print the character 'T'
3. **`z01.PrintRune('F')`** — Otherwise, print the character 'F'

## Testing Your Function

When you run this program:
```bash
go run .
```

You should see:
```
F
F
T
```

**Why this output?**
- `IsNegative(1)` → 1 is not negative → prints `F`
- `IsNegative(0)` → 0 is not negative → prints `F`
- `IsNegative(-1)` → -1 is negative → prints `T`

## Common Mistakes to Avoid

1. **Forgetting the comparison operator:** `if nb` won't work. You need `if nb < 0`
2. **Using `≤` instead of `<`:** Remember, 0 is NOT negative, so use `<`, not `<=`
3. **Printing strings instead of runes:** Use `z01.PrintRune('T')`, not `fmt.Println("true")`
4. **Forgetting to call the function in `main()`:** The function won't run unless you call it

## Quick Recap

- **Functions** are reusable code blocks that perform specific tasks
- **Conditionals** (`if`/`else`) let your program make decisions
- **Comparison operators** (`<`, `>`, etc.) compare values
- **Negative numbers** are any integers less than 0
- Use `z01.PrintRune()` to print individual characters