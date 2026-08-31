package main

/*
import (
	"fmt"
)

// Student struct holds information about an individual student
type Student struct {
	ID         int
	Name       string
	Department string
	GPA        float64
}

// StudentStore manages a slice of students without pointers
type StudentStore struct {
	students []Student
}

// NewStudentStore initializes a StudentStore with initial mock data
func NewStudentStore() StudentStore {
	return StudentStore{
		students: []Student{
			{ID: 1, Name: "John Doe", Department: "Computer Science", GPA: 1.5},
			{ID: 2, Name: "Jane Doe", Department: "Mathematics", GPA: 3.9},
			{ID: 3, Name: "Bob Smith", Department: "Physics", GPA: 2.2},
			{ID: 4, Name: "Charlie Brown", Department: "Chemistry", GPA: 1.8},
			{ID: 5, Name: "Emily Davis", Department: "Biology", GPA: 3.6},
		},
	}
}

// Display prints a student's info in a formatted line
func (s Student) Display() {
	fmt.Printf("[%d] %-16s | Dept: %-18s | GPA: %.2f\n", s.ID, s.Name, s.Department, s.GPA)
}

// reindexIDs updates the ID of each student sequentially from 1 to N
func (ss StudentStore) reindexIDs() StudentStore {
	for i := range ss.students {
		ss.students[i].ID = i + 1
	}
	return ss
}

// DisplayAll prints all students in the system
func (ss StudentStore) DisplayAll() {
	fmt.Println("\n--- All Students ---")
	if len(ss.students) == 0 {
		fmt.Println("No students found.")
		return
	}
	for _, s := range ss.students {
		s.Display()
	}
}

// GetStudentByID searches for a student by ID
func (ss StudentStore) GetStudentByID(id int) (Student, bool) {
	for _, s := range ss.students {
		if s.ID == id {
			return s, true
		}
	}
	return Student{}, false
}

// AddStudent adds a new student and returns the updated store
func (ss StudentStore) AddStudent(name, dept string, gpa float64) StudentStore {
	newID := len(ss.students) + 1
	newStudent := Student{
		ID:         newID,
		Name:       name,
		Department: dept,
		GPA:        gpa,
	}
	ss.students = append(ss.students, newStudent)
	fmt.Printf("\nStudent successfully added with ID: %d\n", newID)
	return ss
}

// DeleteStudent removes a student by ID and re-indexes remaining IDs
func (ss StudentStore) DeleteStudent(id int) (StudentStore, bool) {
	for i, s := range ss.students {
		if s.ID == id {
			ss.students = append(ss.students[:i], ss.students[i+1:]...)
			ss = ss.reindexIDs()
			return ss, true
		}
	}
	return ss, false
}

// DisplayPassedStudents shows students with GPA >= 2.00
func (ss StudentStore) DisplayPassedStudents(passThreshold float64) {
	fmt.Printf("\n--- Passed Students (GPA >= %.2f) ---\n", passThreshold)
	count := 0
	for _, s := range ss.students {
		if s.GPA >= passThreshold {
			s.Display()
			count++
		}
	}
	if count == 0 {
		fmt.Println("No passed students found.")
	}
}

// DisplayFailedStudents shows students with GPA < 2.00
func (ss StudentStore) DisplayFailedStudents(passThreshold float64) {
	fmt.Printf("\n--- Failed Students (GPA < %.2f) ---\n", passThreshold)
	count := 0
	for _, s := range ss.students {
		if s.GPA < passThreshold {
			s.Display()
			count++
		}
	}
	if count == 0 {
		fmt.Println("No failed students found.")
	}
}

// DisplayHighestGPAStudent finds and prints the student with the highest GPA
func (ss StudentStore) DisplayHighestGPAStudent() {
	if len(ss.students) == 0 {
		fmt.Println("\nNo students available.")
		return
	}

	highest := ss.students[0]
	for _, s := range ss.students {
		if s.GPA > highest.GPA {
			highest = s
		}
	}

	fmt.Println("\n--- Student with Highest GPA ---")
	highest.Display()
}

// printMenu displays the system menu
func printMenu() {
	fmt.Println("\n==================================")
	fmt.Println("   STUDENT MANAGEMENT SYSTEM")
	fmt.Println("==================================")
	fmt.Println("1. Show All Students")
	fmt.Println("2. Search Student by ID")
	fmt.Println("3. Add Student")
	fmt.Println("4. Delete Student")
	fmt.Println("5. Show Passed Students")
	fmt.Println("6. Show Failed Students")
	fmt.Println("7. Show Highest GPA Student")
	fmt.Println("8. Exit")
	fmt.Println("==================================")
	fmt.Print("Enter your choice (1-8): ")
}

func main() {
	store := NewStudentStore()
	const passGPA = 2.00

	for {
		printMenu()

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			store.DisplayAll()

		case 2:
			var id int
			fmt.Print("Enter Student ID to search: ")
			fmt.Scan(&id)
			if student, found := store.GetStudentByID(id); found {
				fmt.Println("\nStudent Found:")
				student.Display()
			} else {
				fmt.Printf("\nStudent with ID %d not found.\n", id)
			}

		case 3:
			var name, dept string
			var gpa float64
			fmt.Print("Enter Student Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Department: ")
			fmt.Scan(&dept)
			fmt.Print("Enter GPA: ")
			fmt.Scan(&gpa)
			store = store.AddStudent(name, dept, gpa)

		case 4:
			var id int
			fmt.Print("Enter Student ID to delete: ")
			fmt.Scan(&id)
			var success bool
			store, success = store.DeleteStudent(id)
			if success {
				fmt.Printf("\nStudent with ID %d successfully deleted.\n", id)
			} else {
				fmt.Printf("\nStudent with ID %d not found.\n", id)
			}

		case 5:
			store.DisplayPassedStudents(passGPA)

		case 6:
			store.DisplayFailedStudents(passGPA)

		case 7:
			store.DisplayHighestGPAStudent()

		case 8:
			fmt.Println("\nThank you for using Student Management System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice! Please select an option between 1 and 8.")
		}
	}
}

*/