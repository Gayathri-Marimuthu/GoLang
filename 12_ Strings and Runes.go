// A string in Go is an immutable sequence of UTF-8 encoded bytes. The len() function returns the number of bytes, which works as expected for ASCII text because each character is one byte. However, Unicode characters like Chinese letters or emojis use multiple bytes, so len() may not equal the number of characters. To work with actual characters, Go uses the rune type, which represents a single Unicode code point.

//len([]rune(word)

// Counting Characters Correctly
package main

import "fmt"

func main() {
	word := "你好"

	fmt.Println(len(word))
	fmt.Println(len([]rune(word)))
}

//looping
// name := "Hello"

// for i := 0; i < len(name); i++ {
// 	fmt.Printf("%c ", name[i])
// }

// package main

// import "fmt"

// func main() {
// 	word := "你好"

// 	for i, ch := range word {
// 		fmt.Println(i, string(ch))
// 	}
// }
