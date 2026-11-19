package calculator

import "errors"

// Exported Variable (বড় হাতের বর্ণ দিয়ে শুরু -> অন্য প্যাকেজ থেকে access করা যাবে)
var AppName = "My Custom Calculator"

// Unexported Variable (ছোট হাতের বর্ণ দিয়ে শুরু -> অন্য প্যাকেজ থেকে access করা যাবে না)
var version = "v1.0.0"

// Exported Function (Add) -> Public
func Add(a, b int) int {
	return a + b
}

// Exported Function (Sub) -> Public
func Sub(a, b int) int {
	return a - b
}

// Exported Function (Mul) -> Public
func Mul(a, b int) int {
	return a * b
}

// Exported Function (Div) -> Public
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Unexported Function (internalHelper) -> Private (শুধুমাত্র calculator প্যাকেজের ভেতর ব্যবহারযোগ্য)
func internalHelper(msg string) string {
	return "[LOG]: " + msg
}