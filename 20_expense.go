package main

import "fmt"

type Expense struct {
	expense string
	amount  float64
	date    string
}

// Declare the Expense struct here
// your code goes here

func Total(s []Expense) float64 {
	total := 0.0
	for idx, _ := range s {
		total += s[idx].amount
	}
	return total
}

// Implement the Total method to calculate the total amount spent
// your code goes here

// Implement the getName method on the Expense struct here
// your code goes here

func (s Expense) getName() string {
	return s.expense
}
func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}
