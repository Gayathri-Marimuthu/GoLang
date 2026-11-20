<!-- fmt.Println() -->

Prints values with a newline at the end.

<!-- fmt.Print()  -->

- Prints values without adding a newline.

fmt.Printf()

<!-- - Prints formatted output. -->

eg:fmt.Println(name, age)
fmt.Printf("Name: %s Age: %d", name, age)

<!-- fmt.Sprintf() -->

- It does NOT print anything.Instead, it returns a formatted string.
<!-- import "fmt"

func main() {

    name := "Gayathri"
    age := 32

    result := fmt.Sprintf("Name=%s Age=%d", name, age)

    fmt.Println(result)

}output:Name=Gayathri Age=32 -->

<!-- fmt.Errorf()  -->

- Creates an error with formatting.

<!-- fmt.Scan() -->

-Reads input from keyboard. fmt.Scanln(),fmt.Scanf()

%d-print integer
%s-print string
%f-float
%.2f-two decimal number
%v-default format print the value
%+v-Prints struct with field names. fmt.Printf("%+v", e)
%T-Prints the type of a variable.
%p-Memory address (pointer).
%t-print boolean
%c-Character ASCII
%b-binary
%x-hexadecimal
%o-octal

<!-- What is the difference between Printf() and Sprintf()? -->

Printf() formats and prints to standard output.
Sprintf() formats and returns the string without printing it.
