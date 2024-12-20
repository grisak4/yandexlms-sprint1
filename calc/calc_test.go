package calc_test

import (
	"sprintone/calc"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       float64
		wantErr    bool
	}{
		{"простое сложение", "2+2", 4, false},
		{"простое вычитание", "5-3", 2, false},
		{"простое умножение", "4*3", 12, false},
		{"простое деление", "8/4", 2, false},
		{"деление на ноль", "8/0", 0, true},
		{"приоритет операторов", "2+3*4", 14, false},
		{"приоритет скобок", "(2+3)*4", 20, false},
		{"вложенные скобки", "((2+3)*4)-6/(1+2)", 18, false},
		{"числа с плавающей точкой", "3.5+2.1", 5.6, false},
		{"недопустимый символ", "3+5a", 0, true},
		{"несогласованные скобки", "(2+3*4", 0, true},
		{"пустое выражение", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Calculator(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
