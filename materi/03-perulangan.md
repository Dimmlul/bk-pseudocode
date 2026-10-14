# 03 - Perulangan

## for

Dipake kalau udah tau mau ngulang berapa kali.

```
for i <- 1 to n do
    output(i)
endfor
```

Jumlah ulangan `for i <- a to b` = `b - a + 1`.

## while

Cek kondisi dulu, baru jalan. Bisa aja nggak jalan sama sekali.

```
while x > 0 do
    ...
endwhile
```

Di dalem loop **harus** ada yang ngubah kondisinya, kalau nggak jadi infinite loop.

## repeat-until

Jalan dulu, baru cek. Minimal jalan 1 kali. Berhenti pas kondisinya **true**.

```
repeat
    ...
until x > 50
```

Bedanya sama while: while jalan *selama* kondisi true, repeat jalan *sampai* kondisi true.

## Pola yang sering keluar

**Jumlah / total** (mulai dari 0)

```
jumlah <- 0
for i <- 1 to n do
    input(x)
    jumlah <- jumlah + x
endfor
```

**Perkalian / faktorial** (mulai dari 1, kalau 0 hasilnya 0 terus)

```
hasil <- 1
for i <- 1 to n do
    hasil <- hasil * i
endfor
```

**Rata-rata**: `rata <- jumlah / n`, dan `rata` harus `real`.

**Ambil digit satu-satu**

```
while x > 0 do
    d <- x mod 10       // digit paling kanan
    ...
    x <- x div 10       // buang digit paling kanan
endwhile
```

Dari pola ini bisa jadi:
- jumlah digit: `jumlah <- jumlah + x mod 10`
- balik angka: `b <- b * 10 + x mod 10` (1234 jadi 4321)
- hitung banyak digit: `cacah <- cacah + 1`

**Input terus sampai ketemu angka tertentu** (misal 0)

```
input(x)
while x != 0 do
    saldo <- saldo + x
    input(x)            // input lagi di akhir
endwhile
```

**FPB (Euclid)**

```
while b != 0 do
    r <- a mod b
    a <- b
    b <- r
endwhile
output(a)
```

KPK = a awal x b awal div FPB.

**Fibonacci**

```
a <- 0
b <- 1
for i <- 2 to n do
    c <- a + b
    a <- b
    b <- c
endfor
```

## Loop di dalam loop

```
for i <- 1 to 3 do
    for j <- 1 to i do
        output(i * j)
    endfor
endfor
```

Loop dalem jalan sampai habis dulu tiap 1 putaran loop luar. Total output di atas 1 + 2 + 3 = 6 kali.

Contoh soalnya: [soal/w06.md](../soal/w06.md), [soal/w07.md](../soal/w07.md).
