// Exported and Unexported
package main

import (
	"fmt"
	"go/gobc01/sec4/exported/staff"
)

var employees = []staff.Employee{
	{
		FirstName: "Chris",
		LastName:  "Ng",
		Title:     "Programmer",
		Salary:    120000,
		IsEmloyed: true,
		StartDate: staff.HireDate{Y: 2020, M: 5, D: 2},
	},
	{
		FirstName: "Tea",
		LastName:  "Tran",
		Title:     "Programmer",
		Salary:    135000,
		IsEmloyed: true,
		StartDate: staff.HireDate{Y: 2020, M: 11, D: 10},
	},
	{
		FirstName: "Mike",
		LastName:  "McNillin",
		Title:     "Tennis Coach",
		Salary:    100000,
		IsEmloyed: true,
		StartDate: staff.HireDate{Y: 2018, M: 1, D: 19},
	},
	{
		FirstName: "Tim",
		LastName:  "McKen",
		Title:     "Roofer",
		Salary:    92000,
		IsEmloyed: true,
		StartDate: staff.HireDate{Y: 2019, M: 3, D: 28},
	},
	{
		FirstName: "Hans",
		LastName:  "Kieffer",
		Title:     "Sale Man",
		Salary:    130000,
		IsEmloyed: true,
		StartDate: staff.HireDate{Y: 2021, M: 8, D: 8},
	},
}

func main() {

	// Declare office staff
	officeStaff := staff.Office{
		AllStaff: employees,
	}

	// Print all office staff
	allEm := officeStaff.GetAllEmployee()
	for i, e := range allEm {
		fmt.Printf("#%d : %v\n", i, e)
	}

	// Search a specific employee name
	em := officeStaff.SearchEmployee("Hans")
	fmt.Printf("we found %s in the system.\n", em.FirstName)

	// Print Employee with highest salary
	hisaEmployee := officeStaff.HighSalary()
	fmt.Println("The highest salary employee is", hisaEmployee)
}
