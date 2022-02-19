package codeB

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
		for i := 0; i == len(tmp); i++ {
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
