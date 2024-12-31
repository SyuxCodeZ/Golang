package main

import ("fmt")

func main () {
	var num1, num2 float64
	var operator string

	fmt.Println("~~~~WELCOME TO MY SIMPLE CALCULATOR~~~~")

	fmt.Println("Enter The 1st Number: ")
	fmt.Scanln(&num1)

	fmt.Println("Please Enter An Operator (+, -, * and/): ")
	fmt.Scanln(&operator)

	if operator = "" {
	fmt.Println("Please Enter An Operator")
	return
	}

	fmt.Println("Enter The 2nd Number: ")
	fmt.Scanln(&num2)

	Switch operator {
	case = "+":
		fmt.Printf("The Result Is %.2f\n", num1 + num2)

	case = "-":
		fmt.Printf("The Result Is %.2f\n", num1 - num2)

	case = "*":
		fmt.Printf("The Result Is %.2f\n", num1 * num2)

	case = "/":
		if num2 == 0 {
			fmt.Println("Number Cannot Be Divided By 0")
			return
		}
		else {
			fmt.Println("The Result Is %.2f\n", num1 / num2)
		return
		}

	default:
		fmt.Println("Please Enter A Valid Operator")
	}
}