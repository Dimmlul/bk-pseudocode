# 04 - Tracing

Tracing = jalanin program di kepala/kertas, catat nilai variabel tiap langkah.

## Caranya

1. Bikin tabel, satu kolom per variabel (+ kolom kondisi kalau ada loop).
2. Satu baris = satu putaran loop.
3. Tulis nilai **setelah** barisnya dijalanin.

Contoh FPB dengan a = 42, b = 15:

```
while b != 0 do
    r <- a mod b
    a <- b
    b <- r
endwhile
output(a)
```

| a  | b  | r  | b != 0? |
|----|----|----|---------|
| 42 | 15 | -  | ya      |
| 15 | 12 | 12 | ya      |
| 12 | 3  | 3  | ya      |
| 3  | 0  | 0  | tidak   |

Output: 3

## Latihan

Kerjain dulu, jawabannya di bawah.

**1.** x = 4728

```
a <- x div 1000
b <- (x div 100) mod 10
c <- (x div 10) mod 10
d <- x mod 10
output(a + d, b * c)
```

**2.** t = 7384

```
j <- t div 3600
m <- (t mod 3600) div 60
d <- t mod 60
output(j, m, d)
```

**3.** a = 7, b = 3, c = 10. Hasilnya true atau false?

```
(a) a mod b == 1 and not (c < a) or b > c
(b) not (a > b and b > c)
(c) (a > c or b > c) and a != b
```

**4.**

```
s <- 0
p <- 1
for i <- 1 to 4 do
    s <- s + i
    p <- p * 2
endfor
output(s, p)
```

**5.** x = 1234

```
b <- 0
while x > 0 do
    b <- b * 10 + x mod 10
    x <- x div 10
endwhile
output(b)
```

**6.**

```
x <- 1
n <- 0
repeat
    x <- x * 3
    n <- n + 1
until x > 50
output(n, x)
```

**7.** x = 100. Output A sama B beda nggak?

```
A:                          B:
while x < 50 do             repeat
    x <- x * 2                  x <- x * 2
endwhile                    until x >= 50
output(x)                   output(x)
```

**8.** Cari salahnya. Input hari = 2, jam 3 dan 4.

```
program RataNgoding
kamus
    hari, i, jam, total, rata : integer
algoritma
    input(hari)
    for i <- 1 to hari do
        input(jam)
        total <- total + jam
    endfor
    rata <- total / hari
    output(rata)
endprogram
```

## Jawaban

<details>
<summary>buka</summary>

1. a=4, b=7, c=2, d=8 → `12 14`
2. 7384 = 2 x 3600 + 184, 184 = 3 x 60 + 4 → `2 3 4`
3. (a) `(T and T) or F` = true, (b) `not (T and F)` = true, (c) `(F or F) and T` = false
4. s: 1, 3, 6, 10. p: 2, 4, 8, 16 → `10 16`
5. b: 4, 43, 432, 4321 → `4321`
6. x: 3, 9, 27, 81 → `4 81`
7. Beda. A = 100 (kondisi langsung false, loop nggak jalan). B = 200 (repeat jalan sekali dulu).
8. `total` belum di-set 0, terus `rata` harus real (7 / 2 = 3.5, kalau integer jadi 3).

</details>
