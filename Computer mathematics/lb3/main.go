package main

import (
	"fmt"
	"math"
)

func funnc_1(x float64) float64 {
	// подынтегральная функция
	return math.Pow(x, 3) + math.Pow(x, 2)
}

func funnc_2(x float64) float64 {
	// подынтегральная функция
	return math.Cos(3 * x)
}

func funnc_3(x float64) float64 {
	// подынтегральная функция
	return math.Pow(3, x) + 2
}

func rectangularRule(f func(x float64) float64, a, b float64, n int) float64 {
	// метод центральных прямоугольников
	h := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := a + (float64(i)+0.5)*h
		sum += f(x)
	}
	return sum * h
}

func trapezoidalRule(f func(x float64) float64, a, b float64, n int) float64 {
	// метод трапеций
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}
	return sum * h
}

func parabolicRule(f func(x float64) float64, a, b float64, n int) float64 {
	// метод парабол
	h := (b - a) / float64(n)
	sum := f(a) + f(b)
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}
	return sum * h / 3
}

func rungeRule(f func(f func(x float64) float64, a, b float64, n int) float64, fun func(x float64) float64, a, b float64, n int, eps float64) float64 {
	// метод Рунге
	I1 := f(fun, a, b, n)
	I2 := f(fun, a, b, 2*n)
	R := math.Abs(I2-I1) / 3
	if R < eps {
		return I2 + R
	}
	return rungeRule(f, fun, a, b, 2*n, eps)
}

func Gausse(a, b float64, n int, f func(x float64) float64) float64 {
	switch n {
	case 1:
		x := 0.0
		w := 2.0
		return w * f(((b-a)/2)*x+((b+a)/2)) * (b - a) / 2
	case 2:
		x := [2]float64{-0.5773502692, 0.5773502692}
		w := [2]float64{1.0, 1.0}
		sum := 0.0
		for i := 0; i < n; i++ {
			xi := ((b-a)/2)*x[i] + ((b + a) / 2)
			sum += w[i] * f(xi)
		}
		return sum * (b - a) / 2
	case 3:
		x := [3]float64{-0.7745966692, 0, 0.7745966692}
		w := [3]float64{0.5555555556, 0.8888888889, 0.5555555556}
		sum := 0.0
		for i := 0; i < n; i++ {
			xi := ((b-a)/2)*x[i] + ((b + a) / 2)
			sum += w[i] * f(xi)
		}
		return sum * (b - a) / 2
	case 4:
		x := [4]float64{-0.8611363116, -0.3399810436, 0.3399810436, 0.8611363116}
		w := [4]float64{0.3478548451, 0.6521451549, 0.6521451549, 0.3478548451}
		sum := 0.0
		for i := 0; i < n; i++ {
			xi := ((b-a)/2)*x[i] + ((b + a) / 2)
			sum += w[i] * f(xi)
		}
		return sum * (b - a) / 2
	default:
		return 0
	}
}

