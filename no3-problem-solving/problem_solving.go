package no3problemsolving

import "fmt"

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func NewBebek(energi int, hidup, bisaTerbang bool, suaraTerbang string) Bebek {
	return Bebek{energi, hidup, bisaTerbang, suaraTerbang}
}

func Mati(b *Bebek) {
	b.hidup = false
}

func Terbang(b *Bebek) {
	if b.energi > 0 && b.hidup == true && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			Mati(b)
		}
	}
}

func Makan(b *Bebek) {
	if b.energi > 0 && b.hidup == true {
		b.energi += 1
	}
}

func Test() {
	bebek := NewBebek(120, true, true, "petok")
	fmt.Printf("Bebek awal %+v\n", bebek)
	Terbang(&bebek)
	fmt.Printf("Bebek setelah terbang %+v\n", bebek)
	Makan(&bebek)
	fmt.Printf("Bebek setelah makan %+v\n", bebek)
	Mati(&bebek)
	fmt.Printf("Bebek setalah mati %+v\n", bebek)
}
