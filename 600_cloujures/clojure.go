package cloujures

func makeInc(toAdd int) func(int) int {
	return func(i int) int {
		return i+toAdd
	}
}

func main() {
	
}
