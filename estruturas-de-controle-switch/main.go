package main

func main() {
	j := 10
	switch j {
	case 11:
		println("here: 11")
		break
	default:
		println("here default")
		break
	}

	for i := 0; i <= 10; i++ {

		switch i {
		case 5:
			goto LABELS
		case i:
			println("i: ", i)
			break
		default:
			println("default: ", i)
		}
	}

LABELS:
	f()

}

func f() {
	println("goto fim")
}
