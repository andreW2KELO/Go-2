package main

// В Go есть несколько встроенных типов данных, в том числе Complex, Float, Integer, Ordered, Signed и Unsigned:

// Тип Integer — это целочисленные значения.
// Может быть представлен в виде int, int8, int16, int32 или int64, в зависимости от выбранной точности

// Тип Ordered — это упорядоченные значения.
// Может быть представлен в виде byte, rune, int, int16, int32, int64, float32 или float64

// Тип Complex — это комплексное число из действительной и мнимой частей.
// Может быть представлен в виде complex64 или complex128, в зависимости от выбранной точности

type MyConstraint interface {
	int | int8 | int16 | int32 | int64
}

func MyGeneric[T MyConstraint](x T) {
	// ...
}

func main() {
	// Получится, потому что тип int входит в список ограничений
	// (int | int8 | int16 | int32 | int64)
	MyGeneric[int](1)

	// Не получится, потому что string  — не один из типов интерфейса MyConstraint
	//MyGeneric[string]("hello")
}
