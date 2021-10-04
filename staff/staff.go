package staff

import "fmt"

var OverpaidLimit = 75000
var underpaidLimit = 25000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, employee := range e.AllStaff {
		if employee.Salary > OverpaidLimit {
			overpaid = append(overpaid, employee)
		}
	}

	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee

	myFunction()

	for _, employee := range e.AllStaff {
		if employee.Salary < underpaidLimit {
			underpaid = append(underpaid, employee)
		}
	}

	return underpaid
}

func (e *Office) notVisible() {
	fmt.Println("Hello, world")
}

func myFunction() {
	fmt.Println("I am a function")
}