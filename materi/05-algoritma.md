# 05 - Algoritma W1-W3

Minggu 1-3 belum ngoding, semuanya dikerjain manual di kertas.

## 4 pilar berpikir komputasional

- **Dekomposisi**: mecah masalah gede jadi kecil-kecil
- **Pengenalan pola**: nyari kesamaan / keteraturan
- **Abstraksi**: buang detail yang nggak penting
- **Algoritma**: nyusun langkah-langkah yang urut

## Greedy vs brute force

Greedy = tiap langkah ambil yang paling bagus saat itu. Cepet, tapi belum tentu hasil terbaik.
Brute force = coba semua kemungkinan. Pasti dapet yang terbaik, tapi lama.

## Tower of Hanoi

Pindahin n piringan dari A ke C:
1. pindahin n-1 piringan ke tiang bantu
2. pindahin piringan paling gede ke C
3. pindahin n-1 piringan dari tiang bantu ke C

Langkah minimal = 2ⁿ - 1 (3 piringan = 7 langkah, 4 = 15).

## Biner

Nilai tiap posisi: `... 32 16 8 4 2 1`

- 26 = 16 + 8 + 2 → `11010`
- `101101` = 32 + 8 + 4 + 1 = 45

## Sorting

Contoh data: `5 1 4 2 8`

**Bubble sort**: bandingin 2 angka sebelahan dari kiri ke kanan, tukar kalau kebalik. Ulang sampai nggak ada yang ditukar.

```
pass 1:  (5,1) tukar → 1 5 4 2 8
         (5,4) tukar → 1 4 5 2 8
         (5,2) tukar → 1 4 2 5 8
         (5,8) tetap → 1 4 2 5 8
pass 2:  (4,2) tukar → 1 2 4 5 8
pass 3:  nggak ada yang tukar, selesai
```

**Selection sort**: cari yang paling kecil, taruh di depan.

```
1 | 5 4 2 8     (1 paling kecil)
1 2 | 4 5 8     (2 tukar sama 5)
1 2 4 | 5 8
1 2 4 5 8
```

**Insertion sort**: ambil satu-satu, selipin ke posisi yang pas (kayak nyusun kartu).

```
5 | 1 4 2 8
1 5 | 4 2 8
1 4 5 | 2 8
1 2 4 5 | 8
1 2 4 5 8
```

**Merge sort**: bagi dua terus sampai sisa 1, terus gabungin balik sambil diurutin. Bentuknya kayak cermin.

```
[38 27 43 3 9 82 10 1]
[38 27 43 3]  [9 82 10 1]
[38 27] [43 3] [9 82] [10 1]
[38] [27] [43] [3] [9] [82] [10] [1]
[27 38] [3 43] [9 82] [1 10]
[3 27 38 43]  [1 9 10 82]
[1 3 9 10 27 38 43 82]
```

**Quick sort**: pilih pivot, yang lebih kecil ke kiri, yang lebih gede ke kanan. Ulang di tiap bagian.

```
[3 7 1 6 2 5 4]   pivot 4 → [3 1 2] 4 [7 5 6]
[3 1 2]           pivot 2 → [1] 2 [3]
[7 5 6]           pivot 6 → [5] 6 [7]
hasil: 1 2 3 4 5 6 7
```

## BFS & DFS

BFS pake **queue** (antrean, yang masuk duluan keluar duluan). Jalannya per level.
DFS pake **stack** (tumpukan, yang masuk terakhir keluar duluan). Jalannya masuk sedalam-dalamnya dulu.

Contoh graph: A ke B, C. B ke D, E. C ke F.

```
BFS:
queue [A]      → keluarin A, masukin B C
queue [B C]    → keluarin B, masukin D E
queue [C D E]  → keluarin C, masukin F
queue [D E F]  → D, E, F
urutan: A B C D E F

DFS (stack):
stack [A]      → ambil A, push B C
stack [B C]    → ambil C (paling atas), push F
stack [B F]    → ambil F
stack [B]      → ambil B, push D E
stack [D E]    → ambil E, terus D
urutan: A C F B E D
```

Node yang udah pernah masuk queue/stack jangan dimasukin lagi.

## Dijkstra (jalur terpendek)

1. Jarak titik awal = 0, sisanya ∞.
2. Pilih titik yang belum dikunjungi dengan jarak paling kecil.
3. Update jarak tetangganya kalau lewat titik ini jadi lebih deket. Catat lewat mana.
4. Ulang sampai semua titik dikunjungi.

Jalurnya dibaca mundur dari tujuan.

## Kruskal (minimum spanning tree)

1. Urutin semua sisi dari bobot paling kecil.
2. Ambil satu-satu. Kalau bikin siklus (muter), skip.
3. Berhenti kalau udah dapet (jumlah titik - 1) sisi.

## TSP

Brute force: coba semua rute keliling. Dari n kota ada (n-1)! rute, setengahnya cuma kebalikan (mirror),
jadi yang beda beneran (n-1)!/2. Contoh 5 kota = 12 rute.

Nearest neighbor: dari kota sekarang, pergi ke kota terdekat yang belum dikunjungin. Cepet tapi belum tentu paling pendek.

## Dynamic programming

Hasil yang udah dihitung disimpen di tabel, jadi nggak ngitung ulang.
Fibonacci pake tabel: `0 1 1 2 3 5 8 13 21 ...` (tiap angka = jumlah 2 angka sebelumnya).
