// program FungsiKuadrat
// menghitung f(x) = x² + 2x + 3
// kamus
//     x : integer
// algoritma
//     input(x)
//     output(x*x + 2*x + 3)
// endprogram

package main

import (
	"fmt"
)

func main() {
	var x int // kamus

	fmt.Scan(&x) // input, contoh: 2

	// Go nggak punya operator pangkat, x² ditulis x*x
	fmt.Println(x*x + 2*x + 3) // output: 11
}
