// ________________________________________________________________
// Challenge 3: Employee Data Management
// You are tasked with managing a list of employees with the following details:
// ID, Name, Age, and Salary. Implement a Manager struct that provides the following functionalities:
// AddEmployee: Add a new employee to the list.
// RemoveEmployee: Remove an employee based on their ID.
// GetAverageSalary: Calculate the average salary of all employees.
// FindEmployeeByID: Retrieve an employee's details by their ID.

package main

import (
	"fmt"
	"slices"
)

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	// TODO: Implement this method
	employee := Employee{
		ID:     e.ID,
		Name:   e.Name,
		Age:    e.Age,
		Salary: e.Salary,
	}
	m.Employees = append(m.Employees, employee)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
	// TODO: Implement this method
	for i, employee := range m.Employees {
		if employee.ID == id {
			//m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
			m.Employees = slices.Delete(m.Employees, i, i+1)
			return
		}
	}
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	// TODO: Implement this method
	var salarySumm, count int

	if len(m.Employees) == 0 {
		return 0
	}

	for _, employee := range m.Employees {
		salarySumm += int(employee.Salary)
		count++
	}
	return float64(salarySumm / count)
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	// TODO: Implement this method
	for _, employee := range m.Employees {
		if employee.ID == id {
			return &employee
		}
	}
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	//fmt.Println(manager)
	manager.RemoveEmployee(1)
	//fmt.Println("После удаления 1го", manager)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
