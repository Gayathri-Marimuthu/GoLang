//Generics in Go allow you to write functions, methods, and data structures that work with multiple types without duplicating code.

//syntax : func FunctionName[T constraint](parameter T) T {
// code
//}
// T is simply a placeholder for a type. t can be any type

//Generic Function Example
// func Add[T int | float64](a, b T) T {
// 	return a + b
// }

// Why use Generics? → To avoid writing the same code for different types.
// Main benefit → Reusable, type-safe, and cleaner code

Note:
Suppose you want to find the maximum value.
For integers:
func maxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}
For float:
func maxFloat(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}

The logic is the same, but we have duplicate code.
func max[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}

T can be either int or float64.
T is called a type parameter.


## any means any type.

func Print[T any](value T) {
    fmt.Println(value)
}

        Generic Function
              │
       ┌──────┼──────┐
       ↓      ↓      ↓
      int   string  float
