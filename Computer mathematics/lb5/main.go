package main

import (
	"fmt"
	"math"
)

// Определяем функции уравнений системы и их производные
func f1(x, y float64) float64 {
	return math.Cos(x+1) - y - 0.5
}

func f2(x, y float64) float64 {
	return x*x - y*y - 1
}

func df1dx(x, y float64) float64 {
	return -math.Sin(x + 1)
}

func df1dy(x, y float64) float64 {
	return -1
}

func df2dx(x, y float64) float64 {
	return 2 * x
}

func df2dy(x, y float64) float64 {
	return -2 * y
}

// Функция для решения системы уравнений методом Ньютона
func newtonSystem(f1, f2, df1dx, df1dy, df2dx, df2dy func(float64, float64) float64, x0, y0, eps float64, n int) (float64, float64, bool) {
	var x, y float64

	for i := 1; i <= n; i++ {
		// Определяем матрицу Якоби и ее обратную
		j11, j12, j21, j22 := df1dx(x0, y0), df1dy(x0, y0), df2dx(x0, y0), df2dy(x0, y0)
		detJ := j11*j22 - j12*j21
		jInv11, jInv12, jInv21, jInv22 := j22/detJ, -j12/detJ, -j21/detJ, j11/detJ

		// Рассчитываем значение функций в текущей точке
		f := []float64{f1(x0, y0), f2(x0, y0)}

		// Вычисляем приращение
		delta := []float64{
			jInv11*f[0] + jInv12*f[1],
			jInv21*f[0] + jInv22*f[1],
		}

		// Обновляем значения переменных
		x, y = x0-delta[0], y0-delta[1]

		// Проверяем достижение заданной точности
		if math.Abs(delta[0])+math.Abs(delta[1]) < eps {
			return x, y, true
		}

		// Обновляем начальные значения для следующей итерации
		x0, y0 = x, y
	}

	// Если не удалось достичь заданной точности за n итераций, возвращаем false
	return x, y, false
}

func main() {
	// Задаем начальное приближение, точность решения и максимальное количество итераций
	x0 := 0.7529
	y0 := 0.6581
	eps := 0.0001
	n := 100

	// Решаем систему уравнений методом Ньютона
	x, y, converged := newtonSystem(f1, f2, df1dx, df1dy, df2dx, df2dy, x0, y0, eps, n)
	// Печатаем результаты
	if converged {
		fmt.Printf("Solution converged to x = %f, y = %f\n", x, y)
	} else {
		fmt.Println("Solution did not converge")
	}
}
