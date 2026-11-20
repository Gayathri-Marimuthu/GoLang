// An array is a fixed-size collection of elements of the same data type.
// Size is fixed at compile time.
// All elements must have the same type.
// Index starts from 0.

// syntax : var arrayname[size]datatype
// declaration: var arr[5]int
// initialize: arr := [5]int{1,2,3,4,5}

// find the maximum nuber of array

package main

import "fmt"

func main() {
	numbers := [5]int{30, 50, 90, 70, 80}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Maximum number is", max)

}


numbers = [30, 50, 90, 70, 80]

Start
  │
  ▼
max = numbers[0]
max = 30
  │
  ▼
┌─────────────────────────────┐
│ Loop through each number    │
└─────────────────────────────┘
  │
  ├── num = 30
  │      30 > 30 ? ❌
  │      max = 30
  │
  ├── num = 50
  │      50 > 30 ? ✅
  │      max = 50
  │
  ├── num = 90
  │      90 > 50 ? ✅
  │      max = 90
  │
  ├── num = 70
  │      70 > 90 ? ❌
  │      max = 90
  │
  └── num = 80
         80 > 90 ? ❌
         max = 90

  ▼
Print: Maximum number is 90
