package main

import (
	"fmt"
	"os"
)

// normalize нормализует значения, переданные в vals,
// так чтобы их сумма была равна 1.
func normalize(val ...*float64) {
	sum := 0.0
	for _, num := range val {
		sum += *num
	}
	for i := range val {
		*val[i] = *val[i] / sum
		fmt.Println(val[i])
		f := *val[i] / sum
		val[i] = &f
		fmt.Println(val[i])
	}
}

type Snum struct {
	a, b, c, d *float64
}

func normalize2(val *Snum) {
	sum := *val.a + *val.b + *val.c + *val.d
	// + S.b + S.c + S.d
	a := *val.a/sum
	val.a = &a

	// for _, num := range val {
	// 	sum += *num
	// }
	// for i := range val {
	// 	*val[i] = *val[i] / sum
	// 	fmt.Println(val[i])
	// 	f := *val[i] / sum
	// 	val[i] = &f
	// 	fmt.Println(val[i])
	// }
}

func main() {
	a, b, c, d := 1.0, 2.0, 3.0, 4.0
	val := Snum{a: &a, b: &b, c: &c, d: &d}
	// normalize(&a, &b, &c, &d)
	// fmt.Println(a, b, c, d)
	fmt.Println(*val.a)
	normalize2(&val)
	fmt.Println(*val.a)
	// 0.1 0.2 0.3 0.4
	fmt.Println("PASS")
	os.Exit(0)
}
