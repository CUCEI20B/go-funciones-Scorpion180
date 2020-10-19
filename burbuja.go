package main

func Burbuja(s []int64) {
	var dummy int64
	for i := 1; i < len(s); i++ {
		for j := 0; j < (len(s) - i); j++ {
			if s[j] > s[j+1] {
				dummy = s[j]
				s[j] = s[j+1]
				s[j+1] = dummy
			}
		}
	}
}

func main() {

}
