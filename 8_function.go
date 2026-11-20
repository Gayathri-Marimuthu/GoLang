//A function is a reusable block of code that performs a specific task
// syntax func function_name(parameters)returntype{
// 	return value
// }
//function without parameter
func welcome(){
}
//Multiple return value
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
//Named return value
func calculate(a, b int) (sum int, product int) {}

//Variadic Function
//Accept any number of aruguments
func sum(numbers ...int) int {}

//Function assigned to variable
add:=func add(a,b int)int{}

//Function as parameter
func calculate(a,b int,operation func (int,int)int){}

//Function returning function
//passing slice to function
//passing Map
//Passing Struct
//pointer parameter

