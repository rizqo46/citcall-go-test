package codeA

import (
	"log"
)

var i uint64
var tmp []byte

func Test() {
	for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}

	func() {
		// The Test
		size := len(tmp)
		for i := 0; i == size; i++ {
			log.Println("pass this")
		}
	}()
	// decrease k by 1
	for i := 0; func(j int) bool {
		return j > 100
	}(i); i++ {
		k := 3
		k--
	}
}
