package staff

type HireDate struct {
	Y int
	M int
	D int
}

type Employee struct {
	FirstName string
	LastName  string
	Title     string
	Salary    int
	IsEmloyed bool
	StartDate HireDate
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) GetAllEmployee() []Employee {
	return e.AllStaff
}

func (en *Office) SearchEmployee(n string) Employee {
	var employee Employee
	for _, v := range en.AllStaff {
		if v.FirstName == n {
			employee = v
		}
	}
	return employee
}

func (en *Office) HighSalary() Employee {
	var employee Employee
	var largerNum, temp int

	// Find the largest salary number
	for _, em := range en.AllStaff {
		if em.Salary > temp {
			temp = em.Salary
			largerNum = temp
		}

		if em.Salary == largerNum {
			employee = em
		}
	}
	return employee
}
