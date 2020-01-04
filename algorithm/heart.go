package main

func accumulation(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	print("accumulation:", accumulation(100))
}
