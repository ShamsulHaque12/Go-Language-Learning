package database

import "testing"

func TestAdd_Basic(t *testing.T) {
	got := Add(10, 20)
	want := 30

	if got != want {
		t.Errorf("Add(10, 20) = %d; want %d", got, want)
	}
}

func TestSubtract_Basic(t *testing.T) {
	got := Subtract(20, 8)
	want := 12

	if got != want {
		t.Errorf("Subtract(20, 8) = %d; want %d", got, want)
	}

	got = Subtract(8, 20)
	want = -12

	if got != want {
		t.Errorf("Subtract(8, 20) = %d; want %d", got, want)
	}
}

func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive numbers", 10, 20, 30},
		{"negative numbers", -5, -3, -8},
		{"zero", 0, 10, 10},
		{"mixed numbers", -5, 10, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}
func TestSubtract_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive numbers", 20, 8, 12},
		{"negative result", 8, 20, -12},
		{"subtract zero", 10, 0, 10},
		{"negative numbers", -5, -3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d, want %d",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}
