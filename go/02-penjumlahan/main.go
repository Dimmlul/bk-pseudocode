// program Penjumlahan
// kamus
//     a, b, hasil : integer
// algoritma
//     input(a, b)
//     hasil <- a + b
//     output(hasil)
// endprogram

package main

import (
	"fmt"
)

func main() {
	var a, b int // kamus
	var hasil int

	fmt.Scan(&a, &b) // input, contoh: 3 4
	hasil = a + b
	fmt.Println(hasil) // output: 7
}
