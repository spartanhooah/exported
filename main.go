package main

import (
	"fmt"
	"myapp/staff"
)

var employees = []staff.Employee {
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: true},
}

func main() {
	myStaff := staff.Office {
		AllStaff: employees,
	}

	// fmt.Println(myStaff.All())

	// staff.OverPaidLimit = 10000
	fmt.Println("Overpaid staff:", myStaff.Overpaid())
	fmt.Println("Underpaid staff:", myStaff.Underpaid())
}