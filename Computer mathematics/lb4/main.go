package main

import (
	"fmt"
	"math"
)

// Решает задачу Коши методом Эйлера
// f - функция правой части уравнения
// y0 - начальное значение y
// x0 - начальное значение x
// b - конечное значение x
// h - шаг интегрирования
// E - точность вычисления
func EulerMethod(f func(x, y float64) float64, y0, x0, b, h, E float64) float64 {
	y := y0
	for x := x0; x < b; x += h {
		y += h * f(x, y)
	}
	return y
}

// Решает задачу Коши методом Эйлера-Коши
// f - функция правой части уравнения
// y0 - начальное значение y
// x0 - начальное значение x
// b - конечное значение x
// h - шаг интегрирования
// E - точность вычисления
func EulerCauchyMethod(f func(x, y float64) float64, y0, x0, b, h, E float64) float64 {
	y := y0
	for x := x0; x < b; x += h {
		y += h * f(x+h/2, y+h/2*f(x, y))
	}
	return y
}

// Решает задачу Коши модифицированным методом Эйлера
// f - функция правой части уравнения
// y0 - начальное значение y
// x0 - начальное значение x
// b - конечное значение x
// h - шаг интегрирования
// E - точность вычисления
func ModifiedEulerMethod(f func(x, y float64) float64, y0, x0, b, h, E float64) float64 {
	y := y0
	for x := x0; x < b; x += h {
		k1 := f(x, y)
		k2 := f(x+h, y+h*k1)
		y += h / 2 * (k1 + k2)
	}
	return y
}

// Решает задачу Коши методом Рунге-Кутты 4-го порядка
// f - функция правой части уравнения
// y0 - начальное значение y
// x0 - начальное значение x
// b - конечное значение x
// h - шаг интегрирования
// E - точность вычисления
func RungeKuttaMethodWithAccuracy(f func(x, y float64) float64, y0, x0, b, h, E float64) float64 {
	y := y0
	for x := x0; x < b; x += h {
		k1 := h * f(x, y)
		k2 := h * f(x+h/2, y+k1/2)
		k3 := h * f(x+h/2, y+k2/2)
		k4 := h * f(x+h, y+k3)
		y += (k1 + 2*k2 + 2*k3 + k4) / 6
	}
	y2 := y0
	for x := x0; x < b; x += h / 2 {
		k1 := h / 2 * f(x, y2)
		k2 := h / 2 * f(x+h/4, y2+k1/2)
		k3 := h / 2 * f(x+h/4, y2+k2/2)
		k4 := h / 2 * f(x+h/2, y2+k3)
		y2 += (k1 + 2*k2 + 2*k3 + k4) / 6
	}
	err := math.Abs(y-y2) / 15

	if err <= E {
		return y2
	}

	return RungeKuttaMethodWithAccuracy(f, y0, x0, b, h/2, E)
}

func main() {

	a := 3.1415926535 / 2
	b := (3.1415926535 / 2) + 1
	y0 := 1.0
	h := 0.1

	f := func(x, y float64) float64 {
		return (y / x) + x*math.Sin(x)
	}

	E := 0.0001

	println("Метод Эйлера")
	fmt.Printf("%f", EulerMethod(f, y0, a, b, h, E))
	println("Метод Эйлера-Коши")
	fmt.Printf("%f", EulerCauchyMethod(f, y0, a, b, h, E))
	println("Модифицированный метод Эйлера")
	fmt.Printf("%f", ModifiedEulerMethod(f, y0, a, b, h, E))
	println("Метод Рунге-Кутты 4-го порядка")
	fmt.Printf("%f", RungeKuttaMethodWithAccuracy(f, y0, a, b, h, E))

}
