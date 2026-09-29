package main

/*
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Employee struct {
	Name       string
	BaseSalary float64
	Bonus      float64
	TaxRate    float64
	Position   string
	Active     bool
}

// Function to clean, convert, validate and create Employee struct
func NewEmployee(rawName, rawSalary, rawBonus, rawTax, rawPos, rawActive string) (Employee, error) {
	// 1. Clean strings using strings package
	name := strings.TrimSpace(rawName)
	if name == "" {
		return Employee{}, errors.New("employee name cannot be empty")
	}

	position := strings.ToUpper(strings.TrimSpace(rawPos))
	if position == "" {
		return Employee{}, errors.New("position cannot be empty")
	}

	// 2. Parse numbers using strconv package & validate
	salary, err := strconv.ParseFloat(strings.TrimSpace(rawSalary), 64)
	if err != nil {
		return Employee{}, fmt.Errorf("invalid base salary input: %w", err)
	}
	if salary <= 0 {
		return Employee{}, errors.New("base salary must be greater than 0")
	}

	bonus, err := strconv.ParseFloat(strings.TrimSpace(rawBonus), 64)
	if err != nil {
		return Employee{}, fmt.Errorf("invalid bonus input: %w", err)
	}
	if bonus < 0 {
		return Employee{}, errors.New("bonus cannot be negative")
	}

	taxRate, err := strconv.ParseFloat(strings.TrimSpace(rawTax), 64)
	if err != nil {
		return Employee{}, fmt.Errorf("invalid tax rate input: %w", err)
	}
	if taxRate < 0 || taxRate > 100 {
		return Employee{}, errors.New("tax rate must be between 0 and 100")
	}

	active, err := strconv.ParseBool(strings.TrimSpace(rawActive))
	if err != nil {
		return Employee{}, fmt.Errorf("invalid active status: %w", err)
	}

	return Employee{
		Name:       name,
		BaseSalary: salary,
		Bonus:      bonus,
		TaxRate:    taxRate,
		Position:   position,
		Active:     active,
	}, nil
}

// Method to calculate Gross, Tax, and Net Salary
func (e Employee) CalculateNetSalary() (grossSalary float64, taxAmount float64, netSalary float64) {
	grossSalary = e.BaseSalary + e.Bonus
	taxAmount = grossSalary * (e.TaxRate / 100.0)
	netSalary = grossSalary - taxAmount
	return grossSalary, taxAmount, netSalary
}

// Function to print Employee Details nicely
func PrintEmployeeSummary(emp Employee) {
	gross, tax, net := emp.CalculateNetSalary()

	fmt.Println("\n========================================")
	fmt.Println("        EMPLOYEE SALARY SLIP            ")
	fmt.Println("========================================")
	fmt.Printf("Name         : %s\n", emp.Name)
	fmt.Printf("Position     : %s\n", emp.Position)
	fmt.Printf("Active       : %t\n", emp.Active)
	fmt.Println("----------------------------------------")
	fmt.Printf("Base Salary  : $%.2f\n", emp.BaseSalary)
	fmt.Printf("Bonus        : $%.2f\n", emp.Bonus)
	fmt.Printf("Gross Salary : $%.2f\n", gross)
	fmt.Printf("Tax (%-4.1f%%)  : -$%.2f\n", emp.TaxRate, tax)
	fmt.Println("----------------------------------------")
	fmt.Printf("Net Salary   : $%.2f\n", net)
	fmt.Println("========================================")
}

func main() {
	fmt.Println("====== Employee Salary Calculator ======")

	// Raw input data
	rawName := "   Sujon Haque   "
	rawSalary := " 65000.50 "
	rawBonus := " 5000.00 "
	rawTax := " 12.5 "
	rawPos := "   senior go developer   "
	rawActive := " true "

	fmt.Println("\n--- Raw Inputs ---")
	fmt.Printf("Name    : %q\n", rawName)
	fmt.Printf("Salary  : %q\n", rawSalary)
	fmt.Printf("Bonus   : %q\n", rawBonus)
	fmt.Printf("Tax Rate: %q\n", rawTax)
	fmt.Printf("Position: %q\n", rawPos)
	fmt.Printf("Active  : %q\n", rawActive)

	// Validate & Create Employee using NewEmployee function
	emp, err := NewEmployee(rawName, rawSalary, rawBonus, rawTax, rawPos, rawActive)
	if err != nil {
		fmt.Println("\nValidation Error:", err)
		return
	}

	// Print Employee Salary Summary
	PrintEmployeeSummary(emp)
}

*/
