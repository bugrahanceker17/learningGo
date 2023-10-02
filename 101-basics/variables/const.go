package main

const sabit1 = "sabit 1"

const (
	sabit2 = "sabit 2"
	sabit3 = "sabit 3"
	sabit4 = "sabit 4"
)

// iota bir enum sağlar ve değerleri birer birer arttırır. bu durumda sabit5=0, sabit6=1, sabit7=2 şeklinde olur
const (
	sabit5 = iota
	sabit6
	sabit7
)

func main() {
	println(sabit1, sabit2, sabit3, sabit4, sabit5, sabit6, sabit7)
}
