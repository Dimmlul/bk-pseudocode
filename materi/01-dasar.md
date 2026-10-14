# 01 - Dasar Pseudocode

## Format program

Semua jawaban formatnya kayak gini:

```
program LuasPersegi
kamus
    s, luas : integer
algoritma
    input(s)
    luas <- s * s
    output(luas)
endprogram
```

- `kamus` = tempat deklarasi semua variabel + tipenya. Kalau dipake di algoritma tapi nggak ada di kamus, salah.
- `algoritma` = langkah-langkahnya, dijalanin dari atas ke bawah.
- Konstanta ditulis di kamus juga: `constant PI : real = 3.14`

## Tipe data

- `integer` → bilangan bulat (`5`, `-3`)
- `real` → bilangan desimal (`3.14`)
- `char` → 1 karakter, petik satu (`'A'`)
- `string` → teks, petik dua (`"halo"`)
- `boolean` → `true` / `false`

## Operator

```
+  -  *        biasa
/              bagi, hasilnya real          7 / 2   = 3.5
div            bagi bulat (dibuang komanya) 7 div 2 = 3
mod            sisa bagi                    7 mod 2 = 1

==  !=  <  <=  >  >=      perbandingan, hasilnya boolean
and  or  not              logika
<-                        assignment (masukin nilai ke variabel)
```

Urutan pengerjaan: kurung → `not` → `* / div mod` → `+ -` → perbandingan → `and` → `or`.
Kalau ragu kasih kurung aja.

## div & mod (ini paling sering keluar)

Misal `x = 4728`:

```
x mod 10              = 8      satuan
(x div 10) mod 10     = 2      puluhan
(x div 100) mod 10    = 7      ratusan
x div 1000            = 4      ribuan
(x div 10) mod 100    = 72     dua digit tengah
x div 10              = 472    buang digit terakhir
```

Cek genap: `x mod 2 == 0`. Cek x faktor dari y: `y mod x == 0`.

Konversi detik ke jam/menit/detik:

```
jam   <- t div 3600
menit <- (t mod 3600) div 60
detik <- t mod 60
```

## Char

Char itu sebenernya angka (kode ASCII):

```
'0'..'9'  = 48..57
'A'..'Z'  = 65..90
'a'..'z'  = 97..122
```

Jadi bisa dibandingin: `'A' <= k and k <= 'Z'` artinya k huruf kapital.
Huruf kecil ke kapital tinggal dikurang 32 (`'a'` 97 → `'A'` 65).

## Analisis masalah (W5)

Sebelum nulis algoritma, tulis dulu:

```
Diketahui : rumus / aturan dari soal
Dibaca    : apa aja yang di-input
Proses    : rumusnya
Dicetak   : apa yang di-output
```

Ini ngebantu banget biar kamus-nya nggak ada yang kelewat.
