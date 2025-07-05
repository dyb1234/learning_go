package main

func plus10(p *int) {
	*p += 10
}

func multiply2(p *[]int) {
	for i := 0; i < len(*p); i++ {
		(*p)[i] *= 2
	}
}
