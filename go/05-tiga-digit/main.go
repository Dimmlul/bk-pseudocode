// program TigaDigit
// memecah bilangan 3 digit jadi ratusan, puluhan, satuan
// kamus
//     x, d1, d2, d3 : integer
// algoritma
//     input(x)
//     d1 <- x div 100
//     d2 <- (x div 10) mod 10
//     d3 <- x mod 10
//     output(d1, d2, d3)
// endprogram

package main

import (
	"fmt"
)

func main() {
	var x, d1, d2, d3 int // kamus

	fmt.Scan(&x) // input cukup x aja, contoh: 472

	// int dibagi int di Go hasilnya bulat, jadi / di sini = div
	d1 = x / 100       // ratusan: 472 / 100 = 4
	d2 = (x / 10) % 10 // puluhan: 47 % 10 = 7
	d3 = x % 10        // satuan:  472 % 10 = 2

	fmt.Println(d1, d2, d3) // output: 4 7 2
}
