# 02 - Boolean

## Tabel kebenaran

```
p  q  | p and q | p or q | not p
T  T  |    T    |   T    |   F
T  F  |    F    |   T    |   F
F  T  |    F    |   T    |   T
F  F  |    F    |   F    |   T
```

`and` → true kalau dua-duanya true.
`or` → true kalau salah satu true.

De Morgan (kadang ditanya):
`not (p and q)` sama dengan `not p or not q`,
`not (p or q)` sama dengan `not p and not q`.

## Soal output true/false nggak perlu if

Soal-soal W4-W5 yang outputnya boolean cukup langsung di-output ekspresinya:

```
output(tahun mod 400 == 0 or (tahun mod 4 == 0 and tahun mod 100 != 0))
```

Contoh ekspresi lain yang sering kepake:

```
segitiga   <- a + b > c and a + c > b and b + c > a
kapital    <- 'A' <= k and k <= 'Z'
angka      <- '0' <= k and k <= '9'
lulus      <- (nilai > 55 and tugas) or nilai > 90
diRentang  <- 1 <= x and x <= 10
```

## Hati-hati

- `1 <= x <= 10` itu **salah** di pseudocode. Harus dipecah: `1 <= x and x <= 10`.
- `=` bukan perbandingan. Bandingin pakai `==`, assignment pakai `<-`.
- Kalau ada `mod` sama variabel yang bisa 0, cek dulu: `d != 0 and x mod d == 0`.

Contoh soalnya: [soal/w04.md](../soal/w04.md) nomor 8-14, [soal/w05.md](../soal/w05.md).