func main() {

	n := 10
	eps := 0.0001

	a_b := []float64{0.0, 3.0, 0.0, math.Pi / 3, 2.0, 0.0}

	I_rect := rungeRule(rectangularRule, funnc_1, a_b[0], a_b[1], n, eps)
	I_trap := rungeRule(trapezoidalRule, funnc_1, a_b[0], a_b[1], n, eps)
	I_parab := rungeRule(parabolicRule, funnc_1, a_b[0], a_b[1], n, eps)
	fmt.Printf("Интеграл от x^3 + x^2 на отрезке [%.2f, %.2f] с точностью %.4f\n", a_b[0], a_b[1], eps)
	fmt.Printf("Метод центральных прямоугольников: %.4f\n", I_rect)
	fmt.Printf("Метод трапеций: %.4f\n", I_trap)
	fmt.Printf("Метод парабол: %.4f\n", I_parab)

	I_rect_2 := rungeRule(rectangularRule, funnc_2, a_b[2], a_b[3], n, eps)
	I_trap_2 := rungeRule(trapezoidalRule, funnc_2, a_b[2], a_b[3], n, eps)
	I_parab_2 := rungeRule(parabolicRule, funnc_2, a_b[2], a_b[3], n, eps)
	fmt.Printf("\nИнтеграл от cos(3x) на отрезке [%.2f, %.2f] с точностью %.4f\n", a_b[2], a_b[3], eps)
	fmt.Printf("Метод центральных прямоугольников: %.4f\n", I_rect_2*-1)
	fmt.Printf("Метод трапеций: %.4f\n", I_trap_2*-1)
	fmt.Printf("Метод парабол: %.4f\n", I_parab_2)

	I_rect_3 := rungeRule(rectangularRule, funnc_3, a_b[4], a_b[5], n, eps)
	I_trap_3 := rungeRule(trapezoidalRule, funnc_3, a_b[4], a_b[5], n, eps)
	I_parab_3 := rungeRule(parabolicRule, funnc_3, a_b[4], a_b[5], n, eps)
	fmt.Printf("\nИнтеграл от 3^x + 2 на отрезке [%.2f, %.2f] с точностью %.4f\n", a_b[4], a_b[5], eps)
	fmt.Printf("Метод центральных прямоугольников: %.4f\n", I_rect_3*-1)
	fmt.Printf("Метод трапеций: %.4f\n", I_trap_3*-1)
	fmt.Printf("Метод парабол: %.4f\n", I_parab_3*-1)

	I_gauss_1 := Gausse(a_b[0], a_b[1], 1, funnc_1)
	I_gauss_2 := Gausse(a_b[0], a_b[1], 2, funnc_1)
	I_gauss_3 := Gausse(a_b[0], a_b[1], 3, funnc_1)
	I_gauss_4 := Gausse(a_b[0], a_b[1], 4, funnc_1)

	fmt.Printf("\nИнтеграл от x^3 + x^2  на отрезке [%.2f, %.2f] с точностью %.4f\n", a_b[0], a_b[1], eps)
	fmt.Printf("Метод Гаусса с 1 узлом: %.4f\n", I_gauss_1)
	fmt.Printf("Метод Гаусса с 2 узлами: %.4f\n", I_gauss_2)
	fmt.Printf("Метод Гаусса с 3 узлами: %.4f\n", I_gauss_3)
	fmt.Printf("Метод Гаусса с 4 узлами: %.4f\n", I_gauss_4)

	I_gauss_1_2 := Gausse(a_b[2], a_b[3], 1, funnc_2)
	I_gauss_2_2 := Gausse(a_b[2], a_b[3], 2, funnc_2)
	I_gauss_3_2 := Gausse(a_b[2], a_b[3], 3, funnc_2)
	I_gauss_4_2 := Gausse(a_b[2], a_b[3], 4, funnc_2)

	fmt.Printf("\nИнтеграл от cos(3x) на отрезке [%.2f, %.2f] с точностью %.4f\n", a_b[2], a_b[3], eps)
	fmt.Printf("Метод Гаусса с 1 узлом: %.4f\n", I_gauss_1_2*-1)
	fmt.Printf("Метод Гаусса с 2 узлами: %.4f\n", I_gauss_2_2*-1)
	fmt.Printf("Метод Гаусса с 3 узлами: %.4f\n", I_gauss_3_2*-1)
	fmt.Printf("Метод Гаусса с 4 узлами: %.4f\n", I_gauss_4_2*-1)

	I_gauss_1_3 := Gausse(a_b[4], a_b[5], 1, funnc_3)
	I_gauss_2_3 := Gausse(a_b[4], a_b[5], 2, funnc_3)
	I_gauss_3_3 := Gausse(a_b[4], a_b[5], 3, funnc_3)
	I_gauss_4_3 := Gausse(a_b[4], a_b[5], 4, funnc_3)

	fmt.Printf("\nИнтеграл от 3^x + 2 на отрезке [%.2f, %.2f] с точностью %.4f\n", a_b[4], a_b[5], eps)
	fmt.Printf("Метод Гаусса с 1 узлом: %.4f\n", I_gauss_1_3*-1)
	fmt.Printf("Метод Гаусса с 2 узлами: %.4f\n", I_gauss_2_3*-1)
	fmt.Printf("Метод Гаусса с 3 узлами: %.4f\n", I_gauss_3_3*-1)
	fmt.Printf("Метод Гаусса с 4 узлами: %.4f\n", I_gauss_4_3*-1)

}
