// program Modulo
// kamus
//     a, b, hasil : integer
// algoritma
//     input(a, b)
//     hasil <- a mod b
//     output(hasil)
// endprogram

package main

import (
	"fmt"
)

func main() {
	var a, b int // kamus
	var hasil int

	fmt.Scan(&a, &b) // input, contoh: 17 5

	// mod 0 bikin program error (panic), jadi dicek dulu
	if b == 0 {
		fmt.Println("b tidak boleh 0")
		return
	}

	hasil = a % b      // % di Go = mod di pseudocode
	fmt.Println(hasil) // output: 2
}
