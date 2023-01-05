package main

func main() {
	tr := triangle{
		height: 12,
		base:   2,
	}
	sq := square{
		sideLength: 5,
	}
	printArea(tr)
	printArea(sq)
}
