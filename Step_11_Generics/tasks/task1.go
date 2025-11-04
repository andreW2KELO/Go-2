package main

type MyType interface {
	int | int8 | int16 | int32 | int64 | float64 | float32
}

func Sum[T MyType](x []T) T {
	var amount T
	for _, v := range x {
		amount += v
	}
	return amount
}
