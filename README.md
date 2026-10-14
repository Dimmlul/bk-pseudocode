# bk-pseudocode

Catatan matkul **Berpikir Komputasional & Pengenalan Pemrograman (CBK1BAB4)**.
Fokusnya ke pseudocode karena ujiannya tulis dikertas :D
Kode Go cuma buat nyoba pseudocode-nya beneran jalan.

## Mulai dari mana?

| Mau ngapain | Buka |
|---|---|
| Belajar dari awal | [`materi/01-dasar.md`](materi/01-dasar.md), terus lanjut urut sampai 05 |
| Latihan soal | [`soal/`](soal) (satu file per minggu) |
| Persiapan ujian | [`cheatsheet.md`](cheatsheet.md) |
| Nyoba jalanin di Go | [`go/`](go) |

## Isi folder

```
materi/      catatan teori, dibaca urut 01 → 05
soal/        soal + jawaban pseudocode per minggu (w04, w05, ...)
go/          latihan Go, satu folder = satu program
cheatsheet   ringkasan 1 halaman
```

## Keyword pseudocode

| Keyword | Artinya |
|---|---|
| `program` | awal program + nama programnya |
| `kamus` | daftar variabel + tipe datanya |
| `algoritma` | langkah-langkah yang dijalanin |
| `endprogram` | akhir program |
| `<-` | masukin nilai ke variabel (`a <- 5`) |
| `input()` / `output()` | baca / tampilin data |
| `div` / `mod` | bagi bulat / sisa bagi |
| `for ... endfor` | ulang sebanyak n kali |
| `while ... endwhile` | ulang selama kondisinya true |
| `repeat ... until` | ulang sampai kondisinya true |
| `if ... endif` | percabangan |

Contoh:

```
program Luas
kamus
    s, luas : integer
algoritma
    input(s)
    luas <- s * s
    output(luas)
endprogram
```

## Keyword Go

| Go | Artinya | Kalau di pseudocode |
|---|---|---|
| `package main` | penanda file ini program yang bisa dijalanin | `program` |
| `import "fmt"` | pakai library `fmt` buat input/output | - |
| `func main()` | tempat program mulai jalan | `algoritma` |
| `var a int` | bikin variabel | `kamus` |
| `fmt.Println()` | cetak + pindah baris | `output()` |
| `fmt.Scanln()` | baca input satu baris | `input()` |
| `=` | assignment | `<-` |
| `/` dan `%` | bagi (bulat kalau int) dan sisa | `div` dan `mod` |

Contoh yang sama dalam Go:

```go
package main

import "fmt"

func main() {
    var s, luas int
    fmt.Scanln(&s)
    luas = s * s
    fmt.Println(luas)
}
```

> Satu folder cuma boleh punya satu `func main`, makanya tiap program di `go/` punya folder sendiri.

## Cara jalanin program Go

### 1. Buka terminal di folder repo

Paling gampang lewat VS Code: buka folder `bk-pseudocode`, terus `Terminal > New Terminal`.
Terminalnya otomatis udah ada di root repo, kira-kira kayak gini:

```
PS D:\...\bk-pseudocode>
```

Cek isinya pakai `ls`. Kalau yang keluar `go`, `materi`, `soal`, `README.md`, berarti udah di tempat yang bener.

### 2. Pahami letak foldernya

Program Go **nggak** ada di root, tapi di dalam folder `go`:

```
bk-pseudocode/          <- kamu di sini (root)
├── materi/
├── soal/
└── go/
    ├── 01-hello-world/
    ├── 02-penjumlahan/
    ├── 03-modulo/
    ├── 04-fungsi-kuadrat/
    └── 05-tiga-digit/
```

Jadi kalau langsung `cd 05-tiga-digit` dari root, hasilnya error `Cannot find path`, soalnya harus masuk ke `go` dulu.

### 3. Jalanin programnya

**Cara A: langsung dari root (paling gampang)**

```powershell
go run ./go/05-tiga-digit
```

Terminal bakal diem nunggu input. Ketik angkanya (misal `472`), tekan Enter, keluar `4 7 2`.

**Cara B: masuk ke foldernya dulu**

```powershell
cd go               # masuk ke folder go
cd 05-tiga-digit    # masuk ke folder programnya
go run .            # titik = jalanin program di folder ini
```

Bisa juga sekaligus: `cd go\05-tiga-digit`

Balik lagi ke root:

```powershell
cd ..\..            # mundur 2 folder (05-tiga-digit -> go -> root)
```

### Perintah navigasi

| Perintah | Fungsi |
|---|---|
| `ls` | lihat isi folder sekarang |
| `cd nama-folder` | masuk ke folder |
| `cd ..` | mundur 1 folder |
| `cd ..\..` | mundur 2 folder |
| `pwd` | cek sekarang lagi di folder mana |

Tips: ketik sebagian nama folder terus tekan **Tab**, nanti dilengkapin sendiri.
Misal `cd go\05` + Tab jadi `cd .\go\05-tiga-digit\`.

### Daftar program

| Jalanin dari root | Isi | Contoh input → output |
|---|---|---|
| `go run ./go/01-hello-world` | output pertama | - → `Hello, World!` |
| `go run ./go/02-penjumlahan` | jumlahin 2 angka | `3 4` → `7` |
| `go run ./go/03-modulo` | sisa bagi (`mod`) | `17 5` → `2` |
| `go run ./go/04-fungsi-kuadrat` | f(x) = x² + 2x + 3 | `2` → `11` |
| `go run ./go/05-tiga-digit` | pecah ratusan, puluhan, satuan | `472` → `4 7 2` |

Tiap file ada pseudocode-nya di komentar paling atas.

### Error yang sering muncul

| Error | Artinya | Solusi |
|---|---|---|
| `Cannot find path ...` | foldernya nggak ada di tempat kamu sekarang | cek pakai `ls`, mungkin lupa masuk `go` dulu |
| `go: cannot find main module` | kamu di luar folder repo | `cd` ke folder `bk-pseudocode` dulu |
| `'go' is not recognized` | Go belum ke-install | install dari https://go.dev/dl, terus buka ulang terminal |
| program diem aja | lagi nunggu input | ketik angkanya, tekan Enter |

## Catatan

- Penulisan pseudocode ngikutin gaya di kelas. Kalau di kertas, `<-` biasanya ditulis `←`.
- Kalau ada yang salah atau beda sama versi kelas kalian, buka issue aja.
