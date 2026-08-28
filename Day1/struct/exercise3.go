package main

/*
import "fmt"

type Student struct {
	Name       string
	Age        int
	Email      string
	Department string
	ID         int
	GPA        float64
}

func createStudent(name string, age int, email string, department string, id int, gpa float64) Student {
	return Student{
		Name:       name,
		Age:        age,
		Email:      email,
		Department: department,
		ID:         id,
		GPA:        gpa,
	}
}

func (s Student) introduce() {
	fmt.Println("========== Student Information ==========")
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Email:", s.Email)
	fmt.Println("Department:", s.Department)
	fmt.Println("ID:", s.ID)
	fmt.Println("GPA:", s.GPA)
}

func (s Student) getPassed() bool {
	return s.GPA >= 2.0
}

func main() {
	student1 := createStudent("Shamsul", 23, "xyz@gmail.com", "IT", 123, 3.5)
	student2 := createStudent("Rimon", 22, "xyz@gmail.com", "CSE", 456, 2.0)
	student3 := createStudent("Rofiq", 24, "xyz@gmail.com", "EEE", 789, 3.8)
	student4 := createStudent("Polash", 25, "xyz@gmail.com", "BBA", 101, 1.2)
	student5 := createStudent("Borhan", 26, "xyz@gmail.com", "ACA", 102, 3.4)

	students := []Student{
		student1,
		student2,
		student3,
		student4,
		student5,
	}

	for _, student := range students {
		student.introduce()
		if student.getPassed() {
			fmt.Println("Status: Passed")
		} else {
			fmt.Println("Status: Failed")
		}
		fmt.Println()
	}
}
*/