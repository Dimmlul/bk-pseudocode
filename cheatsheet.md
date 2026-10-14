# Cheatsheet UTS

**Format**

```
program Nama
kamus
    a, b : integer
algoritma
    input(a, b)
    output(a + b)
endprogram
```

**Tipe**: integer, real, char `'A'`, string `"abc"`, boolean

**Operator**: `/` hasil real · `div` bagi bulat · `mod` sisa · `==` banding · `<-` assignment

**Digit** (x = 4728): satuan `x mod 10` · puluhan `(x div 10) mod 10` · ratusan `(x div 100) mod 10` · ribuan `x div 1000`

**Waktu**: jam `t div 3600` · menit `(t mod 3600) div 60` · detik `t mod 60`

**Char**: `'0'`=48 · `'A'`=65 · `'a'`=97 · kapital: `'A' <= k and k <= 'Z'`

**Boolean**: kabisat `t mod 400 == 0 or (t mod 4 == 0 and t mod 100 != 0)` · rentang `a <= x and x <= b`

**Loop**

```
for i <- 1 to n do      while k do        repeat
    ...                     ...               ...
endfor                  endwhile          until k
```

for = tau berapa kali · while = cek dulu (bisa 0x) · repeat = jalan dulu (min 1x), berhenti pas true

**Pola**: jumlah mulai 0 · kali mulai 1 · rata-rata pake real · ambil digit `x mod 10` lalu `x div 10` · FPB Euclid

**Sorting**: bubble = tukar sebelahan · selection = cari min taruh depan · insertion = selipin · merge = bagi dua lalu gabung · quick = pivot

**Graph**: BFS queue (per level) · DFS stack (dalem dulu) · Dijkstra jalur terpendek · Kruskal sisi terkecil tanpa siklus · TSP (n-1)!/2 rute

**Sebelum kumpul, cek**:
- semua variabel ada di kamus
- akumulator udah di-set awal
- tiap loop ada penutupnya
- while ada yang ngubah kondisi
- hasil bagi desimal disimpen di real
