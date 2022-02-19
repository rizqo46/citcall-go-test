# citcall-go-test
Created by Muhammad Abyan Rizqo

Repository link https://github.com/rizqo46/citcall-go-test.

This repo is part of Citcall Golang Test.

----

## No. 1 Pengolahan Data (Data Menupulation)
Dalam mengejakan soal ini digunakan fungsi table generator yang ada di repository pada file [__no1-data-manipulation/data_manipulation.go__](https://github.com/rizqo46/citcall-go-test/blob/main/no1-data-manipulation/data_manipulation.go)  pada fungsi __MakeTable()__. Dimana data diambil langsung dari https://citcall.com/test/countries.json, kemudian data diolah secara langsung menjadi table html. Untuk membuat tabel anda dapat run command berikut.

```
go run main.go
```
atau anda dapat menjalankan fungsi testing __TestMakeTable()__ jika anda menggunakan VsCode. Kemudian akan ada file baru (atau ditimpa bila file sudah ada) yang bernama __table.html__. Tabel hasil jadi sudah ada di [__no1-data-manipulation/table.html__](https://github.com/rizqo46/citcall-go-test/blob/main/no1-data-manipulation/table.html)

Cacatan: file __table.html__ dapat dihapus terlebih dahulu jika memang sudah ada dan ingin melihat fungsi generator berhasil atau tidak.

----

## No. 2 Performa (Performance)
Perhatikan bahwa pada Code A dan Code B hanya berbeda pada baris-baris berikut:

<table>
<tr>
<th> Code A </th>
<th> Code B </th>
</tr>
<tr>
<td>

```golang
...
size := len(tmp)
for i := 0; i == size; i++ {
...
```

</td>
<td>

```golang
...
for i := 0; i == len(tmp); i++ {
...
```

</td>
</tr>
</table>

Dimana pada Code A panjang dari tmp disimpan terlebih dahulu di variable __size__ (dihitung diluar loop) kemudian __size__-lah yang akan masuk pada kondisi loop setelahnya. Sedangkan pada Code B panjang dari tmp tidak dihitung diluar loop. Melainkan dihitung secara langsung.

Perhatikan bahwa pada loop teresbut kondisi awalnya adalah __i := 0__ dan akan berjalan jika __i == len(tmp)__. Namun berdasarkan code berikut:

```golang
...
for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}
...
```
Dapat disimpulkan bahwa __len(tmp) = 1000__. Sehingga kondisi loop 
``` golang 
for i := 0; i == len(tmp); i++ {
```
tidak terpenuhi sejak awal, sehingga loop tersebut tidak pernah berjalan.

Dengan menimbang kasus diatas dimana loop tidak berjalan. Menurut saya Code B lebih baik. Karena pada Code A __len(tmp)__ disimpan terlebih dahulu pada __size__ dan akan menambah memory yang terbuang percuma. Walaupun memory yang digunakan sangat kecil.

----

## No. 3 Problem Solving
```golang
type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func Mati(b Bebek) {
	b.hidup = false
}

func Terbang(b Bebek) {
	if b.energi > 0 && b.hidup == true && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			Mati(b)
		}
	}
}

func Makan(b Bebek) {
	if b.energi > 0 && b.hidup == true {
		b.energi += 1
	}
}
```
Fungsi-fungsi yang ada pada code diatas tidak akan dapat mengubah properti/atribut dari class __Bebek__, ini terjadi karena setiap fungsi __Mati__, __Makan__ dan __Terbang__ mendapatkan input berupa copy dari __Bebek__, kemudian di dalam setiap fungsi tersebut hasil copy dari __Bebek__-lah yang dimodifikasi. Sehingga objek __Bebek__ yang ada di alamat aslinya tetap tidak berubah.

Agar fungsi-fungsi tersebut dapat memodifikasi objek __Bebek__ aslinya maka kita perlu mengubah input dari fungsi-fungsi menjadi pointer dari __Bebek__ __(*Bebek)__. Sehingga fungsi-fungsi tersebut akan menjadi seperti berikut.

```golang
...
func Mati(b *Bebek) {
...

func Terbang(b *Bebek) {
...

func Makan(b *Bebek) {
...
```

Sehingga dapat diambil objek __Bebek__ dari alamat aslinya dan atribut yang ada di objek aslinya dapat berubah. Namun perlu diperhatikan bahwa input fungsi terebut harus berupa pinter dari __Bebek__ __(*Bebek)__.

Dapat dijalankan testing yang ada pada [__no3-problem-solving/problem_solving_test.go__](https://github.com/rizqo46/citcall-go-test/blob/main/no3-problem-solving/problem_solving_test.go) untuk mengetes apakah atribut dari __Bebek__ berhasil diubah. 



