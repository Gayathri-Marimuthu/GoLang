// A map is a collection of key-value pairs.
// Every key must be unique.
// Values can be duplicated.
// Maps are unordered (the iteration order is not guaranteed).

// syntax var mapname map[keytype]valuetype
package main

import "fmt"

func main() {
	// students := make(map[int]string)
	student_name := map[int]string{
		101: "gayu",
		102: "hari",
	}
	fmt.Println(student_name)

	// imp fun delete(),len(),make()
	// value, ok := m[key] key exist
}
