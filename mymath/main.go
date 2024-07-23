package mymath

import "math"

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Abs - функция для вычисления абсолютного значения числа
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Ceil - функция для округления числа вверх
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor - функция для округления числа вниз
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Max - функция для нахождения максимального значения из двух чисел
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min - функция для нахождения минимального значения из двух чисел
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

// Pow - функция для возведения числа в степень
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Sin - функция для вычисления синуса угла
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Cos - функция для вычисления косинуса угла
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Tan - функция для вычисления тангенса угла
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Acos - функция для вычисления арккосинуса
func Acos(x float64) float64 {
	return math.Acos(x)
}

// Asin - функция для вычисления арксинуса
func Asin(x float64) float64 {
	return math.Asin(x)
}

// Atan - функция для вычисления арктангенса
func Atan(x float64) float64 {
	return math.Atan(x)
}

// Atan2 - функция для вычисления арктангенса двух аргументов
func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Log - функция для вычисления натурального логарифма
func Log(x float64) float64 {
	return math.Log(x)
}

// Log10 - функция для вычисления десятичного логарифма
func Log10(x float64) float64 {
	return math.Log10(x)
}

// Log2 - функция для вычисления двоичного логарифма
func Log2(x float64) float64 {
	return math.Log2(x)
}

// Exp - функция для вычисления экспоненциальной функции
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Signbit - функция для проверки бита знака числа
func Signbit(x float64) bool {
	return math.Signbit(x)
}

// Inf - функция для получения положительной бесконечности
func Inf(x int) float64 {
	return math.Inf(x)
}

// NaN - функция для получения значения "не число" (NaN)
func NaN() float64 {
	return math.NaN()
}
